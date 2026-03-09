// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"xygo/internal/model/input/adminin"
)

// MemberMoneyLogListReq 余额变动日志列表请求
type MemberMoneyLogListReq struct {
	g.Meta `path:"/admin/member-money-log/list" method:"get" tags:"MemberMoneyLog" summary:"余额变动日志列表"`
	adminin.MemberMoneyLogListInp
}

type MemberMoneyLogListRes struct {
	*adminin.MemberMoneyLogListModel
}

// MemberMoneyLogViewReq 余额变动日志详情请求
type MemberMoneyLogViewReq struct {
	g.Meta `path:"/admin/member-money-log/view" method:"get" tags:"MemberMoneyLog" summary:"余额变动日志详情"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type MemberMoneyLogViewRes struct {
	*adminin.MemberMoneyLogViewModel
}

// MemberMoneyLogEditReq 余额变动日志保存请求
type MemberMoneyLogEditReq struct {
	g.Meta `path:"/admin/member-money-log/edit" method:"post" tags:"MemberMoneyLog" summary:"保存余额变动日志"`
	adminin.MemberMoneyLogEditInp
}

type MemberMoneyLogEditRes struct{}

// MemberMoneyLogDeleteReq 余额变动日志删除请求
type MemberMoneyLogDeleteReq struct {
	g.Meta `path:"/admin/member-money-log/delete" method:"post" tags:"MemberMoneyLog" summary:"删除余额变动日志"`
	Id uint64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

type MemberMoneyLogDeleteRes struct{}
