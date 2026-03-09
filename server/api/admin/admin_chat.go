// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package admin

import (
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/model/input/adminin"
)

// ===================== 聊天会话 =====================

type ChatSessionsReq struct {
	g.Meta `path:"/admin/chat/sessions" method:"get" tags:"AdminChat" summary:"会话列表"`
}
type ChatSessionsRes struct {
	List []adminin.ChatSessionItem `json:"list"`
}

type ChatSessionCreateReq struct {
	g.Meta `path:"/admin/chat/session/create" method:"post" tags:"AdminChat" summary:"创建会话"`
	adminin.ChatSessionCreateInp
}
type ChatSessionCreateRes struct {
	*adminin.ChatSessionCreateModel
}

type ChatSessionDeleteReq struct {
	g.Meta `path:"/admin/chat/session/delete" method:"post" tags:"AdminChat" summary:"删除会话"`
	adminin.ChatSessionDeleteInp
}
type ChatSessionDeleteRes struct{}

// ===================== 聊天消息 =====================

type ChatMessagesReq struct {
	g.Meta `path:"/admin/chat/messages" method:"post" tags:"AdminChat" summary:"消息列表"`
	adminin.ChatMessageListInp
}
type ChatMessagesRes struct {
	*adminin.ChatMessageListModel
}

type ChatSendReq struct {
	g.Meta `path:"/admin/chat/send" method:"post" tags:"AdminChat" summary:"发送消息"`
	adminin.ChatSendMessageInp
}
type ChatSendRes struct {
	*adminin.ChatSendMessageModel
}

type ChatReadReq struct {
	g.Meta `path:"/admin/chat/read" method:"post" tags:"AdminChat" summary:"标记已读"`
	adminin.ChatReadInp
}
type ChatReadRes struct{}

// ===================== 联系人 =====================

type ChatContactsReq struct {
	g.Meta `path:"/admin/chat/contacts" method:"get" tags:"AdminChat" summary:"联系人列表"`
}
type ChatContactsRes struct {
	List []adminin.ChatContactItem `json:"list"`
}

// ===================== 群聊管理 =====================

type ChatGroupUpdateReq struct {
	g.Meta `path:"/admin/chat/group/update" method:"post" tags:"AdminChat" summary:"编辑群聊"`
	adminin.ChatGroupUpdateInp
}
type ChatGroupUpdateRes struct{}

// ===================== 未读总数 =====================

type ChatUnreadTotalReq struct {
	g.Meta `path:"/admin/chat/unread-total" method:"get" tags:"AdminChat" summary:"未读总数"`
}
type ChatUnreadTotalRes struct {
	Total int `json:"total"`
}
