// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

import axios from 'axios'

/**
 * 站点信息接口（公开接口，不使用baseURL）
 */

export interface SiteInfo {
  group: string
  items: Record<string, string>
  siteName: string
  siteSubtitle: string
  icp: string
  timezone: string
  description: string
  themeColor: string
  logo: string
  closed: string
  openMemberCenter: boolean
}

/**
 * 获取站点基础信息（公开接口，不加/admin前缀）
 */
export async function fetchSiteIndex(): Promise<SiteInfo> {
  const { data } = await axios.get<{ code: number; data: SiteInfo; message?: string; msg?: string }>(
    '/site/index'
  )
  if (data.code === 0) {
    return data.data
  }
  throw new Error(data.message || data.msg || '获取站点信息失败')
}
