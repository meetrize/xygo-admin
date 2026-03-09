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
  <div>
    <div class="flex justify-end">
      <div
        @click="$emit('close')"
        class="flex-cc c-p size-7.5 !transition-all duration-200 rounded hover:bg-g-300/80"
      >
        <ArtSvgIcon icon="ri:close-fill" class="block text-xl text-g-600" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  defineEmits<{
    close: []
  }>()
</script>
