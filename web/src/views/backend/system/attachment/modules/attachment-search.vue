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
  <ArtSearchBar
    ref="searchBarRef"
    v-model="formData"
    :items="formItems"
    :showExpand="false"
    @reset="handleReset"
    @search="handleSearch"
  />
</template>

<script setup lang="ts">
  interface Props {
    modelValue: Record<string, any>
  }
  interface Emits {
    (e: 'update:modelValue', value: Record<string, any>): void
    (e: 'search'): void
    (e: 'reset'): void
  }
  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  const searchBarRef = ref()
  const formData = computed({
    get: () => props.modelValue,
    set: (val) => emit('update:modelValue', val)
  })

  const formItems = computed(() => [
    {
      label: '文件分类',
      key: 'topic',
      type: 'select',
      props: {
        clearable: true,
        placeholder: '请选择文件分类',
        options: [
          { label: '图片', value: 'image' },
          { label: '视频', value: 'video' },
          { label: '音频', value: 'audio' },
          { label: '文档', value: 'doc' },
          { label: '压缩包', value: 'archive' },
          { label: '其他', value: 'other' }
        ]
      }
    },
    {
      label: '存储方式',
      key: 'storage',
      type: 'select',
      props: {
        clearable: true,
        placeholder: '请选择存储方式',
        options: [
          { label: '本地存储', value: 'local' },
          { label: '阿里云OSS', value: 'aliyun-oss' },
          { label: '腾讯云COS', value: 'tencent-cos' },
          { label: '七牛云', value: 'qiniu' }
        ]
      }
    }
  ])

  function handleReset() {
    emit('reset')
  }

  function handleSearch() {
    emit('search')
  }
</script>
