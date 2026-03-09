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

// AdminPostFields 岗位管理字段定义
// 对应表：xy_admin_post
var AdminPostFields = []ResourceField{
	// ========== 主表字段（与entity.AdminPost完全对应） ==========
	{Name: "id", Label: "岗位ID", Type: "number", Source: "admin_post"},
	{Name: "code", Label: "岗位编码", Type: "string", Source: "admin_post"},
	{Name: "name", Label: "岗位名称", Type: "string", Source: "admin_post"},
	{Name: "sort", Label: "排序", Type: "number", Source: "admin_post"},
	{Name: "status", Label: "状态", Type: "number", Source: "admin_post"},
	{Name: "remark", Label: "备注", Type: "string", Source: "admin_post"},
	{Name: "created_by", Label: "创建人ID", Type: "number", Source: "admin_post"},
	{Name: "updated_by", Label: "更新人ID", Type: "number", Source: "admin_post"},
	{Name: "create_time", Label: "创建时间", Type: "datetime", Source: "admin_post"},
	{Name: "update_time", Label: "更新时间", Type: "datetime", Source: "admin_post"},
}

func init() {
	Register("admin_post", "岗位管理", "admin", AdminPostFields)
}
