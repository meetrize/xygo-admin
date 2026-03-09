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
 * 后台登录日志 API
 */
import { adminRequest } from '@/utils/http'

/** 登录日志列表项 */
export interface LoginLogItem {
  id: number
  userId: number
  username: string
  ip: string
  location: string
  browser: string
  os: string
  status: number
  message: string
  createdAt: string
}

/** 登录日志列表查询参数 */
export interface LoginLogListParams {
  page?: number
  pageSize?: number
  username?: string
  ip?: string
  status?: number
  dateRange?: string[]
}

/** 登录日志列表响应 */
export interface LoginLogListResult {
  list: LoginLogItem[]
  total: number
  page: number
  pageSize: number
}

/**
 * 获取登录日志列表
 */
export function getLoginLogList(params: LoginLogListParams) {
  return adminRequest.post<LoginLogListResult>({
    url: '/log/login/list',
    data: params
  })
}

/**
 * 删除登录日志
 */
export function deleteLoginLog(ids: number[]) {
  return adminRequest.post<void>({
    url: '/log/login/delete',
    data: { ids }
  })
}

/**
 * 清空登录日志
 */
export function clearLoginLog() {
  return adminRequest.post<void>({
    url: '/log/login/clear'
  })
}
