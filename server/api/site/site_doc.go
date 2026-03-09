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
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/model/input/adminin"
)

// ===================== 前台文档公开接口 =====================

// DocCategoryTreeReq 获取文档分类树（前台展示用）
type DocCategoryTreeReq struct {
	g.Meta `path:"/site/doc/categoryTree" method:"get" tags:"SiteDoc" summary:"文档分类树"`
}
type DocCategoryTreeRes struct {
	List []*adminin.DocCategoryListItem `json:"list"`
}

// DocListByCategoryReq 按分类获取文档列表（前台展示用）
type DocListByCategoryReq struct {
	g.Meta     `path:"/site/doc/list" method:"get" tags:"SiteDoc" summary:"文档列表"`
	CategoryId uint64 `p:"categoryId" json:"categoryId" dc:"分类ID"`
}
type DocListByCategoryRes struct {
	List []adminin.DocListItem `json:"list"`
}

// DocDetailBySlugReq 按 slug 获取文档详情（前台展示用）
type DocDetailBySlugReq struct {
	g.Meta `path:"/site/doc/detail" method:"get" tags:"SiteDoc" summary:"文档详情"`
	Slug   string `p:"slug" v:"required#文档标识不能为空" json:"slug" dc:"文档URL标识"`
}
type DocDetailBySlugRes struct {
	*adminin.DocDetailModel
}

// DocSearchReq 全文搜索文档（标题 + 内容）
type DocSearchReq struct {
	g.Meta  `path:"/site/doc/search" method:"get" tags:"SiteDoc" summary:"搜索文档"`
	Keyword string `p:"keyword" v:"required#搜索关键词不能为空" json:"keyword" dc:"搜索关键词"`
}
type DocSearchRes struct {
	List []adminin.DocSearchItem `json:"list"`
}
