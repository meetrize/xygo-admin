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
 * 后台操作日志 API
 */
import { adminRequest } from '@/utils/http'

/** 操作日志列表项 */
export interface OperationLogItem {
  id: number
  userId: number
  username: string
  module: string
  title: string
  method: string
  url: string
  ip: string
  location: string
  requestBody: string
  responseBody: string
  errorMessage: string
  status: number
  elapsed: number
  createdAt: string
}

/** 操作日志列表查询参数 */
export interface OperationLogListParams {
  page?: number
  pageSize?: number
  username?: string
  module?: string
  status?: number
  dateRange?: string[]
}

/** 操作日志列表响应 */
export interface OperationLogListResult {
  list: OperationLogItem[]
  total: number
  page: number
  pageSize: number
}

/**
 * 获取操作日志列表
 */
export function getOperationLogList(params: OperationLogListParams) {
  return adminRequest.post<OperationLogListResult>({
    url: '/log/operation/list',
    data: params
  })
}

/**
 * 获取操作日志详情
 */
export function getOperationLogDetail(id: number) {
  return adminRequest.get<OperationLogItem>({
    url: '/log/operation/detail',
    params: { id }
  })
}

/**
 * 删除操作日志
 */
export function deleteOperationLog(ids: number[]) {
  return adminRequest.post<void>({
    url: '/log/operation/delete',
    data: { ids }
  })
}

/**
 * 清空操作日志
 */
export function clearOperationLog() {
  return adminRequest.post<void>({
    url: '/log/operation/clear'
  })
}
