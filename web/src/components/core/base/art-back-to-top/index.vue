<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 返回顶部按钮 -->
<template>
  <Transition
    enter-active-class="tad-300 ease-out"
    leave-active-class="tad-200 ease-in"
    enter-from-class="opacity-0 translate-y-2"
    enter-to-class="opacity-100 translate-y-0"
    leave-from-class="opacity-100 translate-y-0"
    leave-to-class="opacity-0 translate-y-2"
  >
    <div
      v-show="showButton"
      class="fixed right-10 bottom-15 size-9.5 flex-cc c-p border border-g-300 rounded-md tad-300 hover:bg-g-200"
      @click="scrollToTop"
    >
      <ArtSvgIcon icon="ri:arrow-up-wide-line" class="text-g-500 text-lg" />
    </div>
  </Transition>
</template>

<script setup lang="ts">
  import { useCommon } from '@/hooks/core/useCommon'

  defineOptions({ name: 'ArtBackToTop' })

  const { scrollToTop } = useCommon()

  const showButton = ref(false)
  const scrollThreshold = 300

  onMounted(() => {
    const scrollContainer = document.getElementById('app-main')
    if (scrollContainer) {
      const { y } = useScroll(scrollContainer)
      watch(y, (newY: number) => {
        showButton.value = newY > scrollThreshold
      })
    }
  })
</script>
