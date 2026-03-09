<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 函数级性能分析（独立页面） -->
<template>
  <div class="page-content">
    <!-- 顶部操作栏 -->
    <div class="pprof-header">
      <div class="pprof-header__left">
        <ArtSvgIcon icon="ri:code-s-slash-line" class="text-xl text-purple-500" />
        <span class="font-medium text-base">函数级性能分析</span>
        <ElTag v-if="result" size="small" type="info" effect="plain" round class="ml-2">
          CPU采样 {{ result.cpuTime }} | {{ result.timestamp }}
        </ElTag>
      </div>
      <div class="pprof-header__right">
        <ElSelect v-model="cpuSeconds" size="small" style="width: 130px">
          <ElOption :value="1" label="采样 1 秒" />
          <ElOption :value="3" label="采样 3 秒" />
          <ElOption :value="5" label="采样 5 秒" />
          <ElOption :value="10" label="采样 10 秒" />
        </ElSelect>
        <ElInputNumber v-model="topLimit" size="small" :min="5" :max="50" :step="5" controls-position="right" style="width: 100px" />
        <ElButton type="primary" @click="doSample" :loading="loading">
          <ArtSvgIcon icon="ri:pulse-line" class="text-sm mr-1" />
          {{ loading ? `采样中...` : '开始采样' }}
        </ElButton>
      </div>
    </div>

    <!-- 采样进度 -->
    <ElAlert v-if="loading" type="info" :closable="false" class="mb-4">
      <template #title>
        <div class="flex items-center gap-2">
          <ElIcon class="is-loading"><Loading /></ElIcon>
          正在采样 CPU Profile ({{ cpuSeconds }}秒)，采样期间可正常使用系统...
        </div>
      </template>
    </ElAlert>

    <!-- 概览统计 -->
    <ElRow :gutter="16" class="mb-4" v-if="result">
      <ElCol :xs="24" :sm="12" :md="6" class="mb-4">
        <ArtStatsCard icon="ri:fire-line" iconStyle="bg-error" :count="result.cpuTop?.length || 0" description="CPU 热点函数" />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-4">
        <ArtStatsCard icon="ri:database-2-line" iconStyle="bg-info" :count="result.memTop?.length || 0" description="内存热点函数" />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-4">
        <ArtStatsCard icon="ri:timer-line" iconStyle="bg-success" :title="result.cpuTime" description="CPU 采样时长" />
      </ElCol>
      <ElCol :xs="24" :sm="12" :md="6" class="mb-4">
        <ArtStatsCard icon="ri:time-line" iconStyle="bg-warning" :title="topCpuFunc" description="CPU 最热函数" />
      </ElCol>
    </ElRow>

    <!-- 空状态 -->
    <div v-if="!result && !loading" class="pprof-empty">
      <ArtSvgIcon icon="ri:search-eye-line" class="text-6xl mb-4" style="color: var(--el-text-color-placeholder)" />
      <h3>函数级性能分析</h3>
      <p>通过 Go pprof 采样，精确定位 CPU 和内存的热点函数</p>
      <p class="text-xs mt-2">帮助开发者快速发现：哪些函数占用最多 CPU / 分配最多内存</p>
      <ElButton type="primary" class="mt-6" @click="doSample">
        <ArtSvgIcon icon="ri:pulse-line" class="text-sm mr-1" />
        开始第一次采样
      </ElButton>
    </div>

    <!-- 结果区域 -->
    <template v-if="result">
      <!-- CPU 热点 -->
      <ElCard shadow="never" class="mb-4">
        <template #header>
          <div class="flex items-center gap-2">
            <ArtSvgIcon icon="ri:cpu-line" class="text-lg text-orange-500" />
            <span class="font-medium">CPU 热点函数 Top {{ topLimit }}</span>
            <span class="text-xs text-gray-400 ml-2">按自身 CPU 时间排序，帮助定位 CPU 密集型代码</span>
          </div>
        </template>
        <ElTable :data="result.cpuTop || []" stripe size="small" class="w-full" :max-height="500">
          <ElTableColumn label="#" type="index" width="45" />
          <ElTableColumn label="函数名" prop="func" min-width="280" show-overflow-tooltip>
            <template #default="{ row }">
              <code class="text-xs font-mono">{{ row.func }}</code>
            </template>
          </ElTableColumn>
          <ElTableColumn label="自身耗时" prop="flat" width="100" align="right" sortable>
            <template #default="{ row }">
              <span class="font-medium">{{ row.flat }}</span>
            </template>
          </ElTableColumn>
          <ElTableColumn label="自身占比" width="110" align="right" sortable sort-by="flatPct">
            <template #default="{ row }">
              <div class="flex items-center justify-end gap-2">
                <ElProgress :percentage="row.flatPct" :stroke-width="6" :show-text="false" style="width: 50px"
                  :color="row.flatPct > 30 ? '#f56c6c' : row.flatPct > 10 ? '#e6a23c' : '#67c23a'" />
                <span :class="pctClass(row.flatPct)" class="w-12 text-right">{{ row.flatPct }}%</span>
              </div>
            </template>
          </ElTableColumn>
          <ElTableColumn label="累计耗时" prop="cum" width="100" align="right" />
          <ElTableColumn label="累计占比" width="90" align="right">
            <template #default="{ row }">
              <span class="text-gray-500">{{ row.cumPct }}%</span>
            </template>
          </ElTableColumn>
          <ElTableColumn label="源文件" prop="file" width="180" show-overflow-tooltip>
            <template #default="{ row }">
              <span class="text-xs text-gray-400 font-mono">{{ row.file }}</span>
            </template>
          </ElTableColumn>
        </ElTable>
      </ElCard>

      <!-- 内存热点 -->
      <ElCard shadow="never">
        <template #header>
          <div class="flex items-center gap-2">
            <ArtSvgIcon icon="ri:database-2-line" class="text-lg text-blue-500" />
            <span class="font-medium">内存分配热点 Top {{ topLimit }}</span>
            <span class="text-xs text-gray-400 ml-2">按堆内存分配量排序，帮助定位内存泄漏和大量分配</span>
          </div>
        </template>
        <ElTable :data="result.memTop || []" stripe size="small" class="w-full" :max-height="500">
          <ElTableColumn label="#" type="index" width="45" />
          <ElTableColumn label="函数名" prop="func" min-width="280" show-overflow-tooltip>
            <template #default="{ row }">
              <code class="text-xs font-mono">{{ row.func }}</code>
            </template>
          </ElTableColumn>
          <ElTableColumn label="自身分配" prop="flat" width="100" align="right" sortable>
            <template #default="{ row }">
              <span class="font-medium">{{ row.flat }}</span>
            </template>
          </ElTableColumn>
          <ElTableColumn label="自身占比" width="110" align="right" sortable sort-by="flatPct">
            <template #default="{ row }">
              <div class="flex items-center justify-end gap-2">
                <ElProgress :percentage="row.flatPct" :stroke-width="6" :show-text="false" style="width: 50px"
                  :color="row.flatPct > 30 ? '#f56c6c' : row.flatPct > 10 ? '#e6a23c' : '#67c23a'" />
                <span :class="pctClass(row.flatPct)" class="w-12 text-right">{{ row.flatPct }}%</span>
              </div>
            </template>
          </ElTableColumn>
          <ElTableColumn label="累计分配" prop="cum" width="100" align="right" />
          <ElTableColumn label="累计占比" width="90" align="right">
            <template #default="{ row }">
              <span class="text-gray-500">{{ row.cumPct }}%</span>
            </template>
          </ElTableColumn>
          <ElTableColumn label="源文件" prop="file" width="180" show-overflow-tooltip>
            <template #default="{ row }">
              <span class="text-xs text-gray-400 font-mono">{{ row.file }}</span>
            </template>
          </ElTableColumn>
        </ElTable>
      </ElCard>
    </template>
  </div>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import { getPprofTop, type PprofTopResult } from '@/api/backend/monitor'
  import ArtSvgIcon from '@/components/core/base/art-svg-icon/index.vue'

  defineOptions({ name: 'PprofAnalysis' })

  const loading = ref(false)
  const result = ref<PprofTopResult | null>(null)
  const cpuSeconds = ref(3)
  const topLimit = ref(15)

  const topCpuFunc = computed(() => {
    if (!result.value?.cpuTop?.length) return '--'
    return result.value.cpuTop[0].func
  })

  const doSample = async () => {
    loading.value = true
    try {
      const res = await getPprofTop({ seconds: cpuSeconds.value, limit: topLimit.value })
      result.value = res as any
      ElMessage.success('采样完成')
    } catch (e: any) {
      console.error('pprof 采样失败:', e)
      ElMessage.error(e?.message || '采样失败，请检查后端服务')
    } finally {
      loading.value = false
    }
  }

  const pctClass = (pct: number) => {
    if (pct > 30) return 'text-red-500 font-medium'
    if (pct > 10) return 'text-orange-500 font-medium'
    return 'text-green-600'
  }
</script>

<style scoped>
  @reference '@styles/core/tailwind.css';

  .pprof-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: 12px;
    margin-bottom: 16px;
  }

  .pprof-header__left {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .pprof-header__right {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .pprof-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 80px 20px;
    text-align: center;
    color: var(--el-text-color-secondary);
  }

  .pprof-empty h3 {
    font-size: 18px;
    font-weight: 600;
    color: var(--el-text-color-primary);
    margin-bottom: 8px;
  }

  .pprof-empty p {
    font-size: 13px;
    max-width: 400px;
    line-height: 1.6;
  }
</style>
