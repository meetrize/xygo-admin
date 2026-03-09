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
 * 菜单管理 API
 * @module api/backend/system/menu
 */
import { adminRequest } from '@/utils/http'
import { AppRouteRecord } from '@/types/router'

/**
 * 获取菜单路由（前端使用，按角色过滤）
 */
export async function fetchGetMenuList() {
  const res = await adminRequest.get<{ list?: AppRouteRecord[] }>({
    url: '/menu/routes'
  })

  const list = (res as any)?.list ?? res
  return (Array.isArray(list) ? list : []) as AppRouteRecord[]
}

/**
 * 获取菜单树（管理后台菜单配置）
 */
export async function fetchGetMenuTree() {
  const res = await adminRequest.get<{ list?: any[] }>({
    url: '/menu/tree'
  })

  const list = (res as any)?.list ?? res
  return (Array.isArray(list) ? list : []) as any[]
}

/**
 * 保存菜单（新增/编辑）
 */
export function fetchSaveMenu(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/menu/save',
    params
  })
}

/**
 * 删除菜单
 */
export function fetchDeleteMenu(id: number) {
  return adminRequest.post({
    url: '/menu/delete',
    params: { id }
  })
}
