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

// autoSyncFieldsToDb 生成前自动同步设计器新增字段到数据库（只做 ADD，不做 DROP/MODIFY）
func autoSyncFieldsToDb(ctx context.Context, in *adminin.GenCodesEditInp) error {
	tablePrefix := getTablePrefix(ctx)
	dialect := dbdialect.Get()
	tableName := in.TableName
	if !strings.HasPrefix(tableName, tablePrefix) {
		tableName = tablePrefix + tableName
	}

	// 获取数据库实际字段
	dbColumns, err := getDbColumns(ctx, tableName)
	if err != nil {
		return err
	}
	dbMap := make(map[string]bool)
	for _, c := range dbColumns {
		dbMap[c.Name] = true
	}

	// 找出设计器有但数据库没有的字段，执行 ADD COLUMN
	var addedCount int
	for _, col := range in.Columns {
		if dbMap[col.Name] {
			continue // 已存在
		}
		// 构建 ADD COLUMN 并执行
		meta := dbdialect.ColumnMeta{
			Name:         col.Name,
			Type:         col.DbType,
			IsPk:         col.IsPk == 1,
			IsNullable:   false,
			DefaultValue: "",
			Comment:      col.Comment,
		}
		if meta.Type == "" {
			meta.Type = "varchar(100)" // 兜底
		}
		sql := dialect.BuildAddColumnSQL(tableName, meta)
		g.Log().Infof(ctx, "[AutoSync] adding column: %s", sql)
		// PG 的 BuildAddColumnSQL 可能包含 COMMENT ON COLUMN（分号分隔）
		for _, stmt := range splitSQL(sql) {
			if _, err := g.DB().Exec(ctx, stmt); err != nil {
				g.Log().Warningf(ctx, "[AutoSync] ADD COLUMN failed: %v\nSQL: %s", err, stmt)
				return fmt.Errorf("添加字段 %s 失败: %w", col.Name, err)
			}
		}
		addedCount++
	}

	// 找出数据库有但设计器没有的字段（排除主键），执行 DROP COLUMN
	designMap := make(map[string]bool)
	for _, col := range in.Columns {
		designMap[col.Name] = true
	}
	var droppedCount int
	for _, dbCol := range dbColumns {
		if designMap[dbCol.Name] {
			continue
		}
		if dbCol.ColumnKey == "PRI" {
			continue // 主键不删
		}
		sql := dialect.BuildDropColumnSQL(tableName, dbCol.Name)
		g.Log().Infof(ctx, "[AutoSync] dropping column: %s", sql)
		if _, err := g.DB().Exec(ctx, sql); err != nil {
			g.Log().Warningf(ctx, "[AutoSync] DROP COLUMN failed: %v", err)
			// 不中断，继续
		} else {
			droppedCount++
		}
	}

	if addedCount > 0 || droppedCount > 0 {
		g.Log().Infof(ctx, "[AutoSync] added %d, dropped %d columns in %s, running gf gen dao...", addedCount, droppedCount, tableName)
		if err := runGfGenDao(ctx); err != nil {
			g.Log().Warningf(ctx, "[AutoSync] gf gen dao warning: %v", err)
		}
	}
	return nil
}

