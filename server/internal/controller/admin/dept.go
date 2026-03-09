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
	"xygo/internal/service"
)

// DeptList 部门列表（树形结构）
func (c *ControllerV1) DeptList(ctx context.Context, req *api.DeptListReq) (res *api.DeptListRes, err error) {
	list, err := service.AdminDept().List(ctx, &req.DeptListInp)
	if err != nil {
		return nil, err
	}

	res = &api.DeptListRes{
		DeptListModel: adminin.DeptListModel{
			List: list,
		},
	}
	return
}

// DeptDetail 部门详情
func (c *ControllerV1) DeptDetail(ctx context.Context, req *api.DeptDetailReq) (res *api.DeptDetailRes, err error) {
	detail, err := service.AdminDept().Detail(ctx, uint64(req.Id))
	if err != nil {
		return nil, err
	}

	res = &api.DeptDetailRes{
		DeptListItem: *detail,
	}
	return
}

// DeptSave 部门保存（新增/编辑）
func (c *ControllerV1) DeptSave(ctx context.Context, req *api.DeptSaveReq) (res *api.DeptSaveRes, err error) {
	id, err := service.AdminDept().Save(ctx, &req.DeptSaveInp)
	if err != nil {
		return nil, err
	}

	res = &api.DeptSaveRes{Id: id}
	return
}

// DeptDelete 部门删除
func (c *ControllerV1) DeptDelete(ctx context.Context, req *api.DeptDeleteReq) (res *api.DeptDeleteRes, err error) {
	err = service.AdminDept().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.DeptDeleteRes{}, nil
}
