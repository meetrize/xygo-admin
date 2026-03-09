<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 颜色选择器组件 -->
<template>
  <div class="art-color-picker">
    <ElInput
      v-model="colorValue"
      :placeholder="placeholder"
      clearable
      @change="handleChange"
    >
      <template #append>
        <div class="color-picker-wrapper">
          <ElColorPicker
            v-model="colorValue"
            :predefine="predefineColors"
            :show-alpha="showAlpha"
            :color-format="colorFormat"
            @change="handleChange"
          
          />
        </div>
      </template>
    </ElInput>
  </div>
</template>

<script setup lang="ts">
defineOptions({ name: 'ArtColorPicker' })

interface Props {
  modelValue?: string
  placeholder?: string
  showAlpha?: boolean
  colorFormat?: 'hex' | 'rgb' | 'hsl' | 'hsv'
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: '请选择颜色',
  showAlpha: false,
  colorFormat: 'hex'
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
  'change': [value: string]
}>()

const colorValue = computed({
  get: () => props.modelValue || '',
  set: (val) => {
    emit('update:modelValue', val || '')
  }
})

// 预定义颜色
const predefineColors = [
  '#409EFF', // 主题蓝
  '#67C23A', // 成功绿
  '#E6A23C', // 警告橙
  '#F56C6C', // 危险红
  '#909399', // 信息灰
  '#303133', // 主要文字
  '#606266', // 常规文字
  '#909399', // 次要文字
  '#C0C4CC', // 占位文字
  '#DCDFE6', // 一级边框
  '#E4E7ED', // 二级边框
  '#EBEEF5', // 三级边框
  '#F2F6FC', // 基础白色
  '#000000', // 黑色
  '#FFFFFF', // 白色
]

const handleChange = (value: string | null) => {
  emit('change', value || '')
}
</script>

<style scoped lang="scss">
.art-color-picker {
  width: 100%;
  max-width: 400px;
}

.color-picker-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 4px;
}
</style>
