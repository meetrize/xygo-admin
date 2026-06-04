package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	adminin "xygo/internal/model/input/adminin"
)

// UranTestableListReq 测试记帐列表请求
type UranTestableListReq struct {
	g.Meta `path:"/admin/uran-testable/list" method:"get" tags:"UranTestable" summary:"测试记帐列表"`
	adminin.UranTestableListInp
}

type UranTestableListRes struct {
	*adminin.UranTestableListModel
}

// UranTestableViewReq 测试记帐详情请求
type UranTestableViewReq struct {
	g.Meta `path:"/admin/uran-testable/view" method:"get" tags:"UranTestable" summary:"测试记帐详情"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type UranTestableViewRes struct {
	*adminin.UranTestableViewModel
}

// UranTestableEditReq 测试记帐保存请求
type UranTestableEditReq struct {
	g.Meta `path:"/admin/uran-testable/edit" method:"post" tags:"UranTestable" summary:"保存测试记帐"`
	adminin.UranTestableEditInp
}

type UranTestableEditRes struct{}

// UranTestableDeleteReq 测试记帐删除请求
type UranTestableDeleteReq struct {
	g.Meta `path:"/admin/uran-testable/delete" method:"post" tags:"UranTestable" summary:"删除测试记帐"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type UranTestableDeleteRes struct{}
