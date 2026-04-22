// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MigrationDao is the data access object for the table xy_migration.
type MigrationDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MigrationColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MigrationColumns defines and stores column names for the table xy_migration.
type MigrationColumns struct {
	Id         string //
	Version    string //
	Name       string //
	ExecutedAt string //
	Checksum   string //
	Success    string //
}

// migrationColumns holds the columns for the table xy_migration.
var migrationColumns = MigrationColumns{
	Id:         "id",
	Version:    "version",
	Name:       "name",
	ExecutedAt: "executed_at",
	Checksum:   "checksum",
	Success:    "success",
}

// NewMigrationDao creates and returns a new DAO object for table data access.
func NewMigrationDao(handlers ...gdb.ModelHandler) *MigrationDao {
	return &MigrationDao{
		group:    "default",
		table:    "xy_migration",
		columns:  migrationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MigrationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MigrationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MigrationDao) Columns() MigrationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MigrationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MigrationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MigrationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
