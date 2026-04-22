// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserTablegenDao is the data access object for the table xy_user_tablegen.
type UserTablegenDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  UserTablegenColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// UserTablegenColumns defines and stores column names for the table xy_user_tablegen.
type UserTablegenColumns struct {
	Id            string // 主键
	Inputtext     string // 测试文本
	Inputlongtext string // 长文本
	Feditor       string // 富文本
	Inputjson     string // json
	Inputdatetime string // 日期时间
	Inputenum     string // 状态
	Inputset      string // 集合
	Status        string // 状态:0=禁用,1=启用
	Sort          string // 排序
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// userTablegenColumns holds the columns for the table xy_user_tablegen.
var userTablegenColumns = UserTablegenColumns{
	Id:            "id",
	Inputtext:     "inputtext",
	Inputlongtext: "inputlongtext",
	Feditor:       "feditor",
	Inputjson:     "inputjson",
	Inputdatetime: "inputdatetime",
	Inputenum:     "inputenum",
	Inputset:      "inputset",
	Status:        "status",
	Sort:          "sort",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewUserTablegenDao creates and returns a new DAO object for table data access.
func NewUserTablegenDao(handlers ...gdb.ModelHandler) *UserTablegenDao {
	return &UserTablegenDao{
		group:    "default",
		table:    "xy_user_tablegen",
		columns:  userTablegenColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserTablegenDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserTablegenDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserTablegenDao) Columns() UserTablegenColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserTablegenDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserTablegenDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserTablegenDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
