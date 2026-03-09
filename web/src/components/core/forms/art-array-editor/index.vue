<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 多维数组编辑器组件 -->
<template>
  <div class="art-array-editor" :class="{ compact }">
    <!-- 头部工具栏 -->
    <div class="editor-header">
      <div class="header-title">
        <span class="title-text">数据列表</span>
        <span class="title-count">{{ dataList.length }} 项</span>
      </div>
      <div class="header-actions">
        <ElButton size="small" @click="handleAddRow" type="primary" plain>
          <ArtSvgIcon icon="ri:add-line" :size="14" style="margin-right: 4px;" />
          添加数据
        </ElButton>
        <ElButton size="small" @click="previewVisible = true" v-if="dataList.length > 0">
          <ArtSvgIcon icon="ri:eye-line" :size="14" style="margin-right: 4px;" />
          预览 JSON
        </ElButton>
        <ElButton size="small" @click="handleClearAll" v-if="dataList.length > 0">
          <ArtSvgIcon icon="ri:delete-bin-line" :size="14" style="margin-right: 4px;" />
          清空全部
        </ElButton>
      </div>
    </div>

    <!-- 数据列表 -->
    <div class="editor-content" v-if="dataList.length > 0">
      <VueDraggable 
        v-model="dataList" 
        :animation="200"
        handle=".drag-handle"
        @end="handleDragEnd"
      >
        <div
          v-for="(row, rowIndex) in dataList"
          :key="row.__id"
          class="data-row"
        >
          <!-- 拖拽手柄 -->
          <div class="drag-handle">
            <ArtSvgIcon icon="ri:draggable" :size="16" />
          </div>

          <!-- 序号 -->
          <div v-if="showIndex" class="row-index">
            {{ rowIndex + 1 }}
          </div>

          <!-- 字段列表 -->
          <div class="row-fields">
            <div
              v-for="field in fields"
              :key="field.key"
              class="field-item"
            >
              <div class="field-label">{{ field.label }}</div>
              <div class="field-control">
                <!-- 文本输入 -->
                <ElInput
                  v-if="field.type === 'text' || field.type === 'string'"
                  v-model="row[field.key]"
                  :placeholder="field.placeholder || `请输入${field.label}`"
                  clearable
                  style="height: 32px;"
                />

                <!-- 数字输入 -->
                <ElInputNumber
                  v-else-if="field.type === 'number'"
                  v-model="row[field.key]"
                  :min="field.min"
                  :max="field.max"
                  style="width: 100%; height: 32px;"
                />

                <!-- 下拉单选 -->
                <ElSelect
                  v-else-if="field.type === 'select'"
                  v-model="row[field.key]"
                  :placeholder="field.placeholder || `请选择${field.label}`"
                  clearable
                  style="width: 100%; height: 32px;"
                >
                  <ElOption
                    v-for="opt in field.options || []"
                    :key="opt.value"
                    :label="opt.label"
                    :value="opt.value"
                  />
                </ElSelect>

                <!-- 下拉多选 -->
                <ElSelect
                  v-else-if="field.type === 'selects'"
                  v-model="row[field.key]"
                  :placeholder="field.placeholder || `请选择${field.label}`"
                  multiple
                  clearable
                  style="width: 100%; height: 32px;"
                >
                  <ElOption
                    v-for="opt in field.options || []"
                    :key="opt.value"
                    :label="opt.label"
                    :value="opt.value"
                  />
                </ElSelect>

                <!-- 开关 -->
                <ElSwitch
                  v-else-if="field.type === 'switch'"
                  v-model="row[field.key]"
                />

                <!-- 日期 -->
                <ElDatePicker
                  v-else-if="field.type === 'date'"
                  v-model="row[field.key]"
                  type="date"
                  value-format="YYYY-MM-DD"
                  style="width: 100%; height: 32px;"
                />

                <!-- 颜色 -->
                <ElColorPicker
                  v-else-if="field.type === 'color'"
                  v-model="row[field.key]"
                />

                <!-- 图标 -->
                <ArtIconSelector
                  v-else-if="field.type === 'icon'"
                  v-model="row[field.key]"
                />

                <!-- 单图上传 -->
                <ArtImageUpload
                  v-else-if="field.type === 'image'"
                  v-model="row[field.key]"
                  :max-size="field.maxSize || 5"
                  :compact="compact"
                />

                <!-- 多图上传 -->
                <ArtImageUpload
                  v-else-if="field.type === 'images'"
                  v-model="row[field.key]"
                  multiple
                  :limit="field.limit || 9"
                  :max-size="field.maxSize || 5"
                  :compact="compact"
                />

                <!-- 单文件上传 -->
                <ArtFileUpload
                  v-else-if="field.type === 'file'"
                  v-model="row[field.key]"
                  :accept="field.accept"
                  :max-size="field.maxSize || 10"
                  :compact="compact"
                />

                <!-- 多文件上传 -->
                <ArtFileUpload
                  v-else-if="field.type === 'files'"
                  v-model="row[field.key]"
                  multiple
                  :limit="field.limit || 10"
                  :accept="field.accept"
                  :max-size="field.maxSize || 10"
                  :compact="compact"
                />

                <!-- 密码输入 -->
                <ElInput
                  v-else-if="field.type === 'password'"
                  v-model="row[field.key]"
                  type="password"
                  show-password
                  :placeholder="field.placeholder || `请输入${field.label}`"
                  clearable
                  style="height: 32px;"
                />

                <!-- 多行文本 -->
                <ElInput
                  v-else-if="field.type === 'textarea'"
                  v-model="row[field.key]"
                  type="textarea"
                  :rows="field.rows || 3"
                  :placeholder="field.placeholder || `请输入${field.label}`"
                />

                <!-- 日期时间 -->
                <ElDatePicker
                  v-else-if="field.type === 'datetime'"
                  v-model="row[field.key]"
                  type="datetime"
                  value-format="YYYY-MM-DD HH:mm:ss"
                  style="width: 100%; height: 32px;"
                />

                <!-- 时间选择 -->
                <ElTimePicker
                  v-else-if="field.type === 'time'"
                  v-model="row[field.key]"
                  value-format="HH:mm:ss"
                  style="width: 100%; height: 32px;"
                />

                <!-- 年份选择 -->
                <ElDatePicker
                  v-else-if="field.type === 'year'"
                  v-model="row[field.key]"
                  type="year"
                  value-format="YYYY"
                  style="width: 100%; height: 32px;"
                />

                <!-- 单选框组 -->
                <ElRadioGroup
                  v-else-if="field.type === 'radio'"
                  v-model="row[field.key]"
                >
                  <ElRadio
                    v-for="opt in field.options || []"
                    :key="opt.value"
                    :label="opt.value"
                  >
                    {{ opt.label }}
                  </ElRadio>
                </ElRadioGroup>

                <!-- 复选框 -->
                <ElCheckbox
                  v-else-if="field.type === 'checkbox'"
                  v-model="row[field.key]"
                >
                  {{ field.placeholder || field.label }}
                </ElCheckbox>

                <!-- 颜色选择器（使用自定义组件） -->
                <ArtColorPicker
                  v-else-if="field.type === 'color'"
                  v-model="row[field.key]"
                />

                <!-- 默认文本 -->
                <ElInput
                  v-else
                  v-model="row[field.key]"
                  :placeholder="field.placeholder || `请输入${field.label}`"
                  clearable
                  style="height: 32px;"
                />
              </div>
            </div>
          </div>

          <!-- 操作按钮 -->
          <div class="row-actions">
            <ElButton
              size="small"
              text
              type="danger"
              @click="handleDeleteRow(rowIndex)"
            >
              <ArtSvgIcon icon="ri:delete-bin-line" :size="16" />
            </ElButton>
          </div>
        </div>
      </VueDraggable>
    </div>

    <!-- 空状态 -->
    <div v-else class="empty-state">
      <ArtSvgIcon icon="ri:database-line" :size="48" class="empty-icon" />
      <div class="empty-text">暂无数据</div>
      <ElButton size="small" type="primary" @click="handleAddRow">
        <ArtSvgIcon icon="ri:add-line" :size="14" style="margin-right: 4px;" />
        添加第一条数据
      </ElButton>
    </div>

    <!-- JSON 预览弹窗 -->
    <ElDialog v-model="previewVisible" title="JSON 预览" width="640px">
      <ElInput
        type="textarea"
        :value="previewJson"
        :rows="14"
        readonly
      />
      <template #footer>
        <ElButton @click="previewVisible = false">关闭</ElButton>
      </template>
    </ElDialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, nextTick } from 'vue'
