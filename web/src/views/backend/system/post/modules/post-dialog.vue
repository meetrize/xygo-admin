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
    :title="dialogType === 'add' ? '添加岗位' : '编辑岗位'"
    width="500px"
    align-center
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="90px">
      <ElFormItem label="岗位编码" prop="code">
        <ElInput v-model="formData.code" placeholder="请输入岗位编码（如：CEO、PM）" />
      </ElFormItem>
      <ElFormItem label="岗位名称" prop="name">
        <ElInput v-model="formData.name" placeholder="请输入岗位名称" />
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
    postData?: Partial<any>
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
    code: '',
    name: '',
    sort: 0,
    status: 1,
    remark: ''
  })

  const rules: FormRules = {
    code: [
      { required: true, message: '请输入岗位编码', trigger: 'blur' },
      { min: 2, max: 64, message: '长度在 2 到 64 个字符', trigger: 'blur' }
    ],
    name: [
      { required: true, message: '请输入岗位名称', trigger: 'blur' },
      { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
    ]
  }

  const initFormData = () => {
    const isEdit = props.type === 'edit' && props.postData
    const row = props.postData

    if (isEdit && row) {
      formData.id = row.id || 0
      formData.code = row.code || ''
      formData.name = row.name || ''
      formData.sort = row.sort || 0
      formData.status = row.status || 1
      formData.remark = row.remark || ''
    } else {
      formData.id = 0
      formData.code = ''
      formData.name = ''
      formData.sort = 0
      formData.status = 1
      formData.remark = ''
    }
  }

  watch(
    () => [props.visible, props.type, props.postData],
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
