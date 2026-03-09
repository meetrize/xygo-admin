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

// ==================== 积分变动日志 ====================

// MemberScoreLogListInp 积分变动日志列表入参
type MemberScoreLogListInp struct {
	form.PageReq
	ScoreStart string `json:"scoreStart" dc:"变动积分开始值"`
	ScoreEnd string `json:"scoreEnd" dc:"变动积分结束值"`
	// 关联表搜索字段
	MemberUsername string `json:"member_username" dc:"用户名"`
	MemberNickname string `json:"member_nickname" dc:"昵称"`
}

// MemberScoreLogListItem 积分变动日志列表项
type MemberScoreLogListItem struct {
	Id uint64 `json:"id" dc:""`
	MemberId uint64 `json:"memberId" dc:"会员ID"`
	Score int `json:"score" dc:"变动积分"`
	Before int `json:"before" dc:"变动前积分"`
	After int `json:"after" dc:"变动后积分"`
	Memo string `json:"memo" dc:"变动说明"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
	// 关联表字段（来自 LeftJoin）
	MemberNickname string `json:"member_nickname" dc:"Membernickname"`
	MemberUsername string `json:"member_username" dc:"Memberusername"`
	MemberAvatar string `json:"member_avatar" dc:"Memberavatar"`
}

// MemberScoreLogListModel 积分变动日志列表出参
type MemberScoreLogListModel struct {
	List []MemberScoreLogListItem `json:"list"`
	form.PageRes
}

// MemberScoreLogViewModel 积分变动日志详情出参
type MemberScoreLogViewModel struct {
	Id uint64 `json:"id" dc:""`
	MemberId uint64 `json:"memberId" dc:"会员ID"`
	Score int `json:"score" dc:"变动积分"`
	Before int `json:"before" dc:"变动前积分"`
	After int `json:"after" dc:"变动后积分"`
	Memo string `json:"memo" dc:"变动说明"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
}

// MemberScoreLogEditInp 积分变动日志编辑入参
type MemberScoreLogEditInp struct {
	Id uint64 `json:"id" dc:""`
	MemberId uint64 `json:"memberId" v:"required#会员ID不能为空" dc:"会员ID"`
	Score int `json:"score" v:"required#变动积分不能为空" dc:"变动积分"`
	Before int `json:"before" v:"required#变动前积分不能为空" dc:"变动前积分"`
	After int `json:"after" v:"required#变动后积分不能为空" dc:"变动后积分"`
	Memo string `json:"memo" v:"required#变动说明不能为空" dc:"变动说明"`
}
