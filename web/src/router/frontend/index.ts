/**
 * Frontend 前台路由配置（对齐 BuildAdmin 路径隔离模式）
 *
 * 所有前台页面在 FrontendLayout (/) 下静态注册。
 * 菜单 API 仅驱动导航栏 UI，不做动态路由注册。
 * 后台页面全部在 /admin 下，前后台路由互不干扰。
 */

import { AppRouteRecordRaw } from '@/utils/router'

export const frontendRoutes: AppRouteRecordRaw[] = [
  {
    path: '/',
    name: 'FrontendLayout',
    component: () => import('@/views/frontend/layouts/FrontendLayout.vue'),
    meta: { title: '首页' },
    children: [
      // 首页
      {
        path: '',
        name: 'FrontendHome',
        component: () => import('@/views/frontend/index/index.vue'),
        meta: { title: '首页' }
      },
      {
        path: 'docs',
        name: 'FrontendDocs',
        component: () => import('@/views/frontend/docs/index.vue'),
        meta: { title: '文档中心' }
      },
      // // 案例
      // {
      //   path: 'cases',
      //   name: 'FrontendCases',
      //   component: () => import('@/views/frontend/cases/index.vue'),
      //   meta: { title: '客户案例' }
      // },
      // // 社区
      // {
      //   path: 'community',
      //   name: 'FrontendCommunity',
      //   component: () => import('@/views/frontend/community/index.vue'),
      //   meta: { title: '开发者社区' }
      // },
      // // 社区帖子详情
      // {
      //   path: 'community/:id',
      //   name: 'CommunityDetail',
      //   component: () => import('@/views/frontend/community/detail.vue'),
      //   meta: { title: '帖子详情' }
      // },
      // // 社区发帖
      // {
      //   path: 'community-publish',
      //   name: 'CommunityPublish',
      //   component: () => import('@/views/frontend/community/publish.vue'),
      //   meta: { title: '发布提问', requiresAuth: true }
      // },
      // // 社区搜索
      // {
      //   path: 'community-search',
      //   name: 'CommunitySearch',
      //   component: () => import('@/views/frontend/community/search.vue'),
      //   meta: { title: '搜索结果' }
      // },
      // 会员登录
      {
        path: 'user/login',
        name: 'MemberLogin',
        component: () => import('@/views/frontend/member/login.vue'),
        meta: { title: '登录' }
      },
      // 会员注册
      {
        path: 'user/register',
        name: 'MemberRegister',
        component: () => import('@/views/frontend/member/register.vue'),
        meta: { title: '注册' }
      },
      // 用户中心
      {
        path: 'user',
        name: 'MemberCenter',
        component: () => import('@/views/frontend/member/center.vue'),
        meta: { title: '用户中心', requiresAuth: true }
      },
    ]
  }
]

export default frontendRoutes
