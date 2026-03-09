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
	"testing"
)

// TestMysqlDialect_QuoteIdentifier 测试 MySQL 标识符引号
func TestMysqlDialect_QuoteIdentifier(t *testing.T) {
	d := &mysqlDialect{}
	tests := []struct {
		in, want string
	}{
		{"table_name", "`table_name`"},
		{"col", "`col`"},
	}
	for _, tt := range tests {
		got := d.QuoteIdentifier(tt.in)
		if got != tt.want {
			t.Errorf("QuoteIdentifier(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// TestPgDialect_QuoteIdentifier 测试 PG 标识符引号
func TestPgDialect_QuoteIdentifier(t *testing.T) {
	d := &pgDialect{}
	tests := []struct {
		in, want string
	}{
		{"table_name", `"table_name"`},
		{"col", `"col"`},
	}
	for _, tt := range tests {
		got := d.QuoteIdentifier(tt.in)
		if got != tt.want {
			t.Errorf("QuoteIdentifier(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// TestMysqlDialect_DateFormat 测试 MySQL DATE_FORMAT
func TestMysqlDialect_DateFormat(t *testing.T) {
	d := &mysqlDialect{}
	got := d.DateFormat("FROM_UNIXTIME(created_at)", "%%Y-%%m-%%d %%H")
	want := "DATE_FORMAT(FROM_UNIXTIME(created_at), '%%Y-%%m-%%d %%H')"
	if got != want {
		t.Errorf("DateFormat = %q, want %q", got, want)
	}
}

// TestPgDialect_DateFormat 测试 PG 日期格式转换
func TestPgDialect_DateFormat(t *testing.T) {
	d := &pgDialect{}
	got := d.DateFormat("to_timestamp(created_at)", "%%Y-%%m-%%d %%H")
	want := "to_char(to_timestamp(created_at), 'YYYY-MM-DD HH24')"
	if got != want {
		t.Errorf("DateFormat = %q, want %q", got, want)
	}
}

// TestPgDialect_FromUnixtime 测试 PG 时间戳转换
func TestPgDialect_FromUnixtime(t *testing.T) {
	d := &pgDialect{}
	got := d.FromUnixtime("created_at")
	want := "to_timestamp(created_at)"
	if got != want {
		t.Errorf("FromUnixtime = %q, want %q", got, want)
	}
}

// TestMysqlDialect_UnixTimestampNow 测试 MySQL 当前时间戳
func TestMysqlDialect_UnixTimestampNow(t *testing.T) {
	d := &mysqlDialect{}
	got := d.UnixTimestampNow()
	if got != "UNIX_TIMESTAMP()" {
		t.Errorf("UnixTimestampNow = %q, want UNIX_TIMESTAMP()", got)
	}
}

// TestPgDialect_UnixTimestampNow 测试 PG 当前时间戳
func TestPgDialect_UnixTimestampNow(t *testing.T) {
	d := &pgDialect{}
	got := d.UnixTimestampNow()
	if got != "extract(epoch from now())::bigint" {
		t.Errorf("UnixTimestampNow = %q, want extract(epoch from now())::bigint", got)
	}
}

// TestMysqlTypeToPgType 测试 MySQL -> PG 类型映射
func TestMysqlTypeToPgType(t *testing.T) {
	tests := []struct {
		mysqlType string
		pgType    string
	}{
		{"tinyint", "smallint"},
		{"int", "integer"},
		{"bigint", "bigint"},
		{"bigint unsigned", "bigint"},
		{"varchar(100)", "varchar(100)"},
		{"text", "text"},
		{"longtext", "text"},
		{"datetime", "timestamp"},
		{"json", "jsonb"},
		{"float", "real"},
		{"double", "double precision"},
		{"decimal(10,2)", "numeric(10,2)"},
		{"blob", "bytea"},
		{"enum", "varchar(50)"},
	}
	for _, tt := range tests {
		got := mysqlTypeToPgType(tt.mysqlType)
		if got != tt.pgType {
			t.Errorf("mysqlTypeToPgType(%q) = %q, want %q", tt.mysqlType, got, tt.pgType)
		}
	}
}

// TestMysqlTypeToPgSerial 测试 PG SERIAL 类型映射
func TestMysqlTypeToPgSerial(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"int", "SERIAL"},
		{"bigint", "BIGSERIAL"},
		{"bigint unsigned", "BIGSERIAL"},
		{"tinyint", "SERIAL"},
	}
	for _, tt := range tests {
		got := mysqlTypeToPgSerial(tt.in)
		if got != tt.want {
			t.Errorf("mysqlTypeToPgSerial(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// TestPgDialect_TypeToGoType 测试 PG 类型到 Go 类型映射
func TestPgDialect_TypeToGoType(t *testing.T) {
	d := &pgDialect{}
	tests := []struct {
		dataType, columnType, goType string
	}{
		{"int4", "int4", "int"},
		{"int8", "int8", "int64"},
		{"varchar", "varchar(100)", "string"},
		{"text", "text", "string"},
		{"bool", "bool", "int"},
		{"jsonb", "jsonb", "*gjson.Json"},
		{"timestamp", "timestamp", "*gtime.Time"},
		{"timestamptz", "timestamptz", "*gtime.Time"},
		{"bytea", "bytea", "[]byte"},
		{"uuid", "uuid", "string"},
		{"numeric", "numeric(10,2)", "float64"},
	}
	for _, tt := range tests {
		got := d.TypeToGoType(tt.dataType, tt.columnType)
		if got != tt.goType {
			t.Errorf("TypeToGoType(%q, %q) = %q, want %q", tt.dataType, tt.columnType, got, tt.goType)
		}
	}
}

// TestPgDialect_NullCoalesce 测试 PG COALESCE
func TestPgDialect_NullCoalesce(t *testing.T) {
	d := &pgDialect{}
	got := d.NullCoalesce("col", "''")
	want := "COALESCE(col, '')"
	if got != want {
		t.Errorf("NullCoalesce = %q, want %q", got, want)
	}
}

// TestMysqlDialect_NullCoalesce 测试 MySQL IFNULL
func TestMysqlDialect_NullCoalesce(t *testing.T) {
	d := &mysqlDialect{}
	got := d.NullCoalesce("col", "''")
	want := "IFNULL(col, '')"
	if got != want {
		t.Errorf("NullCoalesce = %q, want %q", got, want)
	}
}

// TestPgDialect_BuildColumnDef 测试 PG 列定义构建
func TestPgDialect_BuildColumnDef(t *testing.T) {
	d := &pgDialect{}

	// 主键列
	pkCol := ColumnMeta{Name: "id", Type: "int", IsPk: true}
	got := d.BuildColumnDef(pkCol)
	if got != `  "id" SERIAL NOT NULL` {
		t.Errorf("BuildColumnDef(pk) = %q", got)
	}

	// bigint 主键
	pkBig := ColumnMeta{Name: "id", Type: "bigint", IsPk: true}
	got = d.BuildColumnDef(pkBig)
	if got != `  "id" BIGSERIAL NOT NULL` {
		t.Errorf("BuildColumnDef(bigint pk) = %q", got)
	}

	// 普通列
	normalCol := ColumnMeta{Name: "name", Type: "varchar(100)", IsNullable: false, DefaultValue: "", Comment: "姓名"}
	got = d.BuildColumnDef(normalCol)
	if got != `  "name" varchar(100) NOT NULL DEFAULT ''` {
		t.Errorf("BuildColumnDef(normal) = %q", got)
	}
}

// TestMysqlDialect_CreateTableSQL 测试 MySQL 建表 SQL
func TestMysqlDialect_CreateTableSQL(t *testing.T) {
	d := &mysqlDialect{}
	colDefs := []string{"  `id` int NOT NULL AUTO_INCREMENT", "  `name` varchar(100) NOT NULL DEFAULT ''", "  PRIMARY KEY (`id`)"}
	got := d.CreateTableSQL("xy_test", colDefs, "测试表")
	if !contains(got, "ENGINE=InnoDB") || !contains(got, "COMMENT='测试表'") || !contains(got, "`xy_test`") {
		t.Errorf("CreateTableSQL = %q", got)
	}
}

// TestPgDialect_CreateTableSQL 测试 PG 建表 SQL
func TestPgDialect_CreateTableSQL(t *testing.T) {
	d := &pgDialect{}
	colDefs := []string{`  "id" SERIAL NOT NULL`, `  "name" varchar(100) NOT NULL DEFAULT ''`}
	got := d.CreateTableSQL("xy_test", colDefs, "测试表")
	if !contains(got, `"xy_test"`) || !contains(got, "COMMENT ON TABLE") || contains(got, "ENGINE") {
		t.Errorf("CreateTableSQL = %q", got)
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && containsStr(s, substr)
}

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && searchStr(s, substr)
}

func searchStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// TestMysqlFmtToPgFmt 测试日期格式转换
func TestMysqlFmtToPgFmt(t *testing.T) {
	tests := []struct {
		mysql, pg string
	}{
		{"%%Y-%%m-%%d %%H", "YYYY-MM-DD HH24"},
		{"%Y-%m-%d %H:%i:%s", "YYYY-MM-DD HH24:MI:SS"},
	}
	for _, tt := range tests {
		got := mysqlFmtToPgFmt(tt.mysql)
		if got != tt.pg {
			t.Errorf("mysqlFmtToPgFmt(%q) = %q, want %q", tt.mysql, got, tt.pg)
		}
	}
}
