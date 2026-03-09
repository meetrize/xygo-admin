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
 * 后台 Store 模块统一导出
 * 
 * BuildAdmin 风格：所有后台状态管理统一从此处导出
 */

// 导出各 Store
export { useUserStore } from './user'
export { useMenuStore } from './menu'
export { useSettingStore } from './setting'
export { useWorktabStore } from './worktab'
export { useTableStore } from './table'
export { useSiteStore } from './site'
