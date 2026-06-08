// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UranExample is the golang structure of table xy_uran_example for DAO operations like Where/Data.
type UranExample struct {
	g.Meta    `orm:"table:xy_uran_example, do:true"`
	Id        any // 主键
	Status    any // 状态:0=禁用,1=启用
	Sort      any // 排序
	CreatedAt any // 创建时间
	UpdatedAt any // 更新时间
	Tname     any // 标题
	Timage    any // 图片
	Ttest     any // 测试
}
