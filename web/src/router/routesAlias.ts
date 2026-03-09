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
 * 后台路由基础路径（通过 VITE_ADMIN_PATH 环境变量可配置）
 * 所有后台管理页面都在此前缀下，前后台路由靠路径前缀严格隔离
 */
export const ADMIN_BASE_PATH: string = import.meta.env.VITE_ADMIN_PATH || '/admin'

/** 后台登录页路径（跟随 ADMIN_BASE_PATH 动态计算） */
export const ADMIN_LOGIN_PATH = `${ADMIN_BASE_PATH}/login`

/**
 * 公共路由别名
 * 存放系统级公共路由路径，如布局容器等
 */
export enum RoutesAlias {
  Layout = '/index/index', // 布局容器
}
