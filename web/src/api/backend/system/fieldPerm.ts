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
 * 字段权限管理 API
 */
import { adminRequest } from '@/utils/http'

// 查询字段权限列表
export function getFieldPermListApi(params: {
  roleId: number
  module?: string
  resource?: string
}) {
  return adminRequest.get<any>({
    url: '/fieldPerm/list',
    params
  })
}

// 批量保存字段权限
export function batchSaveFieldPermApi(data: {
  roleId: number
  resource: string
  fields: Array<{
    fieldName: string
    fieldLabel: string
    permType: number
  }>
}) {
  return adminRequest.post<any>({
    url: '/fieldPerm/batchSave',
    data
  })
}

// 获取角色字段权限映射
export function getFieldPermByRoleApi(params: {
  roleId: number
  resource?: string
}) {
  return adminRequest.get<any>({
    url: '/fieldPerm/getByRole',
    params
  })
}

// 获取资源字段列表
export function getResourceFieldsApi(params: {
  resource: string
}) {
  return adminRequest.get<any>({
    url: '/fieldPerm/resourceFields',
    params
  })
}
