package admin

import (
	"context"

	api "xygo/api/admin"
	"xygo/internal/service"
)

// UranTestableList 测试记帐列表
func (c *ControllerV1) UranTestableList(ctx context.Context, req *api.UranTestableListReq) (res *api.UranTestableListRes, err error) {
	result, err := service.UranTestable().List(ctx, &req.UranTestableListInp)
	if err != nil {
		return nil, err
	}
	return &api.UranTestableListRes{result}, nil
}

// UranTestableView 测试记帐详情
func (c *ControllerV1) UranTestableView(ctx context.Context, req *api.UranTestableViewReq) (res *api.UranTestableViewRes, err error) {
	result, err := service.UranTestable().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.UranTestableViewRes{result}, nil
}

// UranTestableEdit 保存测试记帐
func (c *ControllerV1) UranTestableEdit(ctx context.Context, req *api.UranTestableEditReq) (res *api.UranTestableEditRes, err error) {
	err = service.UranTestable().Edit(ctx, &req.UranTestableEditInp)
	return &api.UranTestableEditRes{}, err
}

// UranTestableDelete 删除测试记帐
func (c *ControllerV1) UranTestableDelete(ctx context.Context, req *api.UranTestableDeleteReq) (res *api.UranTestableDeleteRes, err error) {
	err = service.UranTestable().Delete(ctx, req.Id)
	return &api.UranTestableDeleteRes{}, err
}
