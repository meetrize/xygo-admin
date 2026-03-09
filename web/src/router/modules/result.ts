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

export const resultRoutes: AppRouteRecord = {
  path: '/result',
  name: 'Result',
  component: '/index/index',
  meta: {
    title: 'menus.result.title',
    icon: 'ri:checkbox-circle-line'
  },
  children: [
    {
      path: 'success',
      name: 'ResultSuccess',
      component: '/common/result/success',
      meta: {
        title: 'menus.result.success',
        icon: 'ri:checkbox-circle-line',
        keepAlive: true
      }
    },
    {
      path: 'fail',
      name: 'ResultFail',
      component: '/common/result/fail',
      meta: {
        title: 'menus.result.fail',
        icon: 'ri:close-circle-line',
        keepAlive: true
      }
    }
  ]
}
