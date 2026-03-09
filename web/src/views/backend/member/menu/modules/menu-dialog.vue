<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 会员菜单编辑对话框（对齐 BuildAdmin user/rule/popupForm） -->
<template>
  <ElDialog
    :title="dialogTitle"
    :model-value="visible"
    @update:model-value="handleCancel"
    width="700px"
    align-center
    @closed="handleClosed"
  >
    <ElForm ref="formRef" :model="form" :rules="currentRules" label-width="100px">

      <!-- 上级菜单规则 -->
      <ElFormItem label="上级菜单" v-if="form.type !== 'button'">
        <ElTreeSelect
          v-model="form.pid"
          :data="menuTreeOptions"
          check-strictly
          placeholder="点击选择"
          clearable
          style="width: 100%"
        />
      </ElFormItem>

      <!-- 规则类型（6 种） -->
      <ElFormItem label="规则类型" prop="type">
        <ElRadioGroup v-model="form.type">
          <ElRadioButton v-for="opt in typeOptions" :key="opt.value" :value="opt.value">
            <ElTooltip :content="opt.tip" placement="top" :show-after="300">
              <span>{{ opt.label }}</span>
            </ElTooltip>
          </ElRadioButton>
        </ElRadioGroup>
      </ElFormItem>

      <!-- 规则标题 -->
      <ElFormItem label="规则标题" prop="title">
        <ElInput v-model="form.title" placeholder="请输入规则标题" />
      </ElFormItem>

      <!-- 规则名称（非 button） -->
      <ElFormItem v-if="form.type !== 'button'" label="规则名称" prop="name">
        <ElInput v-model="form.name" placeholder="将注册为 web 端路由名称，同时作为 server 端 API 权鉴使用" />
      </ElFormItem>

      <!-- 路由路径（非 button） -->
      <ElFormItem v-if="form.type !== 'button'" label="路由路径" prop="path">
        <ElInput v-model="form.path" placeholder="路由路径" />
      </ElFormItem>

      <!-- 图标（非 button） -->
      <ElFormItem v-if="form.type !== 'button'" label="规则图标">
        <ArtIconSelector v-model="form.icon" />
      </ElFormItem>

      <!-- 组件路径（非 button） -->
      <ElFormItem v-if="form.type !== 'button'" label="组件路径">
        <ElInput v-model="form.component" placeholder="Vue 组件路径，如 frontend/member/center" />
      </ElFormItem>

      <!-- 菜单类型 tab/link/iframe（仅 menu/nav/nav_user_menu 显示） -->
      <ElFormItem v-if="showMenuType" label="菜单类型">
        <ElRadioGroup v-model="form.menuType">
          <ElRadioButton value="tab">标签卡</ElRadioButton>
          <ElRadioButton value="link">链接(站外)</ElRadioButton>
          <ElRadioButton value="iframe">Iframe</ElRadioButton>
        </ElRadioGroup>
      </ElFormItem>

      <!-- 链接地址（仅 link/iframe 显示） -->
      <ElFormItem v-if="showUrl" label="链接地址">
        <ElInput v-model="form.url" placeholder="https://example.com" />
      </ElFormItem>

      <!-- 未登录有效（仅 route/nav/button 显示，menu_dir/menu/nav_user_menu 强制为 0） -->
      <ElFormItem v-if="showNoLoginValid" label="未登录有效">
        <ElRadioGroup v-model="form.noLoginValid">
          <ElRadioButton :value="0">游客无效</ElRadioButton>
          <ElRadioButton :value="1">游客有效</ElRadioButton>
        </ElRadioGroup>
        <div class="form-item-tip">游客没有会员分组，通过本选项设置当前规则是否对游客有效（可见）</div>
      </ElFormItem>

      <!-- 权限标识（仅 button 显示） -->
      <ElFormItem v-if="form.type === 'button'" label="权限标识">
        <ElInput v-model="form.permission" placeholder="如：user:profile:view" />
      </ElFormItem>

      <!-- 扩展属性 -->
      <ElFormItem v-if="form.type !== 'button'" label="扩展属性">
        <ElSelect v-model="form.extend" style="width: 100%">
          <ElOption value="none" label="无" />
          <ElOption value="add_rules_only" label="只添加为路由" />
          <ElOption value="add_menu_only" label="只添加为菜单" />
        </ElSelect>
      </ElFormItem>

      <!-- 排序 + 状态 -->
      <ElFormItem label="排序">
        <div class="flex items-center gap-8 w-full">
          <ElInputNumber v-model="form.sort" :min="0" :max="9999" style="width: 160px" />
          <div class="flex items-center gap-2">
            <span class="text-sm text-gray-500">状态</span>
            <ElSwitch v-model="form.status" :active-value="1" :inactive-value="0" />
          </div>
        </div>
      </ElFormItem>

      <!-- 备注 -->
      <ElFormItem label="规则备注">
        <ElInput v-model="form.remark" placeholder="请输入规则备注" />
      </ElFormItem>
    </ElForm>

    <template #footer>
      <ElButton @click="handleCancel">取 消</ElButton>
      <ElButton type="primary" @click="handleSubmit" :loading="submitting">确 定</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
