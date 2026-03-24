// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	api "xygo/api/admin"
	"xygo/internal/dao"
	"xygo/internal/library/dbdialect"
	"xygo/internal/service"
)

type sMonitor struct{}

func init() {
	service.RegisterMonitor(&sMonitor{})
}

var opLogTable = dao.AdminOperationLog.Table()
var opLogCols = dao.AdminOperationLog.Columns()

// GetPerformanceStats 获取性能统计聚合数据
func (s *sMonitor) GetPerformanceStats(ctx context.Context, startDate, endDate string) (*api.MonitorStatsRes, error) {
	res := &api.MonitorStatsRes{}

	// 构建时间条件
	timeWhere := buildTimeWhere(startDate, endDate)

	// 1. 概览统计
	summary, err := getSummary(ctx, timeWhere)
	if err != nil {
		return nil, err
	}
	res.Summary = summary

	// 2. 请求趋势（按小时聚合）
	trend, err := getTrend(ctx, timeWhere)
	if err != nil {
		return nil, err
	}
	res.Trend = trend

	// 3. 模块耗时排行 Top 10
	moduleTop, err := getModuleTop(ctx, timeWhere, 10)
	if err != nil {
		return nil, err
	}
	res.ModuleTop = moduleTop

	// 4. 错误分布
	errorDist, err := getErrorDist(ctx, timeWhere)
	if err != nil {
		return nil, err
	}
	res.ErrorDist = errorDist

	return res, nil
}

// GetSlowApiTop 获取慢接口排行
func (s *sMonitor) GetSlowApiTop(ctx context.Context, startDate, endDate string, limit int) (*api.MonitorSlowTopRes, error) {
	if limit <= 0 {
		limit = 20
	}

	timeWhere := buildTimeWhere(startDate, endDate)

	type slowRow struct {
		Url        string  `json:"url"`
		Method     string  `json:"method"`
		Module     string  `json:"module"`
		AvgElapsed float64 `json:"avgElapsed"`
		MaxElapsed int     `json:"maxElapsed"`
		Count      int     `json:"count"`
	}

	var rows []slowRow
	sql := fmt.Sprintf(
		`SELECT url, method, module,
		 ROUND(AVG(elapsed), 1) as avgElapsed,
		 MAX(elapsed) as maxElapsed,
		 COUNT(*) as count
		 FROM %s WHERE 1=1 %s
		 GROUP BY url, method, module
		 ORDER BY avgElapsed DESC
		 LIMIT %d`,
		opLogTable, timeWhere, limit,
	)

	err := g.DB().Ctx(ctx).Raw(sql).Scan(&rows)
	if err != nil {
		return nil, err
	}

	list := make([]api.SlowApiItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, api.SlowApiItem{
			Url:        r.Url,
			Method:     r.Method,
			Module:     r.Module,
			AvgElapsed: r.AvgElapsed,
			MaxElapsed: r.MaxElapsed,
			Count:      r.Count,
		})
	}

	return &api.MonitorSlowTopRes{List: list}, nil
}

// --------- 内部辅助函数 ---------

func buildTimeWhere(startDate, endDate string) string {
	dialect := dbdialect.Get()
	where := ""
	if startDate != "" {
		// created_at 为 bigint 时间戳，将日期字符串转为 Unix 时间戳
		t, err := time.ParseInLocation("2006-01-02", startDate, time.Local)
		if err == nil {
			where += fmt.Sprintf(" AND %s >= %d", dialect.QuoteIdentifier(opLogCols.CreatedAt), t.Unix())
		}
	}
	if endDate != "" {
		t, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
		if err == nil {
			// 当天 23:59:59
			where += fmt.Sprintf(" AND %s <= %d", dialect.QuoteIdentifier(opLogCols.CreatedAt), t.Unix()+86399)
		}
	}
	return where
}

