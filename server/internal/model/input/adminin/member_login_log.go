// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package adminin

import (
	"xygo/internal/model/input/form"
)

// ==================== 登录日志 ====================

// MemberLoginLogListInp 登录日志列表入参
type MemberLoginLogListInp struct {
	form.PageReq
	Status *int `json:"status" dc:"状态:0=成功,1=失败"`
	// 关联表搜索字段
	MemberUsername string `json:"member_username" dc:"用户名"`
}

// MemberLoginLogListItem 登录日志列表项
type MemberLoginLogListItem struct {
	Id uint64 `json:"id" dc:"ID"`
	MemberId uint64 `json:"memberId" dc:"会员ID"`
	Username string `json:"username" dc:"用户名"`
	Ip string `json:"ip" dc:"登录IP"`
	UserAgent string `json:"userAgent" dc:"User-Agent"`
	Status int `json:"status" dc:"状态:0=成功,1=失败"`
	Message string `json:"message" dc:"提示信息"`
	CreatedAt uint64 `json:"createdAt" dc:"登录时间"`
	// 关联表字段（来自 LeftJoin）
	MemberUsername string `json:"member_username" dc:"Memberusername"`
}

// MemberLoginLogListModel 登录日志列表出参
type MemberLoginLogListModel struct {
	List []MemberLoginLogListItem `json:"list"`
	form.PageRes
}

// MemberLoginLogViewModel 登录日志详情出参
type MemberLoginLogViewModel struct {
	Id uint64 `json:"id" dc:"ID"`
	MemberId uint64 `json:"memberId" dc:"会员ID"`
	Username string `json:"username" dc:"用户名"`
	Ip string `json:"ip" dc:"登录IP"`
	UserAgent string `json:"userAgent" dc:"User-Agent"`
	Status int `json:"status" dc:"状态:0=成功,1=失败"`
	Message string `json:"message" dc:"提示信息"`
	CreatedAt uint64 `json:"createdAt" dc:"登录时间"`
}
