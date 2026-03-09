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

// AdminMenuFields 菜单管理字段定义
// 对应表：xy_admin_menu
var AdminMenuFields = []ResourceField{
	// ========== 主表字段（与entity.AdminMenu完全对应） ==========
	{Name: "id", Label: "菜单ID", Type: "number", Source: "admin_menu"},
	{Name: "parent_id", Label: "上级菜单ID", Type: "number", Source: "admin_menu"},
	{Name: "type", Label: "类型", Type: "number", Source: "admin_menu"},
	{Name: "title", Label: "菜单名称", Type: "string", Source: "admin_menu"},
	{Name: "name", Label: "路由name", Type: "string", Source: "admin_menu"},
	{Name: "path", Label: "路由路径", Type: "string", Source: "admin_menu"},
	{Name: "component", Label: "组件路径", Type: "string", Source: "admin_menu"},
	{Name: "icon", Label: "图标", Type: "string", Source: "admin_menu"},
	{Name: "hidden", Label: "是否隐藏", Type: "number", Source: "admin_menu"},
	{Name: "keep_alive", Label: "是否缓存", Type: "number", Source: "admin_menu"},
	{Name: "redirect", Label: "重定向地址", Type: "string", Source: "admin_menu"},
	{Name: "frame_src", Label: "iframe地址", Type: "string", Source: "admin_menu"},
	{Name: "perms", Label: "权限点列表", Type: "string", Source: "admin_menu"},
	{Name: "is_frame", Label: "是否内嵌", Type: "number", Source: "admin_menu"},
	{Name: "affix", Label: "是否固定标签", Type: "number", Source: "admin_menu"},
	{Name: "show_badge", Label: "是否显示徽章", Type: "number", Source: "admin_menu"},
	{Name: "badge_text", Label: "徽章文本", Type: "string", Source: "admin_menu"},
	{Name: "active_path", Label: "激活高亮路径", Type: "string", Source: "admin_menu"},
	{Name: "hide_tab", Label: "是否隐藏标签", Type: "number", Source: "admin_menu"},
	{Name: "is_full_page", Label: "是否全屏页面", Type: "number", Source: "admin_menu"},
	{Name: "sort", Label: "排序", Type: "number", Source: "admin_menu"},
	{Name: "status", Label: "状态", Type: "number", Source: "admin_menu"},
	{Name: "remark", Label: "备注", Type: "string", Source: "admin_menu"},
	{Name: "created_by", Label: "创建人ID", Type: "number", Source: "admin_menu"},
	{Name: "updated_by", Label: "更新人ID", Type: "number", Source: "admin_menu"},
	{Name: "create_time", Label: "创建时间", Type: "datetime", Source: "admin_menu"},
	{Name: "update_time", Label: "更新时间", Type: "datetime", Source: "admin_menu"},
}

func init() {
	Register("admin_menu", "菜单管理", "admin", AdminMenuFields)
}
