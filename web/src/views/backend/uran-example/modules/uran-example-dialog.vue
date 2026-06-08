<!-- 悠然示例 编辑弹窗 -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="type === 'add' ? '新增悠然示例' : '编辑悠然示例'"
    width="600px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <!-- 主键隐藏，不在表单中显示 -->
      <ElFormItem label="状态" prop="status">
        <ElRadioGroup v-model="formData.status">
          <ElRadio :value="0">禁用</ElRadio>
          <ElRadio :value="1">启用</ElRadio>
        </ElRadioGroup>
      </ElFormItem>
      <ElFormItem label="排序" prop="sort">
        <ElInputNumber v-model="formData.sort" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="标题" prop="tname">
        <ElInput v-model="formData.tname" placeholder="请输入标题" />
      </ElFormItem>
      <ElFormItem label="图片" prop="timage">
        <ArtFileSelector v-model="formData.timage" file-type="image" />
      </ElFormItem>
      <ElFormItem label="测试" prop="ttest">
        <ElInput v-model="formData.ttest" placeholder="请输入测试" />
      </ElFormItem>
    </ElForm>

    <template #footer>
      <ElButton @click="handleClose">取消</ElButton>
      <ElButton type="primary" :loading="loading" @click="handleSubmit">确定</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'
  import type { DialogType } from '@/types'
  import ArtFileSelector from '@/components/core/forms/art-file-selector/index.vue'

  const props = defineProps<{
    visible: boolean
    type: DialogType
    editData?: Record<string, any>
  }>()

  const emit = defineEmits<{
    (e: 'update:visible', v: boolean): void
    (e: 'submit', data: Record<string, any>): void
  }>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: (val: boolean) => emit('update:visible', val)
  })

  const formRef = ref<FormInstance>()
  const loading = ref(false)

  const defaultForm = (): Record<string, any> => ({
    id: 0,
    status: 0,
    sort: 0,
    tname: '',
    timage: '',
    ttest: ''
  })

  const formData = reactive(defaultForm())

  const rules = reactive<FormRules>({
    status: [{ required: true, message: '状态不能为空', trigger: 'blur' }],
    sort: [{ required: true, message: '排序不能为空', trigger: 'blur' }],
    tname: [{ required: true, message: '标题不能为空', trigger: 'blur' }],
    timage: [{ required: true, message: '图片不能为空', trigger: 'blur' }],
    ttest: [{ required: true, message: '测试不能为空', trigger: 'blur' }]
  })

  watch(
    () => props.visible,
    (val) => {
      if (val && props.type === 'edit' && props.editData) {
        Object.assign(formData, props.editData)
      } else if (val) {
        Object.assign(formData, defaultForm())
      }
    }
  )

  const handleSubmit = async () => {
    if (!formRef.value) return
    await formRef.value.validate()
    emit('submit', { ...formData })
  }

  const handleClose = () => {
    formRef.value?.resetFields()
    Object.assign(formData, defaultForm())
    dialogVisible.value = false
  }
</script>
