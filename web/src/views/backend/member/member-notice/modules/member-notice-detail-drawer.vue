<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 会员通知 详情抽屉 -->
<template>
  <ElDrawer
    v-model="visible"
    title="会员通知详情"
    size="600px"
    :destroy-on-close="true"
  >
    <div v-if="loading" class="flex justify-center py-10">
      <ElIcon class="is-loading" :size="24"><Loading /></ElIcon>
    </div>
    <ElDescriptions v-else-if="detail" :column="1" border>
      <ElDescriptionsItem label="Id">{{ detail.id ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="通知标题">{{ detail.title ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="通知类型">
        <ElTag :type="(({ 'system': 'success', 'announce': 'danger', 'feature': 'warning', 'maintain': 'info' } as Record<string, any>)[String(detail.type)] || 'info') as any" size="small">{{ ({ 'system': '系统通知', 'announce': '公告', 'feature': '功能更新', 'maintain': '维护通知',  })[String(detail.type)] || detail.type }}</ElTag>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="目标">
        <ElTag :type="(({ 'all': 'success', 'group': 'danger' } as Record<string, any>)[String(detail.target)] || 'info') as any" size="small">{{ ({ 'all': '全部会员', 'group': '指定分组',  })[String(detail.target)] || detail.target }}</ElTag>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="目标分组ID">{{ detail.targetId ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="发送者">{{ detail.sender ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="状态">
        <ElTag :type="(({ '0': 'success', '1': 'danger' } as Record<string, any>)[String(detail.status)] || 'info') as any" size="small">{{ ({ '0': '草稿', '1': '已发布',  })[String(detail.status)] || detail.status }}</ElTag>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="创建时间">{{ formatTimestamp(detail.createdAt) }}</ElDescriptionsItem>
    </ElDescriptions>
  </ElDrawer>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import { fetchMemberNoticeView } from '@/api/backend/member/member-notice'
  import { formatTimestamp } from '@/utils/time'

  const visible = defineModel<boolean>({ default: false })

  interface Props {
    viewId?: number
  }

  const props = defineProps<Props>()

  const loading = ref(false)
  const detail = ref<Record<string, any> | null>(null)

  watch(visible, async (val) => {
    if (val && props.viewId) {
      loading.value = true
      try {
        detail.value = await fetchMemberNoticeView(props.viewId) as any
      } catch { detail.value = null }
      loading.value = false
    }
  })
</script>
