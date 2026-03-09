// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package adminin

import "xygo/internal/model/input/form"

// ==================== 聊天会话 ====================

// ChatSessionListInp 会话列表入参
type ChatSessionListInp struct{}

// ChatSessionItem 会话列表项
type ChatSessionItem struct {
	Id              uint64 `json:"id"`
	Type            int    `json:"type"`            // 1=单聊,2=群聊
	Name            string `json:"name"`            // 会话名称（单聊=对方昵称，群聊=群名）
	Avatar          string `json:"avatar"`          // 头像（单聊=对方头像，群聊=群头像）
	LastMessage     string `json:"lastMessage"`     // 最后一条消息
	LastMessageTime int `json:"lastMessageTime"` // 最后消息时间
	UnreadCount     int    `json:"unreadCount"`     // 未读数
	MemberCount     int    `json:"memberCount"`     // 成员数
	IsMuted         int    `json:"isMuted"`         // 免打扰
}

// ChatSessionCreateInp 创建会话入参
type ChatSessionCreateInp struct {
	Type    int      `json:"type"    v:"required|in:1,2#类型不能为空|类型值无效" dc:"类型:1=单聊,2=群聊"`
	Name    string   `json:"name"    dc:"群名称(群聊时必填)"`
	UserIds []uint64 `json:"userIds" v:"required#成员不能为空" dc:"成员ID列表"`
}

// ChatSessionCreateModel 创建会话出参
type ChatSessionCreateModel struct {
	SessionId uint64 `json:"sessionId"`
}

// ChatSessionDeleteInp 删除会话入参
type ChatSessionDeleteInp struct {
	SessionId uint64 `json:"sessionId" v:"required#会话ID不能为空" dc:"会话ID"`
}

// ==================== 聊天消息 ====================

// ChatMessageListInp 消息列表入参
type ChatMessageListInp struct {
	SessionId uint64 `json:"sessionId" v:"required#会话ID不能为空" dc:"会话ID"`
	LastMsgId uint64 `json:"lastMsgId" dc:"上一页最后消息ID(向上翻页用)"`
	form.PageReq
}

// ChatMessageItem 消息项
type ChatMessageItem struct {
	Id           uint64 `json:"id"`
	SessionId    uint64 `json:"sessionId"`
	SenderId     uint64 `json:"senderId"`
	SenderName   string `json:"senderName"`
	SenderAvatar string `json:"senderAvatar"`
	Type         int    `json:"type"` // 1=文字,2=图片,3=系统消息
	Content      string `json:"content"`
	CreatedAt    int    `json:"createdAt"`
}

// ChatMessageListModel 消息列表出参
type ChatMessageListModel struct {
	List    []ChatMessageItem `json:"list"`
	HasMore bool              `json:"hasMore"`
}

// ChatSendMessageInp 发送消息入参
type ChatSendMessageInp struct {
	SessionId uint64 `json:"sessionId" v:"required#会话ID不能为空" dc:"会话ID"`
	Type      int    `json:"type"      v:"required|in:1,2#消息类型不能为空|消息类型无效" dc:"类型:1=文字,2=图片"`
	Content   string `json:"content"   v:"required#内容不能为空" dc:"消息内容"`
}

// ChatSendMessageModel 发送消息出参
type ChatSendMessageModel struct {
	Id        uint64 `json:"id"`
	CreatedAt int    `json:"createdAt"`
}

// ==================== 已读标记 ====================

// ChatReadInp 标记已读入参
type ChatReadInp struct {
	SessionId uint64 `json:"sessionId" v:"required#会话ID不能为空" dc:"会话ID"`
}

// ==================== 联系人 ====================

// ChatContactItem 联系人项
type ChatContactItem struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	RealName string `json:"realName"`
	DeptId   uint64 `json:"deptId"`
	DeptName string `json:"deptName"`
	PostName string `json:"postName"` // 岗位名称
	Avatar   string `json:"avatar"`
	IsOnline bool   `json:"isOnline"`
}

// ==================== 群聊管理 ====================

// ChatGroupUpdateInp 编辑群聊入参
type ChatGroupUpdateInp struct {
	SessionId uint64   `json:"sessionId" v:"required#会话ID不能为空" dc:"会话ID"`
	Name      string   `json:"name"      dc:"群名称"`
	AddUsers  []uint64 `json:"addUsers"  dc:"新增成员ID列表"`
	DelUsers  []uint64 `json:"delUsers"  dc:"移除成员ID列表"`
}

// ==================== 未读总数 ====================

// ChatUnreadTotalModel 未读总数出参
type ChatUnreadTotalModel struct {
	Total int `json:"total"`
}
