// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UranExampleDao is the data access object for the table xy_uran_example.
type UranExampleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UranExampleColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UranExampleColumns defines and stores column names for the table xy_uran_example.
type UranExampleColumns struct {
	Id        string // 主键
	Status    string // 状态:0=禁用,1=启用
	Sort      string // 排序
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Tname     string // 标题
	Timage    string // 图片
	Ttest     string // 测试
}

// uranExampleColumns holds the columns for the table xy_uran_example.
var uranExampleColumns = UranExampleColumns{
	Id:        "id",
	Status:    "status",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Tname:     "tname",
	Timage:    "timage",
	Ttest:     "ttest",
}

// NewUranExampleDao creates and returns a new DAO object for table data access.
func NewUranExampleDao(handlers ...gdb.ModelHandler) *UranExampleDao {
	return &UranExampleDao{
		group:    "default",
		table:    "xy_uran_example",
		columns:  uranExampleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UranExampleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UranExampleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UranExampleDao) Columns() UranExampleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UranExampleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UranExampleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UranExampleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
