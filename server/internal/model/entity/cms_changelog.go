// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CmsChangelog is the golang structure for table cms_changelog.
type CmsChangelog struct {
	Id             uint64 `json:"id"             orm:"id"               description:"日志ID"`              // 日志ID
	Version        string `json:"version"        orm:"version"          description:"版本号(如v3.0.1)"`      // 版本号(如v3.0.1)
	Title          string `json:"title"          orm:"title"            description:"更新标题"`              // 更新标题
	Content        string `json:"content"        orm:"content"          description:"更新内容(JSON数组)"`      // 更新内容(JSON数组)
	Remark         string `json:"remark"         orm:"remark"           description:"备注"`                // 备注
	RequireReLogin int    `json:"requireReLogin" orm:"require_re_login" description:"是否需要重新登录:0=否,1=是"`  // 是否需要重新登录:0=否,1=是
	ReleaseDate    string `json:"releaseDate"    orm:"release_date"     description:"发布日期(如2025-11-15)"` // 发布日期(如2025-11-15)
	Status         int    `json:"status"         orm:"status"           description:"状态:1=已发布,2=草稿"`     // 状态:1=已发布,2=草稿
	Sort           int    `json:"sort"           orm:"sort"             description:"排序(越大越靠前)"`         // 排序(越大越靠前)
	CreatedBy      uint64 `json:"createdBy"      orm:"created_by"       description:"创建人ID"`             // 创建人ID
	UpdatedBy      uint64 `json:"updatedBy"      orm:"updated_by"       description:"更新人ID"`             // 更新人ID
	CreatedAt      uint64 `json:"createdAt"      orm:"created_at"       description:"创建时间"`              // 创建时间
	UpdatedAt      uint64 `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`              // 更新时间
	DeletedAt      uint64 `json:"deletedAt"      orm:"deleted_at"       description:"删除时间(软删除)"`         // 删除时间(软删除)
}
