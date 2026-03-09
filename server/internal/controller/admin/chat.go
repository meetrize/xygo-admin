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
	"context"

	api "xygo/api/admin"
	"xygo/internal/library/contexts"
	"xygo/internal/service"
)

// ChatSessions 会话列表
func (c *ControllerV1) ChatSessions(ctx context.Context, req *api.ChatSessionsReq) (res *api.ChatSessionsRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.ChatSessionsRes{List: nil}, nil
	}
	list, err := service.Chat().Sessions(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.ChatSessionsRes{List: list}, nil
}

// ChatSessionCreate 创建会话
func (c *ControllerV1) ChatSessionCreate(ctx context.Context, req *api.ChatSessionCreateReq) (res *api.ChatSessionCreateRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, nil
	}
	result, err := service.Chat().CreateSession(ctx, &req.ChatSessionCreateInp, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.ChatSessionCreateRes{ChatSessionCreateModel: result}, nil
}

// ChatSessionDelete 删除会话
func (c *ControllerV1) ChatSessionDelete(ctx context.Context, req *api.ChatSessionDeleteReq) (res *api.ChatSessionDeleteRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.ChatSessionDeleteRes{}, nil
	}
	err = service.Chat().DeleteSession(ctx, &req.ChatSessionDeleteInp, user.Id)
	return &api.ChatSessionDeleteRes{}, err
}

// ChatMessages 消息列表
func (c *ControllerV1) ChatMessages(ctx context.Context, req *api.ChatMessagesReq) (res *api.ChatMessagesRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, nil
	}
	result, err := service.Chat().Messages(ctx, &req.ChatMessageListInp, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.ChatMessagesRes{ChatMessageListModel: result}, nil
}

// ChatSend 发送消息
func (c *ControllerV1) ChatSend(ctx context.Context, req *api.ChatSendReq) (res *api.ChatSendRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, nil
	}
	result, err := service.Chat().SendMessage(ctx, &req.ChatSendMessageInp, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.ChatSendRes{ChatSendMessageModel: result}, nil
}

// ChatRead 标记已读
func (c *ControllerV1) ChatRead(ctx context.Context, req *api.ChatReadReq) (res *api.ChatReadRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.ChatReadRes{}, nil
	}
	err = service.Chat().MarkRead(ctx, req.SessionId, user.Id)
	return &api.ChatReadRes{}, err
}

// ChatContacts 联系人列表
func (c *ControllerV1) ChatContacts(ctx context.Context, req *api.ChatContactsReq) (res *api.ChatContactsRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.ChatContactsRes{List: nil}, nil
	}
	list, err := service.Chat().Contacts(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.ChatContactsRes{List: list}, nil
}

// ChatGroupUpdate 编辑群聊
func (c *ControllerV1) ChatGroupUpdate(ctx context.Context, req *api.ChatGroupUpdateReq) (res *api.ChatGroupUpdateRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.ChatGroupUpdateRes{}, nil
	}
	err = service.Chat().GroupUpdate(ctx, &req.ChatGroupUpdateInp, user.Id)
	return &api.ChatGroupUpdateRes{}, err
}

// ChatUnreadTotal 未读总数
func (c *ControllerV1) ChatUnreadTotal(ctx context.Context, req *api.ChatUnreadTotalReq) (res *api.ChatUnreadTotalRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.ChatUnreadTotalRes{Total: 0}, nil
	}
	total, err := service.Chat().UnreadTotal(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.ChatUnreadTotalRes{Total: total}, nil
}
