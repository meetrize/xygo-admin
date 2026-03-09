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
 * 后台会员管理 API
 */
import { adminRequest } from '@/utils/http'

/** 会员列表项 */
export interface MemberItem {
  id: number
  username: string
  nickname: string
  mobile: string
  email: string
  avatar: string
  gender: number
  level: number
  groupId: number
  groupName: string
  score: number
  money: number
  status: number
  loginCount: number
  lastLoginAt: string
  lastLoginIp: string
  createdAt: string
}

/** 会员列表查询参数 */
export interface MemberListParams {
  page?: number
  pageSize?: number
  username?: string
  mobile?: string
  email?: string
  status?: number
  groupId?: number
}

/** 会员列表响应 */
export interface MemberListResult {
  list: MemberItem[]
  total: number
  page: number
  pageSize: number
}

/** 添加会员参数 */
export interface MemberAddParams {
  username: string
  password: string
  nickname?: string
  mobile?: string
  email?: string
  avatar?: string
  gender?: number
  groupId?: number
  score?: number
  money?: number
  status?: number
  remark?: string
}

/** 编辑会员参数 */
export interface MemberEditParams {
  id: number
  username?: string
  password?: string
  nickname?: string
  mobile?: string
  email?: string
  avatar?: string
  gender?: number
  groupId?: number
  score?: number
  money?: number
  status?: number
  remark?: string
}

/** 会员分组选项 */
export interface MemberGroupOption {
  id: number
  name: string
}

/**
 * 获取会员列表
 */
export function getMemberList(params: MemberListParams) {
  return adminRequest.get<MemberListResult>({
    url: '/admin/member/list',
    params
  })
}

/**
 * 获取会员详情
 */
export function getMemberDetail(id: number) {
  return adminRequest.get<MemberItem>({
    url: '/admin/member/detail',
    params: { id }
  })
}

/**
 * 添加会员
 */
export function addMember(data: MemberAddParams) {
  return adminRequest.post<{ id: number }>({
    url: '/admin/member/add',
    data
  })
}

/**
 * 编辑会员
 */
export function editMember(data: MemberEditParams) {
  return adminRequest.put<void>({
    url: '/admin/member/edit',
    data
  })
}

/**
 * 删除会员
 */
export function deleteMember(ids: number[]) {
  return adminRequest.del<void>({
    url: '/admin/member/delete',
    data: { ids }
  })
}

/**
 * 修改会员状态
 */
export function updateMemberStatus(id: number, status: number) {
  return adminRequest.put<void>({
    url: '/admin/member/status',
    data: { id, status }
  }) 
}

/**
 * 重置会员密码
 */
export function resetMemberPassword(id: number, password: string) {
  return adminRequest.put<void>({
    url: '/admin/member/resetPassword',
    data: { id, password }
  })
}

/**
 * 获取会员分组选项
 */
export function getMemberGroupOptions() {
  return adminRequest.get<{ list: MemberGroupOption[] }>({
    url: '/admin/member/groupOptions'
  })
}
