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

// ===================== 角色列表 =====================

type RoleListReq struct {
	g.Meta `path:"/admin/role/list" method:"get" tags:"AdminRole" summary:"角色列表"`
	adminin.RoleListInp
}

type RoleListRes struct {
	*adminin.RoleListModel
}

// ===================== 角色详情 =====================

type RoleDetailReq struct {
	g.Meta `path:"/admin/role/detail" method:"get" tags:"AdminRole" summary:"角色详情"`
	adminin.RoleDetailInp
}

type RoleDetailRes struct {
	*adminin.RoleDetailModel
}

// ===================== 角色保存 =====================

type RoleSaveReq struct {
	g.Meta `path:"/admin/role/save" method:"post" tags:"AdminRole" summary:"角色保存"`
	adminin.RoleSaveInp
}

type RoleSaveRes struct {
	Id uint64 `json:"id" dc:"角色ID"`
}

// ===================== 角色删除 =====================

type RoleDeleteReq struct {
	g.Meta `path:"/admin/role/delete" method:"post" tags:"AdminRole" summary:"角色删除"`
	adminin.RoleDeleteInp
}

type RoleDeleteRes struct{}

// ===================== 角色菜单 =====================

type RoleMenuIdsReq struct {
	g.Meta `path:"/admin/role/menuIds" method:"get" tags:"AdminRole" summary:"获取角色菜单ID"`
	adminin.RoleMenuIdsInp
}

type RoleMenuIdsRes struct {
	*adminin.RoleMenuIdsModel
}

type RoleBindMenusReq struct {
	g.Meta `path:"/admin/role/bindMenus" method:"post" tags:"AdminRole" summary:"角色绑定菜单"`
	adminin.RoleBindMenusInp
}

type RoleBindMenusRes struct{}

// ===================== 数据范围 =====================

type DataScopeSelectReq struct {
	g.Meta `path:"/admin/role/dataScopeSelect" method:"get" tags:"AdminRole" summary:"数据范围选项"`
}

type DataScopeSelectRes struct {
	List interface{} `json:"list" dc:"数据范围选项列表"`
}

type DataScopeEditReq struct {
	g.Meta `path:"/admin/role/dataScopeEdit" method:"post" tags:"AdminRole" summary:"编辑角色数据权限"`
	adminin.DataScopeEditInp
}

type DataScopeEditRes struct{}

// ===================== 可用资源 =====================

type RoleAvailableResourcesReq struct {
	g.Meta `path:"/admin/role/availableResources" method:"get" tags:"AdminRole" summary:"获取角色可用资源列表"`
	adminin.RoleAvailableResourcesInp
}

type RoleAvailableResourcesRes struct {
	*adminin.RoleAvailableResourcesModel
}
