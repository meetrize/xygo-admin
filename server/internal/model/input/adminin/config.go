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

// ===================== 配置列表 =====================

// ConfigListInp 配置列表入参
type ConfigListInp struct {
	Group string `p:"group" json:"group" v:"required#分组必填" dc:"分组标识"`
}

// ConfigListModel 配置列表出参
type ConfigListModel struct {
	Group string            `json:"group" dc:"分组标识"`
	Items map[string]string `json:"items" dc:"配置键值对"`
	List  []ConfigKVItem    `json:"list" dc:"配置列表"`
}

// ===================== 删除配置 =====================

// ConfigDeleteInp 删除配置入参
type ConfigDeleteInp struct {
	Key string `json:"key" v:"required#配置键必填" dc:"配置键"`
}

// ===================== 配置 Schema =====================

// ConfigSchemaItem 返回给前端的配置项 schema
type ConfigSchemaItem struct {
	Id        uint64      `json:"id"`
	Group     string      `json:"group"`
	GroupName string      `json:"groupName"`
	Name      string      `json:"name"`
	Key       string      `json:"key"`
	Value     string      `json:"value"`
	Type      string      `json:"type"`
	Options   interface{} `json:"options"` // JSON 透传
	Rules     interface{} `json:"rules"`   // JSON 透传
	Sort      int         `json:"sort"`
	Remark    string      `json:"remark"`
	AllowDel  int         `json:"allowDel"` // 是否允许删除：0-否，1-是
}

// ConfigKVItem 简单键值
type ConfigKVItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ConfigSaveItem 提交的键值
type ConfigSaveItem struct {
	Key   string `json:"key" v:"required#配置键必填"`
	Value string `json:"value"`
}

// ConfigSaveInp 保存入参
type ConfigSaveInp struct {
	Group string           `p:"group" json:"group" v:"required#分组必填"`
	Items []ConfigSaveItem `p:"items" json:"items" v:"required#配置列表必填"`
}

// ConfigCreateInp 创建配置项入参
type ConfigCreateInp struct {
	Group     string `p:"group" json:"group" v:"required#分组标识必填"`
	GroupName string `p:"groupName" json:"groupName" v:"required#分组名称必填"`
	Name      string `p:"name" json:"name" v:"required#配置项名称必填"`
	Key       string `p:"key" json:"key" v:"required#配置键必填"`
	Type      string `p:"type" json:"type" v:"required#类型必填"`
	Value     string `p:"value" json:"value"`
	Options   string `p:"options" json:"options"` // JSON 字符串
	Rules     string `p:"rules" json:"rules"`     // JSON 字符串
	Sort      int    `p:"sort" json:"sort"`
	Remark    string `p:"remark" json:"remark"`
}
