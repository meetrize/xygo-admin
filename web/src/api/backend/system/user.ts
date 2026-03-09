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
 * 用户管理 API
 * @module api/backend/system/user
 */
import { adminRequest } from '@/utils/http'

/**
 * 获取用户列表
 */
export function fetchGetUserList(params: any) {
  return adminRequest.get<{ 
    list: Api.SystemManage.UserListItem[]
    page: number
    pageSize: number
    total: number 
  }>({
    url: '/user/list',
    params
  })
}

/**
 * 获取用户详情（编辑用，未脱敏）
 */
export function fetchGetUserDetail(id: number) {
  return adminRequest.get<any>({
    url: '/user/detail',
    params: { id }
  })
}

/**
 * 保存用户（新增/编辑）
 */
export function fetchSaveUser(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/user/save',
    params
  })
}

/**
 * 删除用户
 */
export function fetchDeleteUser(id: number) {
  return adminRequest.post({
    url: '/user/delete',
    params: { id }
  })
}

/**
 * 强制用户下线
 */
export function fetchKickUser(id: number) {
  return adminRequest.post({
    url: '/user/kick',
    params: { id }
  })
}
