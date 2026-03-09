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
 * 前台会员签到 API
 */
import { memberRequest } from '@/utils/http'

/** 签到日历天 */
export interface CheckinDayItem {
  date: string
  checked: boolean
  score: number
}

/** 签到信息 */
export interface CheckinInfo {
  continuousDays: number
  todayChecked: boolean
  todayScore: number
  weekDays: CheckinDayItem[]
}

/** 签到结果 */
export interface CheckinResult {
  score: number
  continuousDays: number
}

/** 获取签到信息 */
export function getCheckinInfo() {
  return memberRequest.get<CheckinInfo>({
    url: '/user/checkin/info'
  })
}

/** 执行签到 */
export function doCheckin() {
  return memberRequest.post<CheckinResult>({
    url: '/user/checkin'
  })
}
