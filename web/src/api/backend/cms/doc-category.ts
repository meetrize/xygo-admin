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
 * CMS 文档分类 API
 * @module api/backend/cms/doc-category
 */
import { adminRequest } from '@/utils/http'

/** 获取文档分类列表（树形） */
export async function fetchDocCategoryList(params?: any) {
  const res = await adminRequest.post<{ list?: any[] }>({
    url: '/cms/docCategory/list',
    params
  })
  const list = (res as any)?.list ?? res
  return (Array.isArray(list) ? list : []) as any[]
}

/** 保存文档分类（新增/编辑） */
export function fetchSaveDocCategory(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/cms/docCategory/save',
    params
  })
}

/** 删除文档分类 */
export function fetchDeleteDocCategory(id: number) {
  return adminRequest.post({
    url: '/cms/docCategory/delete',
    params: { id }
  })
}
