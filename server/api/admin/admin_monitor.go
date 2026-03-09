// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package admin

import (
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/library/monitor"
)

// ===================== 服务器监控 =====================

// MonitorServerReq 服务器信息请求
type MonitorServerReq struct {
	g.Meta `path:"/admin/monitor/server" method:"get" tags:"AdminMonitor" summary:"Get server info"`
}

// MonitorServerRes 服务器信息响应
type MonitorServerRes struct {
	*monitor.ServerInfo
}

// ===================== 性能统计 =====================

// MonitorStatsReq 性能统计请求
type MonitorStatsReq struct {
	g.Meta    `path:"/admin/monitor/stats" method:"post" tags:"AdminMonitor" summary:"Get performance stats"`
	StartDate string `json:"startDate" dc:"开始日期 (YYYY-MM-DD)"`
	EndDate   string `json:"endDate"   dc:"结束日期 (YYYY-MM-DD)"`
}

// MonitorStatsRes 性能统计响应
type MonitorStatsRes struct {
	Summary   *StatsSummary  `json:"summary"`   // 概览统计
	Trend     []TrendItem    `json:"trend"`      // 请求趋势
	ModuleTop []ModuleStats  `json:"moduleTop"`  // 模块耗时排行
	ErrorDist []ErrorDistItem `json:"errorDist"` // 错误分布
}

// StatsSummary 概览统计
type StatsSummary struct {
	TotalRequests int     `json:"totalRequests"` // 总请求数
	AvgElapsed    float64 `json:"avgElapsed"`    // 平均耗时 (ms)
	ErrorCount    int     `json:"errorCount"`    // 错误请求数
	SlowCount     int     `json:"slowCount"`     // 慢接口数 (>200ms)
}

// TrendItem 趋势数据项
type TrendItem struct {
	Time       string  `json:"time"`       // 时间标签
	Count      int     `json:"count"`      // 请求数
	AvgElapsed float64 `json:"avgElapsed"` // 平均耗时
}

// ModuleStats 模块统计
type ModuleStats struct {
	Module     string  `json:"module"`     // 模块名
	AvgElapsed float64 `json:"avgElapsed"` // 平均耗时
	Count      int     `json:"count"`      // 调用次数
}

// ErrorDistItem 错误分布
type ErrorDistItem struct {
	Module     string `json:"module"`     // 模块名
	ErrorCount int    `json:"errorCount"` // 错误数
}

// ===================== 慢接口 Top N =====================

// MonitorSlowTopReq 慢接口排行请求
type MonitorSlowTopReq struct {
	g.Meta    `path:"/admin/monitor/slow-top" method:"post" tags:"AdminMonitor" summary:"Get slow API top N"`
	StartDate string `json:"startDate" dc:"开始日期"`
	EndDate   string `json:"endDate"   dc:"结束日期"`
	Limit     int    `json:"limit" d:"20" dc:"返回条数"`
}

// MonitorSlowTopRes 慢接口排行响应
type MonitorSlowTopRes struct {
	List []SlowApiItem `json:"list"`
}

// SlowApiItem 慢接口项
type SlowApiItem struct {
	Url        string  `json:"url"`        // 接口路径
	Method     string  `json:"method"`     // HTTP 方法
	Module     string  `json:"module"`     // 所属模块
	AvgElapsed float64 `json:"avgElapsed"` // 平均耗时 (ms)
	MaxElapsed int     `json:"maxElapsed"` // 最大耗时 (ms)
	Count      int     `json:"count"`      // 调用次数
}

// ===================== 函数级性能分析 =====================

// MonitorPprofTopReq 函数级性能分析请求
type MonitorPprofTopReq struct {
	g.Meta  `path:"/admin/monitor/pprof-top" method:"get" tags:"AdminMonitor" summary:"Get pprof function top N"`
	Seconds int `json:"seconds" d:"3"  dc:"CPU采样秒数(1-30)"`
	Limit   int `json:"limit"   d:"15" dc:"返回条数"`
}

// MonitorPprofTopRes 函数级性能分析响应
type MonitorPprofTopRes struct {
	*monitor.PprofTopResult
}
