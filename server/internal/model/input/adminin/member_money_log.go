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

// ==================== 余额变动日志 ====================

// MemberMoneyLogListInp 余额变动日志列表入参
type MemberMoneyLogListInp struct {
	form.PageReq
	MoneyStart string `json:"moneyStart" dc:"变动金额开始值"`
	MoneyEnd string `json:"moneyEnd" dc:"变动金额结束值"`
	// 关联表搜索字段
	MemberUsername string `json:"member_username" dc:"用户名"`
	MemberNickname string `json:"member_nickname" dc:"昵称"`
}

// MemberMoneyLogListItem 余额变动日志列表项
type MemberMoneyLogListItem struct {
	Id uint64 `json:"id" dc:""`
	MemberId uint64 `json:"memberId" dc:"会员ID"`
	Money int `json:"money" dc:"变动金额"`
	Before int `json:"before" dc:"变动前余额（分）"`
	After int `json:"after" dc:"变动后余额（分）"`
	Memo string `json:"memo" dc:"变动说明"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
	// 关联表字段（来自 LeftJoin）
	MemberNickname string `json:"member_nickname" dc:"Membernickname"`
	MemberUsername string `json:"member_username" dc:"Memberusername"`
	MemberAvatar string `json:"member_avatar" dc:"Memberavatar"`
}

// MemberMoneyLogListModel 余额变动日志列表出参
type MemberMoneyLogListModel struct {
	List []MemberMoneyLogListItem `json:"list"`
	form.PageRes
}

// MemberMoneyLogViewModel 余额变动日志详情出参
type MemberMoneyLogViewModel struct {
	Id uint64 `json:"id" dc:""`
	MemberId uint64 `json:"memberId" dc:"会员ID"`
	Money int `json:"money" dc:"变动金额"`
	Before int `json:"before" dc:"变动前余额（分）"`
	After int `json:"after" dc:"变动后余额（分）"`
	Memo string `json:"memo" dc:"变动说明"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
}

// MemberMoneyLogEditInp 余额变动日志编辑入参
type MemberMoneyLogEditInp struct {
	Id uint64 `json:"id" dc:""`
	MemberId uint64 `json:"memberId" v:"required#会员ID不能为空" dc:"会员ID"`
	Money int `json:"money" v:"required#变动金额不能为空" dc:"变动金额"`
	Before int `json:"before" v:"required#变动前余额（分）不能为空" dc:"变动前余额（分）"`
	After int `json:"after" v:"required#变动后余额（分）不能为空" dc:"变动后余额（分）"`
	Memo string `json:"memo" v:"required#变动说明不能为空" dc:"变动说明"`
}
