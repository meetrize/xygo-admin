<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 会员通知 编辑弹窗 -->
<template>
  <ElDialog
    v-model="dialogVisible"
    :title="type === 'add' ? '新增会员通知' : '编辑会员通知'"
    width="600px"
    :close-on-click-modal="false"
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <!-- 主键隐藏，不在表单中显示 -->
      <ElFormItem label="通知标题" prop="title">
        <ElInput v-model="formData.title" placeholder="请输入通知标题" />
      </ElFormItem>
      <ElFormItem label="通知内容" prop="content">
        <ArtWangEditor v-model="formData.content" placeholder="请输入通知内容" />
      </ElFormItem>
      <ElFormItem label="通知类型" prop="type">
        <ElSelect v-model="formData.type" placeholder="请选择通知类型" clearable>
          <ElOption :value="'system'" label="系统通知" />
          <ElOption :value="'announce'" label="公告" />
          <ElOption :value="'feature'" label="功能更新" />
          <ElOption :value="'maintain'" label="维护通知" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="目标" prop="target">
        <ElRadioGroup v-model="formData.target">
          <ElRadio :value="'all'">全部会员</ElRadio>
          <ElRadio :value="'group'">指定分组</ElRadio>
        </ElRadioGroup>
      </ElFormItem>
      <ElFormItem label="目标分组ID" prop="targetId">
        <ElSelect
          v-model="formData.targetId"
          filterable
          remote
          :remote-method="(q: string) => loadTargetOptions(q)"
          placeholder="请选择目标分组ID"
          clearable
          :loading="targetLoading"
        >
          <ElOption
            v-for="opt in targetOptions"
            :key="opt.value"
            :label="opt.label"
            :value="opt.value"
          />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="发送者" prop="sender">
        <ElInput v-model="formData.sender" placeholder="请输入发送者" />
      </ElFormItem>
      <ElFormItem label="状态" prop="status">
        <ElRadioGroup v-model="formData.status">
          <ElRadio :value="0">草稿</ElRadio>
          <ElRadio :value="1">已发布</ElRadio>
        </ElRadioGroup>
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
    title: '',
    content: '',
    type: '',
    target: '',
    targetId: 0,
    sender: '',
    status: 0,
  })

  const formData = reactive(defaultForm())

  const rules = reactive<FormRules>({
    title: [{ required: true, message: '通知标题不能为空', trigger: 'blur' }],
    type: [{ required: true, message: '通知类型不能为空', trigger: 'blur' }],
    target: [{ required: true, message: '目标不能为空', trigger: 'blur' }],
    targetId: [{ required: true, message: '目标分组ID不能为空', trigger: 'blur' }],
    sender: [{ required: true, message: '发送者不能为空', trigger: 'blur' }],
    status: [{ required: true, message: '状态不能为空', trigger: 'blur' }],
  })

  // ==================== 远程下拉选项 ====================
  const targetOptions = ref<{ value: any; label: string }[]>([])
  const targetLoading = ref(false)
  const loadTargetOptions = async (query: string) => {
    targetLoading.value = true
    try {
      const res = await adminRequest.get<any>({
        url: '/member/group/list',
        params: { pageSize: 50, name: query || undefined }
      })
      targetOptions.value = (res.list || []).map((item: any) => ({
        value: item.id,
        label: item.name,
      }))
    } catch { /* ignore */ }
    targetLoading.value = false
  }

  watch(() => props.visible, (val) => {
    if (val && props.type === 'edit' && props.editData) {
      Object.assign(formData, props.editData)
      // 编辑时加载已选关联项
      loadTargetOptions('')
    } else if (val) {
      Object.assign(formData, defaultForm())
      loadTargetOptions('')
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
