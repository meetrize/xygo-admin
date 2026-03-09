// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

import { useSettingStore } from '@/store/modules/setting'
import { MenuThemeEnum, MenuTypeEnum } from '@/enums/appEnum'

/**
 * 设置状态管理
 */
export function useSettingsState() {
  const settingStore = useSettingStore()

  // 色弱模式初始化
  const initColorWeak = () => {
    if (settingStore.colorWeak) {
      const el = document.getElementsByTagName('html')[0]
      setTimeout(() => {
        el.classList.add('color-weak')
      }, 100)
    }
  }

  // 菜单布局切换
  const switchMenuLayouts = (type: MenuTypeEnum) => {
    if (type === MenuTypeEnum.LEFT || type === MenuTypeEnum.TOP_LEFT) {
      settingStore.setMenuOpen(true)
    }
    settingStore.switchMenuLayouts(type)
    if (type === MenuTypeEnum.DUAL_MENU) {
      settingStore.switchMenuStyles(MenuThemeEnum.DESIGN)
      settingStore.setMenuOpen(true)
    }
  }

  return {
    // 方法
    initColorWeak,
    switchMenuLayouts
  }
}
