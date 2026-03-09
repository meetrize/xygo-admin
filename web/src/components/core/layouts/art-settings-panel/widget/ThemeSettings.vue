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
  <SectionTitle :title="$t('setting.theme.title')" />
  <div class="setting-box-wrap">
    <div
      class="setting-item"
      v-for="(item, index) in configOptions.themeList"
      :key="item.theme"
      @click="switchThemeStyles(item.theme)"
    >
      <div class="box" :class="{ 'is-active': item.theme === systemThemeMode }">
        <img :src="item.img" />
      </div>
      <p class="name">{{ $t(`setting.theme.list[${index}]`) }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
  import SectionTitle from './SectionTitle.vue'
  import { useSettingStore } from '@/store/modules/setting'
  import { useSettingsConfig } from '../composables/useSettingsConfig'
  import { useTheme } from '@/hooks/core/useTheme'

  const settingStore = useSettingStore()
  const { systemThemeMode } = storeToRefs(settingStore)
  const { configOptions } = useSettingsConfig()
  const { switchThemeStyles } = useTheme()
</script>
