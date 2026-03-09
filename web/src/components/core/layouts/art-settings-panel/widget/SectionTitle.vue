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
  <p
    class="relative mt-7.5 mb-5.5 text-sm text-center text-g-800 before:absolute before:top-[10px] before:left-0 before:w-[50px] before:m-auto before:content-[''] before:border-b before:border-[var(--art-gray-300)] after:absolute after:top-[10px] after:right-0 after:w-[50px] after:m-auto after:content-[''] after:border-b after:border-g-300"
    :style="style"
  >
    {{ title }}
  </p>
</template>

<script setup lang="ts">
  interface Props {
    title: string
    style?: Record<string, any>
  }

  defineProps<Props>()
</script>
