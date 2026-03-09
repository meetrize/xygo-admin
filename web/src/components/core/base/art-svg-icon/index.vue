<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 图标组件 -->
<template>
  <Icon v-if="icon" :icon="icon" v-bind="bindAttrs" class="art-svg-icon inline" />
</template>

<script setup lang="ts">
  import { Icon } from '@iconify/vue'

  defineOptions({ name: 'ArtSvgIcon', inheritAttrs: false })

  interface Props {
    /** Iconify icon name */
    icon?: string
  }

  defineProps<Props>()

  const attrs = useAttrs()

  const bindAttrs = computed<{ class: string; style: string }>(() => ({
    class: (attrs.class as string) || '',
    style: (attrs.style as string) || ''
  }))
</script>