import { VueDraggable } from 'vue-draggable-plus'

defineOptions({ name: 'ArtArrayEditor' })

export interface FieldConfig {
  key: string
  label: string
  type: 'text' | 'string' | 'password' | 'textarea' | 'number' | 'select' | 'selects' | 'switch' | 
        'date' | 'datetime' | 'time' | 'year' | 'color' | 'icon' | 
        'image' | 'images' | 'file' | 'files' | 'radio' | 'checkbox'
  placeholder?: string
  options?: Array<{ label: string; value: any }>
  min?: number
  max?: number
  maxSize?: number  // 文件大小限制（MB）
  limit?: number    // 上传数量限制
  accept?: string   // 文件类型限制
  rows?: number     // textarea 行数
  default?: any
}

interface Props {
  modelValue?: any[]
  fields?: FieldConfig[]
  showIndex?: boolean
  sortable?: boolean
  groupable?: boolean
  compact?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  showIndex: true,
  sortable: true,
  groupable: false,
  compact: true
})

const emit = defineEmits<{
  'update:modelValue': [value: any[]]
  'change': [value: any[]]
}>()

let idCounter = 0
const previewVisible = ref(false)

// 数据列表（添加唯一ID用于 key）
const dataList = ref<Array<any>>([])

// 是否正在更新中（防止循环更新）
let isUpdating = false

