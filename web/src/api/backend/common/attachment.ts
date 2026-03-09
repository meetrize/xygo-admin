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
 * 附件管理 API
 * @module api/backend/common/attachment
 */
import { adminRequest } from '@/utils/http'

/**
 * 获取附件列表
 */
export function fetchAttachmentList(params: {
  page: number
  pageSize: number
  topic?: string
  storage?: string
}) {
  return adminRequest.get<{
    list: any[]
    page: number
    pageSize: number
    total: number
  }>({
    url: '/attachment/list',
    params
  })
}

/**
 * 删除附件
 */
export function fetchDeleteAttachment(id: number) {
  return adminRequest.post({
    url: '/attachment/delete',
    params: { id }
  })
}
