// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// Package queue 轻量消息队列，支持 Redis List / Disk 双驱动
// 生产者通过 Push/DelayPush 投递，消费者实现 Consumer 接口并 Register
package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 驱动类型 ====================

const (
	DriverRedis = "redis"
	DriverDisk  = "disk"
)

// ==================== 消息结构 ====================

// Message 队列消息
type Message struct {
	Topic     string `json:"topic"`
	Body      string `json:"body"`
	Retry     int    `json:"retry"`     // 已重试次数
	MaxRetry  int    `json:"maxRetry"`  // 最大重试次数（0=不重试，由消费者通过 RetryError 控制）
	CreatedAt int64  `json:"createdAt"` // 投递时间戳（秒）
}

// RetryError 消费者返回此错误类型时才会重入队，普通 error 只记日志不重试
type RetryError struct {
	Err error
}

func (e *RetryError) Error() string { return e.Err.Error() }

// NewRetryError 创建可重试错误（消费者 Handle 中使用）
func NewRetryError(err error) *RetryError {
	return &RetryError{Err: err}
}

// ==================== 驱动接口 ====================

// Driver 队列驱动接口
type Driver interface {
	Push(ctx context.Context, topic string, msg *Message) error
	Pop(ctx context.Context, topic string, timeout time.Duration) (*Message, error)
	Len(ctx context.Context, topic string) (int64, error)
	Close() error
}

// ==================== 消费者接口 ====================

// Consumer 消费者接口
type Consumer interface {
	GetTopic() string
	Handle(ctx context.Context, msg *Message) error
}

// ==================== 全局实例 ====================

var (
	driver     Driver
	driverName string
	consumers  = struct {
		sync.RWMutex
		list []Consumer
	}{}
	running    bool
	cancelFunc context.CancelFunc
	mu         sync.Mutex
)

// Init 初始化队列（服务启动时调用）
func Init(ctx context.Context) {
	cfg := g.Cfg().MustGet(ctx, "queue")
	if cfg.IsNil() || cfg.IsEmpty() {
		g.Log().Info(ctx, "[queue] no config found, using redis driver with defaults")
		driverName = DriverRedis
	} else {
		driverName = cfg.Map()["driver"].(string)
		if driverName == "" {
			driverName = DriverRedis
		}
	}

	switch driverName {
	case DriverDisk:
		path := "storage/queue"
		if cfg != nil && !cfg.IsNil() {
			if p, ok := cfg.Map()["diskPath"]; ok && p != nil {
				path = fmt.Sprintf("%v", p)
			}
		}
		driver = NewDiskDriver(path)
	default:
		driver = NewRedisDriver()
		driverName = DriverRedis
	}

	g.Log().Infof(ctx, "[queue] initialized, driver=%s", driverName)
}

// ==================== 延迟队列接口 ====================

// DelayDriver 延迟队列驱动（可选实现）
type DelayDriver interface {
	DelayPush(ctx context.Context, topic string, msg *Message, executeAt int64) error
	PollDelay(ctx context.Context, topic string) (*Message, error) // 取出到期消息
}

// ==================== 生产者 API ====================

// DelayPush 延迟投递消息（delay 秒后才会被消费）
func DelayPush(topic string, body interface{}, delaySec int64) error {
	if driver == nil {
		return fmt.Errorf("queue not initialized")
	}
	dd, ok := driver.(DelayDriver)
	if !ok {
		return fmt.Errorf("current driver '%s' does not support delay queue", driverName)
	}

	bodyStr, err := marshalBody(body)
	if err != nil {
		return err
	}

	msg := &Message{
		Topic:     topic,
		Body:      bodyStr,
		MaxRetry:  3,
		CreatedAt: time.Now().Unix(),
	}
	executeAt := time.Now().Unix() + delaySec
	return dd.DelayPush(context.Background(), topic, msg, executeAt)
}

// Push 投递消息到队列（即时消费）
func Push(topic string, body interface{}) error {
	if driver == nil {
		return fmt.Errorf("queue not initialized")
	}
	bodyStr, err := marshalBody(body)
	if err != nil {
		return err
	}
	msg := &Message{
		Topic:     topic,
		Body:      bodyStr,
		MaxRetry:  3,
		CreatedAt: time.Now().Unix(),
	}
	return driver.Push(context.Background(), topic, msg)
}

func marshalBody(body interface{}) (string, error) {
	switch v := body.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("marshal body failed: %v", err)
		}
		return string(b), nil
	}
}

// ==================== 消费者注册 ====================

// Register 注册消费者（在 init() 中调用）
func Register(c Consumer) {
	consumers.Lock()
	defer consumers.Unlock()
	consumers.list = append(consumers.list, c)
	g.Log().Infof(context.Background(), "[queue] consumer registered: %s", c.GetTopic())
}

// ==================== 消费者启动 ====================

