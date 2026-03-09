// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package site

import (
	"context"

	api "xygo/api/site"
	"xygo/internal/model/input/adminin"
	"xygo/internal/service"
)

// DocCategoryTree 前台文档分类树
func (c *ControllerV1) DocCategoryTree(ctx context.Context, req *api.DocCategoryTreeReq) (res *api.DocCategoryTreeRes, err error) {
	result, err := service.CmsDoc().CategoryList(ctx, &adminin.DocCategoryListInp{Status: 1})
	if err != nil {
		return nil, err
	}
	return &api.DocCategoryTreeRes{List: result.List}, nil
}

// DocListByCategory 前台按分类获取文档列表
func (c *ControllerV1) DocListByCategory(ctx context.Context, req *api.DocListByCategoryReq) (res *api.DocListByCategoryRes, err error) {
	result, err := service.CmsDoc().List(ctx, &adminin.DocListInp{
		CategoryId: req.CategoryId,
		Status:     1, // 前台只显示已发布
	})
	if err != nil {
		return nil, err
	}
	return &api.DocListByCategoryRes{List: result.List}, nil
}

// DocDetailBySlug 前台按 slug 获取文档详情
func (c *ControllerV1) DocDetailBySlug(ctx context.Context, req *api.DocDetailBySlugReq) (res *api.DocDetailBySlugRes, err error) {
	result, err := service.CmsDoc().DetailBySlug(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	return &api.DocDetailBySlugRes{DocDetailModel: result}, nil
}

// DocSearch 前台全文搜索文档
func (c *ControllerV1) DocSearch(ctx context.Context, req *api.DocSearchReq) (res *api.DocSearchRes, err error) {
	list, err := service.CmsDoc().Search(ctx, req.Keyword)
	if err != nil {
		return nil, err
	}
	return &api.DocSearchRes{List: list}, nil
}
