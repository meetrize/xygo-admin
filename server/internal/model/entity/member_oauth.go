// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MemberOauth is the golang structure for table member_oauth.
type MemberOauth struct {
	Id         uint64 `json:"id"         orm:"id"          description:"主键ID"`                                  // 主键ID
	MemberId   uint64 `json:"memberId"   orm:"member_id"   description:"关联会员ID"`                                // 关联会员ID
	Platform   string `json:"platform"   orm:"platform"    description:"平台标识 wechat_mapp/wechat_oa/qq/alipay等"` // 平台标识 wechat_mapp/wechat_oa/qq/alipay等
	Openid     string `json:"openid"     orm:"openid"      description:"平台openid"`                              // 平台openid
	Unionid    string `json:"unionid"    orm:"unionid"     description:"unionid"`                               // unionid
	SessionKey string `json:"sessionKey" orm:"session_key" description:"session_key"`                           // session_key
	Nickname   string `json:"nickname"   orm:"nickname"    description:"平台昵称"`                                  // 平台昵称
	Avatar     string `json:"avatar"     orm:"avatar"      description:"平台头像"`                                  // 平台头像
	Extra      string `json:"extra"      orm:"extra"       description:"扩展JSON"`                                // 扩展JSON
	CreatedAt  uint64 `json:"createdAt"  orm:"created_at"  description:"创建时间"`                                  // 创建时间
	UpdatedAt  uint64 `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`                                  // 更新时间
	DeletedAt  uint64 `json:"deletedAt"  orm:"deleted_at"  description:"deleted time"`                          // deleted time
}
