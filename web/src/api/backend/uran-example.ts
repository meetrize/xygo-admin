/**
 * 悠然示例管理 API
 */
import { adminRequest } from '@/utils/http'

/** 列表 */
export function fetchUranExampleList(params: any) {
  return adminRequest.get<Record<string, any>>({
    url: '/uran-example/list',
    params
  })
}

/** 详情 */
export function fetchUranExampleView(id: number) {
  return adminRequest.get<any>({
    url: '/uran-example/view',
    params: { id }
  })
}

/** 保存(新增/编辑) */
export function fetchUranExampleEdit(params: any) {
  return adminRequest.post<any>({
    url: '/uran-example/edit',
    params
  })
}

/** 删除 */
export function fetchUranExampleDelete(id: number) {
  return adminRequest.post<any>({
    url: '/uran-example/delete',
    params: { id }
  })
}

/** 导出 */
export function fetchUranExampleExport(params?: any) {
  return adminRequest.get<any>({
    url: '/uran-example/export',
    params,
    responseType: 'blob'
  })
}