import type { FormInstance, FormRules } from 'element-plus'
import ArtIconSelector from '@/components/core/forms/art-icon-selector/index.vue'
import { saveMemberMenu, type MemberMenuItem, type RuleType, type MenuOpenType } from '@/api/backend/member/menu'

interface MenuFormData {
  id: number
  pid: number
  title: string
  name: string
  path: string
  component: string
  icon: string
  menuType: MenuOpenType
  url: string
  noLoginValid: number
  extend: string
  remark: string
  type: RuleType
  permission: string
  sort: number
  status: number
}

interface Props {
  visible: boolean
  editData?: MemberMenuItem | null
  parentMenu?: MemberMenuItem | null
  menuTree?: MemberMenuItem[]
}

const props = withDefaults(defineProps<Props>(), {
  editData: null,
  parentMenu: null,
  menuTree: () => []
})

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
  (e: 'submit'): void
}>()

// 规则类型选项
const typeOptions: { value: RuleType; label: string; tip: string }[] = [
  { value: 'route', label: '普通路由', tip: '自动注册为前端路由' },
  { value: 'menu_dir', label: '菜单目录', tip: '自动注册路由，并作为会员中心的菜单目录，此项本身不可跳转' },
  { value: 'menu', label: '菜单项', tip: '自动注册路由，并作为会员中心的菜单项目' },
  { value: 'nav', label: '顶栏菜单', tip: '自动注册路由，并作为站点顶栏的菜单项目' },
  { value: 'nav_user_menu', label: '顶栏下拉', tip: '自动注册路由，并作为顶栏会员菜单下拉项' },
  { value: 'button', label: '按钮', tip: '自动注册为权限节点，可通过 v-auth 快速验权' },
]

const formRef = ref<FormInstance>()
const submitting = ref(false)

const form = reactive<MenuFormData>({
  id: 0,
  pid: 0,
  title: '',
  name: '',
  path: '',
  component: '',
  icon: '',
  menuType: 'tab',
  url: '',
  noLoginValid: 0,
  extend: 'none',
  remark: '',
  type: 'menu',
  permission: '',
  sort: 0,
  status: 1,
})

// 弹窗标题
const dialogTitle = computed(() => {
  const typeLabel = typeOptions.find(o => o.value === form.type)?.label || '菜单'
  return props.editData ? `编辑${typeLabel}` : `添加${typeLabel}`
})

// 条件显示：菜单类型 (tab/link/iframe)
// 仅当 type 不是 menu_dir/button/route 时显示
const showMenuType = computed(() =>
  !['menu_dir', 'button', 'route'].includes(form.type)
)

// 条件显示：链接地址
// 仅当 menuType 是 link/iframe 时显示
const showUrl = computed(() =>
  showMenuType.value && ['link', 'iframe'].includes(form.menuType)
)

