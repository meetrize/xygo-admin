// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package dbdialect

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

// mysqlDialect MySQL 方言实现
type mysqlDialect struct{}

func (d *mysqlDialect) DriverName() string {
	return "mysql"
}

func (d *mysqlDialect) QuoteIdentifier(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

func (d *mysqlDialect) GetDbName(ctx context.Context) (string, error) {
	val, err := g.DB().GetValue(ctx, "SELECT DATABASE()")
	if err != nil {
		return "", err
	}
	if val.IsEmpty() {
		return "xygonew", nil
	}
	return val.String(), nil
}

func (d *mysqlDialect) ListTablesSQL(dbName string) string {
	return fmt.Sprintf(
		`SELECT TABLE_NAME as tableName, TABLE_COMMENT as tableComment
		 FROM information_schema.TABLES
		 WHERE TABLE_SCHEMA = '%s' AND TABLE_TYPE = 'BASE TABLE'
		 ORDER BY TABLE_NAME ASC`, escapeSQLString(dbName))
}

func (d *mysqlDialect) ListColumnsSQL(dbName, tableName string) string {
	return fmt.Sprintf(
		`SELECT
			COLUMN_NAME as columnName,
			COLUMN_TYPE as columnType,
			DATA_TYPE as dataType,
			COLUMN_COMMENT as columnComment,
			COLUMN_KEY as columnKey,
			IS_NULLABLE as isNullable,
			EXTRA as extra,
			ORDINAL_POSITION as ordinalPos
		 FROM information_schema.COLUMNS
		 WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'
		 ORDER BY ORDINAL_POSITION ASC`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *mysqlDialect) ListColumnsSQLForSync(dbName, tableName string) string {
	return fmt.Sprintf(
		`SELECT 
			COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, 
			IFNULL(COLUMN_DEFAULT, '') as COLUMN_DEFAULT, 
			IFNULL(COLUMN_COMMENT, '') as COLUMN_COMMENT,
			IFNULL(COLUMN_KEY, '') as COLUMN_KEY,
			IFNULL(EXTRA, '') as EXTRA
		FROM information_schema.COLUMNS 
		WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'
		ORDER BY ORDINAL_POSITION`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *mysqlDialect) ListColumnsSimpleSQL(dbName, tableName string) string {
	return fmt.Sprintf(
		`SELECT COLUMN_NAME AS columnName, COLUMN_COMMENT AS columnComment
		 FROM information_schema.COLUMNS
		 WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'
		 ORDER BY ORDINAL_POSITION ASC`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *mysqlDialect) TableExistsSQL(dbName, tableName string) string {
	return fmt.Sprintf(
		`SELECT TABLE_NAME AS tableName
		 FROM information_schema.TABLES
		 WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *mysqlDialect) CreateTableSQL(tableName string, colDefs []string, tableComment string) string {
	q := d.QuoteIdentifier
	return fmt.Sprintf(
		"CREATE TABLE %s (\n%s\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='%s'",
		q(tableName),
		strings.Join(colDefs, ",\n"),
		escapeSQLString(tableComment),
	)
}

func (d *mysqlDialect) BuildColumnDef(col ColumnMeta) string {
	q := d.QuoteIdentifier
	def := fmt.Sprintf("  %s %s", q(col.Name), col.Type)

	if col.IsPk {
		def += " NOT NULL AUTO_INCREMENT"
	} else if !col.IsNullable {
		if col.DefaultValue != "" {
			if d.IsNumericType(col.Type) {
				def += fmt.Sprintf(" NOT NULL DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" NOT NULL DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += fmt.Sprintf(" NOT NULL DEFAULT %s", d.GetDefaultForType(col.Type))
		}
	} else {
		if col.DefaultValue != "" {
			if d.IsNumericType(col.Type) {
				def += fmt.Sprintf(" DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += " DEFAULT NULL"
		}
	}

	if col.Comment != "" {
		def += fmt.Sprintf(" COMMENT '%s'", escapeSQLString(col.Comment))
	}
	return def
}

func (d *mysqlDialect) BuildAddColumnSQL(tableName string, col ColumnMeta) string {
	q := d.QuoteIdentifier
	def := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", q(tableName), q(col.Name), col.Type)

	if !col.IsNullable {
		if col.DefaultValue != "" {
			if d.IsNumericType(col.Type) {
				def += fmt.Sprintf(" NOT NULL DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" NOT NULL DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += fmt.Sprintf(" NOT NULL DEFAULT %s", d.GetDefaultForType(col.Type))
		}
	} else {
		if col.DefaultValue != "" {
			if d.IsNumericType(col.Type) {
				def += fmt.Sprintf(" DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += " DEFAULT NULL"
		}
	}

	if col.Comment != "" {
		def += fmt.Sprintf(" COMMENT '%s'", escapeSQLString(col.Comment))
	}
	return def
}

func (d *mysqlDialect) BuildModifyColumnSQL(tableName string, col ColumnMeta) string {
	q := d.QuoteIdentifier
	def := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s", q(tableName), q(col.Name), col.Type)

	if col.IsPk {
		def += " NOT NULL AUTO_INCREMENT"
	} else if !col.IsNullable {
		if col.DefaultValue != "" {
			if d.IsNumericType(col.Type) {
				def += fmt.Sprintf(" NOT NULL DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" NOT NULL DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += fmt.Sprintf(" NOT NULL DEFAULT %s", d.GetDefaultForType(col.Type))
		}
	} else {
		if col.DefaultValue != "" {
			if d.IsNumericType(col.Type) {
				def += fmt.Sprintf(" DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += " DEFAULT NULL"
		}
	}

	if col.Comment != "" {
		def += fmt.Sprintf(" COMMENT '%s'", escapeSQLString(col.Comment))
	}
	return def
}

func (d *mysqlDialect) BuildDropColumnSQL(tableName, colName string) string {
	q := d.QuoteIdentifier
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", q(tableName), q(colName))
}

func (d *mysqlDialect) AutoIncrementPKDef(colType string) string {
	return "NOT NULL AUTO_INCREMENT"
}

func (d *mysqlDialect) UnixTimestampNow() string {
	return "UNIX_TIMESTAMP()"
}

func (d *mysqlDialect) FromUnixtime(col string) string {
	return fmt.Sprintf("FROM_UNIXTIME(%s)", col)
}

func (d *mysqlDialect) DateFormat(expr, mysqlFmt string) string {
	return fmt.Sprintf("DATE_FORMAT(%s, '%s')", expr, mysqlFmt)
}

// TypeToGoType MySQL 类型 -> Go 类型
func (d *mysqlDialect) TypeToGoType(dataType, columnType string) string {
	switch strings.ToLower(dataType) {
	case "tinyint":
		if strings.Contains(columnType, "unsigned") {
			return "uint"
		}
		return "int"
	case "smallint":
		if strings.Contains(columnType, "unsigned") {
			return "uint"
		}
		return "int"
	case "mediumint", "int", "integer":
		if strings.Contains(columnType, "unsigned") {
			return "uint"
		}
		return "int"
	case "bigint":
		if strings.Contains(columnType, "unsigned") {
			return "uint64"
		}
		return "int64"
	case "float":
		return "float32"
	case "double", "decimal":
		return "float64"
	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext", "enum", "set":
		return "string"
	case "date", "datetime", "timestamp":
		return "*gtime.Time"
	case "time":
		return "string"
	case "json":
		return "*gjson.Json"
	case "blob", "tinyblob", "mediumblob", "longblob", "binary", "varbinary":
		return "[]byte"
	default:
		return "string"
	}
}

// TypeToTsType MySQL 类型 -> TypeScript 类型
func (d *mysqlDialect) TypeToTsType(dataType string) string {
	switch strings.ToLower(dataType) {
	case "tinyint", "smallint", "mediumint", "int", "integer", "bigint", "float", "double", "decimal":
		return "number"
	case "json":
		return "any"
	default:
		return "string"
	}
}

func (d *mysqlDialect) IsNumericType(colType string) bool {
	return isNumericTypeCommon(colType)
}

func (d *mysqlDialect) GetDefaultForType(colType string) string {
	return getDefaultForTypeCommon(colType)
}

func (d *mysqlDialect) NullCoalesce(expr, def string) string {
	return fmt.Sprintf("IFNULL(%s, %s)", expr, def)
}

func (d *mysqlDialect) RoundExpr(expr string, decimals int) string {
	return fmt.Sprintf("ROUND(%s, %d)", expr, decimals)
}
