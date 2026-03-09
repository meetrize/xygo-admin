// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package field

// AdminRoleFields 角色管理字段定义
// 对应表：xy_admin_role
var AdminRoleFields = []ResourceField{
	// ========== 主表字段（与entity.AdminRole完全对应） ==========
	{Name: "id", Label: "角色ID", Type: "number", Source: "admin_role"},
	{Name: "name", Label: "角色名称", Type: "string", Source: "admin_role"},
	{Name: "key", Label: "角色标识", Type: "string", Source: "admin_role"},
	{Name: "data_scope", Label: "数据范围", Type: "number", Source: "admin_role", IsSensitive: true},
	{Name: "custom_depts", Label: "自定义部门", Type: "string", Source: "admin_role", IsSensitive: true},
	{Name: "pid", Label: "上级角色ID", Type: "number", Source: "admin_role"},
	{Name: "level", Label: "关系树等级", Type: "number", Source: "admin_role"},
	{Name: "tree", Label: "关系树路径", Type: "string", Source: "admin_role"},
	{Name: "sort", Label: "排序", Type: "number", Source: "admin_role"},
	{Name: "status", Label: "状态", Type: "number", Source: "admin_role"},
	{Name: "remark", Label: "备注", Type: "string", Source: "admin_role"},
	{Name: "created_by", Label: "创建人ID", Type: "number", Source: "admin_role"},
	{Name: "updated_by", Label: "更新人ID", Type: "number", Source: "admin_role"},
	{Name: "create_time", Label: "创建时间", Type: "datetime", Source: "admin_role"},
	{Name: "update_time", Label: "更新时间", Type: "datetime", Source: "admin_role"},
}

func init() {
	Register("admin_role", "角色管理", "admin", AdminRoleFields)
}
