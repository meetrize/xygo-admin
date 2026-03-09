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

// MemberLoginLogList 登录日志列表
func (c *ControllerV1) MemberLoginLogList(ctx context.Context, req *api.MemberLoginLogListReq) (res *api.MemberLoginLogListRes, err error) {
	result, err := service.MemberLoginLog().List(ctx, &req.MemberLoginLogListInp)
	if err != nil {
		return nil, err
	}
	return &api.MemberLoginLogListRes{result}, nil
}

// MemberLoginLogView 登录日志详情
func (c *ControllerV1) MemberLoginLogView(ctx context.Context, req *api.MemberLoginLogViewReq) (res *api.MemberLoginLogViewRes, err error) {
	result, err := service.MemberLoginLog().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.MemberLoginLogViewRes{result}, nil
}

// MemberLoginLogDelete 删除登录日志
func (c *ControllerV1) MemberLoginLogDelete(ctx context.Context, req *api.MemberLoginLogDeleteReq) (res *api.MemberLoginLogDeleteRes, err error) {
	err = service.MemberLoginLog().Delete(ctx, req.Id)
	return &api.MemberLoginLogDeleteRes{}, err
}
