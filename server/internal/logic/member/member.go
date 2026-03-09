// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// =================================================================================
// 会员模块 Logic 层注册
// =================================================================================

package member

import (
	"xygo/internal/service"
)

func init() {
	service.RegisterMemberAuth(NewMemberAuth())
	service.RegisterMemberUser(NewMemberUser())
	service.RegisterAdminMemberGroup(NewAdminMemberGroup())
	service.RegisterAdminMemberMenu(NewAdminMemberMenu())
	service.RegisterMemberCheckin(NewMemberCheckin())
	service.RegisterFrontendNotice(NewFrontendNotice())
}