// StartConsumers 启动所有消费者监听（服务启动后调用）
func StartConsumers(ctx context.Context) {
	mu.Lock()
	defer mu.Unlock()

	if running {
		return
	}

	if driver == nil {
		g.Log().Warning(ctx, "[queue] driver not initialized, skip consumers")
		return
	}

	consumers.RLock()
	list := make([]Consumer, len(consumers.list))
	copy(list, consumers.list)
	consumers.RUnlock()

	if len(list) == 0 {
		g.Log().Info(ctx, "[queue] no consumers registered")
		return
	}

	ctx2, cancel := context.WithCancel(ctx)
	cancelFunc = cancel
	running = true

	for _, c := range list {
		go listenConsumer(ctx2, c)
	}

	// 启动消费速率统计收集器
	go startMetricsCollector(ctx2)

	// 如果驱动支持延迟队列，启动延迟轮询协程
	if dd, ok := driver.(DelayDriver); ok {
		go pollDelayQueues(ctx2, dd, list)
	}

	g.Log().Infof(ctx, "[queue] %d consumers started", len(list))
}

// StopConsumers 停止所有消费者
func StopConsumers() {
	mu.Lock()
	defer mu.Unlock()
	if cancelFunc != nil {
		cancelFunc()
	}
	running = false
	if driver != nil {
		_ = driver.Close()
	}
	g.Log().Info(context.Background(), "[queue] all consumers stopped")
}

func listenConsumer(ctx context.Context, c Consumer) {
	topic := c.GetTopic()
	g.Log().Infof(ctx, "[queue] consumer listening: %s", topic)

	for {
		select {
		case <-ctx.Done():
			g.Log().Infof(ctx, "[queue] consumer stopped: %s", topic)
			return
		default:
		}

		msg, err := driver.Pop(ctx, topic, 2*time.Second)
		if err != nil {
			// 超时或空，继续
			continue
		}
		if msg == nil {
			continue
		}

		handleStart := time.Now()
		if err := c.Handle(ctx, msg); err != nil {
			recordConsumeMetrics(topic, time.Since(handleStart).Milliseconds(), false)
			// 判断是否为可重试错误
			if retryErr, ok := err.(*RetryError); ok {
				g.Log().Warningf(ctx, "[queue] consumer '%s' retry error: %v", topic, retryErr.Err)
				maxRetry := msg.MaxRetry
				if maxRetry <= 0 {
					maxRetry = 3 // 默认最多重试 3 次
				}
				if msg.Retry < maxRetry {
					msg.Retry++
					if pushErr := driver.Push(ctx, topic, msg); pushErr != nil {
						g.Log().Errorf(ctx, "[queue] retry push failed for '%s': %v", topic, pushErr)
					} else {
						g.Log().Infof(ctx, "[queue] message re-queued for '%s', retry=%d/%d", topic, msg.Retry, maxRetry)
					}
				} else {
					// 超过重试次数，进死信队列
					deadTopic := topic + ":dead"
					if pushErr := driver.Push(ctx, deadTopic, msg); pushErr != nil {
						g.Log().Errorf(ctx, "[queue] dead letter push failed for '%s': %v", topic, pushErr)
					} else {
						g.Log().Warningf(ctx, "[queue] message moved to dead letter '%s', retry exhausted", deadTopic)
					}
				}
			} else {
				// 普通错误：只记日志，不重试
				g.Log().Errorf(ctx, "[queue] consumer '%s' handle failed (no retry): %v", topic, err)
			}
		} else {
			recordConsumeMetrics(topic, time.Since(handleStart).Milliseconds(), true)
		}
	}
}

// pollDelayQueues 轮询所有延迟队列，到期消息转入普通队列
func pollDelayQueues(ctx context.Context, dd DelayDriver, list []Consumer) {
	g.Log().Info(ctx, "[queue] delay queue poller started")
	for {
		select {
		case <-ctx.Done():
			g.Log().Info(ctx, "[queue] delay queue poller stopped")
			return
		default:
		}

		moved := 0
		for _, c := range list {
			topic := c.GetTopic()
			msg, err := dd.PollDelay(ctx, topic)
			if err != nil {
				g.Log().Errorf(ctx, "[queue] delay poll error for '%s': %v", topic, err)
				continue
			}
			if msg == nil {
				continue
			}
			// 到期了，转入普通队列
			if err := driver.Push(ctx, topic, msg); err != nil {
				g.Log().Errorf(ctx, "[queue] delay->normal push failed for '%s': %v", topic, err)
			} else {
				g.Log().Infof(ctx, "[queue] delay message ready for '%s', moved to normal queue", topic)
				moved++
			}
		}
		if moved == 0 {
			// 没有到期消息，歇一会
			time.Sleep(1 * time.Second)
		}
	}
}

// ==================== 消费速率统计 ====================

// topicMetrics 每个 topic 的消费指标
type topicMetrics struct {
	consumed  atomic.Int64 // 当前窗口消费条数
	failed    atomic.Int64 // 当前窗口失败条数
	totalMs   atomic.Int64 // 当前窗口总耗时(ms)
	lastRate  float64      // 上一轮每分钟消费速率
	lastAvgMs float64      // 上一轮平均耗时(ms)
}

