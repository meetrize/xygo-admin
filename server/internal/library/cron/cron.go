// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// Package cron 提供定时任务调度引擎封装
// 基于 GoFrame gcron，支持代码注册式任务 + 数据库配置 + 执行日志
package cron

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

// ==================== 策略常量 ====================

const (
	PolicyConcurrent = 1 // 并行执行
	PolicySingleton  = 2 // 单例（同一任务同时只运行一个实例）
	PolicyOnce       = 3 // 单次执行后自动停止
	PolicyTimes      = 4 // 固定次数执行
)

// ==================== 任务接口 ====================

// Task 定时任务接口，所有任务需实现此接口并在 init() 中注册
type Task interface {
	GetName() string                                          // 任务唯一标识（与数据库 name 字段对应）
	Execute(ctx context.Context, params []string) (string, error) // 执行任务，返回输出和错误
}

// ==================== 注册表 ====================

var (
	registry = make(map[string]Task)
	mu       sync.RWMutex
)

// Register 注册任务（在各任务文件 init() 中调用）
func Register(t Task) {
	mu.Lock()
	defer mu.Unlock()
	name := t.GetName()
	if name == "" {
		return
	}
	if _, exists := registry[name]; exists {
		g.Log().Warningf(context.Background(), "[cron] task '%s' already registered, skip", name)
		return
	}
	registry[name] = t
	g.Log().Infof(context.Background(), "[cron] task '%s' registered", name)
}

// GetTask 获取已注册的任务
func GetTask(name string) Task {
	mu.RLock()
	defer mu.RUnlock()
	return registry[name]
}

// GetAllTasks 获取所有已注册的任务名
func GetAllTasks() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	return names
}

// ==================== 调度管理 ====================

// CronJob 运行中的任务配置
type CronJob struct {
	Id      uint64
	Name    string
	Title   string
	Pattern string
	Params  string
	Policy  int
	Count   int
}

// LogCallback 执行日志回调
type LogCallback func(ctx context.Context, cronId uint64, name, title, params string, status int, output, errMsg string, takeMs int)

var logCallback LogCallback

// SetLogCallback 设置执行日志回调（由 logic 层注入）
func SetLogCallback(cb LogCallback) {
	logCallback = cb
}

// StartJob 启动单个任务
func StartJob(job *CronJob) error {
	task := GetTask(job.Name)
	if task == nil {
		return fmt.Errorf("task '%s' not registered", job.Name)
	}

	entryName := genEntryName(job.Id, job.Name)

	// 先停掉旧的（如果有）
	StopJob(job.Id, job.Name)

	var err error
	fn := genExecuteFunc(job, task)

	switch job.Policy {
	case PolicySingleton:
		_, err = gcron.AddSingleton(context.Background(), job.Pattern, fn, entryName)
	case PolicyOnce:
		_, err = gcron.AddOnce(context.Background(), job.Pattern, fn, entryName)
	case PolicyTimes:
		count := job.Count
		if count <= 0 {
			count = 1
		}
		_, err = gcron.AddTimes(context.Background(), job.Pattern, count, fn, entryName)
	default: // PolicyConcurrent
		_, err = gcron.Add(context.Background(), job.Pattern, fn, entryName)
	}

	if err != nil {
		return fmt.Errorf("start job '%s' failed: %v", job.Name, err)
	}

	g.Log().Infof(context.Background(), "[cron] job '%s'(%s) started, pattern=%s, policy=%d",
		job.Title, job.Name, job.Pattern, job.Policy)
	return nil
}

// StopJob 停止单个任务
func StopJob(id uint64, name string) {
	entryName := genEntryName(id, name)
	entry := gcron.Search(entryName)
	if entry != nil {
		entry.Stop()
		gcron.Remove(entryName)
		g.Log().Infof(context.Background(), "[cron] job '%s' stopped", entryName)
	}
}

// StopAll 停止所有任务
func StopAll() {
	entries := gcron.Entries()
	for _, entry := range entries {
		entry.Stop()
	}
	g.Log().Info(context.Background(), "[cron] all jobs stopped")
}

// OnlineExec 手动执行一次任务（不受 cron 表达式约束）
func OnlineExec(job *CronJob) (output string, err error) {
	task := GetTask(job.Name)
	if task == nil {
		return "", fmt.Errorf("task '%s' not registered", job.Name)
	}

	params := parseParams(job.Params)
	start := time.Now()
	output, err = task.Execute(context.Background(), params)
	takeMs := int(time.Since(start).Milliseconds())

	status := 1
	errMsg := ""
	if err != nil {
		status = 2
		errMsg = err.Error()
	}

	// 记录日志
	if logCallback != nil {
		logCallback(context.Background(), job.Id, job.Name, job.Title, job.Params, status, output, errMsg, takeMs)
	}

	return output, err
}

// ==================== 内部辅助 ====================

func genEntryName(id uint64, name string) string {
	return fmt.Sprintf("%s@%d", name, id)
}

func parseParams(params string) []string {
	params = strings.TrimSpace(params)
	if params == "" {
		return nil
	}
	parts := strings.Split(params, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func genExecuteFunc(job *CronJob, task Task) func(ctx context.Context) {
	return func(ctx context.Context) {
		params := parseParams(job.Params)
		start := time.Now()
		output, err := task.Execute(ctx, params)
		takeMs := int(time.Since(start).Milliseconds())

		status := 1
		errMsg := ""
		if err != nil {
			status = 2
			errMsg = err.Error()
			g.Log().Errorf(ctx, "[cron] job '%s' failed, took %dms, err: %v", job.Name, takeMs, err)
		} else {
			g.Log().Infof(ctx, "[cron] job '%s' success, took %dms", job.Name, takeMs)
		}

		// 记录执行日志
		if logCallback != nil {
			logCallback(ctx, job.Id, job.Name, job.Title, job.Params, status, output, errMsg, takeMs)
		}
	}
}
