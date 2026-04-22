package admin

import (
	"context"

	api "xygo/api/admin"
	"xygo/internal/service"
)

// UranContactList 悠然联系人列表
func (c *ControllerV1) UranContactList(ctx context.Context, req *api.UranContactListReq) (res *api.UranContactListRes, err error) {
	result, err := service.UranContact().List(ctx, &req.UranContactListInp)
	if err != nil {
		return nil, err
	}
	return &api.UranContactListRes{result}, nil
}

// UranContactView 悠然联系人详情
func (c *ControllerV1) UranContactView(ctx context.Context, req *api.UranContactViewReq) (res *api.UranContactViewRes, err error) {
	result, err := service.UranContact().View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.UranContactViewRes{result}, nil
}

// UranContactEdit 保存悠然联系人
func (c *ControllerV1) UranContactEdit(ctx context.Context, req *api.UranContactEditReq) (res *api.UranContactEditRes, err error) {
	err = service.UranContact().Edit(ctx, &req.UranContactEditInp)
	return &api.UranContactEditRes{}, err
}

// UranContactDelete 删除悠然联系人
func (c *ControllerV1) UranContactDelete(ctx context.Context, req *api.UranContactDeleteReq) (res *api.UranContactDeleteRes, err error) {
	err = service.UranContact().Delete(ctx, req.Id)
	return &api.UranContactDeleteRes{}, err
}
