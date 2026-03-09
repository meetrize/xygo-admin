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
	"xygo/internal/service"
)

// MemberNoticeList 会员通知列表
func (c *ControllerV1) MemberNoticeList(ctx context.Context, req *api.MemberNoticeListReq) (res *api.MemberNoticeListRes, err error) {
	result, err := service.MemberNotice().List(ctx, &req.MemberNoticeListInp)
	if err != nil {
		return nil, err
	}
	return &api.MemberNoticeListRes{result}, nil
}

// MemberNoticeView 会员通知详情
func (c *ControllerV1) MemberNoticeView(ctx context.Context, req *api.MemberNoticeViewReq) (res *api.MemberNoticeViewRes, err error) {
	result, err := service.MemberNotice().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.MemberNoticeViewRes{result}, nil
}

// MemberNoticeEdit 保存会员通知
func (c *ControllerV1) MemberNoticeEdit(ctx context.Context, req *api.MemberNoticeEditReq) (res *api.MemberNoticeEditRes, err error) {
	err = service.MemberNotice().Edit(ctx, &req.MemberNoticeEditInp)
	return &api.MemberNoticeEditRes{}, err
}

// MemberNoticeDelete 删除会员通知
func (c *ControllerV1) MemberNoticeDelete(ctx context.Context, req *api.MemberNoticeDeleteReq) (res *api.MemberNoticeDeleteRes, err error) {
	err = service.MemberNotice().Delete(ctx, req.Id)
	return &api.MemberNoticeDeleteRes{}, err
}