func getSummary(ctx context.Context, timeWhere string) (*api.StatsSummary, error) {
	dialect := dbdialect.Get()
	type summaryRow struct {
		TotalRequests int     `json:"totalRequests"`
		AvgElapsed    float64 `json:"avgElapsed"`
		ErrorCount    int     `json:"errorCount"`
		SlowCount     int     `json:"slowCount"`
	}

	var row summaryRow
	// status 统一为 smallint/tinyint，两边都用 status = 0
	statusFailCond := "status = 0"
	sql := fmt.Sprintf(
		`SELECT 
		 COUNT(*) as "totalRequests",
		 %s as "avgElapsed",
		 SUM(CASE WHEN %s THEN 1 ELSE 0 END) as "errorCount",
		 SUM(CASE WHEN elapsed > 200 THEN 1 ELSE 0 END) as "slowCount"
		 FROM %s WHERE 1=1 %s`,
		dialect.RoundExpr("AVG(elapsed)", 1),
		statusFailCond,
		opLogTable, timeWhere,
	)

	err := g.DB().Ctx(ctx).Raw(sql).Scan(&row)
	if err != nil {
		return nil, err
	}

	return &api.StatsSummary{
		TotalRequests: row.TotalRequests,
		AvgElapsed:    row.AvgElapsed,
		ErrorCount:    row.ErrorCount,
		SlowCount:     row.SlowCount,
	}, nil
}

func getTrend(ctx context.Context, timeWhere string) ([]api.TrendItem, error) {
	dialect := dbdialect.Get()
	type trendRow struct {
		Time       string  `json:"time"`
		Count      int     `json:"count"`
		AvgElapsed float64 `json:"avgElapsed"`
	}

	// 构建方言兼容的时间表达式
	// MySQL: DATE_FORMAT(FROM_UNIXTIME(created_at), '%Y-%m-%d %H')
	// PG:    to_char(to_timestamp(created_at), 'YYYY-MM-DD HH24')
	timeExpr := dialect.DateFormat(dialect.FromUnixtime("created_at"), "%Y-%m-%d %H")
	roundExpr := dialect.RoundExpr("AVG(elapsed)", 1)

	var rows []trendRow
	sql := `SELECT ` + timeExpr + ` as time, COUNT(*) as count, ` +
		roundExpr + ` as "avgElapsed" FROM ` + opLogTable +
		` WHERE 1=1 ` + timeWhere +
		` GROUP BY ` + timeExpr +
		` ORDER BY time ASC LIMIT 168`

	err := g.DB().Ctx(ctx).Raw(sql).Scan(&rows)
	if err != nil {
		return nil, err
	}

	list := make([]api.TrendItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, api.TrendItem{
			Time:       r.Time,
			Count:      r.Count,
			AvgElapsed: r.AvgElapsed,
		})
	}
	return list, nil
}

func getModuleTop(ctx context.Context, timeWhere string, limit int) ([]api.ModuleStats, error) {
	type moduleRow struct {
		Module     string  `json:"module"`
		AvgElapsed float64 `json:"avgElapsed"`
		Count      int     `json:"count"`
	}

	var rows []moduleRow
	sql := fmt.Sprintf(
		`SELECT module,
		 ROUND(AVG(elapsed), 1) as avgElapsed,
		 COUNT(*) as count
		 FROM %s WHERE 1=1 %s AND module != ''
		 GROUP BY module
		 ORDER BY avgElapsed DESC
		 LIMIT %d`,
		opLogTable, timeWhere, limit,
	)

	err := g.DB().Ctx(ctx).Raw(sql).Scan(&rows)
	if err != nil {
		return nil, err
	}

	list := make([]api.ModuleStats, 0, len(rows))
	for _, r := range rows {
		list = append(list, api.ModuleStats{
			Module:     r.Module,
			AvgElapsed: r.AvgElapsed,
			Count:      r.Count,
		})
	}
	return list, nil
}

func getErrorDist(ctx context.Context, timeWhere string) ([]api.ErrorDistItem, error) {
	type errorRow struct {
		Module     string `json:"module"`
		ErrorCount int    `json:"errorCount"`
	}

	// status 统一为 smallint/tinyint，两边都用 status = 0
	statusFailCond := "status = 0"

	var rows []errorRow
	sql := fmt.Sprintf(
		`SELECT module, COUNT(*) as "errorCount"
		 FROM %s WHERE %s %s AND module != ''
		 GROUP BY module
		 ORDER BY "errorCount" DESC
		 LIMIT 10`,
		opLogTable, statusFailCond, timeWhere,
	)

	err := g.DB().Ctx(ctx).Raw(sql).Scan(&rows)
	if err != nil {
		return nil, err
	}

	list := make([]api.ErrorDistItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, api.ErrorDistItem{
			Module:     r.Module,
			ErrorCount: r.ErrorCount,
		})
	}
	return list, nil
}
