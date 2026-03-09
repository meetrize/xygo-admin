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
    :title="dialogType === 'add' ? (parentDept ? `添加【${parentDept.name}】的子部门` : '添加部门') : '编辑部门'"
    width="500px"
    align-center
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="90px">
      <ElFormItem label="上级部门" v-if="dialogType === 'add' && parentDept">
        <ElInput :model-value="parentDept.name" disabled />
      </ElFormItem>
      <ElFormItem label="部门名称" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入部门名称" />
      </ElFormItem>
      <ElFormItem label="排序" prop="sort">
        <ElInputNumber v-model="formData.sort" :min="0" controls-position="right" style="width: 100%" />
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElSwitch v-model="formData.status" :active-value="1" :inactive-value="0" />
      </ElFormItem>
      <ElFormItem label="备注" prop="remark">
        <ElInput v-model="formData.remark" type="textarea" :rows="3" placeholder="请输入备注" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="dialog-footer">
        <ElButton @click="dialogVisible = false">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">确定</ElButton>
      </div>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'

  interface Props {
    visible: boolean
    type: string
    deptData?: any
    parentDept?: any
  }

  interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit', formData: any): void
  }

  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
  })

  const dialogType = computed(() => props.type)

  const formRef = ref<FormInstance>()

  const formData = reactive({
    id: 0,
    parentId: 0,
    name: '',
    sort: 0,
    status: 1,
    remark: ''
  })

  const rules: FormRules = {
    name: [
      { required: true, message: '请输入部门名称', trigger: 'blur' },
      { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
    ]
  }

  const initFormData = () => {
    const isEdit = props.type === 'edit' && props.deptData

    if (isEdit) {
      const row = props.deptData
      formData.id = row.id || 0
      formData.parentId = row.parentId || 0
      formData.name = row.name || ''
      formData.sort = row.sort || 0
      formData.status = row.status || 1
      formData.remark = row.remark || ''
    } else {
      formData.id = 0
      formData.parentId = props.parentDept?.id || 0
      formData.name = ''
      formData.sort = 0
      formData.status = 1
      formData.remark = ''
    }
  }

  watch(
    () => [props.visible, props.type, props.deptData, props.parentDept],
    ([visible]) => {
      if (visible) {
        initFormData()
        nextTick(() => {
          formRef.value?.clearValidate()
        })
      }
    },
    { immediate: true }
  )

  const handleSubmit = async () => {
    if (!formRef.value) return

    await formRef.value.validate((valid) => {
      if (valid) {
        emit('submit', { ...formData })
        dialogVisible.value = false
      }
    })
  }
</script>
