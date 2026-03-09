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
        <h4 class="text-2xl font-medium">55,231</h4>
        <p>商品总数</p>
      </div>
    </div>

    <ArtBarChart
      :showAxisLabel="false"
      :showAxisLine="false"
      :showSplitLine="false"
      :data="[50, 80, 40, 90, 60, 70]"
      height="7rem"
      barWidth="18px"
    />
  </div>
</template>