// 初始化数据
const initData = () => {
  const value = props.modelValue || []
  isUpdating = true
  dataList.value = value.map((item: any) => ({
    ...item,
    __id: item.__id !== undefined ? item.__id : idCounter++
  }))
  // 使用 nextTick 确保更新完成后再允许 emit
  nextTick(() => {
    isUpdating = false
  })
}

// 监听外部值变化
watch(() => props.modelValue, (newVal) => {
  // 如果正在更新中，跳过（避免循环）
  if (isUpdating) return
  
  // 检查数据是否真的不同（简单比较长度和第一层数据）
  const currentCleanData = dataList.value.map(item => {
    const { __id, ...rest } = item
    return rest
  })
  
  const isSame = newVal && currentCleanData.length === newVal.length &&
    newVal.every((item: any, index: number) => {
      const current = currentCleanData[index]
      if (!current) return false
      // 简单对比所有 key
      const keys = Object.keys(item)
      return keys.every(key => item[key] === current[key])
    })
  
  if (!isSame) {
    initData()
  }
}, { immediate: true })

// 监听数据变化并emit
watch(dataList, (newVal) => {
  // 如果正在更新中，不触发 emit，避免循环
  if (isUpdating) return
  
  const cleanData = newVal.map((item: any) => {
    const { __id, ...rest } = item
    return rest
  })
  emit('update:modelValue', cleanData)
  emit('change', cleanData)
}, { deep: true })

// 添加行
const handleAddRow = () => {
  const newRow: any = { __id: idCounter++ }
  
  // 根据字段配置初始化默认值
  props.fields?.forEach((field: FieldConfig) => {
    if (field.default !== undefined) {
      newRow[field.key] = field.default
    } else {
      switch (field.type) {
        case 'number':
          newRow[field.key] = null
          break
        case 'switch':
        case 'checkbox':
          newRow[field.key] = false
          break
        case 'selects':
        case 'images':
        case 'files':
          newRow[field.key] = []
          break
        case 'text':
        case 'string':
        case 'password':
        case 'textarea':
        case 'select':
        case 'date':
        case 'datetime':
        case 'time':
        case 'year':
        case 'color':
        case 'icon':
        case 'image':
        case 'file':
        case 'radio':
          newRow[field.key] = ''
          break
        default:
          newRow[field.key] = ''
      }
    }
  })
  
  dataList.value.push(newRow)
}

