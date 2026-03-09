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
 * 代码生成器 API
 */
import { adminRequest } from '@/utils/http'

/** 获取选项(生成类型/表单组件/查询方式) */
export function fetchGenCodesSelects() {
  return adminRequest.get<any>({ url: '/genCodes/selects' })
}

/** 获取数据库表列表 */
export function fetchGenCodesTableSelect() {
  return adminRequest.get<any>({ url: '/genCodes/tableSelect' })
}

/** 获取表字段列表 */
export function fetchGenCodesColumnList(tableName: string) {
  return adminRequest.get<any>({ url: '/genCodes/columnList', params: { tableName } })
}

/** 生成记录列表 */
export function fetchGenCodesList(params: any) {
  return adminRequest.get<any>({ url: '/genCodes/list', params })
}

/** 获取配置详情 */
export function fetchGenCodesView(id: number) {
  return adminRequest.get<any>({ url: '/genCodes/view', params: { id } })
}

/** 保存生成配置 */
export function fetchGenCodesEdit(params: any) {
  return adminRequest.post<any>({ url: '/genCodes/edit', params })
}

/** 删除生成配置 */
export function fetchGenCodesDelete(params: { id: number; deleteFiles?: boolean; deleteMenus?: boolean }) {
  return adminRequest.post<any>({ url: '/genCodes/delete', params })
}

/** 预览代码 */
export function fetchGenCodesPreview(params: any) {
  return adminRequest.post<any>({ url: '/genCodes/preview', params })
}

/** 执行生成 */
export function fetchGenCodesBuild(params: any) {
  return adminRequest.post<any>({ url: '/genCodes/build', params })
}

/** 发布前端文件（从临时目录移到正式目录） */
export function fetchGenCodesPublishFrontend() {
  return adminRequest.post<any>({ url: '/genCodes/publishFrontend' })
}

/** 预览字段变更（对比设计器与数据库） */
export function fetchGenCodesSyncFields(params: any) {
  return adminRequest.post<any>({ url: '/genCodes/syncFields', params })
}

/** 执行字段同步DDL */
export function fetchGenCodesExecuteDDL(params: { tableName: string; sqls: string[] }) {
  return adminRequest.post<any>({ url: '/genCodes/executeDDL', params })
}

/** 创建数据表 */
export function fetchGenCodesCreateTable(params: any) {
  return adminRequest.post<any>({ url: '/genCodes/createTable', params })
}
