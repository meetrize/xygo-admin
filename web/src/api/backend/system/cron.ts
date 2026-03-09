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
 * 定时任务 API
 */
import { adminRequest } from '@/utils/http'

// ==================== 类型定义 ====================

export interface CronItem {
  id: number
  groupId: number
  groupName: string
  title: string
  name: string
  params: string
  pattern: string
  policy: number
  count: number
  sort: number
  remark: string
  status: number
  createdAt: number
  updatedAt: number
}

export interface CronGroupItem {
  id: number
  name: string
  sort: number
  remark: string
  status: number
  createdAt: number
  updatedAt: number
}

export interface CronLogItem {
  id: number
  cronId: number
  name: string
  title: string
  params: string
  status: number
  output: string
  errMsg: string
  takeMs: number
  createdAt: number
}

// ==================== 定时任务接口 ====================

export function fetchCronList(params?: any) {
  return adminRequest.get<{ list: CronItem[]; page: number; pageSize: number; total: number }>({
    url: '/cron/list',
    params
  })
}

export function fetchCronSave(params: any) {
  return adminRequest.post<{ id: number }>({ url: '/cron/save', params })
}

export function fetchCronDelete(id: number) {
  return adminRequest.post({ url: '/cron/delete', params: { id } })
}

export function fetchCronStatus(params: { id: number; status: number }) {
  return adminRequest.post({ url: '/cron/status', params })
}

export function fetchCronOnlineExec(id: number) {
  return adminRequest.post<{ output: string }>({ url: '/cron/onlineExec', params: { id } })
}

export function fetchCronRegisteredTasks() {
  return adminRequest.get<{ list: string[] }>({ url: '/cron/registeredTasks' })
}

// ==================== 分组接口 ====================

export function fetchCronGroupList(params?: any) {
  return adminRequest.get<{ list: CronGroupItem[] }>({ url: '/cronGroup/list', params })
}

export function fetchCronGroupSave(params: any) {
  return adminRequest.post<{ id: number }>({ url: '/cronGroup/save', params })
}

export function fetchCronGroupDelete(id: number) {
  return adminRequest.post({ url: '/cronGroup/delete', params: { id } })
}

export function fetchCronGroupSelect() {
  return adminRequest.get<{ list: { id: number; name: string }[] }>({ url: '/cronGroup/select' })
}

// ==================== 执行日志接口 ====================

export function fetchCronLogList(params?: any) {
  return adminRequest.get<{ list: CronLogItem[]; page: number; pageSize: number; total: number }>({
    url: '/cronLog/list',
    params
  })
}

export function fetchCronLogClear(cronId?: number) {
  return adminRequest.post({ url: '/cronLog/clear', params: { cronId } })
}
