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
    <SectionTitle :title="$t('setting.color.title')" class="mt-10" />
    <div class="-mr-4">
      <div class="flex flex-wrap">
        <div
          v-for="color in configOptions.mainColors"
          :key="color"
          class="flex items-center justify-center size-[23px] mr-4 mb-2.5 cursor-pointer rounded-full transition-all duration-200 hover:opacity-85"
          :style="{ background: `${color} !important` }"
          @click="colorHandlers.selectColor(color)"
        >
          <ArtSvgIcon
            icon="ri:check-fill"
            class="text-base !text-white"
            v-show="color === systemThemeColor"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import SectionTitle from './SectionTitle.vue'
  import { useSettingStore } from '@/store/modules/setting'
  import { useSettingsConfig } from '../composables/useSettingsConfig'
  import { useSettingsHandlers } from '../composables/useSettingsHandlers'
  import { storeToRefs } from 'pinia'

  const settingStore = useSettingStore()
  const { systemThemeColor } = storeToRefs(settingStore)
  const { configOptions } = useSettingsConfig()
  const { colorHandlers } = useSettingsHandlers()
</script>
