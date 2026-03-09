// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

import { AppRouteRecord } from '@/types/router'

export const developRoutes: AppRouteRecord = {
  path: '/develop',
  name: 'Develop',
  component: '/index/index',
  meta: {
    title: 'menus.develop.title',
    icon: 'ri:code-box-line',
    roles: ['R_SUPER']
  },
  children: [
    {
      path: 'gen-codes',
      name: 'GenCodes',
      component: '/develop/gen-codes/index',
      meta: {
        title: 'menus.develop.genCodes',
        icon: 'ri:magic-line',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    }
  ]
}
