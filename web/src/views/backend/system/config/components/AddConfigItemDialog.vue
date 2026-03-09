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
  <ElDialog
    :model-value="visible"
    title="添加配置项"
    width="560px"
    destroy-on-close
    @close="handleClose"
  >
    <ElForm
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="96px"
      class="add-config-form"
    >
      <ElFormItem label="所属分组" prop="group">
        <ElSelect v-model="form.group" placeholder="请选择分组" filterable :disabled="lockGroup">
          <ElOption
            v-for="item in groupOptions"
            :key="item.key"
            :label="item.title"
            :value="item.key"
          />
        </ElSelect>
      </ElFormItem>

      <ElFormItem label="配置名称" prop="name">
        <ElInput v-model="form.name" placeholder="如：站点名称" />
      </ElFormItem>

      <ElFormItem label="配置键名" prop="key">
        <ElInput v-model="form.key" placeholder="如：site_name" />
      </ElFormItem>

      <ElFormItem label="配置类型" prop="type">
        <ElSelect v-model="form.type" placeholder="请选择类型" filterable>
          <ElOption
            v-for="item in typeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </ElSelect>
      </ElFormItem>

      <ElFormItem
        label="初始值"
        prop="value"
        v-if="!['array', 'json', 'editor'].includes(form.type)"
      >
        <!-- 开关类型使用 ElSwitch -->
        <ElSwitch
          v-if="form.type === 'switch'"
          :model-value="form.value === '1'"
          @update:model-value="(val: string | number | boolean) => form.value = val ? '1' : '0'"
          active-text="开启"
          inactive-text="关闭"
        />
        <!-- 其他类型使用 textarea -->
        <ElInput
          v-else
          v-model="form.value"
          type="textarea"
          :autosize="{ minRows: 2, maxRows: 4 }"
          placeholder="填入默认值；多选/数组请用逗号或 JSON"
        />
      </ElFormItem>

      <!-- 类型属性设计器 -->
      <ElFormItem
        label="属性设置"
        v-if="showOptionDesigner"
      >
        <div class="option-designer">
          <!-- 文本类 -->
          <template v-if="['text','password','textarea','editor'].includes(form.type)">
            <ElFormItem label="占位符">
              <ElInput v-model="optionsForm.placeholder" placeholder="请输入占位提示" />
            </ElFormItem>
            <ElFormItem label="最小长度">
              <ElInputNumber v-model="optionsForm.minLength" :min="0" style="width:120px;" />
            </ElFormItem>
            <ElFormItem label="最大长度">
              <ElInputNumber v-model="optionsForm.maxLength" :min="0" style="width:120px;" />
            </ElFormItem>
            <ElFormItem v-if="form.type==='password'" label="正则校验">
              <ElInput v-model="optionsForm.pattern" placeholder="如：^(?=.*\\d)(?=.*[a-zA-Z]).{6,}$" />
            </ElFormItem>
            <ElFormItem v-if="form.type==='textarea' || form.type==='editor'" label="行数">
              <ElInputNumber v-model="optionsForm.rows" :min="2" :max="20" style="width:120px;" />
            </ElFormItem>
          </template>

          <!-- 数字 -->
          <template v-else-if="form.type==='number'">
            <ElFormItem label="最小值">
              <ElInputNumber v-model="optionsForm.min" :min="Number.MIN_SAFE_INTEGER" style="width:140px;" />
            </ElFormItem>
            <ElFormItem label="最大值">
              <ElInputNumber v-model="optionsForm.max" :max="Number.MAX_SAFE_INTEGER" style="width:140px;" />
            </ElFormItem>
            <ElFormItem label="步长">
              <ElInputNumber v-model="optionsForm.step" :min="0.0001" :step="0.1" style="width:140px;" />
            </ElFormItem>
          </template>

          <!-- 开关 -->
          <template v-else-if="form.type==='switch'">
            <ElFormItem label="开启值">
              <ElInput v-model="optionsForm.activeValue" placeholder="1" style="width:140px;" />
            </ElFormItem>
            <ElFormItem label="关闭值">
              <ElInput v-model="optionsForm.inactiveValue" placeholder="0" style="width:140px;" />
            </ElFormItem>
            <ElFormItem label="开启文本">
              <ElInput v-model="optionsForm.activeText" placeholder="如：开启" style="width:140px;" />
            </ElFormItem>
            <ElFormItem label="关闭文本">
              <ElInput v-model="optionsForm.inactiveText" placeholder="如：关闭" style="width:140px;" />
            </ElFormItem>
          </template>

          <!-- 选择类 -->
          <template v-else-if="['select','selects','radio','checkbox','remoteSelect','remoteSelects'].includes(form.type)">
            <ElFormItem label="占位符">
              <ElInput v-model="optionsForm.placeholder" placeholder="请选择..." />
            </ElFormItem>
            <div class="options-table">
              <div class="options-table-header">
                <span>选项列表</span>
                <ElButton size="small" type="primary" plain @click="addOption">添加选项</ElButton>
              </div>
              <div v-for="(opt, idx) in optionsForm.options" :key="idx" class="option-row">
                <ElInput v-model="opt.label" placeholder="标签" style="width:45%;" />
                <ElInput v-model="opt.value" placeholder="值" style="width:45%;" />
                <ElButton text type="danger" @click="removeOption(idx)">
                  <ArtSvgIcon icon="ri:delete-bin-line" :size="14" />
                </ElButton>
              </div>
            </div>
          </template>

          <!-- 上传类 -->
          <template v-else-if="['image','images','file','files'].includes(form.type)">
            <ElFormItem label="数量限制">
              <ElInputNumber v-model="optionsForm.limit" :min="1" :max="99" style="width:140px;" />
            </ElFormItem>
            <ElFormItem label="大小(MB)">
              <ElInputNumber v-model="optionsForm.maxSize" :min="1" :max="200" style="width:140px;" />
            </ElFormItem>
            <ElFormItem label="类型(accept)">
              <ElInput v-model="optionsForm.accept" placeholder="如 image/* 或 .png,.jpg" />
            </ElFormItem>
          </template>

          <!-- 城市 -->
          <template v-else-if="form.type==='city'">
            <ElFormItem label="占位符">
              <ElInput v-model="optionsForm.placeholder" placeholder="请选择城市" />
            </ElFormItem>
          </template>

          <!-- 数组 -->
          <template v-else-if="form.type==='array'">
            <div class="options-table-header">
              <span>数组字段（使用字段设计器）</span>
              <ElButton size="small" type="primary" plain @click="openFieldDesigner">字段设计器</ElButton>
            </div>
          </template>
        </div>
      </ElFormItem>

      <ElFormItem label="排序" prop="sort">
        <ElInputNumber v-model="form.sort" :min="0" :max="9999" style="width: 160px;" />
      </ElFormItem>

      <ElFormItem label="允许删除" prop="allowDel">
        <ElRadioGroup v-model="form.allowDel">
          <ElRadio :label="1">允许</ElRadio>
          <ElRadio :label="0">不允许</ElRadio>
        </ElRadioGroup>
      </ElFormItem>

      <ElFormItem label="规则 (可选)">
        <ElCheckboxGroup v-model="rulesSelected">
          <ElCheckbox
            v-for="item in currentRuleOptions"
            :key="item.label"
            :label="item.label"
          >
            {{ item.label }}
          </ElCheckbox>
        </ElCheckboxGroup>
        <div v-if="form.rules" class="rules-json-preview">
          {{ form.rules }}
        </div>
      </ElFormItem>

      <ElFormItem label="备注说明" prop="remark">
        <ElInput
          v-model="form.remark"
          type="textarea"
          :autosize="{ minRows: 2, maxRows: 4 }"
          placeholder="配置项说明"
        />
      </ElFormItem>
    </ElForm>

    <template #footer>
      <ElButton @click="handleClose">取消</ElButton>
      <ElButton type="primary" @click="handleSubmit" :loading="submitting">确定</ElButton>
    </template>
  </ElDialog>

  <!-- 字段设计器（数组类型） -->
  <ElDialog v-model="fieldDesignerVisible" title="数组字段设计器" width="760px">
    <ArtFieldEditor
      v-model="fieldDesignerValue"
      :compact="true"
    />
    <template #footer>
      <ElButton @click="fieldDesignerVisible = false">取消</ElButton>
      <ElButton type="primary" @click="confirmFieldDesigner">确定</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch } from 'vue'
