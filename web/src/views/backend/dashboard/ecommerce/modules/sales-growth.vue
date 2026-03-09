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
  <div class="art-card p-5 pt-3.5 mb-5 h-46 max-sm:mb-4">
    <div class="art-card-header">
      <div class="title">
        <h4 class="text-2xl">
          <span>12%</span>
          <ArtSvgIcon icon="ri:arrow-right-up-line" class="text-success text-lg ml-1" />
        </h4>
        <p>增长</p>
      </div>
    </div>

    <ArtLineChart
      :showAreaColor="true"
      :showAxisLabel="false"
      :showAxisLine="false"
      :showSplitLine="false"
      :data="[50, 85, 65, 95, 75, 130, 180]"
      barWidth="16px"
      height="5rem"
    />
  </div>
</template>
