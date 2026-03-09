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

// AdminUserFields 用户管理字段定义
// 对应表：xy_admin_user
var AdminUserFields = []ResourceField{
	// ========== 主表字段（与entity.AdminUser完全对应） ==========
	{Name: "id", Label: "用户ID", Type: "number", Source: "admin_user"},
	{Name: "username", Label: "登录账号", Type: "string", Source: "admin_user"},
	{Name: "nickname", Label: "昵称", Type: "string", Source: "admin_user"},
	{Name: "password", Label: "密码哈希", Type: "string", Source: "admin_user", IsSensitive: true},
	{Name: "gender", Label: "性别", Type: "number", Source: "admin_user"},
	{Name: "salt", Label: "密码盐", Type: "string", Source: "admin_user", IsSensitive: true},
	{Name: "mobile", Label: "手机号", Type: "string", Source: "admin_user", IsSensitive: true},
	{Name: "email", Label: "邮箱", Type: "string", Source: "admin_user", IsSensitive: true},
	{Name: "avatar", Label: "头像", Type: "string", Source: "admin_user"},
	{Name: "dept_id", Label: "部门ID", Type: "number", Source: "admin_user"},
	{Name: "pid", Label: "上级用户ID", Type: "number", Source: "admin_user"},
	{Name: "is_super", Label: "是否超管", Type: "number", Source: "admin_user", IsSensitive: true},
	{Name: "status", Label: "状态", Type: "number", Source: "admin_user"},
	{Name: "last_login_at", Label: "最后登录时间", Type: "datetime", Source: "admin_user"},
	{Name: "last_login_ip", Label: "最后登录IP", Type: "string", Source: "admin_user", IsSensitive: true},
	{Name: "created_by", Label: "创建人ID", Type: "number", Source: "admin_user"},
	{Name: "updated_by", Label: "更新人ID", Type: "number", Source: "admin_user"},
	{Name: "create_time", Label: "创建时间", Type: "datetime", Source: "admin_user"},
	{Name: "update_time", Label: "更新时间", Type: "datetime", Source: "admin_user"},

	// ========== 关联字段（虚拟，来自JOIN） ==========
	{Name: "dept_name", Label: "部门名称", Type: "string", Source: "admin_dept", IsVirtual: true},
	{Name: "role_names", Label: "角色列表", Type: "array", Source: "admin_role", IsVirtual: true},
	{Name: "post_names", Label: "岗位列表", Type: "array", Source: "admin_post", IsVirtual: true},
}

func init() {
	Register("admin_user", "用户管理", "admin", AdminUserFields)
}
