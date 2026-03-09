// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

/**
 * 代码生成器 - 字段设计器库
 * 定义 designType 类型体系 + 预设字段库
 */

// ==================== 类型定义 ====================

/** 属性配置项（描述式配置） */
export interface PropItem {
  type: 'select' | 'switch' | 'string' | 'number' | 'selects' | 'textarea' | 'hidden'
    | 'remoteTableSelect' | 'remoteColumnSelect' | 'remoteColumnMultiSelect' | 'relationFieldsDesigner'
    | 'optionsEditor'
  value: any
  options?: Record<string, string>
  placeholder?: string
  label?: string
}

/** designType 定义 */
export interface DesignTypeDef {
  name: string                          // 显示名称
  defaultDbType?: string                // 默认数据库字段类型
  defaultGoType?: string                // 默认 Go 类型
  defaultTsType?: string                // 默认 TS 类型
  defaultLength?: number                // 默认长度
  table: Record<string, PropItem>       // 表格属性配置
  form: Record<string, PropItem>        // 表单属性配置
}

/** 预设字段 */
export interface PresetField {
  title: string           // 显示标题
  name: string            // 默认字段名
  designType: string      // 设计类型
  comment: string         // 默认注释
  dbType?: string         // 数据库类型（覆盖 designType 默认）
  goType?: string         // Go 类型
  tsType?: string         // TS 类型
  isList?: number         // 是否在列表显示
  isEdit?: number         // 是否在表单显示
  isPk?: number           // 是否主键
  defaultValue?: string   // 默认值
}

// ==================== 通用属性模板 ====================

/** 表格渲染方式选项 */
const renderOptions: Record<string, string> = {
  none: '无',
  switch: '开关',
  image: '图片',
  images: '多图',
  tag: 'Tag',
  tags: 'Tags',
  url: 'URL',
  datetime: '时间日期',
  color: '颜色',
  icon: '图标',
}

/** 搜索运算符选项 */
const operatorOptions: Record<string, string> = {
  false: '禁用搜索',
  eq: '精确匹配(=)',
  neq: '不等于(!=)',
  like: '模糊匹配(LIKE)',
  gt: '大于(>)',
  gte: '大于等于(>=)',
  lt: '小于(<)',
  lte: '小于等于(<=)',
  between: '区间(BETWEEN)',
  in: '包含(IN)',
}

/** 可排序选项 */
const sortableOptions: Record<string, string> = {
  false: '禁用',
  custom: '启用',
}

/** 验证器类型 */
export const validatorTypes: Record<string, string> = {
  required: '必填',
  mobile: '手机号',
  email: '邮箱',
  url: 'URL',
  number: '数字',
  integer: '整数',
  float: '浮点数',
  date: '日期',
  idNumber: '身份证号',
  account: '账户名',
  password: '密码',
}

/** 基础表格属性 */
function baseTableAttr(operator = 'eq', sortable = 'false', render = 'none'): Record<string, PropItem> {
  return {
    render: { type: 'select', value: render, options: renderOptions, label: '列渲染方式' },
    operator: { type: 'select', value: operator, options: operatorOptions, label: '搜索方式' },
    sortable: { type: 'select', value: sortable, options: sortableOptions, label: '可排序' },
  }
}

/** 基础表单属性 */
function baseFormAttr(): Record<string, PropItem> {
  return {
    validator: { type: 'selects', value: [], options: validatorTypes, label: '验证规则' },
    validatorMsg: { type: 'textarea', value: '', placeholder: '留空自动填入验证器标题', label: '验证消息' },
  }
}

// ==================== designType 定义（26种） ====================

