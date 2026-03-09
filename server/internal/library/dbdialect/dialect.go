// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// Package dbdialect 提供数据库方言抽象层，屏蔽 MySQL 与 PostgreSQL 之间的 SQL 语法差异。
// 所有需要方言差异的地方，统一通过本包获取当前方言实例进行操作，避免在业务代码中硬编码特定数据库语法。
package dbdialect

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

// Dialect 数据库方言接口，定义所有需要方言差异适配的方法。
type Dialect interface {
	// DriverName 返回驱动名称，如 "mysql"、"pgsql"
	DriverName() string

	// QuoteIdentifier 用数据库对应的标识符引号包裹名称
	// MySQL: `name`  PG: "name"
	QuoteIdentifier(name string) string

	// GetDbName 获取当前数据库名
	GetDbName(ctx context.Context) (string, error)

	// ListTablesSQL 返回查询所有用户表的 SQL
	// 返回的结果集需要包含 tableName, tableComment 两列
	ListTablesSQL(dbName string) string

	// ListColumnsSQL 返回查询表所有字段的 SQL
	// 返回的结果集需要包含 columnName, columnType, dataType, columnComment, columnKey, isNullable, extra, ordinalPos 列
	ListColumnsSQL(dbName, tableName string) string

	// ListColumnsSQLForSync 返回用于字段同步的字段查询 SQL
	// 返回的结果集需要包含 COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_DEFAULT, COLUMN_COMMENT, COLUMN_KEY, EXTRA
	ListColumnsSQLForSync(dbName, tableName string) string

	// ListColumnsSimpleSQL 返回简单字段列表查询 SQL（用于字段权限等场景）
	// 返回的结果集需要包含 columnName, columnComment 两列
	ListColumnsSimpleSQL(dbName, tableName string) string

	// TableExistsSQL 返回检查指定表是否存在的 SQL
	// 返回的结果集需要包含 tableName 列
	TableExistsSQL(dbName, tableName string) string

	// CreateTableSQL 根据列定义和表注释，生成建表 SQL
	CreateTableSQL(tableName string, colDefs []string, tableComment string) string

	// BuildColumnDef 构建单个列定义（用于 CREATE TABLE 内部）
	BuildColumnDef(col ColumnMeta) string

	// BuildAddColumnSQL 构建 ALTER TABLE ADD COLUMN SQL
	BuildAddColumnSQL(tableName string, col ColumnMeta) string

	// BuildModifyColumnSQL 构建 ALTER TABLE 修改列 SQL
	// MySQL: ALTER TABLE ... MODIFY COLUMN ...
	// PG:    ALTER TABLE ... ALTER COLUMN ... TYPE ... / SET NOT NULL / SET DEFAULT / COMMENT
	BuildModifyColumnSQL(tableName string, col ColumnMeta) string

	// BuildDropColumnSQL 构建 ALTER TABLE DROP COLUMN SQL
	BuildDropColumnSQL(tableName, colName string) string

	// AutoIncrementDef 返回主键自增的列定义片段
	// MySQL: NOT NULL AUTO_INCREMENT
	// PG:    （列类型改为 SERIAL/BIGSERIAL，无需额外关键字）
	AutoIncrementPKDef(colType string) string

	// UnixTimestampNow 返回获取当前 Unix 时间戳的表达式
	// MySQL: UNIX_TIMESTAMP()
	// PG:    extract(epoch from now())::bigint
	UnixTimestampNow() string

	// FromUnixtime 返回将 Unix 时间戳列转为时间的表达式
	// MySQL: FROM_UNIXTIME(col)
	// PG:    to_timestamp(col)
	FromUnixtime(col string) string

	// DateFormat 返回日期格式化表达式
	// MySQL: DATE_FORMAT(expr, '%Y-%m-%d %H')
	// PG:    to_char(expr, 'YYYY-MM-DD HH24')
	DateFormat(expr, mysqlFmt string) string

	// TypeToGoType 将数据库类型转换为 Go 类型
	TypeToGoType(dataType, columnType string) string

	// TypeToTsType 将数据库类型转换为 TypeScript 类型
	TypeToTsType(dataType string) string

	// IsNumericType 判断是否为数值类型
	IsNumericType(colType string) bool

	// GetDefaultForType 根据类型返回合适的 SQL 默认值表达式
	GetDefaultForType(colType string) string

	// NullCoalesce 返回 NULL 合并函数
	// MySQL: IFNULL(expr, def)
	// PG:    COALESCE(expr, def)
	NullCoalesce(expr, def string) string

	// RoundExpr 返回 ROUND 表达式（两个方言一致，但留接口以备扩展）
	RoundExpr(expr string, decimals int) string
}

