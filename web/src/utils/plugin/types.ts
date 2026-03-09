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
 * 插件系统类型定义
 * @module utils/plugin/types
 */

/**
 * 插件状态
 */
export type PluginStatus = 'installed' | 'enabled' | 'disabled' | 'uninstalled'

/**
 * 菜单项
 */
export interface PluginMenuItem {
  title: string
  path: string
  icon?: string
  sort?: number
  children?: PluginMenuItem[]
}

/**
 * 权限项
 */
export interface PluginPermission {
  name: string
  title: string
}

/**
 * 插件信息
 */
export interface PluginInfo {
  id: string
  name: string
  version: string
  description?: string
  author?: string
  homepage?: string
  license?: string
  requires?: {
    xygo?: string
  }
  dependencies?: string[]
  menus?: PluginMenuItem[]
  permissions?: PluginPermission[]
  status?: PluginStatus
}

/**
 * 插件列表响应
 */
export interface PluginListResponse {
  list: PluginInfo[]
}
