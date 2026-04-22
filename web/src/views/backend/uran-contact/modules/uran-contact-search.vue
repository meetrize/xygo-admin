<!-- 悠然联系人 搜索栏 -->
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
