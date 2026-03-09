<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 操作日志详情抽屉 -->
<template>
  <ElDrawer
    v-model="visible"
    title="操作日志详情"
    size="600px"
    :destroy-on-close="true"
  >
    <div v-if="loading" class="flex justify-center py-10">
      <ElIcon class="is-loading" :size="24"><Loading /></ElIcon>
    </div>
    <ElDescriptions v-else-if="detail" :column="1" border>
      <ElDescriptionsItem label="日志ID">{{ detail.id }}</ElDescriptionsItem>
      <ElDescriptionsItem label="操作人">{{ detail.username }}</ElDescriptionsItem>
      <ElDescriptionsItem label="模块">{{ detail.module }}</ElDescriptionsItem>
      <ElDescriptionsItem label="操作">{{ detail.title }}</ElDescriptionsItem>
      <ElDescriptionsItem label="请求方式">
        <ElTag :type="methodTagType(detail.method)" size="small">{{ detail.method }}</ElTag>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="请求URL">{{ detail.url }}</ElDescriptionsItem>
      <ElDescriptionsItem label="操作IP">{{ detail.ip }}</ElDescriptionsItem>
      <ElDescriptionsItem label="操作地点">{{ detail.location || '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="状态">
        <ElTag :type="detail.status === 1 ? 'success' : 'danger'" size="small">
          {{ detail.status === 1 ? '成功' : '失败' }}
        </ElTag>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="耗时">{{ detail.elapsed }}ms</ElDescriptionsItem>
      <ElDescriptionsItem label="操作时间">{{ detail.createdAt }}</ElDescriptionsItem>
      <ElDescriptionsItem v-if="detail.errorMessage" label="错误信息">
        <span class="text-red-500">{{ detail.errorMessage }}</span>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="请求参数">
        <div class="max-h-[200px] overflow-auto">
          <pre class="text-xs whitespace-pre-wrap break-all bg-gray-50 dark:bg-gray-800 p-2 rounded">{{ formatJson(detail.requestBody) }}</pre>
        </div>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="响应结果">
        <div class="max-h-[200px] overflow-auto">
          <pre class="text-xs whitespace-pre-wrap break-all bg-gray-50 dark:bg-gray-800 p-2 rounded">{{ formatJson(detail.responseBody) }}</pre>
        </div>
      </ElDescriptionsItem>
    </ElDescriptions>
  </ElDrawer>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import {
    getOperationLogDetail,
    type OperationLogItem
  } from '@/api/backend/monitor/operationLog'

  const visible = defineModel<boolean>({ default: false })

  interface Props {
    logId?: number
  }

  const props = defineProps<Props>()

  const loading = ref(false)
  const detail = ref<OperationLogItem | null>(null)

  watch(() => props.logId, async (id) => {
    if (id && visible.value) {
      await fetchDetail(id)
    }
  })

  watch(visible, async (val) => {
    if (val && props.logId) {
      await fetchDetail(props.logId)
    }
  })

  const fetchDetail = async (id: number) => {
    loading.value = true
    try {
      detail.value = await getOperationLogDetail(id) as any
    } finally {
      loading.value = false
    }
  }

  const methodTagType = (method: string) => {
    const map: Record<string, string> = {
      GET: 'info',
      POST: 'success',
      PUT: 'warning',
      DELETE: 'danger'
    }
    return (map[method] || 'info') as any
  }

  const formatJson = (str: string) => {
    if (!str) return '-'
    try {
      return JSON.stringify(JSON.parse(str), null, 2)
    } catch {
      return str
    }
  }
</script>
