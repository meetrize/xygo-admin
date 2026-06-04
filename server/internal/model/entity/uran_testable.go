// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UranTestable is the golang structure for table uran_testable.
type UranTestable struct {
	Id        uint64 `json:"id"        orm:"id"         description:"主键"`           // 主键
	Status    int    `json:"status"    orm:"status"     description:"状态:0=禁用,1=启用"` // 状态:0=禁用,1=启用
	Sort      uint   `json:"sort"      orm:"sort"       description:"排序"`           // 排序
	CreatedAt uint64 `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt uint64 `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
	Stuff     string `json:"stuff"     orm:"stuff"      description:"物件名称"`         // 物件名称
	Price     string `json:"price"     orm:"price"      description:"价格"`           // 价格
}
