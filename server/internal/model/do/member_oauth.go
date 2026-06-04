// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MemberOauth is the golang structure of table xy_member_oauth for DAO operations like Where/Data.
type MemberOauth struct {
	g.Meta     `orm:"table:xy_member_oauth, do:true"`
	Id         any // 主键ID
	MemberId   any // 关联会员ID
	Platform   any // 平台标识 wechat_mapp/wechat_oa/qq/alipay等
	Openid     any // 平台openid
	Unionid    any // unionid
	SessionKey any // session_key
	Nickname   any // 平台昵称
	Avatar     any // 平台头像
	Extra      any // 扩展JSON
	CreatedAt  any // 创建时间
	UpdatedAt  any // 更新时间
	DeletedAt  any // deleted time
}
