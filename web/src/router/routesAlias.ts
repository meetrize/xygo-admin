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
