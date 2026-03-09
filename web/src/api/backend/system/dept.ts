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
 * 部门管理 API
 * @module api/backend/system/dept
 */
import { adminRequest } from '@/utils/http'

/**
 * 获取部门列表（树形）
 */
export async function fetchGetDeptList(params: any) {
  const res = await adminRequest.get<{ list?: any[] }>({
    url: '/dept/list',
    params
  })
  const list = (res as any)?.list ?? res
  return (Array.isArray(list) ? list : []) as any[]
}

/**
 * 保存部门（新增/编辑）
 */
export function fetchSaveDept(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/dept/save',
    params
  })
}

/**
 * 删除部门
 */
export function fetchDeleteDept(id: number) {
  return adminRequest.post({
    url: '/dept/delete',
    params: { id }
  })
}
