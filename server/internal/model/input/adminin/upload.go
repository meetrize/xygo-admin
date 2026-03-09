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

// ===================== 文件上传 =====================

// UploadFileInp 文件上传入参
type UploadFileInp struct {
	Drive string `p:"drive" json:"drive" dc:"可选，覆盖配置中的默认驱动"`
	Topic string `p:"topic" json:"topic" dc:"可选，附件分组/主题标识"`
	// 文件字段 name= file，由框架从 multipart 读取
}

// UploadFileModel 上传结果
type UploadFileModel struct {
	URL          string `json:"url" dc:"可访问地址（相对或绝对）"`
	Path         string `json:"path" dc:"存储路径"`
	Size         int64  `json:"size" dc:"字节"`
	Mime         string `json:"mime" dc:"MIME"`
	Ext          string `json:"ext" dc:"扩展名（含 .）"`
	Drive        string `json:"drive" dc:"实际使用的驱动"`
	AttachmentId uint64 `json:"attachmentId" dc:"附件记录ID"`
}
