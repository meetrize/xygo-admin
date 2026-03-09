<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 图标选择器组件 -->
<template>
  <div class="art-icon-selector">
    <ElInput
      :model-value="modelValue"
      placeholder="请选择图标"
      readonly
      @click="showDialog = true"
    >
      <template #prepend>
        <div class="icon-preview">
          <ArtSvgIcon v-if="modelValue" :icon="modelValue" :size="18" />
          <ArtSvgIcon v-else icon="ri:image-line" :size="18" />
        </div>
      </template>
      <template #append>
        <ElButton v-if="modelValue && clearable" @click.stop="handleClear">
          <ArtSvgIcon icon="ri:close-line" :size="16" />
        </ElButton>
        <div v-else class="search-icon-append">
          <ArtSvgIcon icon="ri:search-line" :size="16" />
        </div>
      </template>
    </ElInput>

    <!-- 图标选择弹窗 -->
    <ElDialog
      v-model="showDialog"
      title="选择图标"
      width="900px"
      :close-on-click-modal="false"
      class="icon-selector-dialog"
    >
      <div class="icon-selector-content">
        <!-- 搜索和筛选 -->
        <div class="selector-header">
          <div class="input-group">
            <ElInput
              v-model="searchText"
              placeholder="搜索图标名称..."
              clearable
              class="search-input"
            >
              <template #prefix>
                <ArtSvgIcon icon="ri:search-line" :size="16" />
              </template>
            </ElInput>
            <ElButton type="primary" plain @click="showCustomInput = !showCustomInput">
              <ArtSvgIcon icon="ri:edit-line" :size="16" style="margin-right: 4px;" />
              {{ showCustomInput ? '选择图标' : '自定义输入' }}
            </ElButton>
          </div>
          
          <ElRadioGroup v-if="!showCustomInput" v-model="currentCategory" class="category-tabs" size="default">
            <ElRadioButton label="all">全部</ElRadioButton>
            <ElRadioButton label="awe">Awesome</ElRadioButton>
            <ElRadioButton label="ali">Ali</ElRadioButton>
            <ElRadioButton label="local">本地</ElRadioButton>
          </ElRadioGroup>

          <!-- 自定义输入区域 -->
          <div v-if="showCustomInput" class="custom-input-area">
            <ElInput
              v-model="customIconName"
              placeholder="输入完整图标名称，如：ri:wechat-line"
              clearable
            >
              <template #prepend>
                <ArtSvgIcon v-if="customIconName" :icon="customIconName" :size="18" />
                <ArtSvgIcon v-else icon="ri:quill-pen-line" :size="18" />
              </template>
            </ElInput>
            <div class="custom-hint">
              <span>提示：访问 </span>
              <a href="https://icon-sets.iconify.design/" target="_blank" class="icon-link">
                Iconify 图标库
              </a>
              <span> 查找更多图标</span>
            </div>
          </div>
        </div>

        <!-- 图标列表 -->
        <div v-if="!showCustomInput" class="icon-list-wrapper">
          <ElScrollbar height="500px">
            <div v-if="filteredIcons.length > 0" class="icon-list">
              <div
                v-for="icon in filteredIcons"
                :key="icon"
                class="icon-item"
                :class="{ active: tempSelectedIcon === icon }"
                @click="handleSelectIcon(icon)"
              >
                <ArtSvgIcon :icon="icon" :size="24" />
                <div class="icon-name">{{ icon }}</div>
              </div>
            </div>
            <ElEmpty v-else description="未找到匹配的图标" :image-size="120" />
          </ElScrollbar>
        </div>

        <!-- 自定义输入预览 -->
        <div v-else class="custom-preview-area">
          <div class="preview-title">预览效果</div>
          <div class="preview-box">
            <div class="preview-icon-large">
              <ArtSvgIcon v-if="customIconName" :icon="customIconName" :size="64" />
              <span v-else class="preview-placeholder">输入图标名称查看预览</span>
            </div>
            <div v-if="customIconName" class="preview-icon-name">{{ customIconName }}</div>
          </div>
          <div class="preview-sizes">
            <div class="size-item">
              <span class="size-label">小</span>
              <ArtSvgIcon v-if="customIconName" :icon="customIconName" :size="16" />
            </div>
            <div class="size-item">
              <span class="size-label">中</span>
              <ArtSvgIcon v-if="customIconName" :icon="customIconName" :size="24" />
            </div>
            <div class="size-item">
              <span class="size-label">大</span>
              <ArtSvgIcon v-if="customIconName" :icon="customIconName" :size="32" />
            </div>
          </div>
        </div>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <div class="selected-info">
            <template v-if="modelValue">
              已选择: <strong>{{ modelValue }}</strong>
            </template>
          </div>
          <div>
            <ElButton @click="showDialog = false">取消</ElButton>
            <ElButton type="primary" @click="handleConfirm">确定</ElButton>
          </div>
        </div>
      </template>
    </ElDialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

