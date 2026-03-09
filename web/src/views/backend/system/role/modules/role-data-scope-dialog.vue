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
    v-model="visible"
    title="编辑数据权限"
    width="560px"
    align-center
    @close="handleClose"
  >
    <ElAlert type="info" :closable="false" class="mb-4" show-icon>
      <template #default>
        配置角色的数据访问范围，控制用户只能查看特定范围内的数据
      </template>
    </ElAlert>

    <ElForm ref="formRef" :model="form" :rules="rules" label-width="100px" class="mt-4">
      <ElFormItem label="数据范围" prop="dataScope" class="mb-6">
        <ElSelect
          v-model="form.dataScope"
          placeholder="请选择数据范围"
          style="width: 100%"
        >
          <template v-for="item in dataScopeOptions" :key="item.value || item.key">
            <!-- 分组 -->
            <ElOptionGroup v-if="item.type === 'group'" :label="item.label">
              <ElOption
                v-for="child in item.children"
                :key="child.value"
                :label="child.label"
                :value="child.value"
              />
            </ElOptionGroup>
            <!-- 单项 -->
            <ElOption v-else :label="item.label" :value="item.value" />
          </template>
        </ElSelect>
      </ElFormItem>

      <!-- 自定义部门选择（dataScope=4 时显示） -->
      <ElFormItem
        v-if="form.dataScope === 4"
        label="自定义部门"
        prop="customDepts"
        class="mb-6"
      >
        <ElSelect
          v-model="form.customDepts"
          multiple
          placeholder="请选择部门（暂未接入部门树，后续完善）"
          style="width: 100%"
        >
          <!-- TODO: 后续接入部门树数据 -->
          <ElOption label="示例部门1" :value="1" />
          <ElOption label="示例部门2" :value="2" />
        </ElSelect>
      </ElFormItem>

      <!-- 数据范围说明 -->
      <ElFormItem label="范围说明" class="mb-4">
        <div class="text-sm text-gray-600 leading-relaxed">
          {{ getDataScopeDesc(form.dataScope) }}
        </div>
      </ElFormItem>
    </ElForm>

    <template #footer>
      <ElButton @click="handleClose">取消</ElButton>
      <ElButton type="primary" :loading="loading" @click="handleSubmit">确定</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage } from 'element-plus'
import { fetchGetDataScopeSelect, fetchEditDataScope } from '@/api/backend/system'

type RoleListItem = Api.SystemManage.RoleListItem

interface Props {
  modelValue: boolean
  roleData?: RoleListItem
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
  roleData: undefined
})

const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const dataScopeOptions = ref<any[]>([])

/**
 * 弹窗显示状态双向绑定
 */
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

/**
 * 表单数据
 */
const form = reactive({
  id: 0,
  dataScope: 1,
  customDepts: [] as number[]
})

/**
 * 表单验证规则
 */
const rules = reactive<FormRules>({
  dataScope: [{ required: true, message: '请选择数据范围', trigger: 'change' }],
  customDepts: [
    {
      validator: (rule, value, callback) => {
        if (form.dataScope === 4 && (!value || value.length === 0)) {
          callback(new Error('自定义部门时至少选择一个部门'))
        } else {
          callback()
        }
      },
      trigger: 'change'
    }
  ]
})

/**
 * 数据范围描述映射
 */
const dataScopeDescMap: Record<number, string> = {
  1: '可以查看所有数据，不做任何过滤',
  2: '仅可以看到所属部门下的数据',
  3: '仅可以看到所属以及下级部门的数据',
  4: '可以看到设置的指定部门数据，支持多个',
  5: '仅可以看到自己创建的数据',
  6: '仅可以看到自己和直属下级创建的数据（直属下级是指当前用户往下1级的用户）',
  7: '仅可以看到自己和全部下级创建的数据（全部下级是指当前用户所有的下级用户，包含其下级的下级）'
}

/**
 * 获取数据范围描述
 */
const getDataScopeDesc = (dataScope: number) => {
  return dataScopeDescMap[dataScope] || '未知数据范围'
}

/**
 * 加载数据范围选项
 */
const loadDataScopeOptions = async () => {
  try {
    const res = await fetchGetDataScopeSelect()
    dataScopeOptions.value = res.list || []
  } catch (error) {
    console.error('加载数据范围选项失败', error)
  }
}

/**
 * 监听弹窗打开，初始化表单数据
 */
watch(
  () => props.modelValue,
  async (val) => {
    if (val && props.roleData) {
      form.id = props.roleData.id || 0
      form.dataScope = props.roleData.dataScope || 1
      // 解析自定义部门（JSON 字符串转数组）
      const customDepts = (props.roleData as any).customDepts
      if (customDepts) {
        try {
          form.customDepts = JSON.parse(customDepts)
        } catch {
          form.customDepts = []
        }
      } else {
        form.customDepts = []
      }
      // 加载数据范围选项
      await loadDataScopeOptions()
    }
  },
  { immediate: true }
)

/**
 * 关闭弹窗
 */
const handleClose = () => {
  formRef.value?.resetFields()
  emit('update:modelValue', false)
}

/**
 * 提交表单
 */
const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate()
  loading.value = true

  try {
    await fetchEditDataScope({
      id: form.id,
      dataScope: form.dataScope,
      customDepts: form.dataScope === 4 ? form.customDepts : undefined
    })

    ElMessage.success('数据权限配置成功')
    emit('success')
    emit('update:modelValue', false)
  } catch (error) {
    console.error('保存失败', error)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="scss">
// 遵循 Art 框架风格，无需自定义样式
</style>
