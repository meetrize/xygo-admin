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
 * 示例分类管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchDemoCategoryList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/demo-category/list',
    params
  })
}

/** 详情 */
export function fetchDemoCategoryView(id: number) {
  return adminRequest.get<any>({
    url: '/demo-category/view',
    params: { id }
  })
}

/** 保存(新增/编辑) */
export function fetchDemoCategoryEdit(params: any) {
  return adminRequest.post<any>({
    url: '/demo-category/edit',
    params
  })
}

/** 删除 */
export function fetchDemoCategoryDelete(id: number) {
  return adminRequest.post<any>({
    url: '/demo-category/delete',
    params: { id }
  })
}

/** 导出 */
export function fetchDemoCategoryExport(params?: any) {
  return adminRequest.get<any>({
    url: '/demo-category/export',
    params,
    responseType: 'blob'
  })
}
