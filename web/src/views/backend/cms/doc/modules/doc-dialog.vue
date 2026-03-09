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
    :title="dialogType === 'add' ? '新增文档' : '编辑文档'"
    width="90%"
    top="3vh"
    align-center
    :close-on-click-modal="false"
    destroy-on-close
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="90px">
      <ElRow :gutter="20">
        <ElCol :span="12">
          <ElFormItem label="文档标题" prop="title">
            <ElInput v-model="formData.title" placeholder="请输入文档标题" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="12">
          <ElFormItem label="所属分类" prop="categoryId">
            <ElTreeSelect
              v-model="formData.categoryId"
              :data="categoryTree"
              :props="{ label: 'title', value: 'id', children: 'children' }"
              placeholder="请选择分类"
              check-strictly
              style="width: 100%"
            />
          </ElFormItem>
        </ElCol>
      </ElRow>
      <ElRow :gutter="20">
        <ElCol :span="8">
          <ElFormItem label="URL标识" prop="slug">
            <ElInput v-model="formData.slug" placeholder="留空自动生成" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="8">
          <ElFormItem label="作者">
            <ElInput v-model="formData.author" placeholder="请输入作者" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="4">
          <ElFormItem label="排序">
            <ElInputNumber v-model="formData.sort" :min="0" controls-position="right" style="width: 100%" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="4">
          <ElFormItem label="状态">
            <ElSelect v-model="formData.status" style="width: 100%">
              <ElOption label="已发布" :value="1" />
              <ElOption label="草稿" :value="2" />
              <ElOption label="下架" :value="3" />
            </ElSelect>
          </ElFormItem>
        </ElCol>
      </ElRow>
      <ElRow :gutter="20">
        <ElCol :span="16">
          <ElFormItem label="摘要">
            <ElInput v-model="formData.summary" type="textarea" :rows="2" placeholder="请输入摘要" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="4">
          <ElFormItem label="置顶">
            <ElSwitch v-model="formData.isTop" :active-value="1" :inactive-value="0" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="4">
          <ElFormItem label="标签">
            <ElInput v-model="formData.tags" placeholder='["标签1"]' />
          </ElFormItem>
        </ElCol>
      </ElRow>
      <ElFormItem label="文档内容" prop="content">
        <div style="width: 100%; border: 1px solid var(--el-border-color); border-radius: 4px;">
          <MdEditor
            v-model="formData.content"
            :style="{ height: '500px' }"
            :preview="true"
            :toolbarsExclude="['github', 'prettier', 'mermaid', 'katex']"
            :noPrettier="true"
            :noMermaid="true"
            :noKatex="true"
            @onUploadImg="handleUploadImg"
          />
        </div>
      </ElFormItem>
    </ElForm>
    <template #footer>
      <ElButton @click="dialogVisible = false">取消</ElButton>
      <ElButton type="primary" @click="handleSubmit" :loading="submitting">保存</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import { ref, reactive } from 'vue'
  import type { FormInstance, FormRules } from 'element-plus'
  import { MdEditor } from 'md-editor-v3'
  import 'md-editor-v3/lib/style.css'
  import { adminRequest } from '@/utils/http'

  defineProps<{ categoryTree: any[] }>()
  const emit = defineEmits<{ submit: [data: any] }>()

  const dialogVisible = ref(false)
  const dialogType = ref<'add' | 'edit'>('add')
  const submitting = ref(false)
  const formRef = ref<FormInstance>()

  const defaultForm = () => ({
    id: undefined as number | undefined,
    categoryId: undefined as number | undefined,
    title: '',
    slug: '',
    cover: '',
    summary: '',
    content: '',
    author: '',
    sort: 0,
    status: 1,
    isTop: 0,
    tags: ''
  })

  const formData = reactive(defaultForm())

  const rules: FormRules = {
    title: [{ required: true, message: '请输入文档标题', trigger: 'blur' }],
    categoryId: [{ required: true, message: '请选择分类', trigger: 'change' }],
    content: [{ required: true, message: '请输入文档内容', trigger: 'blur' }]
  }

  /** 图片上传回调 */
  const handleUploadImg = async (files: File[], callback: (urls: string[]) => void) => {
    const urls: string[] = []
    for (const file of files) {
      const formData = new FormData()
      formData.append('file', file)
      try {
        const res = await adminRequest.post<any>({
          url: '/upload/file',
          params: formData,
          headers: { 'Content-Type': 'multipart/form-data' }
        })
        if (res?.url) urls.push(res.url)
      } catch (e) {
        console.error('upload failed', e)
      }
    }
    callback(urls)
  }

  const open = (type: 'add' | 'edit', data?: any) => {
    dialogType.value = type
    Object.assign(formData, defaultForm(), data || {})
    dialogVisible.value = true
    formRef.value?.clearValidate()
  }

  const handleSubmit = async () => {
    if (!formRef.value) return
    await formRef.value.validate(async (valid) => {
      if (valid) {
        submitting.value = true
        try {
          emit('submit', { ...formData })
          dialogVisible.value = false
        } finally {
          submitting.value = false
        }
      }
    })
  }

  defineExpose({ open })
</script>
