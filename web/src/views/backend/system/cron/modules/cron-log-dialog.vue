<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 定时任务执行日志弹窗 -->
<template>
  <ElDialog v-model="dialogVisible" :title="cronTitle ? `执行日志 - ${cronTitle}` : '全部执行日志'" width="90vw" align-center>
    <div class="mb-3 flex-cb">
      <ElSelect v-model="statusFilter" clearable placeholder="全部状态" style="width: 120px" @change="loadList">
        <ElOption label="成功" :value="1" />
        <ElOption label="失败" :value="2" />
      </ElSelect>
      <ElButton type="danger" size="small" @click="handleClear">清空日志</ElButton>
    </div>

    <ElTable :data="list" v-loading="loading" border size="default" max-height="calc(80vh - 160px)">
      <ElTableColumn prop="id" label="ID" width="70" align="center" />
      <ElTableColumn prop="title" label="任务" width="120" show-overflow-tooltip />
      <ElTableColumn prop="name" label="标识" width="100">
        <template #default="{ row }"><code class="text-xs">{{ row.name }}</code></template>
      </ElTableColumn>
      <ElTableColumn prop="status" label="状态" width="70" align="center">
        <template #default="{ row }">
          <ElTag :type="row.status === 1 ? 'success' : 'danger'" size="small">{{ row.status === 1 ? '成功' : '失败' }}</ElTag>
        </template>
      </ElTableColumn>
      <ElTableColumn prop="takeMs" label="耗时" width="90" align="center">
        <template #default="{ row }">{{ row.takeMs }}ms</template>
      </ElTableColumn>
      <ElTableColumn prop="output" label="输出" min-width="200" show-overflow-tooltip />
      <ElTableColumn prop="errMsg" label="错误" min-width="150" show-overflow-tooltip>
        <template #default="{ row }">
          <span v-if="row.errMsg" class="text-red-500">{{ row.errMsg }}</span>
          <span v-else class="text-g-400">-</span>
        </template>
      </ElTableColumn>
      <ElTableColumn prop="createdAt" label="执行时间" width="160">
        <template #default="{ row }">{{ formatTime(row.createdAt) }}</template>
      </ElTableColumn>
    </ElTable>

    <div class="mt-3 flex justify-end">
      <ElPagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :total="pagination.total"
        :page-sizes="[10, 20, 50]"
        layout="total, sizes, prev, pager, next"
        small
        @change="loadList"
      />
    </div>
  </ElDialog>
</template>

<script setup lang="ts">
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { fetchCronLogList, fetchCronLogClear, type CronLogItem } from '@/api/backend/system/cron'
  import { formatTimestamp } from '@/utils/time'

  interface Props {
    visible: boolean
    cronId?: number
    cronTitle?: string
  }
  interface Emits {
    (e: 'update:visible', v: boolean): void
  }
  const props = withDefaults(defineProps<Props>(), { cronId: 0, cronTitle: '' })
  const emit = defineEmits<Emits>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: v => emit('update:visible', v)
  })

  const loading = ref(false)
  const list = ref<CronLogItem[]>([])
  const statusFilter = ref<number | undefined>(undefined)
  const pagination = reactive({ page: 1, pageSize: 20, total: 0 })

  const formatTime = (ts: number) => formatTimestamp(ts)

  const loadList = async () => {
    loading.value = true
    try {
      const res = await fetchCronLogList({
        cronId: props.cronId || undefined,
        status: statusFilter.value,
        page: pagination.page,
        pageSize: pagination.pageSize
      }) as any
      list.value = res?.list || []
      pagination.total = res?.total || 0
    } catch { /* */ } finally {
      loading.value = false
    }
  }

  const handleClear = async () => {
    const msg = props.cronId ? '确定清空该任务的执行日志？' : '确定清空全部执行日志？'
    await ElMessageBox.confirm(msg, '清空确认', { type: 'warning' })
    await fetchCronLogClear(props.cronId || undefined)
    ElMessage.success('已清空')
    loadList()
  }

  watch(() => props.visible, (val) => {
    if (val) {
      pagination.page = 1
      statusFilter.value = undefined
      loadList()
    }
  })
</script>
