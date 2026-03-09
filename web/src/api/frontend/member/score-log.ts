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
 * 前台会员积分记录 API
 */
import { memberRequest } from '@/utils/http'

/** 积分记录项 */
export interface ScoreLogItem {
  id: number
  score: number
  before: number
  after: number
  memo: string
  createdAt: string
}

/** 积分记录列表响应 */
export interface ScoreLogListResult {
  list: ScoreLogItem[]
  page: number
  pageSize: number
  total: number
}

/** 获取积分记录列表 */
export function getScoreLogList(params?: { page?: number; pageSize?: number }) {
  return memberRequest.get<ScoreLogListResult>({
    url: '/user/score/log',
    params
  })
}
