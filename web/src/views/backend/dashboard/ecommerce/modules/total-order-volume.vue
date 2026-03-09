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
  <div class="art-card h-53 p-5 pt-3.5 mb-5 max-sm:mb-4">
    <div class="art-card-header">
      <div class="title">
        <h4 class="text-2xl font-medium">205,216</h4>
        <p>这个月增长</p>
      </div>
    </div>

    <ArtRingChart
      :data="[
        { value: 30, name: '已完成' },
        { value: 25, name: '处理中' },
        { value: 45, name: '待发货' }
      ]"
      :color="['#4C87F3', '#93F1B4', '#8BD8FC']"
      :radius="['56%', '76%']"
      height="7rem"
      :showLabel="false"
      :borderRadius="0"
    />
  </div>
</template>
