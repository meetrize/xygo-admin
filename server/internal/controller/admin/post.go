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

// PostList 岗位列表
func (c *ControllerV1) PostList(ctx context.Context, req *api.PostListReq) (res *api.PostListRes, err error) {
	list, total, err := service.AdminPost().List(ctx, &req.PostListInp)
	if err != nil {
		return nil, err
	}

	res = &api.PostListRes{
		PostListModel: adminin.PostListModel{
			List: list,
			PageRes: form.PageRes{
				Page:     req.Page,
				PageSize: req.PageSize,
				Total:    total,
			},
		},
	}
	return
}

// PostDetail 岗位详情
func (c *ControllerV1) PostDetail(ctx context.Context, req *api.PostDetailReq) (res *api.PostDetailRes, err error) {
	detail, err := service.AdminPost().Detail(ctx, uint64(req.Id))
	if err != nil {
		return nil, err
	}

	res = &api.PostDetailRes{
		PostListItem: *detail,
	}
	return
}

// PostSave 岗位保存（新增/编辑）
func (c *ControllerV1) PostSave(ctx context.Context, req *api.PostSaveReq) (res *api.PostSaveRes, err error) {
	id, err := service.AdminPost().Save(ctx, &req.PostSaveInp)
	if err != nil {
		return nil, err
	}

	res = &api.PostSaveRes{Id: id}
	return
}

// PostDelete 岗位删除
func (c *ControllerV1) PostDelete(ctx context.Context, req *api.PostDeleteReq) (res *api.PostDeleteRes, err error) {
	err = service.AdminPost().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.PostDeleteRes{}, nil
}
