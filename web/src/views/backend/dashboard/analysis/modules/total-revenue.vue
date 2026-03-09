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
  <div class="art-card h-100 p-5 mb-5 max-sm:mb-4">
    <div class="art-card-header">
      <div class="title">
        <h4>总收入</h4>
      </div>
    </div>
    <ArtBarChart
      height="calc(100% - 30px)"
      :data="revenueData"
      :xAxisData="weekDays"
      :showLegend="true"
      :showAxisLine="false"
      barWidth="18%"
    />
  </div>
</template>

<script setup lang="ts">
  interface RevenueDataItem {
    name: string
    data: number[]
  }

  /**
   * 一周的日期标签
   */
  const weekDays = ref(['周一', '周二', '周三', '周四', '周五', '周六', '周日'])

  /**
   * 总收入数据
   * 对比线上和线下销售的一周收入情况
   */
  const revenueData = ref<RevenueDataItem[]>([
    {
      name: '线上销售',
      data: [12, 13, 5, 15, 10, 15, 18]
    },
    {
      name: '线下销售',
      data: [10, 11, 20, 5, 11, 13, 10]
    }
  ])
</script>
