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
 * 插件管理器
 * 
 * BuildAdmin 风格：插件独立存放，安装时复制到项目目录
 * 
 * @module utils/plugin/PluginManager
 */

import { adminRequest } from '@/utils/http'
import type { PluginInfo, PluginListResponse } from './types'

class PluginManager {
  private plugins: Map<string, PluginInfo> = new Map()
  private loaded: boolean = false

  /**
   * 获取所有插件列表
   */
  async getPlugins(): Promise<PluginInfo[]> {
    if (!this.loaded) {
      await this.loadPlugins()
    }
    return Array.from(this.plugins.values())
  }

  /**
   * 从后端加载插件列表
   */
  async loadPlugins(): Promise<void> {
    try {
      const res = await adminRequest.get<PluginListResponse>({
        url: '/plugin/list'
      })
      
      this.plugins.clear()
      const list = res?.list || []
      list.forEach(plugin => {
        this.plugins.set(plugin.id, plugin)
      })
      this.loaded = true
    } catch (error) {
      console.warn('[PluginManager] Failed to load plugins:', error)
      this.loaded = true
    }
  }

  /**
   * 获取单个插件
   */
  getPlugin(id: string): PluginInfo | undefined {
    return this.plugins.get(id)
  }

  /**
   * 获取已启用的插件
   */
  async getEnabledPlugins(): Promise<PluginInfo[]> {
    const plugins = await this.getPlugins()
    return plugins.filter(p => p.status === 'enabled')
  }

  /**
   * 安装插件
   */
  async install(pluginId: string): Promise<void> {
    await adminRequest.post({
      url: '/plugin/install',
      params: { pluginId }
    })
    await this.loadPlugins()
  }

  /**
   * 卸载插件
   */
  async uninstall(pluginId: string): Promise<void> {
    await adminRequest.post({
      url: '/plugin/uninstall',
      params: { pluginId }
    })
    await this.loadPlugins()
  }

  /**
   * 启用插件
   */
  async enable(pluginId: string): Promise<void> {
    await adminRequest.post({
      url: '/plugin/enable',
      params: { pluginId }
    })
    await this.loadPlugins()
  }

  /**
   * 禁用插件
   */
  async disable(pluginId: string): Promise<void> {
    await adminRequest.post({
      url: '/plugin/disable',
      params: { pluginId }
    })
    await this.loadPlugins()
  }

  /**
   * 刷新插件列表
   */
  async refresh(): Promise<void> {
    this.loaded = false
    await this.loadPlugins()
  }

  /**
   * 清除缓存
   */
  clear(): void {
    this.plugins.clear()
    this.loaded = false
  }
}

// 导出单例
export const pluginManager = new PluginManager()
export default PluginManager
