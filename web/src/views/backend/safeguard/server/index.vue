<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 服务器监控页面 - 展示当前服务器真实指标 -->
<template>
  <div class="page-content">
    <!-- 顶部 4 个 ArtStatsCard -->
    <ElRow :gutter="20" class="mb-5">
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:cpu-line"
          :iconStyle="cpuIconStyle"
          :title="`CPU 使用率 ${info?.cpu.usage ?? '--'}%`"
          :description="`${info?.cpu.cores ?? '-'} 核心`"
        />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:database-2-line"
          :iconStyle="memIconStyle"
          :title="`内存 ${info?.memory.usageRate ?? '--'}%`"
          :description="`${info?.memory.usedStr ?? '-'} / ${info?.memory.totalStr ?? '-'}`"
        />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:hard-drive-2-line"
          :iconStyle="diskIconStyle"
          :title="`磁盘 ${info?.disk.usageRate ?? '--'}%`"
          :description="`${info?.disk.usedStr ?? '-'} / ${info?.disk.totalStr ?? '-'}`"
        />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-5">
        <ArtStatsCard
          icon="ri:route-line"
          iconStyle="bg-info"
          :count="info?.runtime.goroutines"
          description="Goroutines"
        />
      </ElCol>
    </ElRow>

    <div class="grid grid-cols-2 gap-4 mb-4 max-lg:grid-cols-1">
      <!-- 系统信息 -->
      <ElCard shadow="never" class="info-card">
        <template #header>
          <div class="flex items-center gap-2">
            <ArtSvgIcon icon="ri:computer-line" class="text-lg text-blue-500" />
            <span class="font-medium">系统信息</span>
          </div>
        </template>
        <div class="info-grid" v-if="info">
          <div class="info-row">
            <span class="info-label">主机名</span>
            <span class="info-value">{{ info.os.hostname }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">操作系统</span>
            <span class="info-value">{{ info.os.platform }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">系统架构</span>
            <span class="info-value">{{ info.os.arch }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">CPU 型号</span>
            <span class="info-value text-xs">{{ info.cpu.modelName || '-' }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">系统运行时长</span>
            <span class="info-value">{{ info.os.uptime }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">应用运行时长</span>
            <span class="info-value text-blue-500 font-medium">{{ info.os.appTime }}</span>
          </div>
        </div>
        <ElSkeleton :rows="6" animated v-else />
      </ElCard>

      <!-- Go 运行时 -->
      <ElCard shadow="never" class="info-card">
        <template #header>
          <div class="flex items-center gap-2">
            <ArtSvgIcon icon="ri:terminal-box-line" class="text-lg text-green-500" />
            <span class="font-medium">Go 运行时</span>
          </div>
        </template>
        <div class="info-grid" v-if="info">
          <div class="info-row">
            <span class="info-label">Go 版本</span>
            <span class="info-value">{{ info.os.goVer }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">Goroutine 数量</span>
            <span class="info-value font-medium" :class="info.runtime.goroutines > 100 ? 'text-orange-500' : 'text-green-500'">
              {{ info.runtime.goroutines }}
            </span>
          </div>
          <div class="info-row">
            <span class="info-label">堆内存分配</span>
            <span class="info-value">{{ info.runtime.heapAlloc }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">堆内存系统</span>
            <span class="info-value">{{ info.runtime.heapSys }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">栈内存使用</span>
            <span class="info-value">{{ info.runtime.stackInUse }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">GC 次数</span>
            <span class="info-value">{{ info.runtime.numGC }}</span>
          </div>
          <div class="info-row">
            <span class="info-label">上次 GC</span>
            <span class="info-value text-xs">{{ info.runtime.lastGC }}</span>
          </div>
        </div>
        <ElSkeleton :rows="7" animated v-else />
      </ElCard>
    </div>

    <!-- 资源使用环形图 -->
    <div class="grid grid-cols-3 gap-4 mb-4 max-lg:grid-cols-1">
      <ElCard shadow="never" v-for="ring in ringData" :key="ring.title">
        <div class="flex flex-col items-center py-4">
          <ElProgress type="dashboard" :percentage="ring.value" :color="ring.colors" :stroke-width="12" :width="140">
            <template #default="{ percentage }">
              <div class="text-center">
                <div class="text-2xl font-bold">{{ percentage }}%</div>
                <div class="text-xs text-gray-400 mt-1">{{ ring.title }}</div>
              </div>
            </template>
          </ElProgress>
          <div class="text-sm text-gray-500 mt-3">{{ ring.desc }}</div>
        </div>
      </ElCard>
    </div>

  </div>
</template>

<script setup lang="ts">
  import { getServerInfo, type ServerInfo } from '@/api/backend/monitor'
  import { useWebSocket } from '@/hooks/useWebSocket'

  defineOptions({ name: 'SafeguardServer' })

  const info = ref<ServerInfo | null>(null)
  const loading = ref(true)

  const { connected, joinTag, quitTag, onEvent } = useWebSocket()

  // 首次加载使用 HTTP 获取（确保即使 WebSocket 未连接也能显示数据）
  const fetchData = async () => {
    try {
      info.value = await getServerInfo()
    } catch (e) {
      console.error('获取服务器信息失败:', e)
    } finally {
      loading.value = false
    }
  }

  // 监听 WebSocket 推送的监控数据
  onEvent('monitor/server', (data: ServerInfo) => {
    info.value = data
    loading.value = false
  })

  // 进入页面时订阅监控推送，离开时退订
  onMounted(() => {
    fetchData()
    // 订阅 monitor_server 标签组，后端会定时推送数据
    if (connected.value) {
      joinTag('monitor_server')
    } else {
      // 等待连接成功后再订阅
      const stopWatch = watch(connected, (val) => {
        if (val) {
          joinTag('monitor_server')
          stopWatch()
        }
      })
    }
  })

  onUnmounted(() => {
    quitTag('monitor_server')
  })

  // 图标样式根据使用率变色
  const getIconStyle = (rate: number | undefined) => {
    if (rate === undefined) return 'bg-info'
    if (rate >= 90) return 'bg-error'
    if (rate >= 70) return 'bg-warning'
    return 'bg-success'
  }

  const cpuIconStyle = computed(() => getIconStyle(info.value?.cpu.usage))
  const memIconStyle = computed(() => getIconStyle(info.value?.memory.usageRate))
  const diskIconStyle = computed(() => getIconStyle(info.value?.disk.usageRate))

  const usageColors = [
    { color: '#67c23a', percentage: 60 },
    { color: '#e6a23c', percentage: 80 },
    { color: '#f56c6c', percentage: 100 }
  ]

  const ringData = computed(() => [
    {
      title: 'CPU',
      value: info.value?.cpu.usage ?? 0,
      colors: usageColors,
      desc: `${info.value?.cpu.cores ?? '-'} 核心`
    },
    {
      title: '内存',
      value: info.value?.memory.usageRate ?? 0,
      colors: usageColors,
      desc: `${info.value?.memory.usedStr ?? '-'} / ${info.value?.memory.totalStr ?? '-'}`
    },
    {
      title: '磁盘',
      value: info.value?.disk.usageRate ?? 0,
      colors: usageColors,
      desc: `${info.value?.disk.usedStr ?? '-'} / ${info.value?.disk.totalStr ?? '-'}`
    }
  ])

</script>

<style scoped>
  @reference '@styles/core/tailwind.css';

  .info-grid {
    @apply divide-y divide-gray-100 dark:divide-gray-700/50;
  }
  .info-row {
    @apply flex items-center justify-between px-5 py-3;
  }
  .info-label {
    @apply text-sm text-gray-500 dark:text-gray-400;
  }
  .info-value {
    @apply text-sm text-gray-800 dark:text-gray-200;
  }
</style>

<style>
.info-card .el-card__body {
  padding: 0;
}
</style>
