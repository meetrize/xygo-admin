// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SmsVariable is the golang structure of table xy_sms_variable for DAO operations like Where/Data.
type SmsVariable struct {
	g.Meta      `orm:"table:xy_sms_variable, do:true"`
	Id          any //
	Title       any // 变量标题
	Name        any // 变量名（如 usermobile）
	SourceType  any // 来源类型：1=字段提取 2=SQL查询 3=内置Helper
	SqlQuery    any // SQL查询语句（source_type=2 时）
	MethodName  any // Helper方法路径（source_type=3 时）
	SharedCount any // 共通数据数
	Status      any // 状态：1=启用 0=禁用
	CreatedBy   any // 创建人ID
	UpdatedBy   any // 更新人ID
	CreateTime  any // 创建时间（Unix秒）
	UpdateTime  any // 更新时间（Unix秒）
}