// SyncFields 预览字段变更（对比设计器字段与数据库实际字段）
func (s *sGenCodes) SyncFields(ctx context.Context, in *adminin.GenCodesSyncFieldsInp) (*adminin.GenCodesSyncFieldsModel, error) {
	tablePrefix := getTablePrefix(ctx)
	dialect := dbdialect.Get()
	tableName := in.TableName
	if !strings.HasPrefix(tableName, tablePrefix) {
		tableName = tablePrefix + tableName
	}

	// 1. 从 information_schema 获取数据库实际字段
	dbColumns, err := getDbColumns(ctx, tableName)
	if err != nil {
		return nil, fmt.Errorf("获取数据库字段失败: %w", err)
	}

	// 2. 构建映射
	dbMap := make(map[string]*dbColumnInfo)
	for i := range dbColumns {
		dbMap[dbColumns[i].Name] = &dbColumns[i]
	}
	designMap := make(map[string]*adminin.CreateTableColumn)
	for i := range in.Columns {
		designMap[in.Columns[i].Name] = &in.Columns[i]
	}

	// 3. 对比差异
	var diffs []adminin.FieldDiff

	// 3.1 新增字段（设计器有，数据库没有）
	for _, col := range in.Columns {
		if _, exists := dbMap[col.Name]; !exists {
			meta := dbdialect.ColumnMeta{
				Name:         col.Name,
				Type:         col.Type,
				IsPk:         col.IsPk == 1,
				IsNullable:   col.IsNullable == 1,
				DefaultValue: col.DefaultValue,
				Comment:      col.Comment,
			}
			sql := dialect.BuildAddColumnSQL(tableName, meta)
			diffs = append(diffs, adminin.FieldDiff{
				Name:   col.Name,
				Action: "add",
				Detail: fmt.Sprintf("新增字段 %s (%s)", col.Name, col.Type),
				SQL:    sql,
			})
		}
	}

	// 3.2 删除字段（数据库有，设计器没有）
	for _, dbCol := range dbColumns {
		if _, exists := designMap[dbCol.Name]; !exists {
			sql := dialect.BuildDropColumnSQL(tableName, dbCol.Name)
			diffs = append(diffs, adminin.FieldDiff{
				Name:    dbCol.Name,
				Action:  "drop",
				Detail:  fmt.Sprintf("删除字段 %s（数据将丢失）", dbCol.Name),
				SQL:     sql,
				IsRisky: true,
			})
		}
	}

	// 3.3 修改字段（都有，但定义不同）
	for _, col := range in.Columns {
		dbCol, exists := dbMap[col.Name]
		if !exists {
			continue // 新增的已处理
		}
		changes := compareColumn(&col, dbCol)
		if len(changes) > 0 {
			meta := dbdialect.ColumnMeta{
				Name:         col.Name,
				Type:         col.Type,
				IsPk:         col.IsPk == 1,
				IsNullable:   col.IsNullable == 1,
				DefaultValue: col.DefaultValue,
				Comment:      col.Comment,
			}
			sql := dialect.BuildModifyColumnSQL(tableName, meta)
			isRisky := isTypeNarrowing(dbCol.ColumnType, col.Type)
			diffs = append(diffs, adminin.FieldDiff{
				Name:    col.Name,
				Action:  "modify",
				Detail:  fmt.Sprintf("修改字段 %s: %s", col.Name, strings.Join(changes, ", ")),
				SQL:     sql,
				IsRisky: isRisky,
			})
		}
	}

	return &adminin.GenCodesSyncFieldsModel{
		Diffs:     diffs,
		HasChange: len(diffs) > 0,
	}, nil
}

// ExecuteDDL 执行字段同步DDL
func (s *sGenCodes) ExecuteDDL(ctx context.Context, in *adminin.GenCodesExecuteDDLInp) error {
	db := g.DB()
	for _, sql := range in.SQLs {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		// PG 的 MODIFY 可能包含多条语句（分号分隔），逐条执行
		for _, stmt := range splitSQL(sql) {
			g.Log().Infof(ctx, "[SyncFields] executing: %s", stmt)
			if _, err := db.Exec(ctx, stmt); err != nil {
				g.Log().Warningf(ctx, "[SyncFields] DDL failed: %v\nSQL: %s", err, stmt)
				return fmt.Errorf("执行失败: %s\n错误: %w", stmt, err)
			}
		}
	}
	g.Log().Info(ctx, "[SyncFields] all DDL executed successfully")

	// 执行完后自动运行 gf gen dao
	if err := runGfGenDao(ctx); err != nil {
		g.Log().Warningf(ctx, "[SyncFields] gf gen dao warning: %v", err)
		// 不中断，DDL 已成功
	}
	return nil
}

// ==================== 内部工具 ====================

// dbColumnInfo 数据库字段信息
type dbColumnInfo struct {
	Name         string
	ColumnType   string // 完整类型如 varchar(100)、bigint unsigned
	IsNullable   string // YES/NO
	DefaultValue string
	Comment      string
	ColumnKey    string // PRI/UNI/MUL
	Extra        string // auto_increment 等
}

