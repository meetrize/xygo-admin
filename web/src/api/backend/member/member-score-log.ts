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
 * 积分变动日志管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchMemberScoreLogList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/member-score-log/list',
    params
  })
}

/** 详情 */
export function fetchMemberScoreLogView(id: number) {
  return adminRequest.get<any>({
    url: '/member-score-log/view',
    params: { id }
  })
}

/** 保存(新增/编辑) */
export function fetchMemberScoreLogEdit(params: any) {
  return adminRequest.post<any>({
    url: '/member-score-log/edit',
    params
  })
}

/** 删除 */
export function fetchMemberScoreLogDelete(id: number) {
  return adminRequest.post<any>({
    url: '/member-score-log/delete',
    params: { id }
  })
}
