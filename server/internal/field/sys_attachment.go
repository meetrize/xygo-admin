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

// SysAttachmentFields 附件管理字段定义
// 对应表：xy_sys_attachment
var SysAttachmentFields = []ResourceField{
	// ========== 主表字段（与entity.SysAttachment完全对应） ==========
	{Name: "id", Label: "附件ID", Type: "number", Source: "sys_attachment"},
	{Name: "topic", Label: "分组/主题标识", Type: "string", Source: "sys_attachment"},
	{Name: "user_id", Label: "上传用户ID", Type: "number", Source: "sys_attachment"},
	{Name: "url", Label: "文件路径", Type: "string", Source: "sys_attachment"},
	{Name: "width", Label: "宽度", Type: "number", Source: "sys_attachment"},
	{Name: "height", Label: "高度", Type: "number", Source: "sys_attachment"},
	{Name: "name", Label: "原始名称", Type: "string", Source: "sys_attachment"},
	{Name: "size", Label: "文件大小", Type: "number", Source: "sys_attachment"},
	{Name: "mimetype", Label: "MIME类型", Type: "string", Source: "sys_attachment"},
	{Name: "quote", Label: "引用次数", Type: "number", Source: "sys_attachment"},
	{Name: "storage", Label: "存储方式", Type: "string", Source: "sys_attachment"},
	{Name: "sha1", Label: "SHA1摘要", Type: "string", Source: "sys_attachment"},
	{Name: "create_time", Label: "创建时间", Type: "datetime", Source: "sys_attachment"},
	{Name: "update_time", Label: "更新时间", Type: "datetime", Source: "sys_attachment"},
}

func init() {
	Register("sys_attachment", "附件管理", "system", SysAttachmentFields)
}