import { ElMessage } from 'element-plus'
import ArtFieldEditor from '@/components/core/forms/art-field-editor/index.vue'

interface GroupOption {
  key: string
  title: string
}

interface FormState {
  group: string
  name: string
  key: string
  type: string
  value: string
  options: string
  rules: string
  sort: number
  allowDel: number
  remark: string
}

const props = defineProps<{
  visible: boolean
  groupOptions: GroupOption[]
  defaultGroup?: string
  lockGroup?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:visible', v: boolean): void
  (e: 'confirm', payload: FormState & { optionsParsed?: any }): void
}>()

const typeOptions = [
  { label: '字符串', value: 'text' },
  { label: '密码', value: 'password' },
  { label: '多行文本', value: 'textarea' },
  { label: '数字', value: 'number' },
  { label: '单选下拉', value: 'select' },
  { label: '多选下拉', value: 'selects' },
  { label: '远程单选', value: 'remoteSelect' },
  { label: '远程多选', value: 'remoteSelects' },
  { label: '单选框', value: 'radio' },
  { label: '复选框', value: 'checkbox' },
  { label: '开关', value: 'switch' },
  { label: '颜色', value: 'color' },
  { label: '图标', value: 'icon' },
  { label: '单图上传', value: 'image' },
  { label: '多图上传', value: 'images' },
  { label: '单文件上传', value: 'file' },
  { label: '多文件上传', value: 'files' },
  { label: '城市选择', value: 'city' },
  { label: '富文本', value: 'editor' },
  { label: '数组', value: 'array' },
  { label: 'JSON/对象', value: 'json' },
  { label: '时间', value: 'time' },
  { label: '日期', value: 'date' },
  { label: '日期时间', value: 'datetime' },
  { label: '年份', value: 'year' },
]