export const designTypes: Record<string, DesignTypeDef> = {
  pk: {
    name: '主键',
    defaultDbType: 'bigint unsigned',
    defaultGoType: 'uint64',
    defaultTsType: 'number',
    table: {
      ...baseTableAttr('eq', 'custom'),
      width: { type: 'number', value: 70, label: '列宽' },
    },
    form: {},
  },
  string: {
    name: '字符串',
    defaultDbType: 'varchar(255)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('like'),
    form: baseFormAttr(),
  },
  number: {
    name: '数字',
    defaultDbType: 'int(10)',
    defaultGoType: 'int',
    defaultTsType: 'number',
    table: baseTableAttr('eq', 'custom'),
    form: {
      ...baseFormAttr(),
      step: { type: 'number', value: 1, label: '步长' },
    },
  },
  float: {
    name: '浮点数',
    defaultDbType: 'decimal(10,2)',
    defaultGoType: 'float64',
    defaultTsType: 'number',
    table: baseTableAttr('eq', 'custom'),
    form: {
      ...baseFormAttr(),
      step: { type: 'number', value: 0.01, label: '步长' },
    },
  },
  switch: {
    name: '开关',
    defaultDbType: 'tinyint(1)',
    defaultGoType: 'int',
    defaultTsType: 'number',
    table: {
      ...baseTableAttr('eq'),
      render: { type: 'select', value: 'switch', options: renderOptions, label: '列渲染方式' },
    },
    form: {
      ...baseFormAttr(),
      'dict-options': { type: 'optionsEditor', value: '', label: '选项列表', placeholder: '' },
    },
  },
  radio: {
    name: '单选框',
    defaultDbType: 'tinyint',
    defaultGoType: 'int',
    defaultTsType: 'number',
    table: {
      ...baseTableAttr('eq'),
      render: { type: 'select', value: 'tag', options: renderOptions, label: '列渲染方式' },
    },
    form: {
      ...baseFormAttr(),
      'dict-options': { type: 'optionsEditor', value: '', label: '选项列表', placeholder: '' },
    },
  },
  checkbox: {
    name: '复选框',
    defaultDbType: 'varchar(255)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('in'),
      render: { type: 'select', value: 'tags', options: renderOptions, label: '列渲染方式' },
    },
    form: {
      ...baseFormAttr(),
      'dict-options': { type: 'optionsEditor', value: '', label: '选项列表', placeholder: '' },
    },
  },
  select: {
    name: '下拉选择',
    defaultDbType: 'varchar(50)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('eq'),
      render: { type: 'select', value: 'tag', options: renderOptions, label: '列渲染方式' },
    },
    form: {
      ...baseFormAttr(),
      'dict-options': { type: 'optionsEditor', value: '', label: '选项列表', placeholder: '' },
    },
  },
  selects: {
    name: '下拉多选',
    defaultDbType: 'varchar(255)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('in'),
      render: { type: 'select', value: 'tags', options: renderOptions, label: '列渲染方式' },
    },
    form: {
      ...baseFormAttr(),
      'dict-options': { type: 'optionsEditor', value: '', label: '选项列表', placeholder: '' },
    },
  },
  textarea: {
    name: '文本域',
    defaultDbType: 'varchar(500)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('false'),
    form: {
      ...baseFormAttr(),
      rows: { type: 'number', value: 3, label: '行数' },
    },
  },
  password: {
    name: '密码',
    defaultDbType: 'varchar(64)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('false'),
    form: {
      ...baseFormAttr(),
      validator: { type: 'selects', value: ['password'], options: validatorTypes, label: '验证规则' },
    },
  },
  datetime: {
    name: '日期时间',
    defaultDbType: 'datetime',
    defaultGoType: '*gtime.Time',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('between', 'custom', 'datetime'),
      width: { type: 'number', value: 170, label: '列宽' },
    },
    form: {
      ...baseFormAttr(),
      validator: { type: 'selects', value: ['date'], options: validatorTypes, label: '验证规则' },
    },
  },
  date: {
    name: '日期',
    defaultDbType: 'date',
    defaultGoType: '*gtime.Time',
    defaultTsType: 'string',
    table: baseTableAttr('between', 'custom'),
    form: {
      ...baseFormAttr(),
      validator: { type: 'selects', value: ['date'], options: validatorTypes, label: '验证规则' },
    },
  },
  time: {
    name: '时间',
    defaultDbType: 'time',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('eq', 'custom'),
    form: baseFormAttr(),
  },
  timestamp: {
    name: '时间戳',
    defaultDbType: 'bigint',
    defaultGoType: 'int64',
    defaultTsType: 'number',
    table: {
      ...baseTableAttr('between', 'custom', 'datetime'),
      width: { type: 'number', value: 170, label: '列宽' },
    },
    form: baseFormAttr(),
  },
  image: {
    name: '图片上传',
    defaultDbType: 'varchar(255)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('false'),
      render: { type: 'select', value: 'image', options: renderOptions, label: '列渲染方式' },
    },
    form: baseFormAttr(),
  },
  images: {
    name: '多图上传',
    defaultDbType: 'varchar(1500)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('false'),
      render: { type: 'select', value: 'images', options: renderOptions, label: '列渲染方式' },
    },
    form: baseFormAttr(),
  },
  file: {
    name: '文件上传',
    defaultDbType: 'varchar(255)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('false'),
    form: baseFormAttr(),
  },
  files: {
    name: '多文件上传',
    defaultDbType: 'varchar(1500)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('false'),
    form: baseFormAttr(),
  },
  editor: {
    name: '富文本',
    defaultDbType: 'text',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('false'),
    form: baseFormAttr(),
  },
  color: {
    name: '颜色选择',
    defaultDbType: 'varchar(30)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('false'),
      render: { type: 'select', value: 'color', options: renderOptions, label: '列渲染方式' },
    },
    form: baseFormAttr(),
  },
  icon: {
    name: '图标选择',
    defaultDbType: 'varchar(50)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: {
      ...baseTableAttr('false'),
      render: { type: 'select', value: 'icon', options: renderOptions, label: '列渲染方式' },
    },
    form: baseFormAttr(),
  },
  city: {
    name: '城市选择',
    defaultDbType: 'varchar(100)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('false'),
    form: baseFormAttr(),
  },
  remoteSelect: {
    name: '远程下拉',
    defaultDbType: 'int(10) unsigned',
    defaultGoType: 'uint',
    defaultTsType: 'number',
    table: baseTableAttr('eq'),
    form: {
      ...baseFormAttr(),
      'remote-table': { type: 'remoteTableSelect', value: '', label: '关联数据表', placeholder: '选择关联表' },
      'remote-pk': { type: 'remoteColumnSelect', value: 'id', label: '关联主键', placeholder: '选择主键字段' },
      'remote-field': { type: 'remoteColumnSelect', value: '', label: '显示Label', placeholder: '自动推导，可手动选择' },
      'relation-fields-config': { type: 'relationFieldsDesigner', value: '', label: '关联字段配置', placeholder: '' },
    },
  },
  remoteSelects: {
    name: '远程多选',
    defaultDbType: 'varchar(255)',
    defaultGoType: 'string',
    defaultTsType: 'string',
    table: baseTableAttr('in'),
    form: {
      ...baseFormAttr(),
      'remote-table': { type: 'remoteTableSelect', value: '', label: '关联数据表', placeholder: '选择关联表' },
      'remote-pk': { type: 'remoteColumnSelect', value: 'id', label: '关联主键', placeholder: '选择主键字段' },
      'remote-field': { type: 'remoteColumnSelect', value: '', label: '显示Label', placeholder: '自动推导，可手动选择' },
      'relation-fields-config': { type: 'relationFieldsDesigner', value: '', label: '关联字段配置', placeholder: '' },
    },
  },
  weigh: {
    name: '权重(拖拽排序)',
    defaultDbType: 'int(10)',
    defaultGoType: 'int',
    defaultTsType: 'number',
    table: baseTableAttr('eq', 'custom'),
    form: {
      ...baseFormAttr(),
      step: { type: 'number', value: 1, label: '步长' },
    },
  },
}

