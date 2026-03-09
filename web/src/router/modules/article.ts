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

export const cmsRoutes: AppRouteRecord = {
  path: '/cms',
  name: 'Cms',
  component: '/index/index',
  meta: {
    title: '内容管理',
    icon: 'ri:article-line',
    roles: ['R_SUPER', 'R_ADMIN']
  },
  children: [
    {
      path: 'doc-category',
      name: 'CmsDocCategory',
      component: '/cms/doc-category',
      meta: {
        title: '文档分类',
        icon: 'ri:folder-line',
        keepAlive: true,
        authList: [
          { title: '新增/编辑', authMark: 'save' },
          { title: '删除', authMark: 'delete' }
        ]
      }
    },
    {
      path: 'doc',
      name: 'CmsDoc',
      component: '/cms/doc',
      meta: {
        title: '文档管理',
        icon: 'ri:file-text-line',
        keepAlive: true,
        authList: [
          { title: '新增/编辑', authMark: 'save' },
          { title: '删除', authMark: 'delete' }
        ]
      }
    }
  ]
}

// 兼容旧名称
export const articleRoutes = cmsRoutes
