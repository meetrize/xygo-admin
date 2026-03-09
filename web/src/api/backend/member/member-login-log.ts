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
 * 登录日志管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchMemberLoginLogList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/member-login-log/list',
    params
  })
}

/** 详情 */
export function fetchMemberLoginLogView(id: number) {
  return adminRequest.get<any>({
    url: '/member-login-log/view',
    params: { id }
  })
}

/** 删除 */
export function fetchMemberLoginLogDelete(id: number) {
  return adminRequest.post<any>({
    url: '/member-login-log/delete',
    params: { id }
  })
}
