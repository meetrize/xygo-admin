// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package consts

// ============================================
// 字段权限常量
// ============================================

// 字段权限类型（PermType）
const (
	FieldPermHidden   = 0 // 不可见（字段隐藏，前端不显示）
	FieldPermReadonly = 1 // 只读（字段可见但不可编辑）
	FieldPermEditable = 2 // 可编辑（字段可见且可编辑）
)

// FieldPermNameMap 字段权限名称映射
var FieldPermNameMap = map[int]string{
	FieldPermHidden:   "不可见",
	FieldPermReadonly: "只读",
	FieldPermEditable: "可编辑",
}

// FieldPermDescMap 字段权限描述映射
var FieldPermDescMap = map[int]string{
	FieldPermHidden:   "字段隐藏，前端不显示",
	FieldPermReadonly: "字段可见但不可编辑，表单项禁用",
	FieldPermEditable: "字段可见且可编辑，拥有完全权限",
}

// ============================================
// 字段权限选项（用于前端选择）
// ============================================

// FieldPermOption 字段权限选项
type FieldPermOption struct {
	Label string `json:"label"` // 显示名称
	Value int    `json:"value"` // 权限值
	Desc  string `json:"desc"`  // 描述信息
}

// FieldPermOptions 字段权限选项列表
var FieldPermOptions = []FieldPermOption{
	{
		Label: FieldPermNameMap[FieldPermHidden],
		Value: FieldPermHidden,
		Desc:  FieldPermDescMap[FieldPermHidden],
	},
	{
		Label: FieldPermNameMap[FieldPermReadonly],
		Value: FieldPermReadonly,
		Desc:  FieldPermDescMap[FieldPermReadonly],
	},
	{
		Label: FieldPermNameMap[FieldPermEditable],
		Value: FieldPermEditable,
		Desc:  FieldPermDescMap[FieldPermEditable],
	},
}

// ============================================
// 资源字段定义（预定义常见资源的字段列表）
// ============================================

// ResourceField 资源字段定义
// ⚠️ 已废弃：字段定义已迁移到 internal/field/ 目录
// 保留此定义用于兼容性，新模块请使用 field.Register() 注册
type ResourceField struct {
	FieldName   string `json:"fieldName"`   // 字段名称
	FieldLabel  string `json:"fieldLabel"`  // 字段显示名称
	IsSensitive bool   `json:"isSensitive"` // 是否敏感字段（建议限制）
}

// ResourceFieldsMap 资源字段映射（已废弃）
// ⚠️ 新字段定义请在 internal/field/ 目录中创建对应文件
var ResourceFieldsMap = map[string][]ResourceField{
	"admin_user": {
		{FieldName: "id", FieldLabel: "用户ID", IsSensitive: false},
		{FieldName: "username", FieldLabel: "用户名", IsSensitive: false},
		{FieldName: "nickname", FieldLabel: "昵称", IsSensitive: false},
		{FieldName: "mobile", FieldLabel: "手机号", IsSensitive: true},
		{FieldName: "email", FieldLabel: "邮箱", IsSensitive: true},
		{FieldName: "dept_id", FieldLabel: "部门ID", IsSensitive: false},
		{FieldName: "avatar", FieldLabel: "头像", IsSensitive: false},
		{FieldName: "gender", FieldLabel: "性别", IsSensitive: false},
		{FieldName: "status", FieldLabel: "状态", IsSensitive: false},
		{FieldName: "is_super", FieldLabel: "是否超管", IsSensitive: true},
		{FieldName: "last_login_at", FieldLabel: "最后登录时间", IsSensitive: false},
		{FieldName: "last_login_ip", FieldLabel: "最后登录IP", IsSensitive: true},
		{FieldName: "create_time", FieldLabel: "创建时间", IsSensitive: false},
	},
	"admin_dept": {
		{FieldName: "id", FieldLabel: "部门ID", IsSensitive: false},
		{FieldName: "name", FieldLabel: "部门名称", IsSensitive: false},
		{FieldName: "parent_id", FieldLabel: "上级部门", IsSensitive: false},
		{FieldName: "leader", FieldLabel: "负责人", IsSensitive: false},
		{FieldName: "phone", FieldLabel: "联系电话", IsSensitive: true},
		{FieldName: "email", FieldLabel: "邮箱", IsSensitive: true},
		{FieldName: "status", FieldLabel: "状态", IsSensitive: false},
		{FieldName: "sort", FieldLabel: "排序", IsSensitive: false},
		{FieldName: "remark", FieldLabel: "备注", IsSensitive: false},
	},
	"admin_role": {
		{FieldName: "id", FieldLabel: "角色ID", IsSensitive: false},
		{FieldName: "name", FieldLabel: "角色名称", IsSensitive: false},
		{FieldName: "key", FieldLabel: "角色标识", IsSensitive: false},
		{FieldName: "data_scope", FieldLabel: "数据范围", IsSensitive: true},
		{FieldName: "custom_depts", FieldLabel: "自定义部门", IsSensitive: true},
		{FieldName: "status", FieldLabel: "状态", IsSensitive: false},
		{FieldName: "sort", FieldLabel: "排序", IsSensitive: false},
		{FieldName: "remark", FieldLabel: "备注", IsSensitive: false},
	},
	// 可根据实际业务扩展更多资源...
}

// GetResourceFields 获取指定资源的字段列表
func GetResourceFields(resource string) []ResourceField {
	if fields, ok := ResourceFieldsMap[resource]; ok {
		return fields
	}
	return []ResourceField{}
}

// ============================================
// 字段权限默认配置
// ============================================

// DefaultFieldPerm 字段权限默认值
// 未配置的字段默认权限（可根据业务需求调整）
const DefaultFieldPerm = FieldPermEditable // 默认可编辑

// SensitiveFieldDefaultPerm 敏感字段默认权限
// 对于标记为敏感的字段，如果未配置，使用此默认值
const SensitiveFieldDefaultPerm = FieldPermReadonly // 默认只读
