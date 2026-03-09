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
  <div class="space-y-5 mb-5">
    <div class="text-2xl font-medium mt-5 max-sm:text-2xl max-sm:mt-3">数组编辑器</div>
    <div class="text-g-800">
      ArtArrayEditor 是一个强大的多维数组编辑器，支持多种数据类型，可灵活配置字段，支持拖拽排序。
    </div>

    <!-- 动态字段配置示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">示例0：动态字段配置（使用字段编辑器）</div>
      <div class="text-sm text-g-600 mb-4">
        先配置字段，然后数组编辑器会根据配置动态渲染表单
      </div>
      
      <div class="grid grid-cols-2 gap-4">
        <!-- 左侧：字段配置编辑器 -->
        <div>
          <div class="text-sm font-medium mb-2">字段配置编辑器</div>
          <ArtFieldEditor v-model="customFields" />
        </div>
        
        <!-- 右侧：数组编辑器 -->
        <div>
          <div class="text-sm font-medium mb-2">数据编辑器</div>
          <ArtArrayEditor
            v-model="customData"
            :fields="customFields"
            :show-index="true"
          />
          
          <div class="mt-4 p-4 bg-g-100 dark:bg-g-800 rounded">
            <div class="text-sm text-g-600 mb-2">当前配置：</div>
            <pre class="text-xs text-g-700 font-mono">{{ JSON.stringify(customFields, null, 2) }}</pre>
          </div>
          
          <div class="mt-4 p-4 bg-g-100 dark:bg-g-800 rounded">
            <div class="text-sm text-g-600 mb-2">当前数据：</div>
            <pre class="text-xs text-g-700 font-mono">{{ JSON.stringify(customData, null, 2) }}</pre>
          </div>
        </div>
      </div>
    </div>

    <!-- 基础示例 - 用户列表 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">示例1：用户列表</div>
      <ArtArrayEditor
        v-model="userList"
        :fields="userFields"
        :show-index="true"
      />
      <div class="mt-4 p-4 bg-g-100 dark:bg-g-800 rounded">
        <div class="text-sm text-g-600 mb-2">当前数据：</div>
        <pre class="text-xs text-g-700 font-mono">{{ JSON.stringify(userList, null, 2) }}</pre>
      </div>
    </div>

    <!-- 示例2 - 权限配置 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">示例2：权限配置</div>
      <ArtArrayEditor
        v-model="permissionList"
        :fields="permissionFields"
        :show-index="true"
      />
    </div>

    <!-- 示例3 - 菜单配置 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">示例3：菜单配置</div>
      <ArtArrayEditor
        v-model="menuList"
        :fields="menuFields"
        :show-index="true"
      />
    </div>

    <!-- 示例4 - 商品管理（包含图片文件） -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">示例4：商品管理（包含图片上传）</div>
      <ArtArrayEditor
        v-model="productList"
        :fields="productFields"
        :show-index="true"
      />
    </div>

    <!-- 示例5 - 文档管理（包含文件上传） -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">示例5：文档管理（包含文件上传）</div>
      <ArtArrayEditor
        v-model="documentList"
        :fields="documentFields"
        :show-index="true"
      />
    </div>

    <!-- 代码示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">代码示例</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-600 mb-2">基础使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;template&gt;
  &lt;ArtArrayEditor 
    v-model="dataList" 
    :fields="fieldConfig"
    :show-index="true"
  /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const dataList = ref([])

