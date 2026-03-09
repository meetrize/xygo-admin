<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 积分变动日志 搜索栏 -->
<template>
  <ArtSearchBar
    v-model="formFilters"
    :items="formItems"
    @reset="handleReset"
    @search="handleSearch"
  />
</template>

<script setup lang="ts">
  const props = defineProps<{ modelValue: Record<string, any> }>()
  const emit = defineEmits<{
    (e: 'update:modelValue', v: Record<string, any>): void
    (e: 'search', params: Record<string, any>): void
    (e: 'reset'): void
  }>()

  const formFilters = reactive({ ...props.modelValue })

  const formItems = computed(() => [
    {
      label: '变动积分',
      key: 'scoreStart',
      span: 8,
      render: () => h('div', { class: 'art-number-range' }, [
        h('input', { type: 'number', value: formFilters.scoreStart ?? '', placeholder: '最小值', class: 'art-number-range__input', onInput: (e: any) => formFilters.scoreStart = e.target.value === '' ? undefined : Number(e.target.value) }),
        h('span', { class: 'art-number-range__separator' }, '至'),
        h('input', { type: 'number', value: formFilters.scoreEnd ?? '', placeholder: '最大值', class: 'art-number-range__input', onInput: (e: any) => formFilters.scoreEnd = e.target.value === '' ? undefined : Number(e.target.value) }),
      ])
    },
    {
      label: '用户名',
      key: 'member_username',
      type: 'input',
      props: { clearable: true }
    },
    {
      label: '昵称',
      key: 'member_nickname',
      type: 'input',
      props: { clearable: true }
    },
  ])

  const handleSearch = () => {
    const params: Record<string, any> = { ...formFilters }
    emit('update:modelValue', params)
    emit('search', params)
  }

  const handleReset = () => {
    Object.keys(formFilters).forEach(k => (formFilters[k] = undefined))
    emit('update:modelValue', { ...formFilters })
    emit('reset')
  }
</script>
<style lang="scss">
.art-number-range {
  display: inline-flex;
  align-items: center;
  width: 100%;
  height: 32px;
  padding: 0 8px;
  background-color: var(--el-fill-color-blank);
  border: 1px solid var(--el-border-color);
  border-radius: var(--el-border-radius-base);
  transition: border-color 0.2s;
  &:hover { border-color: var(--el-border-color-hover); }
  &:focus-within { border-color: var(--el-color-primary); }

  &__input {
    flex: 1;
    min-width: 0;
    height: 100%;
    padding: 0 4px;
    font-size: 13px;
    color: var(--el-text-color-regular);
    text-align: center;
    background: transparent;
    border: none;
    outline: none;
    appearance: textfield;
    -moz-appearance: textfield;
    &::-webkit-inner-spin-button,
    &::-webkit-outer-spin-button { appearance: none; margin: 0; }
    &::placeholder { color: var(--el-text-color-placeholder); }
  }

  &__separator {
    flex-shrink: 0;
    padding: 0 6px;
    font-size: 13px;
    color: var(--el-text-color-placeholder);
  }
}
</style>
