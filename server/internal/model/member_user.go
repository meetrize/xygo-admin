// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package model

// MemberUser 会员登录态中携带的用户信息（用于前台认证）
type MemberUser struct {
	// 基础信息
	Id       uint64 `json:"id"       description:"会员ID"`
	Username string `json:"username" description:"用户名"`
	Nickname string `json:"nickname" description:"昵称"`
	Avatar   string `json:"avatar"   description:"头像"`
	Email    string `json:"email"    description:"邮箱"`
	Mobile   string `json:"mobile"   description:"手机号"`

	// 会员属性
	Gender  int     `json:"gender"  description:"性别：0=未知 1=男 2=女"`
	Level   uint    `json:"level"   description:"会员等级"`
	GroupId uint64  `json:"groupId" description:"会员分组ID"`
	Score   int     `json:"score"   description:"积分"`
	Money   float64 `json:"money"   description:"余额"`

	// 登录信息
	LoginAt int64 `json:"loginAt" description:"登录时间"`
}

// MemberClaims 会员 JWT 载荷
type MemberClaims struct {
	*MemberUser
}
