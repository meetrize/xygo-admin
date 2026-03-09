// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package security

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/library/cache"
)

// LoginProtectConfig 登录防暴力破解配置
type LoginProtectConfig struct {
	Enabled    bool // 是否启用
	MaxRetries int  // 最大连续失败次数
	LockMinutes int // 锁定时长（分钟）
}

// getLoginProtectConfig 从配置文件读取防暴力配置
func getLoginProtectConfig(ctx context.Context) LoginProtectConfig {
	enabled := g.Cfg().MustGet(ctx, "security.loginProtect.enabled").Bool()
	maxRetries := g.Cfg().MustGet(ctx, "security.loginProtect.maxRetries").Int()
	lockMinutes := g.Cfg().MustGet(ctx, "security.loginProtect.lockMinutes").Int()

	if maxRetries <= 0 {
		maxRetries = 5
	}
	if lockMinutes <= 0 {
		lockMinutes = 15
	}

	return LoginProtectConfig{
		Enabled:     enabled,
		MaxRetries:  maxRetries,
		LockMinutes: lockMinutes,
	}
}

// 缓存 key 前缀
const (
	loginFailKeyPrefix = "xygo:login:fail:" // 失败次数: xygo:login:fail:{ip}:{username}
	loginLockKeyPrefix = "xygo:login:lock:" // 锁定标记: xygo:login:lock:{ip}:{username}
)

// CheckLoginLocked 检查是否被锁定
// 返回: locked（是否锁定）, remainMinutes（剩余锁定分钟数）
func CheckLoginLocked(ctx context.Context, ip, username string) (locked bool, remainMinutes int) {
	cfg := getLoginProtectConfig(ctx)
	if !cfg.Enabled {
		return false, 0
	}

	lockKey := fmt.Sprintf("%s%s:%s", loginLockKeyPrefix, ip, username)
	val, err := cache.Instance().Get(ctx, lockKey)
	if err != nil || val.IsEmpty() {
		return false, 0
	}

	return true, val.Int()
}

// RecordLoginFail 记录一次登录失败
// 返回: locked（是否触发锁定）, remainRetries（剩余尝试次数）
func RecordLoginFail(ctx context.Context, ip, username string) (locked bool, remainRetries int) {
	cfg := getLoginProtectConfig(ctx)
	if !cfg.Enabled {
		return false, -1
	}

	failKey := fmt.Sprintf("%s%s:%s", loginFailKeyPrefix, ip, username)
	lockKey := fmt.Sprintf("%s%s:%s", loginLockKeyPrefix, ip, username)

	// 获取当前失败次数
	val, _ := cache.Instance().Get(ctx, failKey)
	count := val.Int() + 1

	// 更新失败次数（有效期与锁定时长相同）
	duration := time.Duration(cfg.LockMinutes) * time.Minute
	_ = cache.Instance().Set(ctx, failKey, count, duration)

	// 达到阈值 → 锁定
	if count >= cfg.MaxRetries {
		_ = cache.Instance().Set(ctx, lockKey, cfg.LockMinutes, duration)
		g.Log().Warningf(ctx, "登录防护：IP=%s 账号=%s 连续失败%d次，已锁定%d分钟", ip, username, count, cfg.LockMinutes)
		return true, 0
	}

	return false, cfg.MaxRetries - count
}

// ClearLoginFail 登录成功后清除失败记录
func ClearLoginFail(ctx context.Context, ip, username string) {
	cfg := getLoginProtectConfig(ctx)
	if !cfg.Enabled {
		return
	}

	failKey := fmt.Sprintf("%s%s:%s", loginFailKeyPrefix, ip, username)
	lockKey := fmt.Sprintf("%s%s:%s", loginLockKeyPrefix, ip, username)

	_, _ = cache.Instance().Remove(ctx, failKey)
	_, _ = cache.Instance().Remove(ctx, lockKey)
}
