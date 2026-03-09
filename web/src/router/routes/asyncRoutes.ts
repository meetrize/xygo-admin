// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// 权限文档：https://www.artd.pro/docs/zh/guide/in-depth/permission.html
import { AppRouteRecord } from '@/types/router'
import { routeModules } from '../modules'

/**
 * 动态路由（需要权限才能访问的路由）
 * 用于渲染菜单以及根据菜单权限动态加载路由，如果没有权限无法访问
 */
export const asyncRoutes: AppRouteRecord[] = routeModules