const formRef = ref()
const submitting = ref(false)
const rulesSelected = ref<string[]>([])
const currentRuleOptions = ref<Array<{ label: string; value: any }>>([])
const fieldDesignerVisible = ref(false)
const fieldDesignerValue = ref<any[]>([])
const optionsForm = reactive<any>({
  placeholder: '',
  min: undefined,
  max: undefined,
  step: undefined,
  minLength: undefined,
  maxLength: undefined,
  pattern: '',
  rows: undefined,
  limit: undefined,
  maxSize: undefined,
  accept: '',
  options: [],
  // 开关类型
  activeValue: '1',
  inactiveValue: '0',
  activeText: '',
  inactiveText: '',
})

// 是否展示属性设计器
const showOptionDesigner = computed(() => {
  return [
    'text','password','textarea','editor',
    'number',
    'switch',
    'select','selects','radio','checkbox','remoteSelect','remoteSelects',
    'image','images','file','files',
    'city',
    'array'
  ].includes(form.type)
})
const form = reactive<FormState>({
  group: props.defaultGroup || '',
  name: '',
  key: '',
  type: 'text',
  value: '',
  options: '',
  rules: '',
  sort: 100,
  allowDel: 1,
  remark: ''
})

const rules = {
  group: [{ required: true, message: '请选择分组', trigger: 'change' }],
  name: [{ required: true, message: '请输入配置名称', trigger: 'blur' }],
  key: [{ required: true, message: '请输入配置键名', trigger: 'blur' }],
  type: [{ required: true, message: '请选择配置类型', trigger: 'change' }],
}

const resetForm = () => {
  form.group = props.defaultGroup || (props.groupOptions[0]?.key || '')
  form.name = ''
  form.key = ''
  form.type = 'text'
  form.value = ''
  form.options = ''
  form.rules = ''
  form.sort = 100
  form.allowDel = 1
  form.remark = ''
  rulesSelected.value = []
  fieldDesignerValue.value = []
  resetOptionsForm()
  updateRuleOptions()
}

const handleClose = () => {
  emit('update:visible', false)
}

const addOption = () => {
  if (!Array.isArray(optionsForm.options)) optionsForm.options = []
  optionsForm.options.push({ label: '', value: '' })
}

const removeOption = (idx: number) => {
  if (!Array.isArray(optionsForm.options)) return
  optionsForm.options.splice(idx, 1)
}

const resetOptionsForm = () => {
  optionsForm.placeholder = ''
  optionsForm.min = undefined
  optionsForm.max = undefined
  optionsForm.step = undefined
  optionsForm.minLength = undefined
  optionsForm.maxLength = undefined
  optionsForm.pattern = ''
  optionsForm.rows = undefined
  optionsForm.limit = undefined
  optionsForm.maxSize = undefined
  optionsForm.accept = ''
  optionsForm.options = []
  optionsForm.activeValue = '1'
  optionsForm.inactiveValue = '0'
  optionsForm.activeText = ''
  optionsForm.inactiveText = ''
}

