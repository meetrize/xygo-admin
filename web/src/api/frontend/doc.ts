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
 * 前台文档公开 API（无需登录）
 * @module api/frontend/doc
 */
import { siteRequest } from '@/utils/http'

/** 获取文档分类树 */
export async function fetchDocCategoryTree() {
  const res = await siteRequest.get<{ list?: any[] }>({
    url: '/doc/categoryTree'
  })
  const list = (res as any)?.list ?? res
  return (Array.isArray(list) ? list : []) as any[]
}

/** 按分类获取文档列表 */
export async function fetchDocListByCategory(categoryId?: number) {
  const res = await siteRequest.get<{ list?: any[] }>({
    url: '/doc/list',
    params: categoryId ? { categoryId } : {}
  })
  const list = (res as any)?.list ?? res
  return (Array.isArray(list) ? list : []) as any[]
}

/** 按 slug 获取文档详情 */
export function fetchDocDetailBySlug(slug: string) {
  return siteRequest.get<any>({
    url: '/doc/detail',
    params: { slug }
  })
}

/** 全文搜索文档 */
export async function fetchDocSearch(keyword: string) {
  const res = await siteRequest.get<{ list?: any[] }>({
    url: '/doc/search',
    params: { keyword }
  })
  const list = (res as any)?.list ?? res
  return (Array.isArray(list) ? list : []) as any[]
}
