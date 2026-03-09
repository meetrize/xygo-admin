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
 * 前台会员通知 API
 */
import { memberRequest } from '@/utils/http'

/** 通知项 */
export interface NoticeItem {
  id: number
  title: string
  content: string
  type: string
  sender: string
  isRead: boolean
  createdAt: string
}

/** 通知列表响应 */
export interface NoticeListResult {
  list: NoticeItem[]
  page: number
  pageSize: number
  total: number
  unread: number
}

/** 获取通知列表 */
export function getNoticeList(params?: { page?: number; pageSize?: number }) {
  return memberRequest.get<NoticeListResult>({
    url: '/user/notice/list',
    params
  })
}

/** 标记通知已读 */
export function markNoticeRead(noticeId: number) {
  return memberRequest.post<void>({
    url: '/user/notice/read',
    data: { noticeId }
  })
}

/** 全部已读 */
export function markAllNoticeRead() {
  return memberRequest.post<void>({
    url: '/user/notice/read-all'
  })
}
