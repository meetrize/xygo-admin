<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<template>
  <div class="box-border w-full h-full" v-loading="isLoading">
    <iframe
      ref="iframeRef"
      :src="iframeUrl"
      frameborder="0"
      class="w-full h-full min-h-[calc(100vh-120px)] border-none"
      @load="handleIframeLoad"
    ></iframe>
  </div>
</template>

<script setup lang="ts">
  import { IframeRouteManager } from '@/router/core'

  defineOptions({ name: 'IframeView' })

  const route = useRoute()
  const isLoading = ref(true)
  const iframeUrl = ref('')
  const iframeRef = ref<HTMLIFrameElement | null>(null)

  /**
   * 初始化 iframe URL
   * 从路由配置中获取对应的外部链接地址
   */
  onMounted(() => {
    const iframeRoute = IframeRouteManager.getInstance().findByPath(route.path)

    if (iframeRoute?.meta) {
      iframeUrl.value = iframeRoute.meta.link || ''
    }
  })

  /**
   * 处理 iframe 加载完成事件
   * 隐藏加载状态
   */
  const handleIframeLoad = (): void => {
    isLoading.value = false
  }
</script>
