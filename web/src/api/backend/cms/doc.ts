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
 * CMS 文档内容 API
 * @module api/backend/cms/doc
 */
import { adminRequest } from '@/utils/http'

/** 获取文档列表（分页） */
export function fetchDocList(params: any) {
  return adminRequest.post<any>({
    url: '/cms/doc/list',
    params
  })
}

/** 获取文档详情 */
export function fetchDocDetail(id: number) {
  return adminRequest.get<any>({
    url: '/cms/doc/detail',
    params: { id }
  })
}

/** 保存文档（新增/编辑） */
export function fetchSaveDoc(params: any) {
  return adminRequest.post<{ id: number }>({
    url: '/cms/doc/save',
    params
  })
}

/** 删除文档 */
export function fetchDeleteDoc(id: number) {
  return adminRequest.post({
    url: '/cms/doc/delete',
    params: { id }
  })
}