defineOptions({ name: 'ArtIconSelector' })

interface Props {
  modelValue?: string
  clearable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: '',
  clearable: true
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
  'change': [value: string]
}>()

const showDialog = ref(false)
const searchText = ref('')
const currentCategory = ref('all')
const tempSelectedIcon = ref('')
const showCustomInput = ref(false)
const customIconName = ref('')

// 图标库数据 - Remix Icon 常用图标
const remixIcons = [
  // 系统图标
  'ri:home-line',
  'ri:home-fill',
  'ri:dashboard-line',
  'ri:dashboard-fill',
  'ri:settings-3-line',
  'ri:settings-3-fill',
  'ri:menu-line',
  'ri:menu-fill',
  'ri:close-line',
  'ri:close-fill',
  'ri:search-line',
  'ri:search-fill',
  'ri:add-line',
  'ri:add-fill',
  'ri:subtract-line',
  'ri:subtract-fill',
  'ri:delete-bin-line',
  'ri:delete-bin-fill',
  'ri:edit-line',
  'ri:edit-fill',
  'ri:save-line',
  'ri:save-fill',
  'ri:download-line',
  'ri:download-fill',
  'ri:upload-line',
  'ri:upload-fill',
  'ri:refresh-line',
  'ri:refresh-fill',
  'ri:loader-line',
  'ri:loader-fill',
  
  // 箭头图标
  'ri:arrow-up-line',
  'ri:arrow-down-line',
  'ri:arrow-left-line',
  'ri:arrow-right-line',
  'ri:arrow-up-circle-line',
  'ri:arrow-down-circle-line',
  'ri:arrow-left-circle-line',
  'ri:arrow-right-circle-line',
  
  // 用户相关
  'ri:user-line',
  'ri:user-fill',
  'ri:user-add-line',
  'ri:user-add-fill',
  'ri:user-settings-line',
  'ri:user-settings-fill',
  'ri:team-line',
  'ri:team-fill',
  'ri:account-circle-line',
  'ri:account-circle-fill',
  'ri:admin-line',
  'ri:admin-fill',
  
  // 文件图标
  'ri:file-line',
  'ri:file-fill',
  'ri:file-text-line',
  'ri:file-text-fill',
  'ri:file-list-line',
  'ri:file-list-fill',
  'ri:folder-line',
  'ri:folder-fill',
  'ri:folder-open-line',
  'ri:folder-open-fill',
  'ri:file-copy-line',
  'ri:file-copy-fill',
  
  // 消息通知
  'ri:notification-line',
  'ri:notification-fill',
  'ri:message-line',
  'ri:message-fill',
  'ri:mail-line',
  'ri:mail-fill',
  'ri:chat-1-line',
  'ri:chat-1-fill',
  'ri:chat-3-line',
  'ri:chat-3-fill',
  'ri:feedback-line',
  'ri:feedback-fill',
  
  // 状态图标
  'ri:checkbox-circle-line',
  'ri:checkbox-circle-fill',
  'ri:close-circle-line',
  'ri:close-circle-fill',
  'ri:error-warning-line',
  'ri:error-warning-fill',
  'ri:information-line',
  'ri:information-fill',
  'ri:question-line',
  'ri:question-fill',
  'ri:alert-line',
  'ri:alert-fill',
  
  // 媒体图标
  'ri:image-line',
  'ri:image-fill',
  'ri:image-add-line',
  'ri:image-add-fill',
  'ri:video-line',
  'ri:video-fill',
  'ri:camera-line',
  'ri:camera-fill',
  'ri:play-line',
  'ri:play-fill',
  'ri:pause-line',
  'ri:pause-fill',
  
  // 商业图标
  'ri:shopping-cart-line',
  'ri:shopping-cart-fill',
  'ri:shopping-bag-line',
  'ri:shopping-bag-fill',
  'ri:store-line',
  'ri:store-fill',
  'ri:bank-card-line',
  'ri:bank-card-fill',
  'ri:money-dollar-circle-line',
  'ri:money-dollar-circle-fill',
  'ri:wallet-line',
  'ri:wallet-fill',
  
  // 工具图标
  'ri:lock-line',
  'ri:lock-fill',
  'ri:lock-unlock-line',
  'ri:lock-unlock-fill',
  'ri:key-line',
  'ri:key-fill',
  'ri:shield-line',
  'ri:shield-fill',
  'ri:eye-line',
  'ri:eye-fill',
  'ri:eye-off-line',
  'ri:eye-off-fill',
  
  // 时间日期
  'ri:calendar-line',
  'ri:calendar-fill',
  'ri:calendar-event-line',
  'ri:calendar-event-fill',
  'ri:time-line',
  'ri:time-fill',
  'ri:timer-line',
  'ri:timer-fill',
  
  // 社交图标
  'ri:wechat-line',
  'ri:wechat-fill',
  'ri:qq-line',
  'ri:qq-fill',
  'ri:github-line',
  'ri:github-fill',
  'ri:twitter-line',
  'ri:twitter-fill',
  'ri:facebook-line',
  'ri:facebook-fill',
  
  // 其他
  'ri:star-line',
  'ri:star-fill',
  'ri:heart-line',
  'ri:heart-fill',
  'ri:thumb-up-line',
  'ri:thumb-up-fill',
  'ri:bookmark-line',
  'ri:bookmark-fill',
  'ri:flag-line',
  'ri:flag-fill',
  'ri:link-line',
  'ri:link-fill',
  'ri:external-link-line',
  'ri:external-link-fill',
  'ri:code-line',
  'ri:code-fill',
  'ri:terminal-line',
  'ri:terminal-fill',
  'ri:bug-line',
  'ri:bug-fill',
  'ri:copyright-line',
  'ri:copyright-fill',
  'ri:gift-line',
  'ri:gift-fill',
  'ri:trophy-line',
  'ri:trophy-fill',
  'ri:fire-line',
  'ri:fire-fill',
  'ri:flashlight-line',
  'ri:flashlight-fill',
  'ri:lightbulb-line',
  'ri:lightbulb-fill',
  
  // 地图导航
  'ri:map-pin-line',
  'ri:map-pin-fill',
  'ri:map-line',
  'ri:map-fill',
  'ri:navigation-line',
  'ri:navigation-fill',
  'ri:compass-line',
  'ri:compass-fill',
  'ri:global-line',
  'ri:global-fill',
  
  // 设备
  'ri:smartphone-line',
  'ri:smartphone-fill',
  'ri:computer-line',
  'ri:computer-fill',
  'ri:tablet-line',
  'ri:tablet-fill',
  'ri:tv-line',
  'ri:tv-fill',
  'ri:printer-line',
  'ri:printer-fill',
  
  // 网络
  'ri:wifi-line',
  'ri:wifi-fill',
  'ri:signal-wifi-line',
  'ri:signal-wifi-fill',
  'ri:cloud-line',
  'ri:cloud-fill',
  'ri:upload-cloud-line',
  'ri:upload-cloud-fill',
  'ri:download-cloud-line',
  'ri:download-cloud-fill',
  
  // 编辑器
  'ri:bold',
  'ri:italic',
  'ri:underline',
  'ri:font-size',
  'ri:font-color',
  'ri:align-left',
  'ri:align-center',
  'ri:align-right',
  'ri:list-unordered',
  'ri:list-ordered',
  'ri:indent-increase',
  'ri:indent-decrease',
  'ri:table-line',
  'ri:table-fill'
]

