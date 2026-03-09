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
  <SectionTitle :title="$t('setting.menu.title')" />
  <div class="setting-box-wrap">
    <div
      class="setting-item"
      v-for="item in menuThemeList"
      :key="item.theme"
      @click="switchMenuStyles(item.theme)"
    >
      <div
        class="box"
        :class="{ 'is-active': item.theme === menuThemeType }"
        :style="{
          cursor: disabled ? 'no-drop' : 'pointer'
        }"
      >
        <img :src="item.img" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import AppConfig from '@/config'
  import SectionTitle from './SectionTitle.vue'
  import { MenuTypeEnum, type MenuThemeEnum } from '@/enums/appEnum'
  import { useSettingStore } from '@/store/modules/setting'

  const menuThemeList = AppConfig.themeList
  const settingStore = useSettingStore()
  const { menuThemeType, menuType, isDark } = storeToRefs(settingStore)
  const isTopMenu = computed(() => menuType.value === MenuTypeEnum.TOP)
  const isDualMenu = computed(() => menuType.value === MenuTypeEnum.DUAL_MENU)

  const disabled = computed(() => isTopMenu.value || isDualMenu.value || isDark.value)

  // 菜单样式切换
  const switchMenuStyles = (theme: MenuThemeEnum) => {
    if (isDualMenu.value || isTopMenu.value || isDark.value) {
      return
    }
    settingStore.switchMenuStyles(theme)
  }
</script>
