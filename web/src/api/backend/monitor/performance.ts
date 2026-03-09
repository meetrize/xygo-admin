// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

/**
 * 性能分析 API
 */
import { adminRequest } from '@/utils/http'

/** 概览统计 */
export interface StatsSummary {
  totalRequests: number
  avgElapsed: number
  errorCount: number
  slowCount: number
}

/** 趋势项 */
export interface TrendItem {
  time: string
  count: number
  avgElapsed: number
}

/** 模块统计 */
export interface ModuleStats {
  module: string
  avgElapsed: number
  count: number
}

/** 错误分布 */
export interface ErrorDistItem {
  module: string
  errorCount: number
}

/** 性能统计响应 */
export interface PerformanceStatsResult {
  summary: StatsSummary
  trend: TrendItem[]
  moduleTop: ModuleStats[]
  errorDist: ErrorDistItem[]
}

/** 慢接口项 */
export interface SlowApiItem {
  url: string
  method: string
  module: string
  avgElapsed: number
  maxElapsed: number
  count: number
}

/** 获取性能统计 */
export function getPerformanceStats(params: { startDate?: string; endDate?: string }) {
  return adminRequest.post<PerformanceStatsResult>({
    url: '/monitor/stats',
    data: params
  })
}

/** 获取慢接口排行 */
export function getSlowApiTop(params: { startDate?: string; endDate?: string; limit?: number }) {
  return adminRequest.post<{ list: SlowApiItem[] }>({
    url: '/monitor/slow-top',
    data: params
  })
}
