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

// pgDialect PostgreSQL 方言实现
type pgDialect struct{}

func (d *pgDialect) DriverName() string {
	return "pgsql"
}

func (d *pgDialect) QuoteIdentifier(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}

func (d *pgDialect) GetDbName(ctx context.Context) (string, error) {
	val, err := g.DB().GetValue(ctx, "SELECT current_database()")
	if err != nil {
		return "", err
	}
	if val.IsEmpty() {
		return "xygonew", nil
	}
	return val.String(), nil
}

func (d *pgDialect) ListTablesSQL(dbName string) string {
	// PG 使用 table_schema = 'public'（或可配置）来查用户表
	return fmt.Sprintf(
		`SELECT table_name AS "tableName",
		        COALESCE(obj_description((quote_ident(table_schema)||'.'||quote_ident(table_name))::regclass, 'pg_class'), '') AS "tableComment"
		 FROM information_schema.tables
		 WHERE table_catalog = '%s'
		   AND table_schema = 'public'
		   AND table_type = 'BASE TABLE'
		 ORDER BY table_name ASC`, escapeSQLString(dbName))
}

func (d *pgDialect) ListColumnsSQL(dbName, tableName string) string {
	// PG 的 information_schema.columns 字段名及含义与 MySQL 略有不同
	// data_type 是标准类型名，udt_name 是底层类型名（如 int4, varchar）
	// PG 没有 COLUMN_TYPE（如 varchar(100)），需要拼接
	// PG 没有 COLUMN_KEY（需从 pg_constraint 获取主键信息）
	// PG 没有 EXTRA（自增通过 column_default 中的 nextval 判断）
	return fmt.Sprintf(
		`SELECT
			c.column_name AS "columnName",
			CASE
				WHEN c.character_maximum_length IS NOT NULL THEN c.udt_name || '(' || c.character_maximum_length || ')'
				WHEN c.numeric_precision IS NOT NULL AND c.udt_name NOT IN ('int2','int4','int8','float4','float8')
					THEN c.udt_name || '(' || c.numeric_precision || ',' || COALESCE(c.numeric_scale, 0) || ')'
				ELSE c.udt_name
			END AS "columnType",
			c.udt_name AS "dataType",
			COALESCE(pgd.description, '') AS "columnComment",
			CASE WHEN pk.column_name IS NOT NULL THEN 'PRI' ELSE '' END AS "columnKey",
			c.is_nullable AS "isNullable",
			CASE WHEN c.column_default LIKE 'nextval%%' THEN 'auto_increment' ELSE '' END AS "extra",
			c.ordinal_position AS "ordinalPos"
		 FROM information_schema.columns c
		 LEFT JOIN pg_catalog.pg_statio_all_tables st
			ON st.schemaname = c.table_schema AND st.relname = c.table_name
		 LEFT JOIN pg_catalog.pg_description pgd
			ON pgd.objoid = st.relid AND pgd.objsubid = c.ordinal_position
		 LEFT JOIN (
			SELECT kcu.column_name, kcu.table_name, kcu.table_schema
			FROM information_schema.table_constraints tc
			JOIN information_schema.key_column_usage kcu
				ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema
			WHERE tc.constraint_type = 'PRIMARY KEY'
		 ) pk ON pk.column_name = c.column_name AND pk.table_name = c.table_name AND pk.table_schema = c.table_schema
		 WHERE c.table_catalog = '%s'
		   AND c.table_schema = 'public'
		   AND c.table_name = '%s'
		 ORDER BY c.ordinal_position ASC`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *pgDialect) ListColumnsSQLForSync(dbName, tableName string) string {
	// 返回与 MySQL 对齐的列名，供字段同步使用
	return fmt.Sprintf(
		`SELECT
			c.column_name AS "COLUMN_NAME",
			CASE
				WHEN c.character_maximum_length IS NOT NULL THEN c.udt_name || '(' || c.character_maximum_length || ')'
				WHEN c.numeric_precision IS NOT NULL AND c.udt_name NOT IN ('int2','int4','int8','float4','float8')
					THEN c.udt_name || '(' || c.numeric_precision || ',' || COALESCE(c.numeric_scale, 0) || ')'
				ELSE c.udt_name
			END AS "COLUMN_TYPE",
			c.is_nullable AS "IS_NULLABLE",
			COALESCE(c.column_default, '') AS "COLUMN_DEFAULT",
			COALESCE(pgd.description, '') AS "COLUMN_COMMENT",
			CASE WHEN pk.column_name IS NOT NULL THEN 'PRI' ELSE '' END AS "COLUMN_KEY",
			CASE WHEN c.column_default LIKE 'nextval%%' THEN 'auto_increment' ELSE '' END AS "EXTRA"
		 FROM information_schema.columns c
		 LEFT JOIN pg_catalog.pg_statio_all_tables st
			ON st.schemaname = c.table_schema AND st.relname = c.table_name
		 LEFT JOIN pg_catalog.pg_description pgd
			ON pgd.objoid = st.relid AND pgd.objsubid = c.ordinal_position
		 LEFT JOIN (
			SELECT kcu.column_name, kcu.table_name, kcu.table_schema
			FROM information_schema.table_constraints tc
			JOIN information_schema.key_column_usage kcu
				ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema
			WHERE tc.constraint_type = 'PRIMARY KEY'
		 ) pk ON pk.column_name = c.column_name AND pk.table_name = c.table_name AND pk.table_schema = c.table_schema
		 WHERE c.table_catalog = '%s'
		   AND c.table_schema = 'public'
		   AND c.table_name = '%s'
		 ORDER BY c.ordinal_position`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *pgDialect) ListColumnsSimpleSQL(dbName, tableName string) string {
	return fmt.Sprintf(
		`SELECT
			c.column_name AS "columnName",
			COALESCE(pgd.description, '') AS "columnComment"
		 FROM information_schema.columns c
		 LEFT JOIN pg_catalog.pg_statio_all_tables st
			ON st.schemaname = c.table_schema AND st.relname = c.table_name
		 LEFT JOIN pg_catalog.pg_description pgd
			ON pgd.objoid = st.relid AND pgd.objsubid = c.ordinal_position
		 WHERE c.table_catalog = '%s'
		   AND c.table_schema = 'public'
		   AND c.table_name = '%s'
		 ORDER BY c.ordinal_position ASC`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *pgDialect) TableExistsSQL(dbName, tableName string) string {
	return fmt.Sprintf(
		`SELECT table_name AS "tableName"
		 FROM information_schema.tables
		 WHERE table_catalog = '%s'
		   AND table_schema = 'public'
		   AND table_name = '%s'`, escapeSQLString(dbName), escapeSQLString(tableName))
}

func (d *pgDialect) CreateTableSQL(tableName string, colDefs []string, tableComment string) string {
	q := d.QuoteIdentifier
	sql := fmt.Sprintf("CREATE TABLE %s (\n%s\n)", q(tableName), strings.Join(colDefs, ",\n"))
	// PG 的表注释需要单独的 COMMENT ON TABLE 语句，这里拼在建表后用分号分隔
	if tableComment != "" {
		sql += fmt.Sprintf(";\nCOMMENT ON TABLE %s IS '%s'", q(tableName), escapeSQLString(tableComment))
	}
	return sql
}

func (d *pgDialect) BuildColumnDef(col ColumnMeta) string {
	q := d.QuoteIdentifier
	var def string

	if col.IsPk {
		// PG 主键自增用 SERIAL 或 BIGSERIAL
		serialType := mysqlTypeToPgSerial(col.Type)
		def = fmt.Sprintf("  %s %s NOT NULL", q(col.Name), serialType)
	} else {
		pgType := mysqlTypeToPgType(col.Type)
		def = fmt.Sprintf("  %s %s", q(col.Name), pgType)

		if !col.IsNullable {
			if col.DefaultValue != "" {
				if d.IsNumericType(pgType) {
					def += fmt.Sprintf(" NOT NULL DEFAULT %s", col.DefaultValue)
				} else {
					def += fmt.Sprintf(" NOT NULL DEFAULT '%s'", escapeSQLString(col.DefaultValue))
				}
			} else {
				def += fmt.Sprintf(" NOT NULL DEFAULT %s", d.GetDefaultForType(pgType))
			}
		} else {
			if col.DefaultValue != "" {
				if d.IsNumericType(pgType) {
					def += fmt.Sprintf(" DEFAULT %s", col.DefaultValue)
				} else {
					def += fmt.Sprintf(" DEFAULT '%s'", escapeSQLString(col.DefaultValue))
				}
			} else {
				def += " DEFAULT NULL"
			}
		}
	}
	// PG 列注释不能 inline，需要单独的 COMMENT ON COLUMN，由调用方处理
	return def
}

func (d *pgDialect) BuildAddColumnSQL(tableName string, col ColumnMeta) string {
	q := d.QuoteIdentifier
	pgType := mysqlTypeToPgType(col.Type)
	def := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", q(tableName), q(col.Name), pgType)

	if !col.IsNullable {
		if col.DefaultValue != "" {
			if d.IsNumericType(pgType) {
				def += fmt.Sprintf(" NOT NULL DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" NOT NULL DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += fmt.Sprintf(" NOT NULL DEFAULT %s", d.GetDefaultForType(pgType))
		}
	} else {
		if col.DefaultValue != "" {
			if d.IsNumericType(pgType) {
				def += fmt.Sprintf(" DEFAULT %s", col.DefaultValue)
			} else {
				def += fmt.Sprintf(" DEFAULT '%s'", escapeSQLString(col.DefaultValue))
			}
		} else {
			def += " DEFAULT NULL"
		}
	}

	// 如果有注释，追加 COMMENT ON COLUMN（用分号分隔）
	if col.Comment != "" {
		def += "; " + buildCommentOnColumnSQL(tableName, col.Name, col.Comment)
	}
	return def
}

func (d *pgDialect) BuildModifyColumnSQL(tableName string, col ColumnMeta) string {
	q := d.QuoteIdentifier
	pgType := mysqlTypeToPgType(col.Type)

	// PG 修改列需要拆成多条 ALTER COLUMN 子句
	var parts []string

	// 改类型
	parts = append(parts,
		fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s USING %s::%s",
			q(tableName), q(col.Name), pgType, q(col.Name), pgType))

	// 改 NOT NULL
	if col.IsPk || !col.IsNullable {
		parts = append(parts,
			fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s SET NOT NULL",
				q(tableName), q(col.Name)))
	} else {
		parts = append(parts,
			fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s DROP NOT NULL",
				q(tableName), q(col.Name)))
	}

	// 改默认值
	if col.DefaultValue != "" {
		if d.IsNumericType(pgType) {
			parts = append(parts,
				fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s SET DEFAULT %s",
					q(tableName), q(col.Name), col.DefaultValue))
		} else {
			parts = append(parts,
				fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s SET DEFAULT '%s'",
					q(tableName), q(col.Name), escapeSQLString(col.DefaultValue)))
		}
	}

	// 改注释
	if col.Comment != "" {
		parts = append(parts, buildCommentOnColumnSQL(tableName, col.Name, col.Comment))
	}

	return strings.Join(parts, ";\n")
}

