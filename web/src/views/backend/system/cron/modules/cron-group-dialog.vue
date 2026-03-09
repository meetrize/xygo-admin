<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 定时任务分组管理弹窗 -->
<template>
  <ElDialog v-model="dialogVisible" title="分组管理" width="600px" align-center>
    <div class="mb-4 flex-cb">
      <ElButton type="primary" size="small" @click="handleAdd">新增分组</ElButton>
    </div>
    <ElTable :data="list" v-loading="loading" border size="small" max-height="400">
      <ElTableColumn prop="id" label="ID" width="60" align="center" />
      <ElTableColumn prop="name" label="分组名称" min-width="120" />
      <ElTableColumn prop="sort" label="排序" width="70" align="center" />
      <ElTableColumn prop="status" label="状态" width="70" align="center">
        <template #default="{ row }">
          <ElTag :type="row.status === 1 ? 'success' : 'danger'" size="small">{{ row.status === 1 ? '启用' : '禁用' }}</ElTag>
        </template>
      </ElTableColumn>
      <ElTableColumn label="操作" width="120" align="right">
        <template #default="{ row }">
          <ElButton size="small" text type="primary" @click="handleEdit(row)">编辑</ElButton>
          <ElButton size="small" text type="danger" @click="handleDelete(row)">删除</ElButton>
        </template>
      </ElTableColumn>
    </ElTable>

    <!-- 内嵌编辑表单 -->
    <ElDialog v-model="showForm" :title="editId ? '编辑分组' : '新增分组'" width="400px" append-to-body>
      <ElForm :model="formData" label-width="80px">
        <ElFormItem label="名称"><ElInput v-model="formData.name" /></ElFormItem>
        <ElFormItem label="排序"><ElInputNumber v-model="formData.sort" :min="0" /></ElFormItem>
        <ElFormItem label="备注"><ElInput v-model="formData.remark" /></ElFormItem>
        <ElFormItem label="状态"><ElSwitch v-model="formData.status" :active-value="1" :inactive-value="0" /></ElFormItem>
      </ElForm>
      <template #footer>
        <ElButton @click="showForm = false">取消</ElButton>
        <ElButton type="primary" :loading="saving" @click="handleSave">保存</ElButton>
      </template>
    </ElDialog>
  </ElDialog>
</template>

<script setup lang="ts">
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { fetchCronGroupList, fetchCronGroupSave, fetchCronGroupDelete, type CronGroupItem } from '@/api/backend/system/cron'

  interface Emits {
    (e: 'update:visible', v: boolean): void
    (e: 'change'): void
  }
  const props = defineProps<{ visible: boolean }>()
  const emit = defineEmits<Emits>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: v => emit('update:visible', v)
  })

  const loading = ref(false)
  const list = ref<CronGroupItem[]>([])
  const showForm = ref(false)
  const saving = ref(false)
  const editId = ref(0)
  const formData = reactive({ name: '', sort: 0, remark: '', status: 1 })

  const loadList = async () => {
    loading.value = true
    try {
      const res = await fetchCronGroupList({}) as any
      list.value = res?.list || []
    } catch { /* */ } finally {
      loading.value = false
    }
  }

  const handleAdd = () => {
    editId.value = 0
    Object.assign(formData, { name: '', sort: 0, remark: '', status: 1 })
    showForm.value = true
  }

  const handleEdit = (row: CronGroupItem) => {
    editId.value = row.id
    Object.assign(formData, { name: row.name, sort: row.sort, remark: row.remark, status: row.status })
    showForm.value = true
  }

  const handleSave = async () => {
    saving.value = true
    try {
      await fetchCronGroupSave({ id: editId.value || undefined, ...formData })
      ElMessage.success('保存成功')
      showForm.value = false
      loadList()
      emit('change')
    } catch { /* */ } finally {
      saving.value = false
    }
  }

  const handleDelete = async (row: CronGroupItem) => {
    await ElMessageBox.confirm(`确定删除分组「${row.name}」？`, '删除确认', { type: 'warning' })
    await fetchCronGroupDelete(row.id)
    ElMessage.success('删除成功')
    loadList()
    emit('change')
  }

  watch(() => props.visible, (val) => { if (val) loadList() })
</script>
