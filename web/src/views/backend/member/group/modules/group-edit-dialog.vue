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
    v-model="visible"
    :title="dialogType === 'add' ? '新增分组' : '编辑分组'"
    width="500px"
    :close-on-click-modal="false"
  >
    <ElForm ref="formRef" :model="form" :rules="rules" label-width="100px">
      <ElFormItem label="分组名称" prop="name">
        <ElInput v-model="form.name" placeholder="请输入分组名称" />
      </ElFormItem>
      <ElFormItem label="排序" prop="sort">
        <ElInputNumber v-model="form.sort" :min="0" :max="9999" />
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElSwitch v-model="form.status" :active-value="1" :inactive-value="0" />
      </ElFormItem>
      <ElFormItem label="备注" prop="remark">
        <ElInput v-model="form.remark" type="textarea" :rows="3" placeholder="请输入备注" />
      </ElFormItem>
    </ElForm>

    <template #footer>
      <ElButton @click="visible = false">取消</ElButton>
      <ElButton type="primary" :loading="submitLoading" @click="handleSubmit">确定</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'
  import { saveMemberGroup, type MemberGroupItem } from '@/api/backend/member/group'

  interface Props {
    modelValue: boolean
    dialogType: 'add' | 'edit'
    groupData?: MemberGroupItem
  }

  interface Emits {
    (e: 'update:modelValue', value: boolean): void
    (e: 'success'): void
  }

  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  const visible = computed({
    get: () => props.modelValue,
    set: (val) => emit('update:modelValue', val)
  })

  const formRef = ref<FormInstance>()
  const submitLoading = ref(false)

  const form = reactive({
    id: 0,
    name: '',
    sort: 0,
    status: 1,
    remark: ''
  })

  const rules: FormRules = {
    name: [
      { required: true, message: '请输入分组名称', trigger: 'blur' },
      { min: 2, max: 32, message: '长度在 2 到 32 个字符', trigger: 'blur' }
    ]
  }

  // 监听弹窗打开，初始化表单
  watch(
    () => props.modelValue,
    (val) => {
      if (val) {
        initForm()
        nextTick(() => {
          formRef.value?.clearValidate()
        })
      }
    }
  )

  const initForm = () => {
    if (props.dialogType === 'edit' && props.groupData) {
      form.id = props.groupData.id
      form.name = props.groupData.name
      form.sort = props.groupData.sort || 0
      form.status = props.groupData.status
      form.remark = props.groupData.remark || ''
    } else {
      form.id = 0
      form.name = ''
      form.sort = 0
      form.status = 1
      form.remark = ''
    }
  }

  const handleSubmit = async () => {
    if (!formRef.value) return

    const valid = await formRef.value.validate().catch(() => false)
    if (!valid) return

    submitLoading.value = true
    try {
      await saveMemberGroup(form)
      ElMessage.success(form.id ? '编辑成功' : '新增成功')
      visible.value = false
      emit('success')
    } catch (error) {
      console.error('保存分组失败:', error)
    } finally {
      submitLoading.value = false
    }
  }
</script>