const fieldConfig = [
  {
    key: 'name',
    label: '姓名',
    type: 'text',
    placeholder: '请输入姓名'
  },
  {
    key: 'role',
    label: '角色',
    type: 'select',
    options: [
      { label: '管理员', value: 'admin' },
      { label: '用户', value: 'user' }
    ]
  }
]
&lt;/script&gt;</code></pre>
          </div>
        </div>
      </div>
    </div>

    <!-- 字段类型说明 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">支持的字段类型</div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead class="bg-g-100 dark:bg-g-800">
            <tr>
              <th class="px-4 py-2 text-left">类型</th>
              <th class="px-4 py-2 text-left">组件</th>
              <th class="px-4 py-2 text-left">说明</th>
              <th class="px-4 py-2 text-left">额外配置</th>
            </tr>
          </thead>
          <tbody>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">text</td>
              <td class="px-4 py-2">ElInput</td>
              <td class="px-4 py-2">文本输入框</td>
              <td class="px-4 py-2">placeholder</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">number</td>
              <td class="px-4 py-2">ElInputNumber</td>
              <td class="px-4 py-2">数字输入框</td>
              <td class="px-4 py-2">min, max</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">select</td>
              <td class="px-4 py-2">ElSelect</td>
              <td class="px-4 py-2">下拉单选</td>
              <td class="px-4 py-2">options[]</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">selects</td>
              <td class="px-4 py-2">ElSelect</td>
              <td class="px-4 py-2">下拉多选</td>
              <td class="px-4 py-2">options[]</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">switch</td>
              <td class="px-4 py-2">ElSwitch</td>
              <td class="px-4 py-2">开关</td>
              <td class="px-4 py-2">-</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">date</td>
              <td class="px-4 py-2">ElDatePicker</td>
              <td class="px-4 py-2">日期选择</td>
              <td class="px-4 py-2">-</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">color</td>
              <td class="px-4 py-2">ElColorPicker</td>
              <td class="px-4 py-2">颜色选择</td>
              <td class="px-4 py-2">-</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">icon</td>
              <td class="px-4 py-2">ArtIconSelector</td>
              <td class="px-4 py-2">图标选择</td>
              <td class="px-4 py-2">-</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- API -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">API</div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead class="bg-g-100 dark:bg-g-800">
            <tr>
              <th class="px-4 py-2 text-left">参数</th>
              <th class="px-4 py-2 text-left">说明</th>
              <th class="px-4 py-2 text-left">类型</th>
              <th class="px-4 py-2 text-left">默认值</th>
            </tr>
          </thead>
          <tbody>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">modelValue</td>
              <td class="px-4 py-2">数组数据</td>
              <td class="px-4 py-2">any[]</td>
              <td class="px-4 py-2">[]</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">fields</td>
              <td class="px-4 py-2">字段配置</td>
              <td class="px-4 py-2">FieldConfig[]</td>
              <td class="px-4 py-2">[]</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">showIndex</td>
              <td class="px-4 py-2">显示序号</td>
              <td class="px-4 py-2">boolean</td>
              <td class="px-4 py-2">true</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">sortable</td>
              <td class="px-4 py-2">支持拖拽排序</td>
              <td class="px-4 py-2">boolean</td>
              <td class="px-4 py-2">true</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { FieldConfig } from '@/components/core/forms/art-array-editor/index.vue'

defineOptions({ name: 'ArrayEditorPage' })

// 示例0：动态字段配置
const customFields = ref<FieldConfig[]>([
  {
    key: 'name',
    label: '姓名',
    type: 'text',
    placeholder: '请输入姓名'
  },
  {
    key: 'role',
    label: '角色',
    type: 'select',
    options: [
      { label: '管理员', value: 'admin' },
      { label: '用户', value: 'user' },
      { label: '访客', value: 'guest' }
    ]
  },
  {
    key: 'tags',
    label: '标签',
    type: 'selects',
    options: [
      { label: 'Vue', value: 'vue' },
      { label: 'React', value: 'react' },
      { label: 'Angular', value: 'angular' }
    ]
  },
  {
    key: 'avatar',
    label: '头像',
    type: 'image',
    maxSize: 5
  }
])

const customData = ref<any[]>([])

// 示例1：用户列表
const userList = ref([
  { name: '张三', role: 'admin', tags: ['vue', 'react'], active: true },
  { name: '李四', role: 'user', tags: ['angular'], active: false }
])

