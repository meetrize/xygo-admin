// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package adminin

import "xygo/internal/model/input/form"

// ==================== 文档分类（树形） ====================

// DocCategoryListInp 文档分类列表入参
type DocCategoryListInp struct {
	Title  string `p:"title"  json:"title"  dc:"按分类名称模糊搜索"`
	Status int    `p:"status" d:"-1" json:"status" dc:"状态过滤:1正常,2禁用,-1全部"`
}

// DocCategoryListItem 文档分类列表项（树形）
type DocCategoryListItem struct {
	Id        uint64                 `json:"id"        dc:"分类ID"`
	Pid       uint64                 `json:"pid"       dc:"父分类ID"`
	Title     string                 `json:"title"     dc:"分类名称"`
	Icon      string                 `json:"icon"      dc:"图标"`
	Sort      int                    `json:"sort"      dc:"排序"`
	Status    int                    `json:"status"    dc:"状态"`
	Remark    string                 `json:"remark"    dc:"备注"`
	CreatedAt uint64                 `json:"createdAt" dc:"创建时间"`
	UpdatedAt uint64                 `json:"updatedAt" dc:"更新时间"`
	Children  []*DocCategoryListItem `json:"children,omitempty" dc:"子分类"`
}

// DocCategoryListModel 文档分类列表响应
type DocCategoryListModel struct {
	List []*DocCategoryListItem `json:"list" dc:"数据列表"`
}

// DocCategorySaveInp 文档分类新增/编辑入参
type DocCategorySaveInp struct {
	Id     uint64 `p:"id"     json:"id"     dc:"分类ID（为空表示新增）"`
	Pid    uint64 `p:"pid"    d:"0"         json:"pid"    dc:"父分类ID，0表示顶级"`
	Title  string `p:"title"  v:"required#分类名称不能为空" json:"title"  dc:"分类名称"`
	Icon   string `p:"icon"                 json:"icon"   dc:"图标"`
	Sort   int    `p:"sort"   d:"0"         json:"sort"   dc:"排序"`
	Status int    `p:"status" d:"1"         json:"status" dc:"状态:1正常,2禁用"`
	Remark string `p:"remark"               json:"remark" dc:"备注"`
}

// DocCategoryDeleteInp 文档分类删除入参
type DocCategoryDeleteInp struct {
	Id uint64 `p:"id" v:"required#分类ID不能为空" json:"id" dc:"分类ID"`
}

// ==================== 文档内容 ====================

// DocListInp 文档列表入参
type DocListInp struct {
	form.PageReq
	CategoryId uint64 `p:"categoryId" json:"categoryId" dc:"分类ID"`
	Title      string `p:"title"      json:"title"      dc:"标题模糊搜索"`
	Status     int    `p:"status" d:"-1" json:"status"  dc:"状态:1已发布,2草稿,3下架,-1全部"`
}

// DocListItem 文档列表项
type DocListItem struct {
	Id           uint64 `json:"id"`
	CategoryId   uint64 `json:"categoryId"`
	CategoryName string `json:"categoryName"` // 关联查询
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Cover        string `json:"cover"`
	Summary      string `json:"summary"`
	Author       string `json:"author"`
	Views        int    `json:"views"`
	Sort         int    `json:"sort"`
	Status       int    `json:"status"`
	IsTop        int    `json:"isTop"`
	Tags         string `json:"tags"`
	CreatedBy    uint64 `json:"createdBy"`
	CreatedAt    uint64 `json:"createdAt"`
	UpdatedAt    uint64 `json:"updatedAt"`
}

// DocListModel 文档列表响应
type DocListModel struct {
	List []DocListItem `json:"list"`
	form.PageRes
}

// DocDetailInp 文档详情入参
type DocDetailInp struct {
	Id uint64 `p:"id" v:"required#文档ID不能为空" json:"id" dc:"文档ID"`
}

// DocDetailModel 文档详情响应
type DocDetailModel struct {
	Id           uint64 `json:"id"`
	CategoryId   uint64 `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Cover        string `json:"cover"`
	Summary      string `json:"summary"`
	Content      string `json:"content"`
	Author       string `json:"author"`
	Views        int    `json:"views"`
	Sort         int    `json:"sort"`
	Status       int    `json:"status"`
	IsTop        int    `json:"isTop"`
	Tags         string `json:"tags"`
	CreatedBy    uint64 `json:"createdBy"`
	UpdatedBy    uint64 `json:"updatedBy"`
	CreatedAt    uint64 `json:"createdAt"`
	UpdatedAt    uint64 `json:"updatedAt"`
}

// DocSearchItem 搜索结果项（返回标题、slug、摘要片段，不返回全文）
type DocSearchItem struct {
	Id           uint64 `json:"id"`
	CategoryId   uint64 `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	Author       string `json:"author"`
	Views        int    `json:"views"`
	MatchType    string `json:"matchType"` // "title" 或 "content"
}

// DocSaveInp 文档新增/编辑入参
type DocSaveInp struct {
	Id         uint64 `p:"id"         json:"id"         dc:"文档ID（为空表示新增）"`
	CategoryId uint64 `p:"categoryId" v:"required#分类不能为空" json:"categoryId" dc:"分类ID"`
	Title      string `p:"title"      v:"required#标题不能为空" json:"title"      dc:"文档标题"`
	Slug       string `p:"slug"       json:"slug"       dc:"URL标识"`
	Cover      string `p:"cover"      json:"cover"      dc:"封面图"`
	Summary    string `p:"summary"    json:"summary"    dc:"摘要"`
	Content    string `p:"content"    json:"content"    dc:"文档内容(Markdown)"`
	Author     string `p:"author"     json:"author"     dc:"作者"`
	Sort       int    `p:"sort"       d:"0"             json:"sort"       dc:"排序"`
	Status     int    `p:"status"     d:"1"             json:"status"     dc:"状态:1已发布,2草稿,3下架"`
	IsTop      int    `p:"isTop"      d:"0"             json:"isTop"      dc:"是否置顶"`
	Tags       string `p:"tags"       json:"tags"       dc:"标签(JSON数组)"`
}

// DocDeleteInp 文档删除入参
type DocDeleteInp struct {
	Id uint64 `p:"id" v:"required#文档ID不能为空" json:"id" dc:"文档ID"`
}
