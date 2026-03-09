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
 * 会员通知管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchMemberNoticeList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/member-notice/list',
    params
  })
}

/** 详情 */
export function fetchMemberNoticeView(id: number) {
  return adminRequest.get<any>({
    url: '/member-notice/view',
    params: { id }
  })
}

/** 保存(新增/编辑) */
export function fetchMemberNoticeEdit(params: any) {
  return adminRequest.post<any>({
    url: '/member-notice/edit',
    params
  })
}

/** 删除 */
export function fetchMemberNoticeDelete(id: number) {
  return adminRequest.post<any>({
    url: '/member-notice/delete',
    params: { id }
  })
}

/** 导出 */
export function fetchMemberNoticeExport(params?: any) {
  return adminRequest.get<any>({
    url: '/member-notice/export',
    params,
    responseType: 'blob'
  })
}
