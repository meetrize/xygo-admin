<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 历史记录列表 -->
<template>
  <div class="history-list">
    <ElTable :data="list" border stripe v-loading="loading" size="small">
      <ElTableColumn prop="tableName" label="数据表" width="200">
        <template #default="{ row }">
          <span class="font-medium">{{ row.tableName }}</span>
        </template>
      </ElTableColumn>
      <ElTableColumn prop="tableComment" label="表注释" width="150" />
      <ElTableColumn prop="varName" label="实体名" width="150">
        <template #default="{ row }">
          <code class="text-xs px-1.5 py-0.5 rounded bg-fill-lighter">{{ row.varName }}</code>
        </template>
      </ElTableColumn>
      <ElTableColumn prop="genType" label="类型" width="100" align="center">
        <template #default="{ row }">
          <ElTag :type="row.genType === 10 ? 'primary' : 'success'" size="small" effect="light" round>
            {{ row.genType === 10 ? '普通列表' : '树表' }}
          </ElTag>
        </template>
      </ElTableColumn>
      <ElTableColumn prop="status" label="状态" width="100" align="center">
        <template #default="{ row }">
          <ElTag :type="row.status === 1 ? 'success' : 'warning'" size="small" effect="light" round>
            {{ row.status === 1 ? '已生成' : '未生成' }}
          </ElTag>
        </template>
      </ElTableColumn>
      <ElTableColumn prop="createdAt" label="创建时间" width="180">
        <template #default="{ row }">
          <span class="text-xs text-color-g-500">{{ formatTime(row.createdAt) }}</span>
        </template>
      </ElTableColumn>
      <ElTableColumn label="操作" width="180" fixed="right">
        <template #default="{ row }">
          <ElButton size="small" type="primary" link @click="handleSelect(row)">
            <ArtSvgIcon icon="ri:settings-3-line" class="text-sm mr-1" />
            继续配置
          </ElButton>
          <ElButton size="small" type="danger" link @click="handleDelete(row)">
            <ArtSvgIcon icon="ri:delete-bin-line" class="text-sm mr-1" />
            删除
          </ElButton>
        </template>
      </ElTableColumn>
    </ElTable>

    <div v-if="!loading && list.length === 0" class="empty-hint">
      <ArtSvgIcon icon="ri:inbox-line" class="text-4xl" style="color: var(--el-text-color-placeholder)" />
      <p>暂无生成记录</p>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { fetchGenCodesList, fetchGenCodesDelete, fetchGenCodesView } from '@/api/backend/develop/genCodes'
  import { ElCheckbox, ElNotification } from 'element-plus'
  import ArtSvgIcon from '@/components/core/base/art-svg-icon/index.vue'
  import { formatTimestamp } from '@/utils/time'

  const emit = defineEmits<{
    (e: 'select', record: any): void
    (e: 'deleted'): void
  }>()

  const loading = ref(false)
  const list = ref<any[]>([])

  const formatTime = (ts: number) => formatTimestamp(ts)

  const fetchData = async () => {
    loading.value = true
    try {
      const res = await fetchGenCodesList({ page: 1, pageSize: 50 })
      list.value = res.list || []
    } finally {
      loading.value = false
    }
  }

  const handleSelect = async (row: any) => {
    try {
      const detail = await fetchGenCodesView(row.id)
      emit('select', detail)
    } catch (e) {
      console.error(e)
    }
  }

  // 删除选项状态
  const deleteOptions = reactive({
    deleteFiles: false,
    deleteMenus: false,
  })

  const handleDelete = async (row: any) => {
    // 重置选项
    deleteOptions.deleteFiles = false
    deleteOptions.deleteMenus = false

    try {
      await ElMessageBox({
        title: '删除确认',
        message: () => h('div', { class: 'delete-confirm-content' }, [
          h('p', { style: 'margin-bottom: 12px; color: var(--el-text-color-regular)' },
            `确定删除「${row.tableName}」(${row.varName}) 的生成配置吗？`),
          h('div', { style: 'display: flex; flex-direction: column; gap: 8px; padding-left: 4px' }, [
            h(ElCheckbox, {
              modelValue: deleteOptions.deleteFiles,
              'onUpdate:modelValue': (v: any) => { deleteOptions.deleteFiles = !!v },
              label: '同时删除已生成的文件（后端 + 前端）',
              style: 'font-size: 13px'
            }),
            h(ElCheckbox, {
              modelValue: deleteOptions.deleteMenus,
              'onUpdate:modelValue': (v: any) => { deleteOptions.deleteMenus = !!v },
              label: '同时删除已创建的菜单和权限',
              style: 'font-size: 13px'
            }),
          ]),
          h('p', {
            style: 'margin-top: 10px; font-size: 12px; color: var(--el-text-color-placeholder)'
          }, '提示: 仅删除配置记录不会影响已生成的代码和菜单')
        ]),
        confirmButtonText: '确认删除',
        cancelButtonText: '取消',
        type: 'warning',
        showCancelButton: true,
      })

      await fetchGenCodesDelete({
        id: row.id,
        deleteFiles: deleteOptions.deleteFiles,
        deleteMenus: deleteOptions.deleteMenus,
      })

      const parts = ['配置记录']
      if (deleteOptions.deleteFiles) parts.push('生成文件')
      if (deleteOptions.deleteMenus) parts.push('菜单权限')
      ElNotification({
        title: '删除成功',
        message: `已删除: ${parts.join(' + ')}`,
        type: 'success',
        duration: 5000,
      })

      fetchData()
      emit('deleted')
    } catch { /* cancel */ }
  }

  onMounted(() => fetchData())
</script>

<style scoped>
  @reference '@styles/core/tailwind.css';

  .empty-hint {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 40px 0;
    color: var(--el-text-color-placeholder);
  }

  .empty-hint p {
    margin-top: 12px;
    font-size: 13px;
  }
</style>
