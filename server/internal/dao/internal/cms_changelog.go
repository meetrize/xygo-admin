// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsChangelogDao is the data access object for the table xy_cms_changelog.
type CmsChangelogDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  CmsChangelogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// CmsChangelogColumns defines and stores column names for the table xy_cms_changelog.
type CmsChangelogColumns struct {
	Id             string // 日志ID
	Version        string // 版本号(如v3.0.1)
	Title          string // 更新标题
	Content        string // 更新内容(JSON数组)
	Remark         string // 备注
	RequireReLogin string // 是否需要重新登录:0=否,1=是
	ReleaseDate    string // 发布日期(如2025-11-15)
	Status         string // 状态:1=已发布,2=草稿
	Sort           string // 排序(越大越靠前)
	CreatedBy      string // 创建人ID
	UpdatedBy      string // 更新人ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间(软删除)
}

// cmsChangelogColumns holds the columns for the table xy_cms_changelog.
var cmsChangelogColumns = CmsChangelogColumns{
	Id:             "id",
	Version:        "version",
	Title:          "title",
	Content:        "content",
	Remark:         "remark",
	RequireReLogin: "require_re_login",
	ReleaseDate:    "release_date",
	Status:         "status",
	Sort:           "sort",
	CreatedBy:      "created_by",
	UpdatedBy:      "updated_by",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewCmsChangelogDao creates and returns a new DAO object for table data access.
func NewCmsChangelogDao(handlers ...gdb.ModelHandler) *CmsChangelogDao {
	return &CmsChangelogDao{
		group:    "default",
		table:    "xy_cms_changelog",
		columns:  cmsChangelogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CmsChangelogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CmsChangelogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CmsChangelogDao) Columns() CmsChangelogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CmsChangelogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CmsChangelogDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CmsChangelogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