// 删除行
const handleDeleteRow = (index: number) => {
  ElMessageBox.confirm(
    '确定要删除这条数据吗？',
    '删除确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    dataList.value.splice(index, 1)
    ElMessage.success('删除成功')
  }).catch(() => {})
}

// 清空全部
const handleClearAll = () => {
  ElMessageBox.confirm(
    `确定要清空全部 ${dataList.value.length} 条数据吗？`,
    '清空确认',
    {
      confirmButtonText: '确定清空',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    dataList.value = []
    ElMessage.success('已清空')
  }).catch(() => {})
}

// 拖拽结束
const handleDragEnd = () => {
  // dataList 的变化会自动触发 watch
}

// 字段配置
const fields = computed(() => props.fields || [])

// 预览 JSON
const previewJson = computed(() => {
  const cleanData = dataList.value.map((item: any) => {
    const { __id, ...rest } = item
    return rest
  })
  return JSON.stringify(cleanData, null, 2)
})
</script>

<style scoped lang="scss">
.art-array-editor {
  width: 100%;
  border: 1px solid var(--art-gray-300);
  border-radius: 8px;
  overflow: hidden;
  background: var(--default-box-color);
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: var(--art-gray-50);
  border-bottom: 1px solid var(--art-gray-300);
}

.header-title {
  display: flex;
  align-items: center;
  gap: 8px;

  .title-text {
    font-size: 14px;
    font-weight: 600;
    color: var(--art-gray-800);
  }

  .title-count {
    font-size: 12px;
    color: var(--art-gray-500);
    background: var(--art-gray-200);
    padding: 2px 8px;
    border-radius: 12px;
  }
}

.header-actions {
  display: flex;
  gap: 8px;
}

.editor-content {
  padding: 12px;
  max-height: 600px;
  overflow-y: auto;
}

.data-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px;
  margin-bottom: 8px;
  background: var(--default-box-color);
  border: 1px solid var(--art-gray-200);
  border-radius: 6px;
  transition: all 0.2s;

  &:hover {
    border-color: var(--theme-color);
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);

    .drag-handle {
      opacity: 1;
    }
  }

  &:last-child {
    margin-bottom: 0;
  }
}

.drag-handle {
  flex-shrink: 0;
  width: 20px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--art-gray-400);
  cursor: move;
  opacity: 0;
  transition: all 0.2s;

  &:hover {
    color: var(--theme-color);
  }
}

.row-index {
  flex-shrink: 0;
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--theme-color-alpha-10);
  color: var(--theme-color);
  border-radius: 4px;
  font-size: 12px;
  font-weight: 600;
}

.row-fields {
  flex: 1;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 8px;
  min-width: 0;
}

