package admin

import (
	"context"

	api "xygo/api/admin"
	"xygo/internal/service"
)

// UranExampleList 悠然示例列表
func (c *ControllerV1) UranExampleList(ctx context.Context, req *api.UranExampleListReq) (res *api.UranExampleListRes, err error) {
	result, err := service.UranExample().List(ctx, &req.UranExampleListInp)
	if err != nil {
		return nil, err
	}
	return &api.UranExampleListRes{result}, nil
}

// UranExampleView 悠然示例详情
func (c *ControllerV1) UranExampleView(ctx context.Context, req *api.UranExampleViewReq) (res *api.UranExampleViewRes, err error) {
	result, err := service.UranExample().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.UranExampleViewRes{result}, nil
}

// UranExampleEdit 保存悠然示例
func (c *ControllerV1) UranExampleEdit(ctx context.Context, req *api.UranExampleEditReq) (res *api.UranExampleEditRes, err error) {
	err = service.UranExample().Edit(ctx, &req.UranExampleEditInp)
	return &api.UranExampleEditRes{}, err
}

// UranExampleDelete 删除悠然示例
func (c *ControllerV1) UranExampleDelete(ctx context.Context, req *api.UranExampleDeleteReq) (res *api.UranExampleDeleteRes, err error) {
	err = service.UranExample().Delete(ctx, req.Id)
	return &api.UranExampleDeleteRes{}, err
}