// 根据分类过滤图标
const filteredIcons = computed(() => {
  let icons = remixIcons
  
  // 根据搜索文本过滤
  if (searchText.value) {
    const search = searchText.value.toLowerCase()
    icons = icons.filter(icon => icon.toLowerCase().includes(search))
  }
  
  // 根据分类过滤
  if (currentCategory.value !== 'all') {
    // 这里可以根据实际情况添加分类过滤逻辑
    // 目前所有图标都是 remix icon
  }
  
  return icons
})

// 监听弹窗打开，同步当前值
watch(showDialog, (val) => {
  if (val) {
    tempSelectedIcon.value = props.modelValue
    customIconName.value = props.modelValue
    showCustomInput.value = false
  }
})

// 监听自定义输入，同步到临时选中值
watch(customIconName, (val) => {
  if (showCustomInput.value && val) {
    tempSelectedIcon.value = val
  }
})

// 选择图标
const handleSelectIcon = (icon: string) => {
  tempSelectedIcon.value = icon
  customIconName.value = icon
}

// 确认选择
const handleConfirm = () => {
  const iconValue = showCustomInput.value ? customIconName.value : tempSelectedIcon.value
  emit('update:modelValue', iconValue)
  emit('change', iconValue)
  showDialog.value = false
}

// 清空选择
const handleClear = () => {
  emit('update:modelValue', '')
  emit('change', '')
}
</script>

