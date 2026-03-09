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
 * 会员状态管理模块（前台用户）
 *
 * 与 userStore（后台管理员）完全独立
 *
 * ## 主要功能
 *
 * - 会员登录状态管理
 * - 会员信息存储
 * - 会员 Token 管理（使用 Xy-User-Token）
 * - 会员登出逻辑
 *
 * ## 使用场景
 *
 * - 前台门户登录
 * - 用户中心
 * - 会员专属功能
 *
 * ## 持久化
 *
 * - 使用 localStorage 存储
 * - 存储键：member
 * - 与后台管理员登录状态独立
 *
 * @module store/modules/member
 */
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { router } from '@/router'

/**
 * 会员信息类型
 */
export interface MemberInfo {
  id: number
  username: string
  nickname: string
  avatar: string
  mobile: string
  email: string
  gender: number
  level: number
  groupId: number
  score: number
  money: number
  lastLoginAt?: string
  lastLoginIp?: string
}

/**
 * 会员状态管理
 * 管理前台会员登录状态、个人信息、Token 等
 */
export const useMemberStore = defineStore(
  'memberStore',
  () => {
    // 登录状态
    const isLogin = ref(false)
    // 会员信息
    const info = ref<Partial<MemberInfo>>({})
    // 访问令牌（使用 Xy-User-Token）
    const token = ref('')

    // 计算属性：获取会员信息
    const getMemberInfo = computed(() => info.value)

    // 计算属性：是否已登录
    const getIsLogin = computed(() => isLogin.value && !!token.value)

    /**
     * 设置会员信息
     * @param newInfo 新的会员信息
     */
    const setMemberInfo = (newInfo: MemberInfo) => {
      info.value = newInfo
    }

    /**
     * 设置登录状态
     * @param status 登录状态
     */
    const setLoginStatus = (status: boolean) => {
      isLogin.value = status
    }

    /**
     * 设置令牌
     * @param newToken 访问令牌
     */
    const setToken = (newToken: string) => {
      token.value = newToken
      if (newToken) {
        isLogin.value = true
      }
    }

    /**
     * 获取令牌
     */
    const getToken = (): string => {
      return token.value
    }

    /**
     * 退出登录
     * 清空会员相关状态并跳转到前台登录页
     */
    const logOut = async (options: { redirect?: boolean } = {}) => {
      const { redirect = true } = options

      // 先保存 token 再清空，避免退出接口触发二次 401
      const oldToken = token.value

      // 立即清空本地状态
      info.value = {}
      isLogin.value = false
      token.value = ''

      // 再调后端退出接口（清除服务端 Token 缓存），token 已清空不会触发重复 401
      if (oldToken) {
        try {
          const { memberLogout } = await import('@/api/frontend/member/auth')
          await memberLogout()
        } catch { /* ignore */ }
      }

      // 跳转到前台登录页
      if (redirect) {
        const currentRoute = router.currentRoute.value
        if (currentRoute.path !== '/user/login') {
          const redirectPath = currentRoute.fullPath
          router.push({
            path: '/user/login',
            query: redirectPath ? { redirect: redirectPath } : undefined
          })
        }
      }
    }

    return {
      isLogin,
      info,
      token,
      getMemberInfo,
      getIsLogin,
      setMemberInfo,
      setLoginStatus,
      setToken,
      getToken,
      logOut
    }
  },
  {
    persist: {
      key: 'member',
      storage: localStorage
    }
  }
)