// ==================== 预设字段库 ====================

/** 常用字段 */
export const commonFields: PresetField[] = [
  {
    title: '主键',
    name: 'id',
    designType: 'pk',
    comment: '主键',
    dbType: 'bigint(20) unsigned',
    goType: 'uint64',
    tsType: 'number',
    isPk: 1,
    isList: 1,
    isEdit: 0,
  },
  {
    title: '状态',
    name: 'status',
    designType: 'switch',
    comment: '状态:0=禁用,1=启用',
    dbType: 'tinyint(1)',
    goType: 'int',
    tsType: 'number',
    isList: 1,
    isEdit: 1,
  },
  {
    title: '排序权重',
    name: 'sort',
    designType: 'weigh',
    comment: '排序权重',
    dbType: 'int(10)',
    goType: 'int',
    tsType: 'number',
    isList: 1,
    isEdit: 1,
  },
  {
    title: '创建时间',
    name: 'created_at',
    designType: 'timestamp',
    comment: '创建时间',
    dbType: 'bigint(16)',
    goType: 'int64',
    tsType: 'number',
    isList: 1,
    isEdit: 0,
  },
  {
    title: '更新时间',
    name: 'updated_at',
    designType: 'timestamp',
    comment: '更新时间',
    dbType: 'bigint(16)',
    goType: 'int64',
    tsType: 'number',
    isList: 1,
    isEdit: 0,
  },
  {
    title: '备注',
    name: 'remark',
    designType: 'textarea',
    comment: '备注',
    dbType: 'varchar(255)',
    goType: 'string',
    tsType: 'string',
    isList: 0,
    isEdit: 1,
  },
  {
    title: '远程下拉(关联)',
    name: 'remote_id',
    designType: 'remoteSelect',
    comment: '关联ID',
    dbType: 'int(10) unsigned',
    goType: 'uint',
    tsType: 'number',
    isList: 0,
    isEdit: 1,
  },
]

