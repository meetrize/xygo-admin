// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// Package queues 内置消息队列消费者
package queues

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	"xygo/internal/library/queue"
)

const TopicOperationLog = "operation_log"

func init() {
	queue.Register(&OperationLogConsumer{})
}

// OperationLogConsumer 操作日志异步写入消费者
type OperationLogConsumer struct{}

func (c *OperationLogConsumer) GetTopic() string { return TopicOperationLog }

func (c *OperationLogConsumer) Handle(ctx context.Context, msg *queue.Message) error {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Body), &data); err != nil {
		g.Log().Errorf(ctx, "[queue:operation_log] unmarshal failed: %v", err)
		return nil // 格式错误不重试
	}
	if len(data) == 0 {
		g.Log().Warning(ctx, "[queue:operation_log] empty data, skip")
		return nil
	}

	_, err := dao.AdminOperationLog.Ctx(ctx).Data(data).Insert()
	if err != nil {
		g.Log().Errorf(ctx, "[queue:operation_log] insert failed: %v", err)
		return queue.NewRetryError(err)
	}
	return nil
}
