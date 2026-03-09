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

// UserListReq 管理员列表请求（仅封装 HTTP Meta，业务入参使用 adminin.UserListInp）
type UserListReq struct {
	g.Meta `path:"/admin/user/list" method:"get" tags:"AdminUser" summary:"Admin user list"`
	adminin.UserListInp
}

// UserListRes 管理员列表响应（业务出参使用 adminin.UserListModel）
type UserListRes struct {
	adminin.UserListModel
}

// UserDetailReq 用户详情请求（编辑时获取未脱敏数据）
type UserDetailReq struct {
	g.Meta `path:"/admin/user/detail" method:"get" tags:"AdminUser" summary:"Admin user detail"`
	Id     uint64 `json:"id" p:"id" v:"required#请指定用户ID" dc:"用户ID"`
}

// UserDetailRes 用户详情响应
type UserDetailRes struct {
	*adminin.UserDetailModel
}

// UserSaveReq 保存用户请求（新增/编辑）
type UserSaveReq struct {
	g.Meta `path:"/admin/user/save" method:"post" tags:"AdminUser" summary:"Save admin user"`
	adminin.UserSaveInp
}

// UserSaveRes 保存用户响应
type UserSaveRes struct {
	Id uint `json:"id"`
}

// UserDeleteReq 删除用户请求
type UserDeleteReq struct {
	g.Meta `path:"/admin/user/delete" method:"post" tags:"AdminUser" summary:"Delete admin user"`
	adminin.UserDeleteInp
}

// UserDeleteRes 删除用户响应
type UserDeleteRes struct{}

// UserKickReq 强制用户下线请求
type UserKickReq struct {
	g.Meta `path:"/admin/user/kick" method:"post" tags:"AdminUser" summary:"Kick user offline"`
	Id     uint64 `json:"id" v:"required#请指定用户ID" dc:"用户ID"`
}

// UserKickRes 强制用户下线响应
type UserKickRes struct{}
