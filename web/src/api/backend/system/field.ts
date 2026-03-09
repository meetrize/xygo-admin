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
 * 字段权限 API
 * @module api/backend/system/field
 */
import { adminRequest } from '@/utils/http'

/**
 * 获取所有资源列表（用于菜单resource字段选择）
 */
export function fetchResourceList() {
  return adminRequest.get<{ list: Array<{ code: string; label: string }> }>({
    url: '/field/resourceList'
  })
}
