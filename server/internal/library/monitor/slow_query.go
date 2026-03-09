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

	"github.com/gogf/gf/v2/frame/g"
)

// InitPerformanceMonitor 初始化性能监控配置
// - 慢查询：GoFrame 的 database.debug=true 已自动输出每条 SQL 耗时
//   本函数仅读取配置并打印启用状态，实际监控由 SlowApiMonitor 中间件完成
// - 慢接口：通过 middleware.SlowApiMonitor 在请求维度监控耗时
func InitPerformanceMonitor(ctx context.Context) {
	// 慢查询监控状态
	slowQueryEnabled := g.Cfg().MustGet(ctx, "performance.slowQuery.enabled").Bool()
	slowQueryThreshold := g.Cfg().MustGet(ctx, "performance.slowQuery.threshold").Int64()
	if slowQueryThreshold <= 0 {
		slowQueryThreshold = 200
	}
	if slowQueryEnabled {
		g.Log().Infof(ctx, "[性能监控] 慢查询告警已启用，阈值: %dms（DB debug 日志中 SQL 耗时超此值需关注）", slowQueryThreshold)
	}

	// 慢接口监控状态
	slowApiEnabled := g.Cfg().MustGet(ctx, "performance.slowApi.enabled").Bool()
	slowApiThreshold := g.Cfg().MustGet(ctx, "performance.slowApi.threshold").Int64()
	if slowApiThreshold <= 0 {
		slowApiThreshold = 1000
	}
	if slowApiEnabled {
		g.Log().Infof(ctx, "[性能监控] 慢接口告警已启用，阈值: %dms", slowApiThreshold)
	}
}
