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
 * 后台会员菜单管理 API（对齐 BuildAdmin user_rule）
 */
import { adminRequest } from '@/utils/http'

/** 规则类型 */
export type RuleType = 'route' | 'menu_dir' | 'menu' | 'nav' | 'nav_user_menu' | 'button'

/** 菜单打开方式 */
export type MenuOpenType = 'tab' | 'link' | 'iframe'

/** 会员菜单树节点 */
export interface MemberMenuItem {
  id: number
  pid: number
  title: string
  name: string
  path: string
  component: string
  icon: string
  menuType: MenuOpenType
  url: string
  noLoginValid: number
  extend: string
  remark: string
  type: RuleType
  permission: string
  sort: number
  status: number
  createdAt: string
  updatedAt: string
  children?: MemberMenuItem[]
}

/** 会员菜单树查询参数 */
export interface MemberMenuTreeParams {
  status?: number
}

/** 会员菜单树响应 */
export interface MemberMenuTreeResult {
  list: MemberMenuItem[]
}

/** 保存会员菜单参数 */
export interface MemberMenuSaveParams {
  id?: number
  pid?: number
  title: string
  name?: string
  path?: string
  component?: string
  icon?: string
  menuType?: MenuOpenType
  url?: string
  noLoginValid?: number
  extend?: string
  remark?: string
  type: RuleType
  permission?: string
  sort?: number
  status?: number
}

/**
 * 获取会员菜单树
 */
export function getMemberMenuTree(params?: MemberMenuTreeParams) {
  return adminRequest.get<MemberMenuTreeResult>({
    url: '/member/menu/tree',
    params
  })
}

/**
 * 保存会员菜单
 */
export function saveMemberMenu(data: MemberMenuSaveParams) {
  return adminRequest.post<{ id: number }>({
    url: '/member/menu/save',
    data
  })
}

/**
 * 删除会员菜单
 */
export function deleteMemberMenu(id: number) {
  return adminRequest.post<void>({
    url: '/member/menu/delete',
    data: { id }
  })
}