func (d *pgDialect) BuildDropColumnSQL(tableName, colName string) string {
	q := d.QuoteIdentifier
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", q(tableName), q(colName))
}

func (d *pgDialect) AutoIncrementPKDef(colType string) string {
	// PG 使用 SERIAL/BIGSERIAL，无需此字段
	return "NOT NULL"
}

func (d *pgDialect) UnixTimestampNow() string {
	return "extract(epoch from now())::bigint"
}

func (d *pgDialect) FromUnixtime(col string) string {
	return fmt.Sprintf("to_timestamp(%s)", col)
}

func (d *pgDialect) DateFormat(expr, mysqlFmt string) string {
	pgFmt := mysqlFmtToPgFmt(mysqlFmt)
	return fmt.Sprintf("to_char(%s, '%s')", expr, pgFmt)
}

// TypeToGoType PG 类型 -> Go 类型
func (d *pgDialect) TypeToGoType(dataType, columnType string) string {
	switch strings.ToLower(dataType) {
	// 整型
	case "int2", "smallint", "smallserial":
		return "int"
	case "int4", "integer", "int", "serial":
		return "int"
	case "int8", "bigint", "bigserial":
		return "int64"
	// 浮点
	case "float4", "real":
		return "float32"
	case "float8", "double precision", "double":
		return "float64"
	case "numeric", "decimal":
		return "float64"
	// 字符串
	case "varchar", "character varying", "char", "character", "text", "citext", "name":
		return "string"
	// 时间
	case "timestamp", "timestamptz", "timestamp without time zone", "timestamp with time zone":
		return "*gtime.Time"
	case "date":
		return "*gtime.Time"
	case "time", "timetz", "time without time zone", "time with time zone":
		return "string"
	// JSON
	case "json", "jsonb":
		return "*gjson.Json"
	// 布尔
	case "bool", "boolean":
		return "int" // 保持与 MySQL tinyint(1) 一致的映射
	// 二进制
	case "bytea":
		return "[]byte"
	// UUID
	case "uuid":
		return "string"
	// 数组
	case "array":
		return "string"
	default:
		return "string"
	}
}

