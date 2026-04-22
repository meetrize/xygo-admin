<!-- 悠然联系人 编辑弹窗 -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="type === 'add' ? '新增悠然联系人' : '编辑悠然联系人'"
    width="600px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <!-- 主键隐藏，不在表单中显示 -->
      <ElFormItem label="姓名" prop="username">
        <ElInput v-model="formData.username" placeholder="请输入姓名" />
      </ElFormItem>
      <ElFormItem label="电话" prop="phone">
        <ElInput v-model="formData.phone" placeholder="请输入电话" />
      </ElFormItem>
      <ElFormItem label="数字" prop="age">
        <ElInputNumber v-model="formData.age" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="头像" prop="avatar">
        <ArtFileSelector v-model="formData.avatar" file-type="image" />
      </ElFormItem>
      <ElFormItem label="排序权重" prop="sort">
        <ElInputNumber v-model="formData.sort" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="备注" prop="remark">
        <ElInput v-model="formData.remark" placeholder="请输入备注" />
      </ElFormItem>
      <ElFormItem label="开关" prop="switchField">
        <ElInputNumber v-model="formData.switchField" controls-position="right" />
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
  import { adminRequest } from '@/utils/http'
  import ArtFileSelector from '@/components/core/forms/art-file-selector/index.vue'
  import ArtIconSelector from '@/components/core/forms/art-icon-selector/index.vue'
  import ArtWangEditor from '@/components/core/forms/art-wang-editor/index.vue'

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
    username: '',
    phone: '',
    age: 0,
    avatar: '',
    sort: 0,
    remark: '',
    switchField: 0,
  })

  const formData = reactive(defaultForm())

  const rules = reactive<FormRules>({
    username: [{ required: true, message: '姓名不能为空', trigger: 'blur' }],
    phone: [{ required: true, message: '电话不能为空', trigger: 'blur' }],
    age: [{ required: true, message: '数字不能为空', trigger: 'blur' }],
    avatar: [{ required: true, message: '头像不能为空', trigger: 'blur' }],
    sort: [{ required: true, message: '排序权重不能为空', trigger: 'blur' }],
    remark: [{ required: true, message: '备注不能为空', trigger: 'blur' }],
    switchField: [{ required: true, message: '开关不能为空', trigger: 'blur' }],
  })

  watch(() => props.visible, (val) => {
    if (val && props.type === 'edit' && props.editData) {
      Object.assign(formData, props.editData)
    } else if (val) {
      Object.assign(formData, defaultForm())
    }
  })

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
