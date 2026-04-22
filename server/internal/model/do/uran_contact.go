// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UranContact is the golang structure of table xy_uran_contact for DAO operations like Where/Data.
type UranContact struct {
	g.Meta      `orm:"table:xy_uran_contact, do:true"`
	Id          any // 主键
	Username    any // 姓名
	Phone       any // 电话
	Age         any // 数字
	Avatar      any // 头像
	Sort        any // 排序权重
	Remark      any // 备注
	SwitchField any // 开关:0=关,1=开
}
