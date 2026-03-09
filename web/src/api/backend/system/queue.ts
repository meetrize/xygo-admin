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
 * 消息队列 API
 */
import { adminRequest } from '@/utils/http'

export interface TopicStats {
  topic: string
  pending: number
  deadSize: number
}

/** 获取队列统计 */
export function fetchQueueStats() {
  return adminRequest.get<{ driver: string; topics: TopicStats[] }>({ url: '/queue/stats' })
}

/** 获取已注册 Topic 列表 */
export function fetchQueueTopics() {
  return adminRequest.get<{ list: string[] }>({ url: '/queue/topics' })
}

/** 测试投递消息（delaySec>0 为延迟投递） */
export function fetchQueuePushTest(params: { topic: string; body: string; delaySec?: number }) {
  return adminRequest.post({ url: '/queue/pushTest', params })
}
