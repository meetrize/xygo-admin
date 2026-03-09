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
import { WEB_LINKS } from '@/utils/constants'

export const helpRoutes: AppRouteRecord[] = [
  {
    name: 'Document',
    path: '',
    component: '',
    meta: {
      title: 'menus.help.document',
      icon: 'ri:bill-line',
      link: WEB_LINKS.DOCS,
      isIframe: false,
      keepAlive: false
    }
  },
  {
    name: 'LiteVersion',
    path: '',
    component: '',
    meta: {
      title: 'menus.help.liteVersion',
      icon: 'ri:bus-2-line',
      link: WEB_LINKS.LiteVersion,
      isIframe: false,
      keepAlive: false
    }
  },
  {
    name: 'OldVersion',
    path: '',
    component: '',
    meta: {
      title: 'menus.help.oldVersion',
      icon: 'ri:subway-line',
      link: WEB_LINKS.OldVersion,
      isIframe: false,
      keepAlive: false
    }
  },
  {
    name: 'ChangeLog',
    path: '/change/log',
    component: '/change/log',
    meta: {
      title: 'menus.plan.log',
      showTextBadge: `v${__APP_VERSION__}`,
      icon: 'ri:gamepad-line',
      keepAlive: false
    }
  }
]
