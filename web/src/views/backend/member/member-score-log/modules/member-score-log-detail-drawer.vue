<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 积分变动日志 详情抽屉 -->
<template>
  <ElDrawer
    v-model="visible"
    title="积分变动日志详情"
    size="600px"
    :destroy-on-close="true"
  >
    <div v-if="loading" class="flex justify-center py-10">
      <ElIcon class="is-loading" :size="24"><Loading /></ElIcon>
    </div>
    <ElDescriptions v-else-if="detail" :column="1" border>
      <ElDescriptionsItem label="Id">{{ detail.id ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="会员ID">{{ detail.memberId ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="变动积分">{{ detail.score ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="变动前积分">{{ detail.before ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="变动后积分">{{ detail.after ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="变动说明">{{ detail.memo ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="创建时间">{{ formatTimestamp(detail.createdAt) }}</ElDescriptionsItem>
    </ElDescriptions>
  </ElDrawer>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import { fetchMemberScoreLogView } from '@/api/backend/member/member-score-log'
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
        detail.value = await fetchMemberScoreLogView(props.viewId) as any
      } catch { detail.value = null }
      loading.value = false
    }
  })
</script>