// ColumnMeta 列元信息，用于 DDL 构建
type ColumnMeta struct {
	Name         string // 列名
	Type         string // 完整类型，如 varchar(100)、bigint unsigned
	IsPk         bool   // 是否主键
	IsNullable   bool   // 是否可空
	DefaultValue string // 默认值
	Comment      string // 字段注释
}

// ---------------------------------------------------------------------------
// 全局方言获取
// ---------------------------------------------------------------------------

// Get 根据当前数据库配置，返回对应的方言实例。
// 调用方只需 dbdialect.Get() 即可。
func Get() Dialect {
	dbType := detectDriver()
	switch dbType {
	case "pgsql":
		return &pgDialect{}
	default:
		return &mysqlDialect{}
	}
}

// MustGet 同 Get，但会在日志中记录一次方言选择（仅首次）。
func MustGet() Dialect {
	return Get()
}

// detectDriver 从 GoFrame 数据库配置中检测当前驱动类型。
func detectDriver() string {
	cfg := g.DB().GetConfig()
	if cfg == nil {
		return "mysql"
	}
	link := strings.ToLower(cfg.Link)
	if strings.HasPrefix(link, "pgsql:") || strings.HasPrefix(link, "postgresql:") {
		return "pgsql"
	}
	typ := strings.ToLower(cfg.Type)
	if typ == "pgsql" || typ == "postgresql" {
		return "pgsql"
	}
	return "mysql"
}

// IsPgsql 快速判断当前是否为 PostgreSQL
func IsPgsql() bool {
	return detectDriver() == "pgsql"
}

// IsMysql 快速判断当前是否为 MySQL
func IsMysql() bool {
	return detectDriver() == "mysql"
}

// escapeSQLString 转义 SQL 字符串中的单引号
func escapeSQLString(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

// mysqlFmtToPgFmt 将 MySQL DATE_FORMAT 格式转换为 PG to_char 格式
func mysqlFmtToPgFmt(mysqlFmt string) string {
	r := strings.NewReplacer(
		"%%Y", "YYYY", "%Y", "YYYY",
		"%%m", "MM", "%m", "MM",
		"%%d", "DD", "%d", "DD",
		"%%H", "HH24", "%H", "HH24",
		"%%i", "MI", "%i", "MI",
		"%%s", "SS", "%s", "SS",
	)
	return r.Replace(mysqlFmt)
}

// isNumericTypeCommon 通用数值类型判断
func isNumericTypeCommon(t string) bool {
	tl := strings.ToLower(t)
	for _, prefix := range []string{
		"int", "bigint", "smallint", "mediumint", "tinyint",
		"float", "double", "decimal", "numeric", "real",
		"serial", "bigserial", "smallserial",
		"integer",
	} {
		if strings.HasPrefix(tl, prefix) {
			return true
		}
	}
	return false
}

// getDefaultForTypeCommon 通用默认值
func getDefaultForTypeCommon(colType string) string {
	if isNumericTypeCommon(colType) {
		return "0"
	}
	return "''"
}

// buildCommentOnColumnSQL 生成 PG 的 COMMENT ON COLUMN 语句
func buildCommentOnColumnSQL(tableName, colName, comment string) string {
	return fmt.Sprintf("COMMENT ON COLUMN %s.%s IS '%s'",
		quoteIdPg(tableName), quoteIdPg(colName), escapeSQLString(comment))
}

// quoteIdPg PG 标识符引号
func quoteIdPg(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}
