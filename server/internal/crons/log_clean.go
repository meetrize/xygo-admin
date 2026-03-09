// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// Package crons 定时任务实现（所有任务在 init 中注册）
package crons

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	cronlib "xygo/internal/library/cron"
)

func init() {
	cronlib.Register(&LogCleanTask{})
}

// LogCleanTask 日志清理任务：清除指定天数前的登录日志和操作日志
type LogCleanTask struct{}

func (t *LogCleanTask) GetName() string { return "log_clean" }

func (t *LogCleanTask) Execute(ctx context.Context, params []string) (string, error) {
	// 默认保留 30 天
	days := 30
	if len(params) > 0 {
		if d := params[0]; d != "" {
			fmt.Sscanf(d, "%d", &days)
		}
	}
	if days <= 0 {
		days = 30
	}

	cutoff := uint(time.Now().AddDate(0, 0, -days).Unix())

	// 清理登录日志
	loginResult, err := dao.AdminLoginLog.Ctx(ctx).WhereLT("created_at", cutoff).Delete()
	if err != nil {
		return "", fmt.Errorf("clean login log failed: %v", err)
	}
	loginCount, _ := loginResult.RowsAffected()

	// 清理操作日志
	opResult, err := dao.AdminOperationLog.Ctx(ctx).WhereLT("created_at", cutoff).Delete()
	if err != nil {
		return "", fmt.Errorf("clean operation log failed: %v", err)
	}
	opCount, _ := opResult.RowsAffected()

	// 清理定时任务执行日志
	cronResult, err := dao.SysCronLog.Ctx(ctx).WhereLT("created_at", cutoff).Delete()
	if err != nil {
		return "", fmt.Errorf("clean cron log failed: %v", err)
	}
	cronLogCount, _ := cronResult.RowsAffected()

	output := fmt.Sprintf("cleaned: login_log=%d, operation_log=%d, cron_log=%d (older than %d days)",
		loginCount, opCount, cronLogCount, days)
	g.Log().Infof(ctx, "[cron:log_clean] %s", output)
	return output, nil
}
