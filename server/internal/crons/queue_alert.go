// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package crons

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	cronlib "xygo/internal/library/cron"
	"xygo/internal/library/queue"
	"xygo/internal/websocket"
)

// getSuperAdminIds 获取所有超管用户ID（is_super=1 且 status=1）
func getSuperAdminIds(ctx context.Context) []uint64 {
	var users []struct {
		Id uint64 `json:"id"`
	}
	_ = dao.AdminUser.Ctx(ctx).
		Where("is_super", 1).
		Where("status", 1).
		Fields("id").
		Scan(&users)
	ids := make([]uint64, 0, len(users))
	for _, u := range users {
		ids = append(ids, u.Id)
	}
	return ids
}

func init() {
	cronlib.Register(&QueueAlertTask{})
}

// QueueAlertTask 队列积压告警任务
// 参数：threshold（积压阈值，默认100），deadThreshold（死信阈值，默认10）
type QueueAlertTask struct{}

func (t *QueueAlertTask) GetName() string { return "queue_alert" }

func (t *QueueAlertTask) Execute(ctx context.Context, params []string) (string, error) {
	threshold := int64(100)
	deadThreshold := int64(10)
	if len(params) > 0 && params[0] != "" {
		fmt.Sscanf(params[0], "%d", &threshold)
	}
	if len(params) > 1 && params[1] != "" {
		fmt.Sscanf(params[1], "%d", &deadThreshold)
	}

	queue.SetAlertConfig(queue.AlertConfig{
		Enabled:       true,
		Threshold:     threshold,
		DeadThreshold: deadThreshold,
	})

	alerts := queue.CheckAlerts(ctx)
	if len(alerts) == 0 {
		return "all queues normal", nil
	}

	msg := strings.Join(alerts, "\n")
	g.Log().Warningf(ctx, "[cron:queue_alert] %s", msg)

	// 查询所有超管用户ID，直接 WebSocket 推送（不走队列，避免延迟）
	superUserIds := getSuperAdminIds(ctx)
	if len(superUserIds) > 0 {
		resp := &websocket.WsResponse{
			Event: "system/alert",
			Data: map[string]string{
				"title":   "队列积压告警",
				"content": msg,
			},
		}
		for _, uid := range superUserIds {
			websocket.SendToUser("admin", uid, resp)
		}
	}

	return fmt.Sprintf("告警 %d 条，已推送给 %d 位管理员:\n%s", len(alerts), len(superUserIds), msg), nil
}
