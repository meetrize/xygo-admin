<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 性能分析仪表板 -->
<template>
  <div class="page-content">
    <!-- 时间范围筛选 -->
    <div class="perf-header">
      <div class="perf-header__left">
        <ArtSvgIcon icon="ri:bar-chart-box-line" class="text-xl text-blue-500" />
        <span class="font-medium text-base">性能分析</span>
      </div>
      <div class="perf-header__right">
        <ElRadioGroup v-model="dateRange" size="small" @change="refreshAll">
          <ElRadioButton value="today">今日</ElRadioButton>
          <ElRadioButton value="7d">近7天</ElRadioButton>
          <ElRadioButton value="30d">近30天</ElRadioButton>
        </ElRadioGroup>
        <ElButton :icon="Refresh" circle size="small" @click="refreshAll" :loading="loading" />
      </div>
    </div>

    <!-- 顶部 4 个 ArtStatsCard -->
    <ElRow :gutter="20" class="mb-5">
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:send-plane-line"
          iconStyle="bg-info"
          :count="stats?.summary?.totalRequests"
          description="总请求数"
        />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:timer-line"
          iconStyle="bg-success"
          :title="`${stats?.summary?.avgElapsed ?? '--'} ms`"
          description="平均响应时间"
        />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:error-warning-line"
          iconStyle="bg-error"
          :count="stats?.summary?.errorCount"
          description="错误请求"
        />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:speed-line"
          iconStyle="bg-warning"
          :count="stats?.summary?.slowCount"
          description="慢接口 (>200ms)"
        />
      </ElCol>
    </ElRow>

    <!-- 图表行 -->
    <div class="grid grid-cols-3 gap-4 mb-4 max-lg:grid-cols-1">
      <!-- 请求趋势折线图 -->
      <ElCard shadow="never" class="col-span-2 max-lg:col-span-1">
        <template #header>
          <div class="flex items-center gap-2">
            <ArtSvgIcon icon="ri:line-chart-line" class="text-blue-500" />
            <span class="font-medium text-sm">请求趋势</span>
          </div>
        </template>
        <div ref="trendRef" class="h-72"></div>
      </ElCard>

      <!-- 模块耗时排行 -->
      <ElCard shadow="never">
        <template #header>
          <div class="flex items-center gap-2">
            <ArtSvgIcon icon="ri:bar-chart-horizontal-line" class="text-green-500" />
            <span class="font-medium text-sm">模块平均耗时 Top 10</span>
          </div>
        </template>
        <div ref="moduleRef" class="h-72"></div>
      </ElCard>
    </div>

    <!-- 慢接口排行表格 -->
    <ElCard shadow="never">
      <template #header>
        <div class="flex items-center gap-2">
          <ArtSvgIcon icon="ri:speed-line" class="text-orange-500" />
          <span class="font-medium text-sm">慢接口排行</span>
        </div>
      </template>
      <ElTable :data="slowList" v-loading="slowLoading" stripe size="small" class="w-full">
        <ElTableColumn label="#" type="index" width="50" />
        <ElTableColumn label="方法" prop="method" width="80">
          <template #default="{ row }">
            <ElTag :type="methodType(row.method)" size="small">{{ row.method }}</ElTag>
          </template>
        </ElTableColumn>
        <ElTableColumn label="接口路径" prop="url" min-width="250" show-overflow-tooltip />
        <ElTableColumn label="模块" prop="module" width="120" show-overflow-tooltip />
        <ElTableColumn label="平均耗时" prop="avgElapsed" width="120" sortable>
          <template #default="{ row }">
            <span :class="elapsedClass(row.avgElapsed)">{{ row.avgElapsed }} ms</span>
          </template>
        </ElTableColumn>
        <ElTableColumn label="最大耗时" prop="maxElapsed" width="120" sortable>
          <template #default="{ row }">
            <span :class="elapsedClass(row.maxElapsed)">{{ row.maxElapsed }} ms</span>
          </template>
        </ElTableColumn>
        <ElTableColumn label="调用次数" prop="count" width="100" sortable />
      </ElTable>
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import { Refresh } from '@element-plus/icons-vue'
  import { useChart } from '@/hooks/core/useChart'
  import {
    getPerformanceStats,
    getSlowApiTop,
    type PerformanceStatsResult,
    type SlowApiItem
  } from '@/api/backend/monitor'

  defineOptions({ name: 'PerformanceAnalysis' })

  const loading = ref(false)
  const slowLoading = ref(false)
  const stats = ref<PerformanceStatsResult | null>(null)
  const slowList = ref<SlowApiItem[]>([])
  const dateRange = ref<'today' | '7d' | '30d'>('today')

  const { chartRef: trendRef, initChart: initTrendChart } = useChart()
  const { chartRef: moduleRef, initChart: initModuleChart } = useChart()

  const getDateParams = () => {
    const now = new Date()
    const fmt = (d: Date) => d.toISOString().slice(0, 10)
    const endDate = fmt(now)

    let startDate = endDate
    if (dateRange.value === '7d') {
      const d = new Date(now)
      d.setDate(d.getDate() - 7)
      startDate = fmt(d)
    } else if (dateRange.value === '30d') {
      const d = new Date(now)
      d.setDate(d.getDate() - 30)
      startDate = fmt(d)
    }
    return { startDate, endDate }
  }

  const fetchStats = async () => {
    loading.value = true
    try {
      const params = getDateParams()
      stats.value = await getPerformanceStats(params)
      // 等 DOM 和 Transition 动画就绪后再初始化图表
      nextTick(() => {
        setTimeout(() => {
          renderTrendChart()
          renderModuleChart()
        }, 300)
      })
    } catch (e) {
      console.error('获取性能统计失败:', e)
    } finally {
      loading.value = false
    }
  }

  const fetchSlowTop = async () => {
    slowLoading.value = true
    try {
      const params = { ...getDateParams(), limit: 20 }
      const res = await getSlowApiTop(params)
      slowList.value = res?.list ?? []
    } catch (e) {
      console.error('获取慢接口排行失败:', e)
    } finally {
      slowLoading.value = false
    }
  }

  const refreshAll = () => {
    fetchStats()
    fetchSlowTop()
  }

  const renderTrendChart = () => {
    const trend = stats.value?.trend ?? []
    const isEmpty = trend.length === 0
    initTrendChart(
      {
        tooltip: { trigger: 'axis', axisPointer: { type: 'cross' } },
        legend: { data: ['请求量', '平均耗时(ms)'], bottom: 0 },
        grid: { left: 50, right: 50, top: 20, bottom: 40 },
        xAxis: {
          type: 'category',
          data: trend.map((t) => t.time),
          axisLabel: { fontSize: 10, rotate: 30 }
        },
        yAxis: [
          { type: 'value', name: '请求量', position: 'left', splitLine: { show: true } },
          { type: 'value', name: 'ms', position: 'right', splitLine: { show: false } }
        ],
        series: [
          {
            name: '请求量',
            type: 'bar',
            yAxisIndex: 0,
            data: trend.map((t) => t.count),
            itemStyle: { borderRadius: [3, 3, 0, 0] },
            barMaxWidth: 20,
            opacity: 0.7
          },
          {
            name: '平均耗时(ms)',
            type: 'line',
            yAxisIndex: 1,
            data: trend.map((t) => t.avgElapsed),
            smooth: true,
            lineStyle: { width: 2 },
            areaStyle: { opacity: 0.1 }
          }
        ]
      } as any,
      isEmpty
    )
  }

  const renderModuleChart = () => {
    const modules = stats.value?.moduleTop ?? []
    const isEmpty = modules.length === 0
    initModuleChart(
      {
        tooltip: { trigger: 'axis' },
        grid: { left: 100, right: 30, top: 10, bottom: 10 },
        xAxis: { type: 'value', name: 'ms' },
        yAxis: {
          type: 'category',
          data: modules.map((m) => m.module).reverse(),
          axisLabel: { fontSize: 11, width: 80, overflow: 'truncate' }
        },
        series: [
          {
            type: 'bar',
            data: modules
              .map((m) => ({
                value: m.avgElapsed,
                itemStyle: {
                  borderRadius: [0, 4, 4, 0],
                  color: m.avgElapsed > 500 ? '#f56c6c' : m.avgElapsed > 200 ? '#e6a23c' : '#67c23a'
                }
              }))
              .reverse(),
            barMaxWidth: 16,
            label: { show: true, position: 'right', fontSize: 11, formatter: '{c} ms' }
          }
        ]
      } as any,
      isEmpty
    )
  }

  const elapsedClass = (ms: number) => {
    if (ms > 500) return 'text-red-500 font-medium'
    if (ms > 200) return 'text-orange-500 font-medium'
    return 'text-green-600'
  }

  const methodType = (method: string) => {
    const map: Record<string, string> = { POST: 'primary', GET: 'success', PUT: 'warning', DELETE: 'danger' }
    return (map[method] || 'info') as any
  }

  onMounted(() => {
    refreshAll()
  })

  // KeepAlive 缓存下，从其他标签切回时重新加载
  onActivated(() => {
    refreshAll()
  })
</script>

<style scoped>
  .perf-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: 12px;
    margin-bottom: 16px;
  }

  .perf-header__left {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .perf-header__right {
    display: flex;
    align-items: center;
    gap: 10px;
  }
</style>
