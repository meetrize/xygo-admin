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
 * 前台会员菜单状态管理
 * 从后端 API 获取会员可用菜单（按分组权限过滤），驱动前台导航栏和用户中心菜单
 */
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getMemberMenus, type MemberMenuItem } from '@/api/frontend/member/user'

export const useMemberMenuStore = defineStore('memberMenuStore', () => {
  // 原始菜单数据
  const navMenus = ref<MemberMenuItem[]>([])       // 顶栏导航菜单
  const centerMenus = ref<MemberMenuItem[]>([])     // 用户中心菜单
  const routeRules = ref<MemberMenuItem[]>([])      // 普通路由规则
  const loaded = ref(false)

  // 计算属性
  const getNavMenus = computed(() => navMenus.value)
  const getCenterMenus = computed(() => centerMenus.value)
  const getRouteRules = computed(() => routeRules.value)
  const isLoaded = computed(() => loaded.value)

  /**
   * 从后端拉取菜单（登录后调用）
   */
  const fetchMenus = async () => {
    try {
      const data = await getMemberMenus()
      navMenus.value = data.nav || []
      centerMenus.value = data.menus || []
      routeRules.value = data.rules || []
      loaded.value = true
    } catch {
      // 接口失败（如未登录401）不清空已有菜单，保持现状
      loaded.value = true
    }
  }

  /**
   * 清空菜单（退出登录时调用）
   */
  const clearMenus = () => {
    navMenus.value = []
    centerMenus.value = []
    routeRules.value = []
    loaded.value = false
  }

  return {
    navMenus,
    centerMenus,
    routeRules,
    loaded,
    getNavMenus,
    getCenterMenus,
    getRouteRules,
    isLoaded,
    fetchMenus,
    clearMenus,
  }
})