<style scoped lang="scss">
.art-icon-selector {
  width: 100%;
  
  :deep(.el-input) {
    cursor: pointer;
  }
  
  :deep(.el-input__inner) {
    cursor: pointer;
  }
}

.icon-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  color: var(--art-gray-600);
}

.search-icon-append {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 8px;
  color: var(--art-gray-400);
  cursor: pointer;
}

.icon-selector-dialog {
  :deep(.el-dialog__body) {
    padding: 0;
  }
}

.icon-selector-content {
  display: flex;
  flex-direction: column;
  height: 600px;
}

.selector-header {
  padding: 20px 20px 16px;
  border-bottom: 1px solid var(--art-gray-200);
  background: var(--art-gray-50);
  
  .input-group {
    display: flex;
    gap: 12px;
    margin-bottom: 16px;
    
    .search-input {
      flex: 1;
    }
  }
  
  .category-tabs {
    display: flex;
    justify-content: center;
    
    :deep(.el-radio-button__inner) {
      padding: 8px 16px;
    }
  }

  .custom-input-area {
    margin-top: 16px;
    
    .custom-hint {
      margin-top: 8px;
      font-size: 12px;
      color: var(--art-gray-500);
      text-align: center;
      
      .icon-link {
        color: var(--theme-color);
        text-decoration: none;
        
        &:hover {
          text-decoration: underline;
        }
      }
    }
  }
}

.icon-list-wrapper {
  flex: 1;
  padding: 20px;
  overflow: hidden;
}

.icon-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
  gap: 12px;
}

.icon-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 16px 8px;
  border: 1px solid var(--art-gray-200);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
  background: var(--default-box-color);
  
  &:hover {
    border-color: var(--theme-color);
    background: var(--theme-color-alpha-5);
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.08);
  }
  
  &.active {
    border-color: var(--theme-color);
    background: var(--theme-color-alpha-10);
    
    .icon-name {
      color: var(--theme-color);
      font-weight: 600;
    }
  }
  
  .icon-name {
    margin-top: 8px;
    font-size: 11px;
    color: var(--art-gray-600);
    text-align: center;
    line-height: 1.2;
    word-break: break-all;
    max-width: 100%;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    line-clamp: 2;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }
}

.custom-preview-area {
  padding: 40px 20px;
  text-align: center;
  
  .preview-title {
    font-size: 14px;
    font-weight: 600;
    color: var(--art-gray-700);
    margin-bottom: 24px;
  }
  
  .preview-box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px;
    background: var(--art-gray-50);
    border: 2px dashed var(--art-gray-300);
    border-radius: 12px;
    margin-bottom: 32px;
    min-height: 200px;
  }
  
  .preview-icon-large {
    color: var(--theme-color);
    margin-bottom: 16px;
  }
  
  .preview-placeholder {
    font-size: 14px;
    color: var(--art-gray-400);
  }
  
  .preview-icon-name {
    font-size: 13px;
    color: var(--art-gray-600);
    font-family: monospace;
    background: var(--art-gray-100);
    padding: 4px 12px;
    border-radius: 4px;
  }
  
  .preview-sizes {
    display: flex;
    justify-content: center;
    gap: 48px;
    
    .size-item {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 12px;
      
      .size-label {
        font-size: 12px;
        color: var(--art-gray-500);
      }
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  
  .selected-info {
    font-size: 14px;
    color: var(--art-gray-600);
    
    strong {
      color: var(--theme-color);
      margin-left: 4px;
    }
  }
}

// 暗黑模式适配
.dark {
  .icon-preview {
    color: var(--art-gray-400);
  }

  .search-icon-append {
    color: var(--art-gray-500);
  }
  
  .selector-header {
    background: var(--art-gray-900);
    border-bottom-color: var(--art-gray-700);
  }
  
  .icon-item {
    background: var(--art-gray-800);
    border-color: var(--art-gray-700);
    
    &:hover {
      background: var(--theme-color-alpha-10);
    }
    
    &.active {
      background: var(--theme-color-alpha-20);
    }
  }

  .custom-preview-area {
    .preview-box {
      background: var(--art-gray-800);
      border-color: var(--art-gray-700);
    }
    
    .preview-icon-name {
      background: var(--art-gray-700);
    }
  }
}
</style>
