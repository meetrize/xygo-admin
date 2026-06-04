/**
 * 测试记帐管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchUranTestableList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/uran-testable/list',
    params
  })
}

/** 详情 */
export function fetchUranTestableView(id: number) {
  return adminRequest.get<any>({
    url: '/uran-testable/view',
    params: { id }
  })
}

/** 保存(新增/编辑) */
export function fetchUranTestableEdit(params: any) {
  return adminRequest.post<any>({
    url: '/uran-testable/edit',
    params
  })
}

/** 删除 */
export function fetchUranTestableDelete(id: number) {
  return adminRequest.post<any>({
    url: '/uran-testable/delete',
    params: { id }
  })
}

/** 导出 */
export function fetchUranTestableExport(params?: any) {
  return adminRequest.get<any>({
    url: '/uran-testable/export',
    params,
    responseType: 'blob'
  })
}
