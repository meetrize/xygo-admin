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
 * 统一时间格式化工具
 *
 * 支持两种时区模式：
 * - client: 使用浏览器本地时区（默认）
 * - server: 使用后端配置的时区（从 /site/index 获取）
 *
 * 兼容三种输入：
 * - 数字时间戳（秒）：1770908749
 * - 毫秒时间戳：1770908749000
 * - 日期字符串："2026-02-12 22:45:49"
 */

import { useSiteStore } from '@/store/modules/site'

/**
 * 格式化时间为可读字符串
 * @param ts 秒级/毫秒级时间戳 或 日期字符串
 * @param format 格式化选项（默认 datetime）
 *   - 'datetime': 2026-02-12 21:56:43
 *   - 'date': 2026-02-12
 *   - 'time': 21:56:43
 *   - 'year': 2026
 * @returns 格式化后的字符串，ts 为空/0 返回 '-'
 */
export function formatTimestamp(ts: number | string | undefined | null, format: 'datetime' | 'date' | 'time' | 'year' = 'datetime'): string {
  if (ts === undefined || ts === null || ts === '' || ts === 0 || ts === '0') return '-'

  let ms: number

  if (typeof ts === 'string') {
    // 日期字符串（含 - / T 等）→ 直接 Date.parse
    if (ts.includes('-') || ts.includes('/') || ts.includes('T')) {
      const d = new Date(ts.replace(/\//g, '-'))
      if (!isNaN(d.getTime())) {
        ms = d.getTime()
      } else {
        return String(ts)
      }
    } else {
      // 纯数字字符串
      const n = parseInt(ts, 10)
      if (isNaN(n) || n <= 0) return '-'
      ms = n < 1e12 ? n * 1000 : n
    }
  } else {
    if (ts <= 0) return '-'
    // 数字：判断秒级还是毫秒级
    ms = ts < 1e12 ? ts * 1000 : ts
  }

  const siteStore = useSiteStore()
  const mode = siteStore.getTimeZoneMode()
  const timezone = siteStore.getTimezone()

  try {
    if (mode === 'server' && timezone) {
      // 使用后端配置的时区
      const options: Intl.DateTimeFormatOptions = { timeZone: timezone }
      if (format === 'year') {
        options.year = 'numeric'
      } else {
        if (format !== 'time') { options.year = 'numeric'; options.month = '2-digit'; options.day = '2-digit' }
        if (format !== 'date') { options.hour = '2-digit'; options.minute = '2-digit'; options.second = '2-digit'; options.hour12 = false }
      }
      return new Intl.DateTimeFormat('zh-CN', options).format(new Date(ms))
    }

    // client 模式：使用浏览器本地时区
    const d = new Date(ms)
    if (format === 'year') {
      return String(d.getFullYear())
    }
    if (format === 'date') {
      return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}`
    }
    if (format === 'time') {
      return `${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`
    }
    return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`
  } catch {
    return String(ts)
  }
}

function pad(n: number): string {
  return n < 10 ? '0' + n : String(n)
}
