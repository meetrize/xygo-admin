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
  <div class="page-content mb-5" element-loading-text="加载中...">
    <ArtMapChart />
  </div>
</template>

<script setup lang="ts">
  defineOptions({ name: 'TemplateMap' })

  /**
   * 异步加载地图组件
   * 减少初始加载体积
   */
  const ArtMapChart = defineAsyncComponent(
    () => import('@/components/core/charts/art-map-chart/index.vue')
  )
</script>