// 根据类型给 optionsForm 注入默认值（仅当当前为空时）
const setOptionDefaultsByType = () => {
  // 文本类
  if (['text','password','textarea','editor'].includes(form.type)) {
    if (!optionsForm.placeholder) optionsForm.placeholder = '请输入';
    if (optionsForm.minLength === undefined) optionsForm.minLength = form.type === 'password' ? 6 : 0;
    if (optionsForm.maxLength === undefined) optionsForm.maxLength = form.type === 'textarea' ? 300 : 50;
    if (form.type === 'password' && !optionsForm.pattern) {
      optionsForm.pattern = '^(?=.*\\\\d)(?=.*[a-zA-Z]).{6,}$';
    }
    if ((form.type === 'textarea' || form.type === 'editor') && optionsForm.rows === undefined) {
      optionsForm.rows = 4;
    }
    return;
  }

  // 数字
  if (form.type === 'number') {
    if (optionsForm.min === undefined) optionsForm.min = 0;
    if (optionsForm.max === undefined) optionsForm.max = 999999;
    if (optionsForm.step === undefined) optionsForm.step = 1;
    return;
  }

  // 选择类
  if (['select','selects','radio','checkbox','remoteSelect','remoteSelects'].includes(form.type)) {
    if (!optionsForm.placeholder) optionsForm.placeholder = '请选择';
    if (!optionsForm.options || optionsForm.options.length === 0) {
      optionsForm.options = [
        { label: '选项1', value: 'opt1' },
        { label: '选项2', value: 'opt2' }
      ];
    }
    return;
  }

  // 上传类
  if (['image','images','file','files'].includes(form.type)) {
    if (optionsForm.limit === undefined) optionsForm.limit = ['images','files'].includes(form.type) ? 9 : 1;
    if (optionsForm.maxSize === undefined) optionsForm.maxSize = 10;
    if (!optionsForm.accept) optionsForm.accept = '*';
    return;
  }

  // 开关
  if (form.type === 'switch') {
    if (!form.value) form.value = '1'; // 默认开启
    if (!optionsForm.activeValue) optionsForm.activeValue = '1';
    if (!optionsForm.inactiveValue) optionsForm.inactiveValue = '0';
    return;
  }

  // 城市
  if (form.type === 'city') {
    if (!optionsForm.placeholder) optionsForm.placeholder = '请选择城市';
    return;
  }

  // 数组无需默认
}

const openFieldDesigner = () => {
  // 如果已有 options，尝试解析
  if (form.options) {
    try {
      const parsed = JSON.parse(form.options)
      if (parsed?.fields && Array.isArray(parsed.fields)) {
        fieldDesignerValue.value = parsed.fields
      }
    } catch {
      // ignore parse error
    }
  }
  fieldDesignerVisible.value = true
}

const confirmFieldDesigner = () => {
  // 将字段设计结果写回 options
  form.options = JSON.stringify({ fields: fieldDesignerValue.value }, null, 2)
  fieldDesignerVisible.value = false
}

// 当弹窗打开或默认分组变化时同步分组
watch(
  () => [props.visible, props.defaultGroup],
  () => {
    if (props.visible) {
      form.group = props.defaultGroup || (props.groupOptions[0]?.key || '')
      updateRuleOptions()
    }
  },
  { immediate: true }
)

// 根据类型生成推荐规则
const rulePresets: Record<string, Array<{ label: string; value: any }>> = {
  text: [
    { label: '必填', value: { required: true } },
    { label: '最大长度50', value: { max: 50 } }
  ],
  password: [
    { label: '必填', value: { required: true } },
    { label: '最小长度6', value: { min: 6 } },
    { label: '最大长度32', value: { max: 32 } }
  ],
  textarea: [
    { label: '必填', value: { required: true } },
    { label: '最大长度300', value: { max: 300 } }
  ],
  number: [
    { label: '必填', value: { required: true } }
  ],
  select: [
    { label: '必选', value: { required: true } }
  ],
  selects: [
    { label: '必选', value: { required: true } }
  ],
  radio: [
    { label: '必选', value: { required: true } }
  ],
  checkbox: [
    { label: '必选', value: { required: true } }
  ],
  switch: [
    { label: '必选', value: { required: true } }
  ],
  color: [
    { label: '必选', value: { required: true } }
  ],
  icon: [
    { label: '必选', value: { required: true } }
  ],
  image: [
    { label: '必选', value: { required: true } }
  ],
  file: [
    { label: '必选', value: { required: true } }
  ],
  json: [
    { label: '必填', value: { required: true } }
  ],
  array: [
    { label: '必填', value: { required: true } }
  ]
}

