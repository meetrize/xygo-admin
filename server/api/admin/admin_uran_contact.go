package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"xygo/internal/model/input/adminin"
)

// UranContactListReq 悠然联系人列表请求
type UranContactListReq struct {
	g.Meta `path:"/admin/uran-contact/list" method:"get" tags:"UranContact" summary:"悠然联系人列表"`
	adminin.UranContactListInp
}

type UranContactListRes struct {
	*adminin.UranContactListModel
}

// UranContactViewReq 悠然联系人详情请求
type UranContactViewReq struct {
	g.Meta `path:"/admin/uran-contact/view" method:"get" tags:"UranContact" summary:"悠然联系人详情"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type UranContactViewRes struct {
	*adminin.UranContactViewModel
}

// UranContactEditReq 悠然联系人保存请求
type UranContactEditReq struct {
	g.Meta `path:"/admin/uran-contact/edit" method:"post" tags:"UranContact" summary:"保存悠然联系人"`
	adminin.UranContactEditInp
}

type UranContactEditRes struct{}

// UranContactDeleteReq 悠然联系人删除请求
type UranContactDeleteReq struct {
	g.Meta `path:"/admin/uran-contact/delete" method:"post" tags:"UranContact" summary:"删除悠然联系人"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type UranContactDeleteRes struct{}
