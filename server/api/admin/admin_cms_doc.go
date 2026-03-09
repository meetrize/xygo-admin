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
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/model/input/adminin"
)

// ===================== 文档分类管理 =====================

type DocCategoryListReq struct {
	g.Meta `path:"/admin/cms/docCategory/list" method:"post" tags:"CmsDocCategory" summary:"文档分类列表"`
	adminin.DocCategoryListInp
}
type DocCategoryListRes struct {
	*adminin.DocCategoryListModel
}

type DocCategorySaveReq struct {
	g.Meta `path:"/admin/cms/docCategory/save" method:"post" tags:"CmsDocCategory" summary:"新增/编辑文档分类"`
	adminin.DocCategorySaveInp
}
type DocCategorySaveRes struct {
	Id uint64 `json:"id"`
}

type DocCategoryDeleteReq struct {
	g.Meta `path:"/admin/cms/docCategory/delete" method:"post" tags:"CmsDocCategory" summary:"删除文档分类"`
	adminin.DocCategoryDeleteInp
}
type DocCategoryDeleteRes struct{}

// ===================== 文档内容管理 =====================

type DocListReq struct {
	g.Meta `path:"/admin/cms/doc/list" method:"post" tags:"CmsDoc" summary:"文档列表"`
	adminin.DocListInp
}
type DocListRes struct {
	*adminin.DocListModel
}

type DocDetailReq struct {
	g.Meta `path:"/admin/cms/doc/detail" method:"get" tags:"CmsDoc" summary:"文档详情"`
	adminin.DocDetailInp
}
type DocDetailRes struct {
	*adminin.DocDetailModel
}

type DocSaveReq struct {
	g.Meta `path:"/admin/cms/doc/save" method:"post" tags:"CmsDoc" summary:"新增/编辑文档"`
	adminin.DocSaveInp
}
type DocSaveRes struct {
	Id uint64 `json:"id"`
}

type DocDeleteReq struct {
	g.Meta `path:"/admin/cms/doc/delete" method:"post" tags:"CmsDoc" summary:"删除文档"`
	adminin.DocDeleteInp
}
type DocDeleteRes struct{}