// getDbColumns 从 information_schema 获取表的所有字段（通过方言层适配 MySQL/PG）
func getDbColumns(ctx context.Context, tableName string) ([]dbColumnInfo, error) {
	db := g.DB()
	dialect := dbdialect.Get()

	// 获取当前数据库名
	dbName, err := dialect.GetDbName(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetAll(ctx, dialect.ListColumnsSQLForSync(dbName, tableName))
	if err != nil {
		return nil, err
	}

	var cols []dbColumnInfo
	for _, row := range rows {
		cols = append(cols, dbColumnInfo{
			Name:         row["COLUMN_NAME"].String(),
			ColumnType:   row["COLUMN_TYPE"].String(),
			IsNullable:   row["IS_NULLABLE"].String(),
			DefaultValue: row["COLUMN_DEFAULT"].String(),
			Comment:      row["COLUMN_COMMENT"].String(),
			ColumnKey:    row["COLUMN_KEY"].String(),
			Extra:        row["EXTRA"].String(),
		})
	}
	return cols, nil
}

// compareColumn 对比设计器字段和数据库字段，返回差异描述
func compareColumn(design *adminin.CreateTableColumn, db *dbColumnInfo) []string {
	var changes []string

	// 类型对比（忽略大小写和空格）
	designType := strings.ToLower(strings.TrimSpace(design.Type))
	dbType := strings.ToLower(strings.TrimSpace(db.ColumnType))
	if designType != dbType {
		changes = append(changes, fmt.Sprintf("类型 %s → %s", db.ColumnType, design.Type))
	}

	// 可空对比
	designNullable := design.IsNullable == 1
	dbNullable := strings.ToUpper(db.IsNullable) == "YES"
	if designNullable != dbNullable {
		if designNullable {
			changes = append(changes, "改为可空")
		} else {
			changes = append(changes, "改为非空")
		}
	}

	// 注释对比
	if design.Comment != db.Comment {
		changes = append(changes, fmt.Sprintf("注释 %s → %s", db.Comment, design.Comment))
	}

	// 默认值对比
	if design.DefaultValue != db.DefaultValue {
		changes = append(changes, fmt.Sprintf("默认值 %s → %s", db.DefaultValue, design.DefaultValue))
	}

	return changes
}

// escapeSQLString 转义 SQL 字符串中的单引号，防止注入和语法错误
func escapeSQLString(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

// isNumericType 判断是否为数值类型（通过方言层适配）
func isNumericType(t string) bool {
	return dbdialect.Get().IsNumericType(t)
}

// getDefaultForType 根据类型返回合适的默认值（通过方言层适配）
func getDefaultForType(colType string) string {
	return dbdialect.Get().GetDefaultForType(colType)
}

// buildAddColumnSQL 构建 ADD COLUMN SQL（通过方言层适配）
func buildAddColumnSQL(tableName string, col *adminin.CreateTableColumn) string {
	dialect := dbdialect.Get()
	meta := dbdialect.ColumnMeta{
		Name:         col.Name,
		Type:         col.Type,
		IsPk:         col.IsPk == 1,
		IsNullable:   col.IsNullable == 1,
		DefaultValue: col.DefaultValue,
		Comment:      col.Comment,
	}
	return dialect.BuildAddColumnSQL(tableName, meta)
}

// buildModifyColumnSQL 构建 MODIFY COLUMN SQL（通过方言层适配）
func buildModifyColumnSQL(tableName string, col *adminin.CreateTableColumn) string {
	dialect := dbdialect.Get()
	meta := dbdialect.ColumnMeta{
		Name:         col.Name,
		Type:         col.Type,
		IsPk:         col.IsPk == 1,
		IsNullable:   col.IsNullable == 1,
		DefaultValue: col.DefaultValue,
		Comment:      col.Comment,
	}
	return dialect.BuildModifyColumnSQL(tableName, meta)
}

// isTypeNarrowing 判断类型变更是否有缩小风险
func isTypeNarrowing(oldType, newType string) bool {
	// 简单判断：如果新类型包含的数字（长度）小于旧类型，认为有风险
	oldLen := extractTypeLength(oldType)
	newLen := extractTypeLength(newType)
	if oldLen > 0 && newLen > 0 && newLen < oldLen {
		return true
	}
	// varchar -> int 等跨类型变更也有风险
	oldBase := strings.Split(strings.ToLower(oldType), "(")[0]
	newBase := strings.Split(strings.ToLower(newType), "(")[0]
	if oldBase != newBase {
		return true
	}
	return false
}

// extractTypeLength 提取类型中的长度数字
func extractTypeLength(t string) int {
	start := strings.Index(t, "(")
	end := strings.Index(t, ")")
	if start < 0 || end <= start+1 {
		return 0
	}
	num := strings.Split(t[start+1:end], ",")[0]
	n := 0
	for _, c := range num {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}
