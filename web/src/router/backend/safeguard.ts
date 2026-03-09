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

export const safeguardRoutes: AppRouteRecord = {
  path: '/safeguard',
  name: 'Safeguard',
  component: '/index/index',
  meta: {
    title: 'menus.safeguard.title',
    icon: 'ri:shield-check-line',
    keepAlive: false
  },
  children: [
    {
      path: 'server',
      name: 'SafeguardServer',
      component: '/safeguard/server',
      meta: {
        title: 'menus.safeguard.server',
        icon: 'ri:hard-drive-3-line',
        keepAlive: true
      }
    },
    {
      path: 'login-log',
      name: 'LoginLog',
      component: '/safeguard/login-log',
      meta: {
        title: 'menus.safeguard.loginLog',
        icon: 'ri:login-box-line',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
    {
      path: 'operation-log',
      name: 'OperationLog',
      component: '/safeguard/operation-log',
      meta: {
        title: 'menus.safeguard.operationLog',
        icon: 'ri:file-text-line',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    },
    {
      path: 'performance',
      name: 'SafeguardPerformance',
      component: '/safeguard/performance',
      meta: {
        title: 'menus.safeguard.performance',
        icon: 'ri:line-chart-line',
        keepAlive: true,
        roles: ['R_SUPER']
      }
    }
  ]
}
