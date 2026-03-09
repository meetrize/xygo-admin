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

// NoticeList 通知列表
func (c *ControllerV1) NoticeList(ctx context.Context, req *api.NoticeListReq) (res *api.NoticeListRes, err error) {
	result, err := service.Notice().List(ctx, &req.NoticeListInp)
	if err != nil {
		return nil, err
	}
	return &api.NoticeListRes{NoticeListModel: result}, nil
}

// NoticeEdit 编辑通知
func (c *ControllerV1) NoticeEdit(ctx context.Context, req *api.NoticeEditReq) (res *api.NoticeEditRes, err error) {
	user := contexts.GetUser(ctx)
	var senderId uint64
	if user != nil {
		senderId = user.Id
	}
	id, err := service.Notice().Edit(ctx, &req.NoticeEditInp, senderId)
	if err != nil {
		return nil, err
	}
	return &api.NoticeEditRes{Id: id}, nil
}

// NoticeDelete 删除通知
func (c *ControllerV1) NoticeDelete(ctx context.Context, req *api.NoticeDeleteReq) (res *api.NoticeDeleteRes, err error) {
	err = service.Notice().Delete(ctx, req.Id)
	return &api.NoticeDeleteRes{}, err
}

// NoticePull 拉取消息
func (c *ControllerV1) NoticePull(ctx context.Context, req *api.NoticePullReq) (res *api.NoticePullRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.NoticePullRes{}, nil
	}
	result, err := service.Notice().PullMessages(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.NoticePullRes{PullMessagesModel: result}, nil
}

// NoticeUnreadCount 未读数
func (c *ControllerV1) NoticeUnreadCount(ctx context.Context, req *api.NoticeUnreadCountReq) (res *api.NoticeUnreadCountRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.NoticeUnreadCountRes{}, nil
	}
	list, total, err := service.Notice().UnreadCount(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &api.NoticeUnreadCountRes{List: list, Total: total}, nil
}

// NoticeRead 标记已读
func (c *ControllerV1) NoticeRead(ctx context.Context, req *api.NoticeReadReq) (res *api.NoticeReadRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.NoticeReadRes{}, nil
	}
	err = service.Notice().Read(ctx, req.Id, user.Id)
	return &api.NoticeReadRes{}, err
}

// NoticeReadAll 全部已读
func (c *ControllerV1) NoticeReadAll(ctx context.Context, req *api.NoticeReadAllReq) (res *api.NoticeReadAllRes, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return &api.NoticeReadAllRes{}, nil
	}
	err = service.Notice().ReadAll(ctx, req.Type, user.Id)
	return &api.NoticeReadAllRes{}, err
}
