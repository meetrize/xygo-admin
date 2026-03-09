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

// SysConfigFields 系统配置字段定义
// 对应表：xy_sys_config
var SysConfigFields = []ResourceField{
	// ========== 主表字段（与entity.SysConfig完全对应） ==========
	{Name: "id", Label: "主键", Type: "number", Source: "sys_config"},
	{Name: "group", Label: "分组标识", Type: "string", Source: "sys_config"},
	{Name: "group_name", Label: "分组名称", Type: "string", Source: "sys_config"},
	{Name: "name", Label: "配置项显示名", Type: "string", Source: "sys_config"},
	{Name: "key", Label: "配置键", Type: "string", Source: "sys_config"},
	{Name: "value", Label: "配置值", Type: "string", Source: "sys_config", IsSensitive: true},
	{Name: "type", Label: "控件类型", Type: "string", Source: "sys_config"},
	{Name: "options", Label: "组件参数", Type: "json", Source: "sys_config"},
	{Name: "rules", Label: "校验规则", Type: "json", Source: "sys_config"},
	{Name: "sort", Label: "排序", Type: "number", Source: "sys_config"},
	{Name: "remark", Label: "备注", Type: "string", Source: "sys_config"},
	{Name: "allow_del", Label: "允许删除", Type: "number", Source: "sys_config"},
	{Name: "created_by", Label: "创建人", Type: "number", Source: "sys_config"},
	{Name: "updated_by", Label: "更新人", Type: "number", Source: "sys_config"},
	{Name: "create_time", Label: "创建时间", Type: "datetime", Source: "sys_config"},
	{Name: "update_time", Label: "更新时间", Type: "datetime", Source: "sys_config"},
}

func init() {
	Register("sys_config", "系统配置", "system", SysConfigFields)
}
