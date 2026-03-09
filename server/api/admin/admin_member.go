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

// ===================== 会员列表 =====================

type MemberListReq struct {
	g.Meta `path:"/admin/member/list" method:"get" tags:"AdminMember" summary:"会员列表"`
	adminin.MemberListInp
}

type MemberListRes struct {
	*adminin.MemberListModel
}

// ===================== 会员详情 =====================

type MemberDetailReq struct {
	g.Meta `path:"/admin/member/detail" method:"get" tags:"AdminMember" summary:"会员详情"`
	adminin.MemberDetailInp
}

type MemberDetailRes struct {
	*adminin.MemberDetailModel
}

// ===================== 添加会员 =====================

type MemberAddReq struct {
	g.Meta `path:"/admin/member/add" method:"post" tags:"AdminMember" summary:"添加会员"`
	adminin.MemberAddInp
}

type MemberAddRes struct {
	*adminin.MemberAddModel
}

// ===================== 编辑会员 =====================

type MemberEditReq struct {
	g.Meta `path:"/admin/member/edit" method:"put" tags:"AdminMember" summary:"编辑会员"`
	adminin.MemberEditInp
}

type MemberEditRes struct{}

// ===================== 删除会员 =====================

type MemberDeleteReq struct {
	g.Meta `path:"/admin/member/delete" method:"delete" tags:"AdminMember" summary:"删除会员"`
	adminin.MemberDeleteInp
}

type MemberDeleteRes struct{}

// ===================== 修改状态 =====================

type MemberStatusReq struct {
	g.Meta `path:"/admin/member/status" method:"put" tags:"AdminMember" summary:"修改会员状态"`
	adminin.MemberStatusInp
}

type MemberStatusRes struct{}

// ===================== 重置密码 =====================

type MemberResetPasswordReq struct {
	g.Meta `path:"/admin/member/resetPassword" method:"put" tags:"AdminMember" summary:"重置会员密码"`
	adminin.MemberResetPasswordInp
}

type MemberResetPasswordRes struct{}

// ===================== 会员分组选项 =====================

type MemberGroupOptionsReq struct {
	g.Meta `path:"/admin/member/groupOptions" method:"get" tags:"AdminMember" summary:"会员分组选项"`
}

type MemberGroupOptionsRes struct {
	*adminin.MemberGroupOptionsModel
}

// ===================== 会员分组列表 =====================

type MemberGroupListReq struct {
	g.Meta `path:"/admin/member/group/list" method:"get" tags:"AdminMemberGroup" summary:"会员分组列表"`
	adminin.MemberGroupListInp
}

type MemberGroupListRes struct {
	*adminin.MemberGroupListModel
}

// ===================== 会员分组保存 =====================

type MemberGroupSaveReq struct {
	g.Meta `path:"/admin/member/group/save" method:"post" tags:"AdminMemberGroup" summary:"保存会员分组"`
	adminin.MemberGroupSaveInp
}

type MemberGroupSaveRes struct {
	Id uint `json:"id" dc:"分组ID"`
}

// ===================== 会员分组删除 =====================

type MemberGroupDeleteReq struct {
	g.Meta `path:"/admin/member/group/delete" method:"post" tags:"AdminMemberGroup" summary:"删除会员分组"`
	adminin.MemberGroupDeleteInp
}

type MemberGroupDeleteRes struct{}

// ===================== 会员菜单树 =====================

type MemberMenuTreeReq struct {
	g.Meta `path:"/admin/member/menu/tree" method:"get" tags:"AdminMemberMenu" summary:"会员菜单树"`
	adminin.MemberMenuTreeInp
}

type MemberMenuTreeRes struct {
	*adminin.MemberMenuTreeModel
}

// ===================== 会员菜单保存 =====================

type MemberMenuSaveReq struct {
	g.Meta `path:"/admin/member/menu/save" method:"post" tags:"AdminMemberMenu" summary:"保存会员菜单"`
	adminin.MemberMenuSaveInp
}

type MemberMenuSaveRes struct {
	Id uint `json:"id" dc:"菜单ID"`
}

// ===================== 会员菜单删除 =====================

type MemberMenuDeleteReq struct {
	g.Meta `path:"/admin/member/menu/delete" method:"post" tags:"AdminMemberMenu" summary:"删除会员菜单"`
	adminin.MemberMenuDeleteInp
}

type MemberMenuDeleteRes struct{}
