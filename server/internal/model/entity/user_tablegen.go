// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserTablegen is the golang structure for table user_tablegen.
type UserTablegen struct {
	Id            uint64      `json:"id"            orm:"id"            description:"主键"`           // 主键
	Inputtext     string      `json:"inputtext"     orm:"inputtext"     description:"测试文本"`         // 测试文本
	Inputlongtext string      `json:"inputlongtext" orm:"inputlongtext" description:"长文本"`          // 长文本
	Feditor       string      `json:"feditor"       orm:"feditor"       description:"富文本"`          // 富文本
	Inputjson     *gjson.Json `json:"inputjson"     orm:"inputjson"     description:"json"`         // json
	Inputdatetime *gtime.Time `json:"inputdatetime" orm:"inputdatetime" description:"日期时间"`         // 日期时间
	Inputenum     string      `json:"inputenum"     orm:"inputenum"     description:"状态"`           // 状态
	Inputset      string      `json:"inputset"      orm:"inputset"      description:"集合"`           // 集合
	Status        int         `json:"status"        orm:"status"        description:"状态:0=禁用,1=启用"` // 状态:0=禁用,1=启用
	Sort          uint        `json:"sort"          orm:"sort"          description:"排序"`           // 排序
	CreatedAt     uint64      `json:"createdAt"     orm:"created_at"    description:"创建时间"`         // 创建时间
	UpdatedAt     uint64      `json:"updatedAt"     orm:"updated_at"    description:"更新时间"`         // 更新时间
}
