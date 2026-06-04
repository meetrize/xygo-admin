// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberOauthDao is the data access object for the table xy_member_oauth.
type MemberOauthDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MemberOauthColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MemberOauthColumns defines and stores column names for the table xy_member_oauth.
type MemberOauthColumns struct {
	Id         string // 主键ID
	MemberId   string // 关联会员ID
	Platform   string // 平台标识 wechat_mapp/wechat_oa/qq/alipay等
	Openid     string // 平台openid
	Unionid    string // unionid
	SessionKey string // session_key
	Nickname   string // 平台昵称
	Avatar     string // 平台头像
	Extra      string // 扩展JSON
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // deleted time
}

// memberOauthColumns holds the columns for the table xy_member_oauth.
var memberOauthColumns = MemberOauthColumns{
	Id:         "id",
	MemberId:   "member_id",
	Platform:   "platform",
	Openid:     "openid",
	Unionid:    "unionid",
	SessionKey: "session_key",
	Nickname:   "nickname",
	Avatar:     "avatar",
	Extra:      "extra",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewMemberOauthDao creates and returns a new DAO object for table data access.
func NewMemberOauthDao(handlers ...gdb.ModelHandler) *MemberOauthDao {
	return &MemberOauthDao{
		group:    "default",
		table:    "xy_member_oauth",
		columns:  memberOauthColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MemberOauthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MemberOauthDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MemberOauthDao) Columns() MemberOauthColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MemberOauthDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MemberOauthDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MemberOauthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
