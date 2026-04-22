/**
 * 悠然联系人管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchUranContactList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/uran-contact/list',
    params
  })
}

/** 详情 */
export function fetchUranContactView(id: number) {
  return adminRequest.get<any>({
    url: '/uran-contact/view',
    params: { id }
  })
}

/** 保存(新增/编辑) */
export function fetchUranContactEdit(params: any) {
  return adminRequest.post<any>({
    url: '/uran-contact/edit',
    params
  })
}

/** 删除 */
export function fetchUranContactDelete(id: number) {
  return adminRequest.post<any>({
    url: '/uran-contact/delete',
    params: { id }
  })
}

/** 导出 */
export function fetchUranContactExport(params?: any) {
  return adminRequest.get<any>({
    url: '/uran-contact/export',
    params,
    responseType: 'blob'
  })
}
