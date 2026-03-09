<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="dialogType === 'add' ? (parentNode ? `添加【${parentNode.title}】的子分类` : '添加分类') : '编辑分类'"
    width="500px"
    align-center
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="90px">
      <ElFormItem label="上级分类" v-if="dialogType === 'add' && parentNode">
        <ElInput :model-value="parentNode.title" disabled />
      </ElFormItem>
      <ElFormItem label="分类名称" prop="title">
        <ElInput v-model="formData.title" placeholder="请输入分类名称" />
      </ElFormItem>
      <ElFormItem label="图标" prop="icon">
        <ElInput v-model="formData.icon" placeholder="如 ri:folder-line" />
      </ElFormItem>
      <ElFormItem label="排序" prop="sort">
        <ElInputNumber v-model="formData.sort" :min="0" controls-position="right" style="width: 100%" />
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElSwitch v-model="formData.status" :active-value="1" :inactive-value="2" />
      </ElFormItem>
      <ElFormItem label="备注" prop="remark">
        <ElInput v-model="formData.remark" type="textarea" :rows="3" placeholder="请输入备注" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <ElButton @click="dialogVisible = false">取消</ElButton>
      <ElButton type="primary" @click="handleSubmit">确定</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import { ref, reactive } from 'vue'
  import type { FormInstance, FormRules } from 'element-plus'

  const emit = defineEmits<{ submit: [data: any] }>()

  const dialogVisible = ref(false)
  const dialogType = ref<'add' | 'edit'>('add')
  const parentNode = ref<any>(null)
  const formRef = ref<FormInstance>()

  const defaultForm = () => ({
    id: undefined as number | undefined,
    pid: 0,
    title: '',
    icon: '',
    sort: 0,
    status: 1,
    remark: ''
  })

  const formData = reactive(defaultForm())

  const rules: FormRules = {
    title: [{ required: true, message: '请输入分类名称', trigger: 'blur' }]
  }

  const open = (type: 'add' | 'edit', data?: any, parent?: any) => {
    dialogType.value = type
    parentNode.value = parent || null
    Object.assign(formData, defaultForm(), data || {})
    dialogVisible.value = true
    formRef.value?.clearValidate()
  }

  const handleSubmit = async () => {
    if (!formRef.value) return
    await formRef.value.validate((valid) => {
      if (valid) {
        emit('submit', { ...formData })
        dialogVisible.value = false
      }
    })
  }

  defineExpose({ open })
</script>
