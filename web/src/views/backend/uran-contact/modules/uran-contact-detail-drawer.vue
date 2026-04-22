<!-- 悠然联系人 详情抽屉 -->
<template>
  <ElDrawer
    v-model="visible"
    title="悠然联系人详情"
    size="600px"
    :destroy-on-close="true"
  >
    <div v-if="loading" class="flex justify-center py-10">
      <ElIcon class="is-loading" :size="24"><Loading /></ElIcon>
    </div>
    <ElDescriptions v-else-if="detail" :column="1" border>
      <ElDescriptionsItem label="主键">{{ detail.id ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="姓名">{{ detail.username ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="电话">{{ detail.phone ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="数字">{{ detail.age ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="头像">
        <ElImage v-if="detail.avatar" :src="detail.avatar" style="width:80px;height:80px" fit="cover" :preview-src-list="[detail.avatar]" />
        <span v-else>-</span>
      </ElDescriptionsItem>
      <ElDescriptionsItem label="排序权重">{{ detail.sort ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="备注">{{ detail.remark ?? '-' }}</ElDescriptionsItem>
      <ElDescriptionsItem label="开关">{{ detail.switchField ?? '-' }}</ElDescriptionsItem>
    </ElDescriptions>
  </ElDrawer>
</template>

<script setup lang="ts">
  import { Loading } from '@element-plus/icons-vue'
  import { fetchUranContactView } from '@/api/backend/uran-contact'
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
        detail.value = await fetchUranContactView(props.viewId) as any
      } catch { detail.value = null }
      loading.value = false
    }
  })
</script>
