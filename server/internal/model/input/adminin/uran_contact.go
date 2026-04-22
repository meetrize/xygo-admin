package adminin

import (
	"xygo/internal/model/input/form"
)

// ==================== 悠然联系人 ====================

// UranContactListInp 悠然联系人列表入参
type UranContactListInp struct {
	form.PageReq
}

// UranContactListItem 悠然联系人列表项
type UranContactListItem struct {
	Id uint64 `json:"id" dc:"主键"`
	Username string `json:"username" dc:"姓名"`
	Phone string `json:"phone" dc:"电话"`
	Age int `json:"age" dc:"数字"`
	Avatar string `json:"avatar" dc:"头像"`
	Sort int `json:"sort" dc:"排序权重"`
	Remark string `json:"remark" dc:"备注"`
	SwitchField int `json:"switchField" dc:"开关:0=关,1=开"`
}

// UranContactListModel 悠然联系人列表出参
type UranContactListModel struct {
	List []UranContactListItem `json:"list"`
	form.PageRes
}

// UranContactViewModel 悠然联系人详情出参
type UranContactViewModel struct {
	Id uint64 `json:"id" dc:"主键"`
	Username string `json:"username" dc:"姓名"`
	Phone string `json:"phone" dc:"电话"`
	Age int `json:"age" dc:"数字"`
	Avatar string `json:"avatar" dc:"头像"`
	Sort int `json:"sort" dc:"排序权重"`
	Remark string `json:"remark" dc:"备注"`
	SwitchField int `json:"switchField" dc:"开关:0=关,1=开"`
}

// UranContactEditInp 悠然联系人编辑入参
type UranContactEditInp struct {
	Id uint64 `json:"id" dc:"主键"`
	Username string `json:"username" v:"required#姓名不能为空" dc:"姓名"`
	Phone string `json:"phone" v:"required#电话不能为空" dc:"电话"`
	Age int `json:"age" v:"required#数字不能为空" dc:"数字"`
	Avatar string `json:"avatar" v:"required#头像不能为空" dc:"头像"`
	Sort int `json:"sort" v:"required#排序权重不能为空" dc:"排序权重"`
	Remark string `json:"remark" v:"required#备注不能为空" dc:"备注"`
	SwitchField int `json:"switchField" v:"required#开关:0=关,1=开不能为空" dc:"开关:0=关,1=开"`
}