// TypeToTsType PG 类型 -> TypeScript 类型
func (d *pgDialect) TypeToTsType(dataType string) string {
	switch strings.ToLower(dataType) {
	case "int2", "int4", "int8", "smallint", "integer", "bigint",
		"float4", "float8", "real", "double precision", "numeric", "decimal",
		"serial", "bigserial", "smallserial",
		"int", "double":
		return "number"
	case "bool", "boolean":
		return "number" // 前端按数字处理
	case "json", "jsonb":
		return "any"
	default:
		return "string"
	}
}

func (d *pgDialect) IsNumericType(colType string) bool {
	return isNumericTypeCommon(colType)
}

func (d *pgDialect) GetDefaultForType(colType string) string {
	return getDefaultForTypeCommon(colType)
}

func (d *pgDialect) NullCoalesce(expr, def string) string {
	return fmt.Sprintf("COALESCE(%s, %s)", expr, def)
}

func (d *pgDialect) RoundExpr(expr string, decimals int) string {
	// PG 的 ROUND 参数类型必须是 numeric，对 float 列需要 cast
	return fmt.Sprintf("ROUND((%s)::numeric, %d)", expr, decimals)
}

// ---------------------------------------------------------------------------
// MySQL 类型 -> PG 类型 转换辅助
// ---------------------------------------------------------------------------

