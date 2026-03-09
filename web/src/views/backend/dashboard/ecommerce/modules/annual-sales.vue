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
  <div class="art-card p-5 h-[28.2rem] mb-5 max-sm:mb-4">
    <div class="art-card-header">
      <div class="title">
        <h4>年度销售额</h4>
        <p>按季度统计</p>
      </div>
    </div>

    <ArtBarChart
      :showAxisLabel="false"
      :showAxisLine="false"
      :showSplitLine="false"
      :data="[50, 80, 50, 90, 60, 70, 50]"
      barWidth="26px"
      height="calc(100% - 155px)"
    />
    <div class="flex justify-around mt-10">
      <div class="flex-c">
        <div class="flex-cc size-10.5 mr-2.5 text-theme bg-theme/10 rounded-lg">
          <ArtSvgIcon icon="ri:money-cny-circle-line" class="text-xl" />
        </div>
        <div>
          <p class="text-lg">¥200,858</p>
          <span class="text-sm">线上销售</span>
        </div>
      </div>
      <div class="flex-c">
        <div class="flex-cc size-10.5 mr-2.5 text-theme bg-theme/10 rounded-lg">
          <ArtSvgIcon icon="ri:heart-3-line" class="text-xl" />
        </div>
        <div>
          <p class="text-lg">¥102,927</p>
          <span class="text-sm">线下销售</span>
        </div>
      </div>
    </div>
  </div>
</template>
