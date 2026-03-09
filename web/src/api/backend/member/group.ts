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
 * 后台会员分组管理 API
 */
import { adminRequest } from '@/utils/http'

/** 会员分组列表项 */
export interface MemberGroupItem {
  id: number
  name: string
  rules: string
  status: number
  sort: number
  remark: string
  createdAt: string
  updatedAt: string
}

/** 会员分组列表查询参数 */
export interface MemberGroupListParams {
  page?: number
  pageSize?: number
  name?: string
  status?: number
}

/** 会员分组列表响应 */
export interface MemberGroupListResult {
  list: MemberGroupItem[]
  total: number
  page: number
  pageSize: number
}

/** 保存会员分组参数 */
export interface MemberGroupSaveParams {
  id?: number
  name: string
  rules?: string
  sort?: number
  status?: number
  remark?: string
}

/**
 * 获取会员分组列表
 */
export function getMemberGroupList(params: MemberGroupListParams) {
  return adminRequest.get<MemberGroupListResult>({
    url: '/member/group/list',
    params
  })
}

/**
 * 保存会员分组
 */
export function saveMemberGroup(data: MemberGroupSaveParams) {
  return adminRequest.post<{ id: number }>({
    url: '/member/group/save',
    data
  })
}

/**
 * 删除会员分组
 */
export function deleteMemberGroup(id: number) {
  return adminRequest.post<void>({
    url: '/member/group/delete',
    data: { id }
  })
}
