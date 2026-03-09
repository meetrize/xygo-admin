// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package gencodes

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/library/dbdialect"
	"xygo/internal/model/input/adminin"
)

// createTableFromDesign 根据设计器表单创建数据表（通过方言层适配 MySQL/PG）
func createTableFromDesign(ctx context.Context, in *adminin.GenCodesCreateTableInp) (*adminin.GenCodesCreateTableModel, error) {
	tablePrefix := getTablePrefix(ctx)
	dialect := dbdialect.Get()

	// 确保表名有前缀
	tableName := in.TableName
	if !strings.HasPrefix(tableName, tablePrefix) {
		tableName = tablePrefix + tableName
	}

	// 构建列定义
	var colDefs []string
	var pkName string
	// PG 的列注释需要单独执行 COMMENT ON COLUMN，收集起来
	var commentSQLs []string

	for _, col := range in.Columns {
		meta := dbdialect.ColumnMeta{
			Name:         col.Name,
			Type:         col.Type,
			IsPk:         col.IsPk == 1,
			IsNullable:   col.IsNullable == 1,
			DefaultValue: col.DefaultValue,
			Comment:      col.Comment,
		}
		def := dialect.BuildColumnDef(meta)
		colDefs = append(colDefs, def)

		if col.IsPk == 1 {
			pkName = col.Name
		}

		// PG 需要单独的 COMMENT ON COLUMN
		if dbdialect.IsPgsql() && col.Comment != "" {
			commentSQLs = append(commentSQLs,
				fmt.Sprintf(`COMMENT ON COLUMN %s.%s IS '%s'`,
					dialect.QuoteIdentifier(tableName),
					dialect.QuoteIdentifier(col.Name),
					strings.ReplaceAll(col.Comment, "'", "''")))
		}
	}

	if pkName != "" {
		colDefs = append(colDefs, fmt.Sprintf("  PRIMARY KEY (%s)", dialect.QuoteIdentifier(pkName)))
	}

	// 生成建表 SQL
	sql := dialect.CreateTableSQL(tableName, colDefs, in.TableComment)

	// 执行建表（PG 的 CreateTableSQL 可能包含多条语句用分号分隔）
	for _, stmt := range splitSQL(sql) {
		if _, err := g.DB().Exec(ctx, stmt); err != nil {
			return nil, fmt.Errorf("创建数据表失败: %w", err)
		}
	}

	// PG: 执行列注释
	for _, cs := range commentSQLs {
		if _, err := g.DB().Exec(ctx, cs); err != nil {
			g.Log().Warningf(ctx, "[CreateTable] comment failed: %v\nSQL: %s", err, cs)
		}
	}

	return &adminin.GenCodesCreateTableModel{
		TableName: tableName,
	}, nil
}

// splitSQL 按分号拆分 SQL（跳过空语句），用于 PG 的多语句执行
func splitSQL(sql string) []string {
	parts := strings.Split(sql, ";")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
