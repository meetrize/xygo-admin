// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UranExample is the golang structure for table uran_example.
type UranExample struct {
	Id        uint64 `json:"id"        orm:"id"         description:"主键"`           // 主键
	Status    int    `json:"status"    orm:"status"     description:"状态:0=禁用,1=启用"` // 状态:0=禁用,1=启用
	Sort      uint   `json:"sort"      orm:"sort"       description:"排序"`           // 排序
	CreatedAt uint64 `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt uint64 `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
	Tname     string `json:"tname"     orm:"tname"      description:"标题"`           // 标题
	Timage    string `json:"timage"    orm:"timage"     description:"图片"`           // 图片
	Ttest     string `json:"ttest"     orm:"ttest"      description:"测试"`           // 测试
}
