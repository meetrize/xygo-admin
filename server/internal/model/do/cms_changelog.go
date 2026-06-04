// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CmsChangelog is the golang structure of table xy_cms_changelog for DAO operations like Where/Data.
type CmsChangelog struct {
	g.Meta         `orm:"table:xy_cms_changelog, do:true"`
	Id             any // 日志ID
	Version        any // 版本号(如v3.0.1)
	Title          any // 更新标题
	Content        any // 更新内容(JSON数组)
	Remark         any // 备注
	RequireReLogin any // 是否需要重新登录:0=否,1=是
	ReleaseDate    any // 发布日期(如2025-11-15)
	Status         any // 状态:1=已发布,2=草稿
	Sort           any // 排序(越大越靠前)
	CreatedBy      any // 创建人ID
	UpdatedBy      any // 更新人ID
	CreatedAt      any // 创建时间
	UpdatedAt      any // 更新时间
	DeletedAt      any // 删除时间(软删除)
}
