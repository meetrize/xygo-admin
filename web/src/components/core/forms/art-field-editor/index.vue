<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 字段配置编辑器 - 用于配置 ArtArrayEditor 的 fields -->
<template>
  <div class="art-field-editor">
    <div class="editor-header">
      <div class="title">数据字段配置</div>
      <ElSwitch v-model="isRequired" inactive-text="是否必填" />
    </div>

    <div class="fields-list">
      <div
        v-for="(field, index) in fieldsList"
        :key="field._id"
        class="field-item"
      >
        <div class="field-header">
          <div class="field-index">{{ index + 1 }}</div>
          <ElButton
            text
            type="danger"
            @click="handleRemoveField(index)"
            :icon="Delete"
          />
        </div>

        <div class="field-body">
          <!-- 字段名 -->
          <ElFormItem label="字段名(key)">
            <ElInput
              v-model="field.key"
              placeholder="例如: name"
              @change="emitChange"
            />
          </ElFormItem>

          <!-- 标签 -->
          <ElFormItem label="标签(label)">
            <ElInput
              v-model="field.label"
              placeholder="例如: 姓名"
              @change="emitChange"
            />
          </ElFormItem>

          <!-- 类型 -->
          <ElFormItem label="类型(type)">
            <ElSelect
              v-model="field.type"
              placeholder="请选择类型"
              @change="handleTypeChange(field)"
              style="width: 100%"
            >
              <ElOption label="文本" value="text" />
              <ElOption label="密码" value="password" />
              <ElOption label="数字" value="number" />
              <ElOption label="单选下拉" value="select" />
              <ElOption label="多选下拉" value="selects" />
              <ElOption label="开关" value="switch" />
              <ElOption label="复选框" value="checkbox" />
              <ElOption label="单选框组" value="radio" />
              <ElOption label="多行文本" value="textarea" />
              <ElOption label="日期" value="date" />
              <ElOption label="日期时间" value="datetime" />
              <ElOption label="时间" value="time" />
              <ElOption label="年份" value="year" />
              <ElOption label="颜色" value="color" />
              <ElOption label="图标" value="icon" />
              <ElOption label="单图上传" value="image" />
              <ElOption label="多图上传" value="images" />
              <ElOption label="单文件上传" value="file" />
              <ElOption label="多文件上传" value="files" />
            </ElSelect>
          </ElFormItem>

          <!-- 选项配置（仅 select/selects/radio 类型） -->
          <div
            v-if="['select', 'selects', 'radio'].includes(field.type)"
            class="options-config"
          >
            <div class="options-header">
              <div class="label">选项配置</div>
              <ElButton
                size="small"
                type="primary"
                text
                @click="handleAddOption(field)"
              >
                添加选项
              </ElButton>
            </div>

            <div class="options-list">
              <div
                v-for="(option, optIndex) in field.options"
                :key="optIndex"
                class="option-item"
              >
                <ElInput
                  v-model="option.label"
                  placeholder="标签"
                  @change="emitChange"
                  style="flex: 1"
                />
                <ElInput
                  v-model="option.value"
                  placeholder="值"
                  @change="emitChange"
                  style="flex: 1"
                />
                <ElButton
                  text
                  type="danger"
                  @click="handleRemoveOption(field, optIndex)"
                >
                  删除
                </ElButton>
              </div>
            </div>
          </div>

          <!-- 占位符 -->
          <ElFormItem label="占位符(placeholder)">
            <ElInput
              v-model="field.placeholder"
              placeholder="请输入占位符"
              @change="emitChange"
            />
          </ElFormItem>

          <!-- 数字类型的限制 -->
          <div v-if="field.type === 'number'" class="number-config">
            <ElFormItem label="最小值">
              <ElInputNumber
                v-model="field.min"
                @change="emitChange"
                style="width: 100%"
              />
            </ElFormItem>
            <ElFormItem label="最大值">
              <ElInputNumber
                v-model="field.max"
                @change="emitChange"
                style="width: 100%"
              />
            </ElFormItem>
          </div>

          <!-- 文件上传类型的配置 -->
          <div v-if="['image', 'images', 'file', 'files'].includes(field.type)">
            <ElFormItem label="文件大小限制(MB)">
              <ElInputNumber
                v-model="field.maxSize"
                :min="1"
                :max="100"
                @change="emitChange"
                style="width: 100%"
              />
            </ElFormItem>
            <ElFormItem
              v-if="['images', 'files'].includes(field.type)"
              label="数量限制"
            >
              <ElInputNumber
                v-model="field.limit"
                :min="1"
                :max="50"
                @change="emitChange"
                style="width: 100%"
              />
            </ElFormItem>
          </div>

          <!-- 多行文本的行数 -->
          <ElFormItem v-if="field.type === 'textarea'" label="行数">
            <ElInputNumber
              v-model="field.rows"
              :min="2"
              :max="20"
              @change="emitChange"
              style="width: 100%"
            />
          </ElFormItem>
        </div>
      </div>
    </div>

    <div class="add-field-btn">
      <ElButton type="primary" @click="handleAddField" style="width: 100%">
        添加字段
      </ElButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import { Delete } from '@element-plus/icons-vue'
