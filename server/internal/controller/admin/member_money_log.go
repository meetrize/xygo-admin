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

// MemberMoneyLogList 余额变动日志列表
func (c *ControllerV1) MemberMoneyLogList(ctx context.Context, req *api.MemberMoneyLogListReq) (res *api.MemberMoneyLogListRes, err error) {
	result, err := service.MemberMoneyLog().List(ctx, &req.MemberMoneyLogListInp)
	if err != nil {
		return nil, err
	}
	return &api.MemberMoneyLogListRes{result}, nil
}

// MemberMoneyLogView 余额变动日志详情
func (c *ControllerV1) MemberMoneyLogView(ctx context.Context, req *api.MemberMoneyLogViewReq) (res *api.MemberMoneyLogViewRes, err error) {
	result, err := service.MemberMoneyLog().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.MemberMoneyLogViewRes{result}, nil
}

// MemberMoneyLogEdit 保存余额变动日志
func (c *ControllerV1) MemberMoneyLogEdit(ctx context.Context, req *api.MemberMoneyLogEditReq) (res *api.MemberMoneyLogEditRes, err error) {
	err = service.MemberMoneyLog().Edit(ctx, &req.MemberMoneyLogEditInp)
	return &api.MemberMoneyLogEditRes{}, err
}

// MemberMoneyLogDelete 删除余额变动日志
func (c *ControllerV1) MemberMoneyLogDelete(ctx context.Context, req *api.MemberMoneyLogDeleteReq) (res *api.MemberMoneyLogDeleteRes, err error) {
	err = service.MemberMoneyLog().Delete(ctx, req.Id)
	return &api.MemberMoneyLogDeleteRes{}, err
}
