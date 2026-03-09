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

var AdminChatSessionMember = &adminChatSessionMemberDao{
	table: "xy_admin_chat_session_member",
}

type adminChatSessionMemberDao struct {
	table string
}

func (d *adminChatSessionMemberDao) Table() string { return d.table }

func (d *adminChatSessionMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB().Model(d.table).Safe().Ctx(ctx)
}

// Columns 字段名常量
func (d *adminChatSessionMemberDao) Columns() *adminChatSessionMemberColumns {
	return &adminChatSessionMemberCols
}

type adminChatSessionMemberColumns struct {
	Id, SessionId, UserId, Role                 string
	UnreadCount, LastReadMsgId, IsMuted         string
	IsDeleted, JoinedAt                         string
}

var adminChatSessionMemberCols = adminChatSessionMemberColumns{
	Id: "id", SessionId: "session_id", UserId: "user_id", Role: "role",
	UnreadCount: "unread_count", LastReadMsgId: "last_read_msg_id",
	IsMuted: "is_muted", IsDeleted: "is_deleted", JoinedAt: "joined_at",
}
