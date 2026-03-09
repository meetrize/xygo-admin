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

// AdminDeptFields 部门管理字段定义
// 对应表：xy_admin_dept
var AdminDeptFields = []ResourceField{
	// ========== 主表字段（与entity.AdminDept完全对应） ==========
	{Name: "id", Label: "部门ID", Type: "number", Source: "admin_dept"},
	{Name: "parent_id", Label: "上级部门ID", Type: "number", Source: "admin_dept"},
	{Name: "name", Label: "部门名称", Type: "string", Source: "admin_dept"},
	{Name: "sort", Label: "排序", Type: "number", Source: "admin_dept"},
	{Name: "status", Label: "状态", Type: "number", Source: "admin_dept"},
	{Name: "remark", Label: "备注", Type: "string", Source: "admin_dept"},
	{Name: "created_by", Label: "创建人ID", Type: "number", Source: "admin_dept"},
	{Name: "updated_by", Label: "更新人ID", Type: "number", Source: "admin_dept"},
	{Name: "create_time", Label: "创建时间", Type: "datetime", Source: "admin_dept"},
	{Name: "update_time", Label: "更新时间", Type: "datetime", Source: "admin_dept"},
}

func init() {
	Register("admin_dept", "部门管理", "admin", AdminDeptFields)
}