.field-item {
  min-width: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.field-label {
  font-size: 11px;
  color: var(--art-gray-600);
  margin-bottom: 4px;
  font-weight: 500;
  line-height: 1.2;
}

.field-control {
  width: 100%;
  display: flex;
  align-items: center;

  // 所有内部组件统一使用 32px 高度
  :deep(.el-input),
  :deep(.el-select),
  :deep(.el-input-number),
  :deep(.el-date-picker),
  :deep(.el-time-picker) {
    --el-component-size: 32px;
    height: 32px;
  }

  :deep(.el-input__inner) {
    height: 32px;
    line-height: 32px;
    font-size: 13px;
  }

  :deep(.el-input__wrapper) {
    min-height: 32px;
    height: 32px;
    padding: 1px 8px;
  }

  // ElSelect 特殊处理
  :deep(.el-select) {
    height: 32px !important;
    
    .el-input {
      height: 32px !important;
    }

    .el-input__wrapper {
      height: 32px !important;
      min-height: 32px !important;
      box-shadow: 0 0 0 1px var(--el-input-border-color,var(--el-border-color)) inset !important;
    }

    .el-input__inner {
      height: 32px !important;
      line-height: 32px !important;
    }

    .el-select__wrapper {
      height: 32px !important;
      min-height: 32px !important;
    }
  }

  // 颜色选择器精简样式
  :deep(.art-color-picker) {
    .el-input-group__append {
      padding: 0 2px;
    }

    .color-picker-wrapper {
      padding: 0 2px;
    }

    .el-color-picker__trigger {
      width: 24px !important;
      height: 24px !important;
      padding: 2px !important;
    }
  }

  // 图标选择器精简样式
  :deep(.art-icon-selector) {
    .el-input__inner {
      font-size: 12px;
    }

    .icon-preview {
      width: 24px;
      height: 24px;
    }

    .search-icon-append {
      padding: 0 4px;
    }
  }

  // 开关和复选框对齐
  :deep(.el-switch),
  :deep(.el-checkbox) {
    display: flex;
    align-items: center;
    height: 32px;
  }

  // 单选按钮组对齐
  :deep(.el-radio-group) {
    display: flex;
    align-items: center;
    height: 32px;
  }

  // 颜色选择器对齐
  :deep(.el-color-picker) {
    display: flex;
    align-items: center;
  }
}

// 紧凑模式下，对通用组件做更强的收紧以避免破坏布局
.art-array-editor.compact {
  .editor-content {
    padding: 10px;
  }

  .data-row {
    padding: 10px;
    gap: 8px;
    align-items: center;
  }

  .row-fields {
    gap: 8px;
    align-items: center;
  }

  .field-label {
    margin-bottom: 3px;
  }
  
  .field-item {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .field-control {
    min-height: 32px;
    display: flex;
    align-items: center;

    // 统一所有组件高度为32px
    :deep(.el-input),
    :deep(.el-select),
    :deep(.el-input-number),
    :deep(.el-date-picker),
    :deep(.el-time-picker) {
      --el-component-size: 32px;
    }

    :deep(.el-input__wrapper) {
      padding: 1px 8px;
      min-height: 32px;
    }

    :deep(.el-input__inner) {
      height: 30px;
      line-height: 30px;
    }

    // 颜色选择器对齐
    :deep(.art-color-picker) {
      display: flex;
      align-items: center;
      width: 100%;
      
      .el-input__inner {
        font-size: 12px;
      }

      .el-color-picker__trigger {
        width: 22px !important;
        height: 22px !important;
      }
    }

    // 图标选择器对齐
    :deep(.art-icon-selector) {
      display: flex;
      align-items: center;
      width: 100%;
      
      .el-input__wrapper {
        padding: 1px 8px;
      }

      .icon-preview {
        width: 22px;
        height: 22px;
      }
    }

    // ElSelect 对齐
    :deep(.el-select) {
      height: 32px !important;
      
      .el-input {
        height: 32px !important;
      }

      .el-input__wrapper {
        height: 32px !important;
        min-height: 32px !important;
      }

      .el-input__inner {
        height: 32px !important;
        line-height: 32px !important;
      }
    }

    // 开关/复选框/单选按钮组统一高度
    :deep(.el-switch),
    :deep(.el-checkbox),
    :deep(.el-radio-group) {
      height: 32px !important;
      display: flex !important;
      align-items: center !important;
    }
  }
}

.row-actions {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  height: 32px;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  color: var(--art-gray-500);

  .empty-icon {
    margin-bottom: 16px;
    color: var(--art-gray-300);
  }

  .empty-text {
    font-size: 14px;
    margin-bottom: 20px;
    color: var(--art-gray-500);
  }
}

// 暗黑模式
.dark {
  .editor-header {
    background: var(--art-gray-900);
    border-bottom-color: var(--art-gray-700);
  }

  .header-title {
    .title-count {
      background: var(--art-gray-700);
    }
  }

  .data-row {
    background: var(--art-gray-800);
    border-color: var(--art-gray-700);

    &:hover {
      background: var(--art-gray-750);
    }
  }

  .empty-state {
    .empty-icon {
      color: var(--art-gray-700);
    }
  }
}
</style>
