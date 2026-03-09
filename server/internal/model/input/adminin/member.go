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

import "xygo/internal/model/input/form"

// ===================== 会员列表 =====================

// MemberListInp 会员列表入参
type MemberListInp struct {
	form.PageReq
	Username string `json:"username" dc:"用户名（模糊搜索）"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Email    string `json:"email" dc:"邮箱"`
	Status   int    `json:"status" d:"-1" dc:"状态：-1=全部 0=禁用 1=正常"`
	GroupId  int64  `json:"groupId" dc:"会员分组ID"`
}

// MemberListModel 会员列表出参
type MemberListModel struct {
	List []MemberItem `json:"list"`
	form.PageRes
}

// MemberItem 会员列表项
type MemberItem struct {
	Id          int64   `json:"id"`
	Username    string  `json:"username"`
	Nickname    string  `json:"nickname"`
	Mobile      string  `json:"mobile"`
	Email       string  `json:"email"`
	Avatar      string  `json:"avatar"`
	Gender      int     `json:"gender"`
	Level       int     `json:"level"`
	GroupId     int64   `json:"groupId"`
	GroupName   string  `json:"groupName"`
	Score       int     `json:"score"`
	Money       float64 `json:"money"`
	Status      int     `json:"status"`
	LoginCount  int     `json:"loginCount"`
	LastLoginAt string  `json:"lastLoginAt"`
	LastLoginIp string  `json:"lastLoginIp"`
	CreatedAt   string  `json:"createdAt"`
}

// ===================== 会员详情 =====================

// MemberDetailInp 会员详情入参
type MemberDetailInp struct {
	Id int64 `json:"id" v:"required|min:1" dc:"会员ID"`
}

// MemberDetailModel 会员详情出参
type MemberDetailModel struct {
	MemberItem
}

// ===================== 添加会员 =====================

// MemberAddInp 添加会员入参
type MemberAddInp struct {
	Username string  `json:"username" v:"required|length:3,32#请输入用户名|用户名长度3-32位" dc:"用户名"`
	Password string  `json:"password" v:"required|length:6,32#请输入密码|密码长度6-32位" dc:"密码"`
	Nickname string  `json:"nickname" dc:"昵称"`
	Mobile   string  `json:"mobile" v:"phone#手机号格式不正确" dc:"手机号"`
	Email    string  `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Avatar   string  `json:"avatar" dc:"头像"`
	Gender   int     `json:"gender" d:"0" dc:"性别：0=未知 1=男 2=女"`
	GroupId  int64   `json:"groupId" d:"1" dc:"会员分组ID"`
	Score    int     `json:"score" d:"0" dc:"积分"`
	Money    float64 `json:"money" d:"0" dc:"余额"`
	Status   int     `json:"status" d:"1" dc:"状态：0=禁用 1=正常"`
	Remark   string  `json:"remark" dc:"备注"`
}

// MemberAddModel 添加会员出参
type MemberAddModel struct {
	Id int64 `json:"id"`
}

// ===================== 编辑会员 =====================

// MemberEditInp 编辑会员入参
type MemberEditInp struct {
	Id       int64   `json:"id" v:"required|min:1" dc:"会员ID"`
	Username string  `json:"username" v:"length:3,32#用户名长度3-32位" dc:"用户名"`
	Password string  `json:"password" v:"length:6,32#密码长度6-32位" dc:"密码（不修改则留空）"`
	Nickname string  `json:"nickname" dc:"昵称"`
	Mobile   string  `json:"mobile" v:"phone#手机号格式不正确" dc:"手机号"`
	Email    string  `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Avatar   string  `json:"avatar" dc:"头像"`
	Gender   int     `json:"gender" dc:"性别：0=未知 1=男 2=女"`
	GroupId  int64   `json:"groupId" dc:"会员分组ID"`
	Score    int     `json:"score" dc:"积分"`
	Money    float64 `json:"money" dc:"余额"`
	Status   int     `json:"status" dc:"状态：0=禁用 1=正常"`
	Remark   string  `json:"remark" dc:"备注"`
}

// ===================== 删除会员 =====================

// MemberDeleteInp 删除会员入参
type MemberDeleteInp struct {
	Ids []int64 `json:"ids" v:"required|min-length:1" dc:"会员ID列表"`
}

// ===================== 修改状态 =====================

// MemberStatusInp 修改会员状态入参
type MemberStatusInp struct {
	Id     int64 `json:"id" v:"required|min:1" dc:"会员ID"`
	Status int   `json:"status" v:"in:0,1#状态值只能是0或1" dc:"状态：0=禁用 1=正常"`
}

// ===================== 重置密码 =====================

// MemberResetPasswordInp 重置会员密码入参
type MemberResetPasswordInp struct {
	Id       int64  `json:"id" v:"required|min:1" dc:"会员ID"`
	Password string `json:"password" v:"required|length:6,32#请输入新密码|密码长度6-32位" dc:"新密码"`
}

// ===================== 会员分组选项 =====================

// MemberGroupOptionsModel 会员分组选项出参
type MemberGroupOptionsModel struct {
	List []MemberGroupOption `json:"list"`
}

// MemberGroupOption 会员分组选项
type MemberGroupOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
