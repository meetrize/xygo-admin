// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserTablegen is the golang structure of table xy_user_tablegen for DAO operations like Where/Data.
type UserTablegen struct {
	g.Meta        `orm:"table:xy_user_tablegen, do:true"`
	Id            any         // 主键
	Inputtext     any         // 测试文本
	Inputlongtext any         // 长文本
	Feditor       any         // 富文本
	Inputjson     *gjson.Json // json
	Inputdatetime *gtime.Time // 日期时间
	Inputenum     any         // 状态
	Inputset      any         // 集合
	Status        any         // 状态:0=禁用,1=启用
	Sort          any         // 排序
	CreatedAt     any         // 创建时间
	UpdatedAt     any         // 更新时间
}
