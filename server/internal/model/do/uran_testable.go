// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UranTestable is the golang structure of table xy_uran_testable for DAO operations like Where/Data.
type UranTestable struct {
	g.Meta    `orm:"table:xy_uran_testable, do:true"`
	Id        any // 主键
	Status    any // 状态:0=禁用,1=启用
	Sort      any // 排序
	CreatedAt any // 创建时间
	UpdatedAt any // 更新时间
	Stuff     any // 物件名称
	Price     any // 价格
}
