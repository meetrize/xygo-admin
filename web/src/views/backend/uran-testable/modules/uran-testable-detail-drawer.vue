<!-- 测试记帐 详情抽屉 -->
<template>
  <ElDrawer v-model="visible" title="测试记帐详情" size="600px" :destroy-on-close="true">
    <div v-if="loading" class="flex justify-center py-10">
      <ElIcon class="is-loading" :size="24"><Loading /></ElIcon>
    </div>
    <ElDescriptions v-else-if="detail" :column="1" border>
      <ElDescriptionsItem label="主键">{{ detail.id ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="状态">
        <ElTag
          :type="({ '0': 'success', '1': 'danger' } as const)[String(detail.status)] ?? undefined"
          size="small"
          >{{ { '0': '禁用', '1': '启用' }[String(detail.status)] || detail.status }}</ElTag
        >
      </ElDescriptionsItem>
      <ElDescriptionsItem label="排序">{{ detail.sort ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="创建时间">{{
        formatTimestamp(detail.createdAt)
      }}</ElDescriptionsItem>
      <ElDescriptionsItem label="更新时间">{{
        formatTimestamp(detail.updatedAt)
      }}</ElDescriptionsItem>
      <ElDescriptionsItem label="物件名称">{{ detail.stuff ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="价格">{{ detail.price ?? '-' }}</ElDescriptionsItem>
    </ElDescriptions>
  </ElDrawer>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import { fetchUranTestableView } from '@/api/backend/uran-testable'
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
        detail.value = (await fetchUranTestableView(props.viewId)) as any
      } catch {
        detail.value = null
      }
      loading.value = false
    }
  })
</script>
