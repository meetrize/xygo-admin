// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var AdminChatMessage = &adminChatMessageDao{
	table: "xy_admin_chat_message",
}

type adminChatMessageDao struct {
	table string
}

func (d *adminChatMessageDao) Table() string { return d.table }

func (d *adminChatMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB().Model(d.table).Safe().Ctx(ctx)
}

// Columns 字段名常量
func (d *adminChatMessageDao) Columns() *adminChatMessageColumns { return &adminChatMessageCols }

type adminChatMessageColumns struct {
	Id, SessionId, SenderId, Type, Content, CreatedAt string
}

var adminChatMessageCols = adminChatMessageColumns{
	Id: "id", SessionId: "session_id", SenderId: "sender_id",
	Type: "type", Content: "content", CreatedAt: "created_at",
}
