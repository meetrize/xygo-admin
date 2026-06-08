package adminin

import (
	"xygo/internal/model/input/form"
)

// ==================== 悠然示例 ====================

// UranExampleListInp 悠然示例列表入参
type UranExampleListInp struct {
	form.PageReq
	Status *int `json:"status" dc:"状态:0=禁用,1=启用"`
}

// UranExampleListItem 悠然示例列表项
type UranExampleListItem struct {
	Id uint64 `json:"id" dc:"主键"`
	Status int `json:"status" dc:"状态:0=禁用,1=启用"`
	Sort uint `json:"sort" dc:"排序"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
	UpdatedAt uint64 `json:"updatedAt" dc:"更新时间"`
	Tname string `json:"tname" dc:"标题"`
	Timage string `json:"timage" dc:"图片"`
	Ttest string `json:"ttest" dc:"测试"`
}

// UranExampleListModel 悠然示例列表出参
type UranExampleListModel struct {
	List []UranExampleListItem `json:"list"`
	form.PageRes
}

// UranExampleViewModel 悠然示例详情出参
type UranExampleViewModel struct {
	Id uint64 `json:"id" dc:"主键"`
	Status int `json:"status" dc:"状态:0=禁用,1=启用"`
	Sort uint `json:"sort" dc:"排序"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
	UpdatedAt uint64 `json:"updatedAt" dc:"更新时间"`
	Tname string `json:"tname" dc:"标题"`
	Timage string `json:"timage" dc:"图片"`
	Ttest string `json:"ttest" dc:"测试"`
}

// UranExampleEditInp 悠然示例编辑入参
type UranExampleEditInp struct {
	Id uint64 `json:"id" dc:"主键"`
	Status int `json:"status" v:"required#状态:0=禁用,1=启用不能为空" dc:"状态:0=禁用,1=启用"`
	Sort uint `json:"sort" v:"required#排序不能为空" dc:"排序"`
	Tname string `json:"tname" v:"required#标题不能为空" dc:"标题"`
	Timage string `json:"timage" v:"required#图片不能为空" dc:"图片"`
	Ttest string `json:"ttest" v:"required#测试不能为空" dc:"测试"`
}