import type { FieldConfig } from '../art-array-editor/index.vue'

defineOptions({ name: 'ArtFieldEditor' })

interface ExtendedFieldConfig extends FieldConfig {
  _id?: number
}

interface Props {
  modelValue?: FieldConfig[]
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: () => []
})

const emit = defineEmits<{
  'update:modelValue': [value: FieldConfig[]]
  'change': [value: FieldConfig[]]
}>()

let fieldIdCounter = 0
let isUpdating = false

const isRequired = ref(false)
const fieldsList = ref<ExtendedFieldConfig[]>([])

// 初始化字段列表
const initFields = () => {
  isUpdating = true
  const value = props.modelValue || []
  fieldsList.value = value.map(field => ({
    ...field,
    _id: fieldIdCounter++
  }))
  nextTick(() => {
    isUpdating = false
  })
}

// 监听外部值变化
watch(() => props.modelValue, () => {
  // 如果正在更新中，跳过
  if (isUpdating) return
  initFields()
}, { immediate: true, deep: true })

// 添加字段
const handleAddField = () => {
  fieldsList.value.push({
    _id: fieldIdCounter++,
    key: '',
    label: '',
    type: 'text',
    placeholder: ''
  })
  emitChange()
}

// 删除字段
const handleRemoveField = (index: number) => {
  fieldsList.value.splice(index, 1)
  emitChange()
}

// 类型改变
const handleTypeChange = (field: ExtendedFieldConfig) => {
  // 如果是选择类型，初始化 options
  if (['select', 'selects', 'radio'].includes(field.type)) {
    if (!field.options) {
      field.options = []
    }
  }
  // 如果是数字类型，初始化 min/max
  if (field.type === 'number') {
    if (field.min === undefined) field.min = 0
    if (field.max === undefined) field.max = 100
  }
  // 如果是文件上传类型，初始化 maxSize
  if (['image', 'images', 'file', 'files'].includes(field.type)) {
    if (!field.maxSize) field.maxSize = field.type.includes('image') ? 5 : 10
    if (['images', 'files'].includes(field.type) && !field.limit) {
      field.limit = 9
    }
  }
  // 如果是多行文本，初始化 rows
  if (field.type === 'textarea' && !field.rows) {
    field.rows = 3
  }
  emitChange()
}

// 添加选项
const handleAddOption = (field: ExtendedFieldConfig) => {
  if (!field.options) {
    field.options = []
  }
  field.options.push({
    label: '',
    value: ''
  })
  emitChange()
}

// 删除选项
const handleRemoveOption = (field: ExtendedFieldConfig, index: number) => {
  if (field.options) {
    field.options.splice(index, 1)
    emitChange()
  }
}

// 发送更新
const emitChange = () => {
  if (isUpdating) return
  
  const cleanFields = fieldsList.value.map(field => {
    const { _id, ...rest } = field
    return rest
  })
  emit('update:modelValue', cleanFields)
  emit('change', cleanFields)
}
</script>

<style scoped lang="scss">
.art-field-editor {
  width: 100%;
  border: 1px solid var(--art-gray-300);
  border-radius: 8px;
  background: var(--default-box-color);
  padding: 16px;
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--art-gray-200);

  .title {
    font-size: 16px;
    font-weight: 600;
    color: var(--art-gray-800);
  }
}

.fields-list {
  max-height: 600px;
  overflow-y: auto;
  margin-bottom: 16px;
}

.field-item {
  background: var(--art-gray-50);
  border: 1px solid var(--art-gray-200);
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 12px;

  &:last-child {
    margin-bottom: 0;
  }
}

.field-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.field-index {
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--theme-color);
  color: white;
  border-radius: 50%;
  font-size: 14px;
  font-weight: 600;
}

.field-body {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;

  :deep(.el-form-item) {
    margin-bottom: 0;

    .el-form-item__label {
      font-size: 12px;
      color: var(--art-gray-600);
      font-weight: 500;
    }
  }
}

.options-config {
  grid-column: 1 / -1;
  border: 1px dashed var(--art-gray-300);
  border-radius: 6px;
  padding: 12px;
  background: var(--default-box-color);
}

.options-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;

  .label {
    font-size: 13px;
    font-weight: 600;
    color: var(--art-gray-700);
  }
}

.options-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.option-item {
  display: flex;
  gap: 8px;
  align-items: center;
}

.number-config {
  display: contents;
}

.add-field-btn {
  margin-top: 16px;
}

// 暗黑模式
.dark {
  .field-item {
    background: var(--art-gray-800);
    border-color: var(--art-gray-700);
  }

  .options-config {
    background: var(--art-gray-750);
  }
}
</style>