var (
	metricsMap = make(map[string]*topicMetrics)
	metricsMu  sync.RWMutex
)

func getMetrics(topic string) *topicMetrics {
	metricsMu.RLock()
	m := metricsMap[topic]
	metricsMu.RUnlock()
	if m != nil {
		return m
	}
	metricsMu.Lock()
	defer metricsMu.Unlock()
	if m = metricsMap[topic]; m != nil {
		return m
	}
	m = &topicMetrics{}
	metricsMap[topic] = m
	return m
}

func recordConsumeMetrics(topic string, takeMs int64, success bool) {
	m := getMetrics(topic)
	m.consumed.Add(1)
	m.totalMs.Add(takeMs)
	if !success {
		m.failed.Add(1)
	}
}

// startMetricsCollector 每分钟汇总一次速率
func startMetricsCollector(ctx context.Context) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			metricsMu.Lock()
			for _, m := range metricsMap {
				cnt := m.consumed.Swap(0)
				ms := m.totalMs.Swap(0)
				m.failed.Swap(0)
				m.lastRate = float64(cnt) // 条/分钟
				if cnt > 0 {
					m.lastAvgMs = float64(ms) / float64(cnt)
				} else {
					m.lastAvgMs = 0
				}
			}
			metricsMu.Unlock()
		}
	}
}

// ==================== 积压告警 ====================

// AlertConfig 积压告警配置
type AlertConfig struct {
	Enabled       bool  // 是否启用
	Threshold     int64 // 积压阈值
	DeadThreshold int64 // 死信阈值
}

var alertCfg = AlertConfig{
	Enabled:       true,
	Threshold:     100,
	DeadThreshold: 10,
}

// SetAlertConfig 设置告警配置（可在外部调整）
func SetAlertConfig(cfg AlertConfig) {
	alertCfg = cfg
}

// CheckAlerts 检查所有队列积压情况，返回告警列表
func CheckAlerts(ctx context.Context) []string {
	if !alertCfg.Enabled || driver == nil {
		return nil
	}

	consumers.RLock()
	list := make([]Consumer, len(consumers.list))
	copy(list, consumers.list)
	consumers.RUnlock()

	var alerts []string
	for _, c := range list {
		topic := c.GetTopic()
		pending, _ := driver.Len(ctx, topic)
		deadSize, _ := driver.Len(ctx, topic+":dead")

		if alertCfg.Threshold > 0 && pending > alertCfg.Threshold {
			alert := fmt.Sprintf("队列 %s 积压 %d 条（阈值 %d）", topic, pending, alertCfg.Threshold)
			alerts = append(alerts, alert)
			g.Log().Warningf(ctx, "[queue:alert] %s", alert)
		}
		if alertCfg.DeadThreshold > 0 && deadSize > alertCfg.DeadThreshold {
			alert := fmt.Sprintf("队列 %s 死信 %d 条（阈值 %d）", topic, deadSize, alertCfg.DeadThreshold)
			alerts = append(alerts, alert)
			g.Log().Warningf(ctx, "[queue:alert] %s", alert)
		}
	}
	return alerts
}

// ==================== 管理接口 ====================

// TopicStats 队列统计
type TopicStats struct {
	Topic     string  `json:"topic"`
	Pending   int64   `json:"pending"`
	DeadSize  int64   `json:"deadSize"`
	Rate      float64 `json:"rate"`      // 每分钟消费速率
	AvgTakeMs float64 `json:"avgTakeMs"` // 平均耗时(ms)
}

// GetStats 获取所有消费者队列的统计
func GetStats(ctx context.Context) []TopicStats {
	consumers.RLock()
	list := make([]Consumer, len(consumers.list))
	copy(list, consumers.list)
	consumers.RUnlock()

	stats := make([]TopicStats, 0, len(list))
	for _, c := range list {
		topic := c.GetTopic()
		pending, _ := driver.Len(ctx, topic)
		deadSize, _ := driver.Len(ctx, topic+":dead")

		var rate, avgMs float64
		metricsMu.RLock()
		if m, ok := metricsMap[topic]; ok {
			rate = m.lastRate
			avgMs = m.lastAvgMs
		}
		metricsMu.RUnlock()

		stats = append(stats, TopicStats{
			Topic:     topic,
			Pending:   pending,
			DeadSize:  deadSize,
			Rate:      rate,
			AvgTakeMs: avgMs,
		})
	}
	return stats
}

// GetDriverName 获取当前驱动名
func GetDriverName() string {
	return driverName
}

// GetRegisteredTopics 获取已注册的 topic 列表
func GetRegisteredTopics() []string {
	consumers.RLock()
	defer consumers.RUnlock()
	topics := make([]string, 0, len(consumers.list))
	for _, c := range consumers.list {
		topics = append(topics, c.GetTopic())
	}
	return topics
}
