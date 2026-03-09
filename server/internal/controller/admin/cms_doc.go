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

// ===================== 文档分类管理 =====================

// DocCategoryList 文档分类列表
func (c *ControllerV1) DocCategoryList(ctx context.Context, req *api.DocCategoryListReq) (res *api.DocCategoryListRes, err error) {
	result, err := service.CmsDoc().CategoryList(ctx, &req.DocCategoryListInp)
	if err != nil {
		return nil, err
	}
	return &api.DocCategoryListRes{DocCategoryListModel: result}, nil
}

// DocCategorySave 新增/编辑文档分类
func (c *ControllerV1) DocCategorySave(ctx context.Context, req *api.DocCategorySaveReq) (res *api.DocCategorySaveRes, err error) {
	id, err := service.CmsDoc().CategorySave(ctx, &req.DocCategorySaveInp)
	if err != nil {
		return nil, err
	}
	return &api.DocCategorySaveRes{Id: id}, nil
}

// DocCategoryDelete 删除文档分类
func (c *ControllerV1) DocCategoryDelete(ctx context.Context, req *api.DocCategoryDeleteReq) (res *api.DocCategoryDeleteRes, err error) {
	err = service.CmsDoc().CategoryDelete(ctx, req.Id)
	return &api.DocCategoryDeleteRes{}, err
}

// ===================== 文档内容管理 =====================

// DocList 文档列表
func (c *ControllerV1) DocList(ctx context.Context, req *api.DocListReq) (res *api.DocListRes, err error) {
	result, err := service.CmsDoc().List(ctx, &req.DocListInp)
	if err != nil {
		return nil, err
	}
	return &api.DocListRes{DocListModel: result}, nil
}

// DocDetail 文档详情
func (c *ControllerV1) DocDetail(ctx context.Context, req *api.DocDetailReq) (res *api.DocDetailRes, err error) {
	result, err := service.CmsDoc().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.DocDetailRes{DocDetailModel: result}, nil
}

// DocSave 新增/编辑文档
func (c *ControllerV1) DocSave(ctx context.Context, req *api.DocSaveReq) (res *api.DocSaveRes, err error) {
	operatorId := contexts.GetUserId(ctx)
	id, err := service.CmsDoc().Save(ctx, &req.DocSaveInp, operatorId)
	if err != nil {
		return nil, err
	}
	return &api.DocSaveRes{Id: id}, nil
}

// DocDelete 删除文档
func (c *ControllerV1) DocDelete(ctx context.Context, req *api.DocDeleteReq) (res *api.DocDeleteRes, err error) {
	err = service.CmsDoc().Delete(ctx, req.Id)
	return &api.DocDeleteRes{}, err
}