const userFields: FieldConfig[] = [
  {
    key: 'name',
    label: '姓名',
    type: 'text',
    placeholder: '请输入姓名'
  },
  {
    key: 'role',
    label: '角色',
    type: 'select',
    options: [
      { label: '管理员', value: 'admin' },
      { label: '用户', value: 'user' },
      { label: '访客', value: 'guest' }
    ]
  },
  {
    key: 'tags',
    label: '标签',
    type: 'selects',
    options: [
      { label: 'Vue', value: 'vue' },
      { label: 'React', value: 'react' },
      { label: 'Angular', value: 'angular' }
    ]
  },
  {
    key: 'active',
    label: '激活',
    type: 'switch'
  }
]

// 示例2：权限配置
const permissionList = ref([
  { module: 'user', action: 'read', enabled: true },
  { module: 'user', action: 'write', enabled: false }
])

const permissionFields: FieldConfig[] = [
  {
    key: 'module',
    label: '模块',
    type: 'select',
    options: [
      { label: '用户管理', value: 'user' },
      { label: '角色管理', value: 'role' },
      { label: '菜单管理', value: 'menu' }
    ]
  },
  {
    key: 'action',
    label: '操作',
    type: 'select',
    options: [
      { label: '查看', value: 'read' },
      { label: '编辑', value: 'write' },
      { label: '删除', value: 'delete' }
    ]
  },
  {
    key: 'enabled',
    label: '启用',
    type: 'switch'
  }
]

// 示例3：菜单配置
const menuList = ref([
  { title: '首页', icon: 'ri:home-line', path: '/home', sort: 1 }
])

const menuFields: FieldConfig[] = [
  {
    key: 'title',
    label: '菜单名称',
    type: 'text',
    placeholder: '请输入菜单名称'
  },
  {
    key: 'icon',
    label: '图标',
    type: 'icon'
  },
  {
    key: 'path',
    label: '路径',
    type: 'text',
    placeholder: '请输入路径'
  },
  {
    key: 'sort',
    label: '排序',
    type: 'number',
    min: 0,
    max: 999,
    default: 0
  }
]

// 示例4：商品管理
const productList = ref([
  { 
    name: 'iPhone 15', 
    price: 5999, 
    category: 'electronics',
    image: '',
    color: '#000000',
    inStock: true
  }
])

const productFields: FieldConfig[] = [
  {
    key: 'name',
    label: '商品名称',
    type: 'text',
    placeholder: '请输入商品名称'
  },
  {
    key: 'price',
    label: '价格',
    type: 'number',
    min: 0,
    default: 0
  },
  {
    key: 'category',
    label: '分类',
    type: 'select',
    options: [
      { label: '电子产品', value: 'electronics' },
      { label: '服装', value: 'clothing' },
      { label: '食品', value: 'food' }
    ]
  },
  {
    key: 'image',
    label: '商品图片',
    type: 'image',
    maxSize: 5
  },
  {
    key: 'color',
    label: '颜色',
    type: 'color'
  },
  {
    key: 'inStock',
    label: '有货',
    type: 'switch'
  }
]

// 示例5：文档管理
const documentList = ref([
  {
    title: '用户手册',
    type: 'pdf',
    file: '',
    uploadDate: '',
    public: false
  }
])

const documentFields: FieldConfig[] = [
  {
    key: 'title',
    label: '文档标题',
    type: 'text',
    placeholder: '请输入文档标题'
  },
  {
    key: 'type',
    label: '文档类型',
    type: 'select',
    options: [
      { label: 'PDF文档', value: 'pdf' },
      { label: 'Word文档', value: 'word' },
      { label: 'Excel表格', value: 'excel' },
      { label: '其他', value: 'other' }
    ]
  },
  {
    key: 'file',
    label: '文档文件',
    type: 'file',
    accept: '.pdf,.doc,.docx,.xls,.xlsx',
    maxSize: 20
  },
  {
    key: 'uploadDate',
    label: '上传日期',
    type: 'date'
  },
  {
    key: 'public',
    label: '公开',
    type: 'switch'
  }
]
</script>
