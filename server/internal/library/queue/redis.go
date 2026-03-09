// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// RedisDriver 基于 Redis List 的队列驱动
type RedisDriver struct {
	prefix string
}

// NewRedisDriver 创建 Redis 驱动
func NewRedisDriver() *RedisDriver {
	return &RedisDriver{prefix: "xygo:queue:"}
}

func (r *RedisDriver) key(topic string) string {
	return r.prefix + topic
}

func (r *RedisDriver) Push(ctx context.Context, topic string, msg *Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = g.Redis().LPush(ctx, r.key(topic), string(data))
	return err
}

func (r *RedisDriver) Pop(ctx context.Context, topic string, timeout time.Duration) (*Message, error) {
	// BRPOP 阻塞式弹出，超时返回 nil
	result, err := g.Redis().Do(ctx, "BRPOP", r.key(topic), int(timeout.Seconds()))
	if err != nil {
		return nil, err
	}
	if result.IsNil() || result.IsEmpty() {
		return nil, nil
	}

	// BRPOP 返回 [key, value]
	arr := result.Strings()
	if len(arr) < 2 {
		return nil, nil
	}

	var msg Message
	if err := json.Unmarshal([]byte(arr[1]), &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func (r *RedisDriver) Len(ctx context.Context, topic string) (int64, error) {
	// 普通队列长度
	listLen, _ := g.Redis().LLen(ctx, r.key(topic))
	// 延迟队列长度
	delayLen, _ := g.Redis().Do(ctx, "ZCARD", r.delayKey(topic))
	total := listLen
	if !delayLen.IsNil() && !delayLen.IsEmpty() {
		total += delayLen.Int64()
	}
	return total, nil
}

func (r *RedisDriver) Close() error {
	return nil
}

// ==================== 延迟队列（Sorted Set）====================

func (r *RedisDriver) delayKey(topic string) string {
	return r.prefix + "delay:" + topic
}

// DelayPush 延迟投递：消息放入 Sorted Set，score=执行时间戳
func (r *RedisDriver) DelayPush(ctx context.Context, topic string, msg *Message, executeAt int64) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = g.Redis().Do(ctx, "ZADD", r.delayKey(topic), executeAt, string(data))
	return err
}

// PollDelay 取出到期的延迟消息（score <= now），转入普通队列
func (r *RedisDriver) PollDelay(ctx context.Context, topic string) (*Message, error) {
	now := time.Now().Unix()
	// ZRANGEBYSCORE key -inf now LIMIT 0 1
	result, err := g.Redis().Do(ctx, "ZRANGEBYSCORE", r.delayKey(topic), "-inf", now, "LIMIT", 0, 1)
	if err != nil {
		return nil, err
	}
	if result.IsNil() || result.IsEmpty() {
		return nil, nil
	}
	arr := result.Strings()
	if len(arr) == 0 {
		return nil, nil
	}

	raw := arr[0]
	// 从 Sorted Set 移除（用 ZREM 保证原子性）
	removed, err := g.Redis().Do(ctx, "ZREM", r.delayKey(topic), raw)
	if err != nil {
		return nil, err
	}
	if removed.Int64() == 0 {
		// 被其他消费者抢走了
		return nil, nil
	}

	var msg Message
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
