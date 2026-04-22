// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UranContact is the golang structure for table uran_contact.
type UranContact struct {
	Id          uint64 `json:"id"          orm:"id"           description:"主键"`         // 主键
	Username    string `json:"username"    orm:"username"     description:"姓名"`         // 姓名
	Phone       string `json:"phone"       orm:"phone"        description:"电话"`         // 电话
	Age         int    `json:"age"         orm:"age"          description:"数字"`         // 数字
	Avatar      string `json:"avatar"      orm:"avatar"       description:"头像"`         // 头像
	Sort        int    `json:"sort"        orm:"sort"         description:"排序权重"`       // 排序权重
	Remark      string `json:"remark"      orm:"remark"       description:"备注"`         // 备注
	SwitchField int    `json:"switchField" orm:"switch_field" description:"开关:0=关,1=开"` // 开关:0=关,1=开
}
