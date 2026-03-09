<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 定时任务新增/编辑弹窗 -->
<template>
  <ElDialog v-model="dialogVisible" :title="isEdit ? '编辑任务' : '新增任务'" width="720px" align-center :close-on-click-modal="false">
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="100px">
      <ElFormItem label="任务标题" prop="title">
        <ElInput v-model="formData.title" placeholder="请输入任务标题" />
      </ElFormItem>
      <ElFormItem label="任务标识" prop="name">
        <ElSelect v-model="formData.name" placeholder="选择已注册的任务" filterable :disabled="isEdit" class="!w-full">
          <ElOption v-for="t in registeredTasks" :key="t" :label="t" :value="t" />
        </ElSelect>
      </ElFormItem>
      <ElFormItem label="执行周期" prop="pattern">
        <ArtCronDesigner v-model="formData.pattern" />
      </ElFormItem>
      <ElFormItem label="分组">
        <ElSelect v-model="formData.groupId" placeholder="选择分组" clearable class="!w-full">
          <ElOption v-for="g in groupOptions" :key="g.id" :label="g.name" :value="g.id" />
        </ElSelect>
      </ElFormItem>
      <ElRow :gutter="16">
        <ElCol :span="12">
          <ElFormItem label="策略" prop="policy">
            <ElSelect v-model="formData.policy" class="!w-full">
              <ElOption label="并行执行" :value="1" />
              <ElOption label="单例执行" :value="2" />
              <ElOption label="单次执行" :value="3" />
              <ElOption label="固定次数" :value="4" />
            </ElSelect>
          </ElFormItem>
        </ElCol>
        <ElCol :span="12">
          <ElFormItem v-if="formData.policy === 4" label="执行次数">
            <ElInputNumber v-model="formData.count" :min="1" :max="9999" class="!w-full" />
          </ElFormItem>
          <ElFormItem v-else label="排序">
            <ElInputNumber v-model="formData.sort" :min="0" class="!w-full" />
          </ElFormItem>
        </ElCol>
      </ElRow>
      <ElFormItem label="任务参数">
        <ElInput v-model="formData.params" placeholder="多个参数用逗号分隔" />
      </ElFormItem>
      <ElFormItem label="备注">
        <ElInput v-model="formData.remark" type="textarea" :rows="2" placeholder="备注说明" />
      </ElFormItem>
      <ElFormItem label="状态">
        <ElSwitch v-model="formData.status" :active-value="1" :inactive-value="0" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <ElButton @click="dialogVisible = false">取消</ElButton>
      <ElButton type="primary" :loading="submitting" @click="handleSubmit">提交</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'
  import { ElMessage } from 'element-plus'
  import { fetchCronSave, fetchCronRegisteredTasks, type CronItem } from '@/api/backend/system/cron'
  import ArtCronDesigner from '@/components/core/forms/art-cron-designer/index.vue'

  interface Props {
    visible: boolean
    data?: CronItem | null
    groupOptions: { id: number; name: string }[]
  }
  interface Emits {
    (e: 'update:visible', v: boolean): void
    (e: 'success'): void
  }

  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: v => emit('update:visible', v)
  })
  const isEdit = computed(() => !!props.data?.id)

  const formRef = ref<FormInstance>()
  const submitting = ref(false)
  const registeredTasks = ref<string[]>([])

  const formData = reactive({
    id: 0,
    title: '',
    name: '',
    groupId: 0 as number | undefined,
    pattern: '',
    policy: 1,
    count: 1,
    sort: 0,
    params: '',
    remark: '',
    status: 1
  })

  const rules: FormRules = {
    title: [{ required: true, message: '请输入任务标题', trigger: 'blur' }],
    name: [{ required: true, message: '请选择任务标识', trigger: 'change' }],
    pattern: [{ required: true, message: '请输入Cron表达式', trigger: 'blur' }]
  }

  const loadRegisteredTasks = async () => {
    try {
      const res = await fetchCronRegisteredTasks() as any
      registeredTasks.value = res?.list || []
    } catch { /* ignore */ }
  }

  watch(() => props.visible, (val) => {
    if (val) {
      loadRegisteredTasks()
      if (props.data) {
        Object.assign(formData, {
          id: props.data.id,
          title: props.data.title,
          name: props.data.name,
          groupId: props.data.groupId || undefined,
          pattern: props.data.pattern,
          policy: props.data.policy,
          count: props.data.count || 1,
          sort: props.data.sort,
          params: props.data.params,
          remark: props.data.remark,
          status: props.data.status
        })
      } else {
        Object.assign(formData, { id: 0, title: '', name: '', groupId: undefined, pattern: '', policy: 1, count: 1, sort: 0, params: '', remark: '', status: 1 })
      }
      nextTick(() => formRef.value?.clearValidate())
    }
  })

  const handleSubmit = async () => {
    if (!formRef.value) return
    await formRef.value.validate(async (valid) => {
      if (!valid) return
      submitting.value = true
      try {
        await fetchCronSave({ ...formData })
        ElMessage.success(isEdit.value ? '编辑成功' : '新增成功')
        emit('success')
        dialogVisible.value = false
      } catch { /* ignore */ } finally {
        submitting.value = false
      }
    })
  }
</script>
