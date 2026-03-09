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
 * 路由工具函数
 *
 * 提供路由相关的工具函数
 *
 * @module utils/router
 */
import { RouteLocationNormalized, RouteRecordRaw } from 'vue-router'
import AppConfig from '@/config'
import { useSiteStore } from '@/store/modules/site'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import i18n, { $t } from '@/locales'

/** 扩展的路由配置类型 */
export type AppRouteRecordRaw = RouteRecordRaw & {
  hidden?: boolean
}

/** 顶部进度条配置 */
export const configureNProgress = () => {
  NProgress.configure({
    easing: 'ease',
    speed: 600,
    showSpinner: false,
    parent: 'body'
  })
}

/**
 * 设置页面标题，根据路由元信息和系统信息拼接标题
 * @param to 当前路由对象
 */
export const setPageTitle = (to: RouteLocationNormalized): void => {
  const { title } = to.meta
  if (title) {
    setTimeout(() => {
      // 动态获取系统名称
      const siteStore = useSiteStore()
      const systemName = siteStore.getSiteName() || AppConfig.systemInfo.name
      document.title = `${formatMenuTitle(String(title))} - ${systemName}`
    }, 150)
  }
}

/**
 * 格式化菜单标题
 * @param title 菜单标题，可以是 i18n 的 key，也可以是字符串
 * @returns 格式化后的菜单标题
 */
export const formatMenuTitle = (title: string): string => {
  if (title) {
    if (title.startsWith('menus.')) {
      // 使用 te() 方法检查翻译键值是否存在，避免控制台警告
      if (i18n.global.te(title)) {
        return $t(title)
      } else {
        // 如果翻译不存在，返回键值的最后部分作为fallback
        return title.split('.').pop() || title
      }
    }
    return title
  }
  return ''
}
