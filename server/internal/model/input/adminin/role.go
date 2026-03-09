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

// RoleListInp 角色列表查询入参
type RoleListInp struct {
	form.PageReq
	Name   string `p:"name"   json:"name"   dc:"按角色名称模糊搜索"`
	Status int    `p:"status" d:"-1" json:"status" dc:"状态过滤:1启用,0禁用,-1全部"`
}

// RoleListItem 角色列表项
type RoleListItem struct {
	Id          uint            `json:"id"          dc:"角色ID"`
	Name        string          `json:"name"        dc:"角色名称"`
	Key         string          `json:"key"         dc:"角色唯一标识"`
	Pid         uint            `json:"pid"         dc:"上级角色ID"`
	Level       int             `json:"level"       dc:"层级（根为1）"`
	Tree        string          `json:"tree"        dc:"关系树路径，如 0,1,3"`
	DataScope   int             `json:"dataScope"   dc:"数据范围：1-7"`
	CustomDepts string          `json:"customDepts" dc:"自定义部门（JSON）"`
	Sort        int             `json:"sort"        dc:"排序（越小越靠前）"`
	Status      int             `json:"status"      dc:"状态:0禁用,1启用"`
	Remark      string          `json:"remark"      dc:"备注"`
	CreatedAt   uint64          `json:"createdAt"   dc:"创建时间"`
	Children    []*RoleListItem `json:"children,omitempty" dc:"子角色集合"`
}

// RoleListModel 角色列表响应模型
type RoleListModel struct {
	List []RoleListItem `json:"list" dc:"数据列表"`
	form.PageRes
}

// RoleDetailInp 角色详情入参
type RoleDetailInp struct {
	Id uint `p:"id" v:"required#角色ID不能为空" json:"id" dc:"角色ID"`
}

// RoleDetailModel 角色详情模型
type RoleDetailModel struct {
	Id          uint64 `json:"id"          dc:"角色ID"`
	Name        string `json:"name"        dc:"角色名称"`
	Key         string `json:"key"         dc:"角色标识(英文唯一)"`
	Pid         uint   `json:"pid"         dc:"上级角色ID"`
	Level       int    `json:"level"       dc:"层级（根为1）"`
	Tree        string `json:"tree"        dc:"关系树路径，如 0,1,3"`
	Sort        int    `json:"sort"        dc:"排序（越小越靠前）"`
	DataScope   int    `json:"dataScope"   dc:"数据范围:0=全部,1=本部门,2=本部门及子,3=仅本人,4=自定义部门"`
	CustomDepts string `json:"customDepts" dc:"自定义数据范围部门ID列表(JSON数组)"`
	Status      int    `json:"status"      dc:"状态:0=禁用,1=启用"`
	Remark      string `json:"remark"      dc:"备注"`
	CreateTime  int    `json:"create_time"   dc:"创建时间"`
	UpdateTime  int    `json:"update_time"   dc:"更新时间"`
	CreatedBy   uint64 `json:"created_by"   dc:"创建人ID"`
	UpdatedBy   uint64 `json:"updated_by"   dc:"更新人ID"`
}

// RoleSaveInp 角色新增/编辑入参
// - Id 为空/0 表示新增
// - Id > 0 表示编辑
type RoleSaveInp struct {
	Id          uint64 `p:"id"          json:"id"          dc:"角色ID（为空表示新增）"`
	Name        string `p:"name"        v:"required#角色名称不能为空"        json:"name"        dc:"角色名称"`
	Key         string `p:"key"         v:"required#角色标识不能为空"        json:"key"         dc:"角色标识(英文唯一)"`
	Pid         uint   `p:"pid"         d:"0"                                   json:"pid"         dc:"上级角色ID，0表示根角色"`
	Sort        int    `p:"sort"        d:"0"                                   json:"sort"        dc:"排序（越小越靠前）"`
	DataScope   int    `p:"dataScope"   d:"0"                                   json:"dataScope"   dc:"数据范围:0=全部,1=本部门,2=本部门及子,3=仅本人,4=自定义部门"`
	CustomDepts string `p:"customDepts" d:"[]"                                   json:"customDepts" dc:"自定义数据范围部门ID列表(JSON数组)"`
	Status      int    `p:"status"      d:"1"                                    json:"status"      dc:"状态:0=禁用,1=启用"`
	Remark      string `p:"remark"                                             json:"remark"      dc:"备注"`
}

// RoleDeleteInp 角色删除入参
type RoleDeleteInp struct {
	Id uint64 `p:"id" v:"required#角色ID不能为空" json:"id" dc:"角色ID"`
}

// RoleMenuIdsInp 角色已绑定菜单ID列表入参
type RoleMenuIdsInp struct {
	RoleId uint64 `p:"roleId" v:"required#角色ID不能为空" json:"roleId" dc:"角色ID"`
}

// RoleMenuIdsModel 角色已绑定菜单ID列表模型
type RoleMenuIdsModel struct {
	MenuIds []uint64 `json:"menuIds" dc:"已绑定的菜单ID列表"`
}

// RoleBindMenusInp 角色绑定菜单入参
type RoleBindMenusInp struct {
	RoleId  uint64   `p:"roleId"  v:"required#角色ID不能为空" json:"roleId"  dc:"角色ID"`
	MenuIds []uint64 `p:"menuIds" json:"menuIds"                     dc:"菜单ID列表"`
}

// ===================== 数据权限 =====================

// DataScopeEditInp 编辑角色数据权限入参
type DataScopeEditInp struct {
	Id          uint64   `json:"id" v:"required#角色ID不能为空" dc:"角色ID"`
	DataScope   int      `json:"dataScope" v:"required|between:1,7#数据范围不能为空|数据范围值错误" dc:"数据范围：1-7"`
	CustomDepts []uint64 `json:"customDepts" dc:"自定义部门ID列表（dataScope=4时必填）"`
}

// ===================== 可用资源 =====================

// RoleAvailableResourcesInp 获取角色可用资源入参
type RoleAvailableResourcesInp struct {
	RoleId uint64 `json:"roleId" v:"required#角色ID不能为空" dc:"角色ID"`
}

// RoleAvailableResourcesModel 可用资源列表模型
type RoleAvailableResourcesModel struct {
	List []AvailableResource `json:"list" dc:"可用资源列表"`
}

// AvailableResource 可用资源项
type AvailableResource struct {
	Code  string `json:"code" dc:"资源编码（表名）"`
	Label string `json:"label" dc:"资源显示名"`
}
