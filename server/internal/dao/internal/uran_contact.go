// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UranContactDao is the data access object for the table xy_uran_contact.
type UranContactDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UranContactColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UranContactColumns defines and stores column names for the table xy_uran_contact.
type UranContactColumns struct {
	Id          string // 主键
	Username    string // 姓名
	Phone       string // 电话
	Age         string // 数字
	Avatar      string // 头像
	Sort        string // 排序权重
	Remark      string // 备注
	SwitchField string // 开关:0=关,1=开
}

// uranContactColumns holds the columns for the table xy_uran_contact.
var uranContactColumns = UranContactColumns{
	Id:          "id",
	Username:    "username",
	Phone:       "phone",
	Age:         "age",
	Avatar:      "avatar",
	Sort:        "sort",
	Remark:      "remark",
	SwitchField: "switch_field",
}

// NewUranContactDao creates and returns a new DAO object for table data access.
func NewUranContactDao(handlers ...gdb.ModelHandler) *UranContactDao {
	return &UranContactDao{
		group:    "default",
		table:    "xy_uran_contact",
		columns:  uranContactColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UranContactDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UranContactDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UranContactDao) Columns() UranContactColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UranContactDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UranContactDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UranContactDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