// 条件显示：未登录有效
// menu_dir/menu/nav_user_menu 强制为 0，不显示
const showNoLoginValid = computed(() =>
  !['menu_dir', 'menu', 'nav_user_menu'].includes(form.type)
)

// 菜单树选项
const menuTreeOptions = computed(() => {
  const formatTree = (items: MemberMenuItem[]): any[] => {
    return items
      .filter(item => item.type !== 'button')
      .map(item => ({
        label: item.title,
        value: item.id,
        children: item.children ? formatTree(item.children) : undefined
      }))
  }
  return [
    { label: '顶级菜单', value: 0 },
    ...formatTree(props.menuTree || [])
  ]
})

// 验证规则（根据类型动态）
const currentRules = computed<FormRules>(() => {
  const r: FormRules = {
    type: [{ required: true, message: '请选择规则类型', trigger: 'change' }],
    title: [{ required: true, message: '请输入规则标题', trigger: 'blur' }],
  }
  if (form.type !== 'button') {
    r.name = [{ required: true, message: '请输入规则名称', trigger: 'blur' }]
    r.path = [{ required: true, message: '请输入路由路径', trigger: 'blur' }]
  }
  return r
})

// 类型变化时的副作用
watch(() => form.type, (newType) => {
  // menu_dir/menu/nav_user_menu 强制 noLoginValid=0
  if (['menu_dir', 'menu', 'nav_user_menu'].includes(newType)) {
    form.noLoginValid = 0
  }
  // route 强制 menuType=tab
  if (newType === 'route') {
    form.menuType = 'tab'
  }
  // button 清空路由相关
  if (newType === 'button') {
    form.path = ''
    form.component = ''
    form.menuType = 'tab'
    form.url = ''
  }
})

// 初始化表单
const initForm = () => {
  if (props.editData) {
    const d = props.editData
    form.id = d.id
    form.pid = d.pid || 0
    form.title = d.title || ''
    form.name = d.name || ''
    form.path = d.path || ''
    form.component = d.component || ''
    form.icon = d.icon || ''
    form.menuType = d.menuType || 'tab'
    form.url = d.url || ''
    form.noLoginValid = d.noLoginValid ?? 0
    form.extend = d.extend || 'none'
    form.remark = d.remark || ''
    form.type = d.type || 'menu'
    form.permission = d.permission || ''
    form.sort = d.sort || 0
    form.status = d.status ?? 1
  } else {
    form.id = 0
    form.pid = props.parentMenu?.id || 0
    form.title = ''
    form.name = ''
    form.path = ''
    form.component = ''
    form.icon = ''
    form.menuType = 'tab'
    form.url = ''
    form.noLoginValid = 0
    form.extend = 'none'
    form.remark = ''
    form.type = 'menu'
    form.permission = ''
    form.sort = 0
    form.status = 1
  }
}

watch(() => props.visible, (val) => {
  if (val) {
    initForm()
    nextTick(() => formRef.value?.clearValidate())
  }
})

const handleCancel = () => emit('update:visible', false)
const handleClosed = () => formRef.value?.resetFields()

const handleSubmit = async () => {
  if (!formRef.value) return
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  submitting.value = true
  try {
    await saveMemberMenu({
      id: form.id || undefined,
      pid: form.pid,
      title: form.title,
      name: form.name,
      path: form.path,
      component: form.component,
      icon: form.icon,
      menuType: form.menuType,
      url: form.url,
      noLoginValid: form.noLoginValid,
      extend: form.extend,
      remark: form.remark,
      type: form.type,
      permission: form.permission,
      sort: form.sort,
      status: form.status,
    })
    ElMessage.success(form.id ? '编辑成功' : '新增成功')
    emit('update:visible', false)
    emit('submit')
  } catch (error) {
    console.error('保存菜单失败:', error)
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.form-item-tip {
  margin-top: 4px;
  font-size: 12px;
  color: var(--el-text-color-placeholder);
  line-height: 1.4;
}
</style>
