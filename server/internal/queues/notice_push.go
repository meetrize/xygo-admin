// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package queues

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/library/queue"
	"xygo/internal/websocket"
)

const TopicNoticePush = "notice_push"

func init() {
	queue.Register(&NoticePushConsumer{})
}

// NoticePushConsumer 通知推送消费者（WebSocket 推送）
type NoticePushConsumer struct{}

func (c *NoticePushConsumer) GetTopic() string { return TopicNoticePush }

func (c *NoticePushConsumer) Handle(ctx context.Context, msg *queue.Message) error {
	var data struct {
		UserIds []uint64    `json:"userIds"`
		Event   string      `json:"event"`
		Payload interface{} `json:"payload"`
	}
	if err := json.Unmarshal([]byte(msg.Body), &data); err != nil {
		g.Log().Errorf(ctx, "[queue:notice_push] unmarshal failed: %v", err)
		return nil
	}
	if len(data.UserIds) == 0 || data.Event == "" {
		g.Log().Warning(ctx, "[queue:notice_push] missing userIds or event, skip")
		return nil
	}

	resp := &websocket.WsResponse{
		Event: data.Event,
		Data:  data.Payload,
	}

	for _, uid := range data.UserIds {
		websocket.SendToUser("admin", uid, resp)
	}
	return nil
}
