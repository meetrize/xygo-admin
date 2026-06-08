package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	adminin "xygo/internal/model/input/adminin"
)

// UranExampleListReq 悠然示例列表请求
type UranExampleListReq struct {
	g.Meta `path:"/admin/uran-example/list" method:"get" tags:"UranExample" summary:"悠然示例列表"`
	adminin.UranExampleListInp
}

type UranExampleListRes struct {
	*adminin.UranExampleListModel
}

// UranExampleViewReq 悠然示例详情请求
type UranExampleViewReq struct {
	g.Meta `path:"/admin/uran-example/view" method:"get" tags:"UranExample" summary:"悠然示例详情"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type UranExampleViewRes struct {
	*adminin.UranExampleViewModel
}

// UranExampleEditReq 悠然示例保存请求
type UranExampleEditReq struct {
	g.Meta `path:"/admin/uran-example/edit" method:"post" tags:"UranExample" summary:"保存悠然示例"`
	adminin.UranExampleEditInp
}

type UranExampleEditRes struct{}

// UranExampleDeleteReq 悠然示例删除请求
type UranExampleDeleteReq struct {
	g.Meta `path:"/admin/uran-example/delete" method:"post" tags:"UranExample" summary:"删除悠然示例"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type UranExampleDeleteRes struct{}
