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

// ===================== 登录 =====================

type LoginReq struct {
	g.Meta `path:"/admin/auth/login" method:"post" tags:"AdminAuth" summary:"Admin login"`
	adminin.LoginInp
}

type LoginRes struct {
	*adminin.LoginModel
}

// ===================== 登出 =====================

type LogoutReq struct {
	g.Meta `path:"/admin/auth/logout" method:"post" tags:"AdminAuth" summary:"Admin logout"`
}

type LogoutRes struct{}

// ===================== 个人信息 =====================

type ProfileReq struct {
	g.Meta `path:"/admin/auth/profile" method:"get" tags:"AdminAuth" summary:"Get current admin profile"`
}

type ProfileRes struct {
	*adminin.ProfileModel
}

// ===================== 更新个人信息 =====================

type UpdateProfileReq struct {
	g.Meta `path:"/admin/auth/updateProfile" method:"post" tags:"AdminAuth" summary:"Update current user profile"`
	adminin.UpdateProfileInp
}

type UpdateProfileRes struct{}

// ===================== 修改密码 =====================

type ChangePasswordReq struct {
	g.Meta `path:"/admin/auth/changePassword" method:"post" tags:"AdminAuth" summary:"Change current user password"`
	adminin.ChangePasswordInp
}

type ChangePasswordRes struct{}