/** 基础字段 */
export const baseFields: PresetField[] = [
  { title: '字符串', name: 'string_field', designType: 'string', comment: '字符串', dbType: 'varchar(255)', goType: 'string', tsType: 'string' },
  { title: '数字', name: 'number_field', designType: 'number', comment: '数字', dbType: 'int(10)', goType: 'int', tsType: 'number' },
  { title: '浮点数', name: 'float_field', designType: 'float', comment: '浮点数', dbType: 'decimal(10,2)', goType: 'float64', tsType: 'number' },
  { title: '开关', name: 'switch_field', designType: 'switch', comment: '开关:0=关,1=开', dbType: 'tinyint(1)', goType: 'int', tsType: 'number' },
  { title: '单选', name: 'radio_field', designType: 'radio', comment: '单选:0=选项A,1=选项B', dbType: 'tinyint', goType: 'int', tsType: 'number' },
  { title: '复选', name: 'checkbox_field', designType: 'checkbox', comment: '复选', dbType: 'varchar(255)', goType: 'string', tsType: 'string' },
  { title: '下拉选择', name: 'select_field', designType: 'select', comment: '下拉', dbType: 'varchar(50)', goType: 'string', tsType: 'string' },
  { title: '文本域', name: 'textarea_field', designType: 'textarea', comment: '文本域', dbType: 'varchar(500)', goType: 'string', tsType: 'string', isList: 0 },
  { title: '日期', name: 'date_field', designType: 'date', comment: '日期', dbType: 'date', goType: '*gtime.Time', tsType: 'string' },
  { title: '日期时间', name: 'datetime_field', designType: 'datetime', comment: '日期时间', dbType: 'datetime', goType: '*gtime.Time', tsType: 'string' },
  { title: '时间', name: 'time_field', designType: 'time', comment: '时间', dbType: 'time', goType: 'string', tsType: 'string' },
  { title: '密码', name: 'password_field', designType: 'password', comment: '密码', dbType: 'varchar(64)', goType: 'string', tsType: 'string', isList: 0 },
]

/** 高级字段 */
export const seniorFields: PresetField[] = [
  { title: '图片', name: 'image_field', designType: 'image', comment: '图片', dbType: 'varchar(255)', goType: 'string', tsType: 'string' },
  { title: '多图', name: 'images_field', designType: 'images', comment: '多图', dbType: 'varchar(1500)', goType: 'string', tsType: 'string' },
  { title: '文件', name: 'file_field', designType: 'file', comment: '文件', dbType: 'varchar(255)', goType: 'string', tsType: 'string', isList: 0 },
  { title: '多文件', name: 'files_field', designType: 'files', comment: '多文件', dbType: 'varchar(1500)', goType: 'string', tsType: 'string', isList: 0 },
  { title: '富文本', name: 'content', designType: 'editor', comment: '内容', dbType: 'text', goType: 'string', tsType: 'string', isList: 0 },
  { title: '图标选择', name: 'icon_field', designType: 'icon', comment: '图标', dbType: 'varchar(50)', goType: 'string', tsType: 'string' },
  { title: '颜色选择', name: 'color_field', designType: 'color', comment: '颜色', dbType: 'varchar(30)', goType: 'string', tsType: 'string' },
  { title: '远程多选', name: 'remote_ids', designType: 'remoteSelects', comment: '关联IDs', dbType: 'varchar(255)', goType: 'string', tsType: 'string', isList: 0 },
]

// ==================== 工具函数 ====================

/** 获取 designType 的显示名称 */
export function getDesignTypeName(dt: string): string {
  return designTypes[dt]?.name || dt
}

/** 获取 designType 对应的表格属性默认值 */
export function getDefaultTableProps(dt: string): Record<string, any> {
  const def = designTypes[dt]
  if (!def) return {}
  const result: Record<string, any> = {}
  for (const [key, prop] of Object.entries(def.table)) {
    result[key] = prop.value
  }
  return result
}

/** 获取 designType 对应的表单属性默认值 */
export function getDefaultFormProps(dt: string): Record<string, any> {
  const def = designTypes[dt]
  if (!def) return {}
  const result: Record<string, any> = {}
  for (const [key, prop] of Object.entries(def.form)) {
    result[key] = prop.value
  }
  return result
}

/** snake_case -> PascalCase */
export function snakeToPascal(s: string): string {
  return s.split('_').map(p => p.charAt(0).toUpperCase() + p.slice(1)).join('')
}

/** snake_case -> camelCase */
export function snakeToCamel(s: string): string {
  const pascal = snakeToPascal(s)
  return pascal.charAt(0).toLowerCase() + pascal.slice(1)
}
