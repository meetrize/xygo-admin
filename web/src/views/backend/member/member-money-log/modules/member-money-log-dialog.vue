<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 余额变动日志 编辑弹窗 -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="type === 'add' ? '新增余额变动日志' : '编辑余额变动日志'"
    width="600px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <!-- 主键隐藏，不在表单中显示 -->
      <ElFormItem label="会员ID" prop="memberId">
        <ElSelect
          v-model="formData.memberId"
          filterable
          remote
          :remote-method="(q: string) => loadMemberOptions(q)"
          placeholder="请选择会员ID"
          clearable
          :loading="memberLoading"
        >
          <ElOption
            v-for="opt in memberOptions"
            :key="opt.value"
            :label="opt.label"
            :value="opt.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="变动金额" prop="money">
        <ElInputNumber v-model="formData.money" :min="0" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="变动前余额（分）" prop="before">
        <ElInputNumber v-model="formData.before" :min="0" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="变动后余额（分）" prop="after">
        <ElInputNumber v-model="formData.after" :min="0" controls-position="right" />
      </ElFormItem>
      <ElFormItem label="变动说明" prop="memo">
        <ElInput v-model="formData.memo" placeholder="请输入变动说明" />
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
    memberId: 0,
    money: 0,
    before: 0,
    after: 0,
    memo: '',
  })

  const formData = reactive(defaultForm())

  const rules = reactive<FormRules>({
    memberId: [{ required: true, message: '会员ID不能为空', trigger: 'blur' }],
    money: [{ required: true, message: '变动金额不能为空', trigger: 'blur' }],
    before: [{ required: true, message: '变动前余额（分）不能为空', trigger: 'blur' }],
    after: [{ required: true, message: '变动后余额（分）不能为空', trigger: 'blur' }],
    memo: [{ required: true, message: '变动说明不能为空', trigger: 'blur' }],
  })

  // ==================== 远程下拉选项 ====================
  const memberOptions = ref<{ value: any; label: string }[]>([])
  const memberLoading = ref(false)
  const loadMemberOptions = async (query: string) => {
    memberLoading.value = true
    try {
      const res = await adminRequest.get<any>({
        url: '/member/list',
        params: { pageSize: 50, nickname: query || undefined }
      })
      memberOptions.value = (res.list || []).map((item: any) => ({
        value: item.id,
        label: item.nickname,
      }))
    } catch { /* ignore */ }
    memberLoading.value = false
  }

  watch(() => props.visible, (val) => {
    if (val && props.type === 'edit' && props.editData) {
      Object.assign(formData, props.editData)
      // 编辑时加载已选关联项
      loadMemberOptions('')
    } else if (val) {
      Object.assign(formData, defaultForm())
      loadMemberOptions('')
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
