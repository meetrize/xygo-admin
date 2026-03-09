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
	"xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

// ===================== 登录日志 =====================

// LoginLogList 登录日志列表
func (c *ControllerV1) LoginLogList(ctx context.Context, req *api.LoginLogListReq) (res *api.LoginLogListRes, err error) {
	list, total, err := service.AdminLog().LoginLogList(ctx, &req.LoginLogListInp)
	if err != nil {
		return nil, err
	}

	res = new(api.LoginLogListRes)
	res.LoginLogListModel = &adminin.LoginLogListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return
}

// LoginLogDelete 删除登录日志
func (c *ControllerV1) LoginLogDelete(ctx context.Context, req *api.LoginLogDeleteReq) (res *api.LoginLogDeleteRes, err error) {
	err = service.AdminLog().LoginLogDelete(ctx, &req.LoginLogDeleteInp)
	if err != nil {
		return nil, err
	}
	res = &api.LoginLogDeleteRes{}
	return
}

// LoginLogClear 清空登录日志
func (c *ControllerV1) LoginLogClear(ctx context.Context, req *api.LoginLogClearReq) (res *api.LoginLogClearRes, err error) {
	err = service.AdminLog().LoginLogClear(ctx)
	if err != nil {
		return nil, err
	}
	res = &api.LoginLogClearRes{}
	return
}

// ===================== 操作日志 =====================

// OperationLogList 操作日志列表
func (c *ControllerV1) OperationLogList(ctx context.Context, req *api.OperationLogListReq) (res *api.OperationLogListRes, err error) {
	list, total, err := service.AdminLog().OperationLogList(ctx, &req.OperationLogListInp)
	if err != nil {
		return nil, err
	}

	res = new(api.OperationLogListRes)
	res.OperationLogListModel = &adminin.OperationLogListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return
}

// OperationLogDetail 操作日志详情
func (c *ControllerV1) OperationLogDetail(ctx context.Context, req *api.OperationLogDetailReq) (res *api.OperationLogDetailRes, err error) {
	out, err := service.AdminLog().OperationLogDetail(ctx, &req.OperationLogDetailInp)
	if err != nil {
		return nil, err
	}

	res = new(api.OperationLogDetailRes)
	res.OperationLogItem = out
	return
}

// OperationLogDelete 删除操作日志
func (c *ControllerV1) OperationLogDelete(ctx context.Context, req *api.OperationLogDeleteReq) (res *api.OperationLogDeleteRes, err error) {
	err = service.AdminLog().OperationLogDelete(ctx, &req.OperationLogDeleteInp)
	if err != nil {
		return nil, err
	}
	res = &api.OperationLogDeleteRes{}
	return
}

// OperationLogClear 清空操作日志
func (c *ControllerV1) OperationLogClear(ctx context.Context, req *api.OperationLogClearReq) (res *api.OperationLogClearRes, err error) {
	err = service.AdminLog().OperationLogClear(ctx)
	if err != nil {
		return nil, err
	}
	res = &api.OperationLogClearRes{}
	return
}
