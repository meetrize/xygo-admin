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
 * Iframe 路由管理器
 *
 * 负责管理 iframe 类型的路由
 *
 * @module router/core/IframeRouteManager
 * @author Art Design Pro Team
 */

import type { AppRouteRecord } from '@/types/router'

export class IframeRouteManager {
  private static instance: IframeRouteManager
  private iframeRoutes: AppRouteRecord[] = []

  private constructor() {}

  static getInstance(): IframeRouteManager {
    if (!IframeRouteManager.instance) {
      IframeRouteManager.instance = new IframeRouteManager()
    }
    return IframeRouteManager.instance
  }

  /**
   * 添加 iframe 路由
   */
  add(route: AppRouteRecord): void {
    if (!this.iframeRoutes.find((r) => r.path === route.path)) {
      this.iframeRoutes.push(route)
    }
  }

  /**
   * 获取所有 iframe 路由
   */
  getAll(): AppRouteRecord[] {
    return this.iframeRoutes
  }

  /**
   * 根据路径查找 iframe 路由
   */
  findByPath(path: string): AppRouteRecord | undefined {
    return this.iframeRoutes.find((route) => route.path === path)
  }

  /**
   * 清空所有 iframe 路由
   */
  clear(): void {
    this.iframeRoutes = []
  }

  /**
   * 保存到 sessionStorage
   */
  save(): void {
    if (this.iframeRoutes.length > 0) {
      sessionStorage.setItem('iframeRoutes', JSON.stringify(this.iframeRoutes))
    }
  }

  /**
   * 从 sessionStorage 加载
   */
  load(): void {
    try {
      const data = sessionStorage.getItem('iframeRoutes')
      if (data) {
        this.iframeRoutes = JSON.parse(data)
      }
    } catch (error) {
      console.error('[IframeRouteManager] 加载 iframe 路由失败:', error)
      this.iframeRoutes = []
    }
  }
}
