package adminin

import (
	"xygo/internal/model/input/form"
)

// ==================== 测试记帐 ====================

// UranTestableListInp 测试记帐列表入参
type UranTestableListInp struct {
	form.PageReq
	Status *int `json:"status" dc:"状态:0=禁用,1=启用"`
}

// UranTestableListItem 测试记帐列表项
type UranTestableListItem struct {
	Id uint64 `json:"id" dc:"主键"`
	Status int `json:"status" dc:"状态:0=禁用,1=启用"`
	Sort uint `json:"sort" dc:"排序"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
	UpdatedAt uint64 `json:"updatedAt" dc:"更新时间"`
	Stuff string `json:"stuff" dc:"物件名称"`
	Price string `json:"price" dc:"价格"`
}

// UranTestableListModel 测试记帐列表出参
type UranTestableListModel struct {
	List []UranTestableListItem `json:"list"`
	form.PageRes
}

// UranTestableViewModel 测试记帐详情出参
type UranTestableViewModel struct {
	Id uint64 `json:"id" dc:"主键"`
	Status int `json:"status" dc:"状态:0=禁用,1=启用"`
	Sort uint `json:"sort" dc:"排序"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
	UpdatedAt uint64 `json:"updatedAt" dc:"更新时间"`
	Stuff string `json:"stuff" dc:"物件名称"`
	Price string `json:"price" dc:"价格"`
}

// UranTestableEditInp 测试记帐编辑入参
type UranTestableEditInp struct {
	Id uint64 `json:"id" dc:"主键"`
	Status int `json:"status" v:"required#状态:0=禁用,1=启用不能为空" dc:"状态:0=禁用,1=启用"`
	Sort uint `json:"sort" v:"required#排序不能为空" dc:"排序"`
	Stuff string `json:"stuff" v:"required#物件名称不能为空" dc:"物件名称"`
	Price string `json:"price" v:"required#价格不能为空" dc:"价格"`
}