const updateRuleOptions = () => {
  const opts = rulePresets[form.type] || [{ label: '必填', value: { required: true } }]
  currentRuleOptions.value = opts
  // 默认不选中任何规则
  rulesSelected.value = []
  // 同步 options 表单默认值
  resetOptionsForm()
  setOptionDefaultsByType()
}

// 监听类型变化，联动规则选项
watch(
  () => form.type,
  () => {
    updateRuleOptions()
    // 类型变化时自动填充 options 模板
    autoFillOptions()
    // 非数组时清空字段设计器缓存
    if (form.type !== 'array') {
      fieldDesignerValue.value = []
    }
  }
)

// 根据类型自动填充 options 示例
const autoFillOptions = () => {
  // 先确保 optionsForm 拥有默认值
  setOptionDefaultsByType()
  // 根据当前 optionsForm 生成 options JSON
  // 选择类
  if (['select','selects','radio','checkbox','remoteSelect','remoteSelects'].includes(form.type)) {
    form.options = JSON.stringify({
      placeholder: optionsForm.placeholder,
      options: optionsForm.options && optionsForm.options.length ? optionsForm.options : [
        { label: '选项1', value: 'opt1' },
        { label: '选项2', value: 'opt2' }
      ]
    }, null, 2)
    return
  }
  // 上传类
  if (['image','images','file','files'].includes(form.type)) {
    form.options = JSON.stringify({
      limit: optionsForm.limit ?? (['images','files'].includes(form.type) ? 9 : 1),
      maxSize: optionsForm.maxSize ?? 10,
      accept: optionsForm.accept || '*'
    }, null, 2)
    return
  }
  // 数字
  if (form.type === 'number') {
    form.options = JSON.stringify({
      min: optionsForm.min,
      max: optionsForm.max,
      step: optionsForm.step
    }, null, 2)
    return
  }
  // 文本类
  if (['text','password','textarea','editor'].includes(form.type)) {
    form.options = JSON.stringify({
      placeholder: optionsForm.placeholder,
      minLength: optionsForm.minLength,
      maxLength: optionsForm.maxLength,
      pattern: optionsForm.pattern,
      rows: optionsForm.rows
    }, null, 2)
    return
  }
  // 开关
  if (form.type === 'switch') {
    form.options = JSON.stringify({
      activeValue: optionsForm.activeValue || '1',
      inactiveValue: optionsForm.inactiveValue || '0',
      activeText: optionsForm.activeText || '',
      inactiveText: optionsForm.inactiveText || '',
    }, null, 2)
    return
  }

  // 城市
  if (form.type === 'city') {
    form.options = JSON.stringify({
      placeholder: optionsForm.placeholder
    }, null, 2)
    return
  }
  // 数组
  if (form.type === 'array') {
    if (fieldDesignerValue.value?.length) {
      form.options = JSON.stringify({ fields: fieldDesignerValue.value }, null, 2)
    } else {
      form.options = ''
    }
    return
  }
  form.options = ''
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate()
  submitting.value = true
  try {
    // 根据选项表单生成 options
    autoFillOptions()
    let optionsParsed: any = undefined
    if (form.options) {
      try {
        optionsParsed = JSON.parse(form.options)
      } catch (e: any) {
        ElMessage.error('配置项 JSON 解析失败，请检查格式')
        return
      }
    }
    // 规则序列化
    if (rulesSelected.value.length > 0) {
      const selectedRules = currentRuleOptions.value
        .filter(item => rulesSelected.value.includes(item.label))
        .map(item => item.value)
      form.rules = JSON.stringify(selectedRules)
    } else {
      form.rules = ''
    }
    emit('confirm', {
      ...form,
      optionsParsed,
      rules: form.rules
    })
    emit('update:visible', false)
    resetForm()
  } catch (error) {
    console.error(error)
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped lang="scss">
.add-config-form {
  padding: 8px 8px 0;
}
</style>
