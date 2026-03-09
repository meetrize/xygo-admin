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
 * 前台会员余额记录 API
 */
import { memberRequest } from '@/utils/http'

/** 余额记录项 */
export interface MoneyLogItem {
  id: number
  money: number
  before: number
  after: number
  memo: string
  createdAt: string
}

/** 余额记录列表响应 */
export interface MoneyLogListResult {
  list: MoneyLogItem[]
  page: number
  pageSize: number
  total: number
}

/** 获取余额记录列表 */
export function getMoneyLogList(params?: { page?: number; pageSize?: number }) {
  return memberRequest.get<MoneyLogListResult>({
    url: '/user/money/log',
    params
  })
}
