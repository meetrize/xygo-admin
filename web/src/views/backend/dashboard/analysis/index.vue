<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 分析页页面 -->
<template>
  <div>
    <ElRow :gutter="20">
      <ElCol :xl="14" :lg="15" :xs="24">
        <TodaySales />
      </ElCol>
      <ElCol :xl="10" :lg="9" :xs="24">
        <VisitorInsights />
      </ElCol>
    </ElRow>

    <ElRow :gutter="20">
      <ElCol :xl="10" :lg="10" :xs="24">
        <TotalRevenue />
      </ElCol>
      <ElCol :xl="7" :lg="7" :xs="24">
        <CustomerSatisfaction />
      </ElCol>
      <ElCol :xl="7" :lg="7" :xs="24">
        <TargetVsReality />
      </ElCol>
    </ElRow>

    <ElRow :gutter="20">
      <ElCol :xl="10" :lg="10" :xs="24">
        <TopProducts />
      </ElCol>
      <ElCol :xl="7" :lg="7" :xs="24">
        <SalesMappingByCountry />
      </ElCol>
      <ElCol :xl="7" :lg="7" :xs="24">
        <VolumeServiceLevel />
      </ElCol>
    </ElRow>
  </div>
</template>

<script setup lang="ts">
  import TodaySales from './modules/today-sales.vue'
  import VisitorInsights from './modules/visitor-insights.vue'
  import TotalRevenue from './modules/total-revenue.vue'
  import CustomerSatisfaction from './modules/customer-satisfaction.vue'
  import TargetVsReality from './modules/target-vs-reality.vue'
  import TopProducts from './modules/top-products.vue'
  import SalesMappingByCountry from './modules/sales-mapping-by-country.vue'
  import VolumeServiceLevel from './modules/volume-service-level.vue'

  defineOptions({ name: 'Analysis' })
</script>