// mysqlTypeToPgType 将 MySQL 列类型映射为 PostgreSQL 对应类型
func mysqlTypeToPgType(mysqlType string) string {
	t := strings.ToLower(strings.TrimSpace(mysqlType))
	base := strings.Split(t, "(")[0]
	base = strings.TrimSpace(base)
	// 去掉 unsigned
	base = strings.TrimSuffix(base, " unsigned")

	switch base {
	case "tinyint":
		return "smallint"
	case "smallint":
		return "smallint"
	case "mediumint":
		return "integer"
	case "int", "integer":
		return "integer"
	case "bigint":
		return "bigint"
	case "float":
		return "real"
	case "double":
		return "double precision"
	case "decimal", "numeric":
		// 保留精度
		if strings.Contains(t, "(") {
			return "numeric" + t[strings.Index(t, "("):]
		}
		return "numeric"
	case "char":
		if strings.Contains(t, "(") {
			return "char" + t[strings.Index(t, "("):]
		}
		return "char(1)"
	case "varchar":
		if strings.Contains(t, "(") {
			return "varchar" + t[strings.Index(t, "("):]
		}
		return "varchar(255)"
	case "tinytext", "text", "mediumtext", "longtext":
		return "text"
	case "blob", "tinyblob", "mediumblob", "longblob":
		return "bytea"
	case "binary", "varbinary":
		return "bytea"
	case "date":
		return "date"
	case "datetime", "timestamp":
		return "timestamp"
	case "time":
		return "time"
	case "json":
		return "jsonb"
	case "enum":
		return "varchar(50)"
	case "set":
		return "varchar(255)"
	default:
		return t // 原样返回
	}
}

// mysqlTypeToPgSerial 将 MySQL 主键自增类型映射为 PG SERIAL 类型
func mysqlTypeToPgSerial(mysqlType string) string {
	t := strings.ToLower(strings.TrimSpace(mysqlType))
	if strings.Contains(t, "bigint") {
		return "BIGSERIAL"
	}
	return "SERIAL"
}
