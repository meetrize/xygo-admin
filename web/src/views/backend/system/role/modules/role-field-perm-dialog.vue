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
    title="字段权限配置"
    width="900px"
    align-center
    destroy-on-close
    @close="handleClose"
  >
    <ElAlert type="info" :closable="false" class="mb-4" show-icon>
      <template #default>
        配置角色对特定资源字段的访问权限：<strong>不可见</strong>（隐藏字段）、<strong>只读</strong>（禁止编辑）、<strong>可编辑</strong>（完全权限）
      </template>
    </ElAlert>

    <ElForm label-width="100px" class="mt-4 mb-6">
      <ElFormItem label="选择资源" class="mb-0">
        <ElSelect
          v-model="selectedResource"
          placeholder="请选择资源（根据角色菜单权限动态加载）"
          style="width: 100%"
          @change="handleResourceChange"
          :loading="resourceLoading"
        >
          <ElOption 
            v-for="res in availableResources" 
            :key="res.code" 
            :label="res.label" 
            :value="res.code" 
          />
        </ElSelect>
      </ElFormItem>
    </ElForm>

    <ElDivider class="my-6" />

    <div v-if="selectedResource && fieldList.length > 0" v-loading="loading">
      <div class="flex items-center justify-between mb-4">
        <h4 class="m-0">字段权限配置</h4>
        <ElButton size="small" @click="resetAllFields">重置为默认</ElButton>
      </div>

      <ElTable :data="fieldList" border max-height="400" style="width: 100%">
        <ElTableColumn prop="fieldName" label="字段名称" width="150">
          <template #default="{ row }">
            <div class="flex items-center gap-2">
              <code class="text-xs text-blue-600">{{ row.fieldName }}</code>
              <ElTag v-if="row.isSensitive" type="warning" size="small">敏感</ElTag>
            </div>
          </template>
        </ElTableColumn>
        <ElTableColumn prop="fieldLabel" label="字段显示名" width="150" />
        <ElTableColumn label="权限设置" min-width="280">
          <template #default="{ row }">
            <ElRadioGroup v-model="row.permType">
              <ElRadio :label="0">
                <span class="flex items-center gap-1">
                  <ArtSvgIcon icon="ri:eye-off-line" :size="14" />
                  不可见
                </span>
              </ElRadio>
              <ElRadio :label="1">
                <span class="flex items-center gap-1">
                  <ArtSvgIcon icon="ri:lock-line" :size="14" />
                  只读
                </span>
              </ElRadio>
              <ElRadio :label="2">
                <span class="flex items-center gap-1">
                  <ArtSvgIcon icon="ri:edit-line" :size="14" />
                  可编辑
                </span>
              </ElRadio>
            </ElRadioGroup>
          </template>
        </ElTableColumn>
      </ElTable>
    </div>

    <ElEmpty
      v-else-if="selectedResource && !loading"
      description="该资源暂无字段定义"
    />

    <template #footer>
      <ElButton @click="handleClose">取消</ElButton>
      <ElButton
        type="primary"
        :loading="saveLoading"
        :disabled="!selectedResource"
        @click="handleSubmit"
      >
        保存
      </ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
import { ElMessage } from 'element-plus'
import {
  getFieldPermListApi,
  getResourceFieldsApi,
  batchSaveFieldPermApi
} from '@/api/backend/system/fieldPerm'
import { fetchRoleAvailableResources } from '@/api/backend/system'

type RoleListItem = Api.SystemManage.RoleListItem

interface Props {
  modelValue: boolean
  roleData?: RoleListItem
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

interface FieldItem {
  fieldName: string
  fieldLabel: string
  permType: number
  isSensitive?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
  roleData: undefined
})

const emit = defineEmits<Emits>()

const loading = ref(false)
const saveLoading = ref(false)
const resourceLoading = ref(false)
const roleId = ref(0)
const selectedResource = ref('')
const fieldList = ref<FieldItem[]>([])
const availableResources = ref<Array<{ code: string; label: string }>>([])

/**
 * 弹窗显示状态双向绑定
 */
const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

/**
 * 资源切换，加载字段列表
 */
const handleResourceChange = async () => {
  if (!selectedResource.value) {
    fieldList.value = []
    return
  }
  await loadFieldPerms()
}

/**
 * 加载字段权限配置
 */
const loadFieldPerms = async () => {
  loading.value = true
  try {
    // 1. 获取资源的字段列表
    const fieldsRes = await getResourceFieldsApi({ resource: selectedResource.value })
    const resourceFields = fieldsRes.fields || []

    // 2. 获取已配置的权限
    const permsRes = await getFieldPermListApi({
      roleId: roleId.value,
      resource: selectedResource.value
    })
    const configuredPerms = permsRes.list || []

    // 3. 合并：已配置的用配置值，未配置的默认可编辑
    const configuredMap = new Map(
      configuredPerms.map((item: any) => [item.fieldName, item.permType])
    )

    fieldList.value = resourceFields.map((field: any) => ({
      fieldName: field.fieldName,
      fieldLabel: field.fieldLabel,
      isSensitive: field.isSensitive,
      permType: configuredMap.get(field.fieldName) ?? 2 // 默认可编辑
    }))
  } catch (error) {
    console.error('加载字段权限失败', error)
    ElMessage.error('加载字段列表失败')
  } finally {
    loading.value = false
  }
}

/**
 * 重置所有字段为默认（可编辑）
 */
const resetAllFields = () => {
  fieldList.value.forEach((field) => {
    field.permType = 2 // 可编辑
  })
  ElMessage.success('已重置为默认权限')
}

/**
 * 监听弹窗打开
 */
watch(
  () => props.modelValue,
  async (val) => {
    if (val && props.roleData) {
      roleId.value = props.roleData.id || 0
      selectedResource.value = ''
      fieldList.value = []
      
      // ✅ 加载角色可配置的资源列表
      resourceLoading.value = true
      try {
        const res = await fetchRoleAvailableResources(props.roleData.id)
        availableResources.value = res.list || []
        console.log('[字段权限] 可用资源:', availableResources.value)
      } catch (error) {
        console.error('加载资源列表失败:', error)
        ElMessage.error('加载资源列表失败')
      } finally {
        resourceLoading.value = false
      }
    }
  }
)

/**
 * 关闭弹窗
 */
const handleClose = () => {
  selectedResource.value = ''
  fieldList.value = []
  emit('update:modelValue', false)
}

/**
 * 提交保存
 */
const handleSubmit = async () => {
  if (!selectedResource.value) {
    ElMessage.warning('请选择资源')
    return
  }

  if (fieldList.value.length === 0) {
    ElMessage.warning('没有可配置的字段')
    return
  }

  saveLoading.value = true
  try {
    await batchSaveFieldPermApi({
      roleId: roleId.value,
      resource: selectedResource.value,
      fields: fieldList.value.map((f) => ({
        fieldName: f.fieldName,
        fieldLabel: f.fieldLabel,
        permType: f.permType
      }))
    })

    ElMessage.success('字段权限保存成功')
    emit('success')
    emit('update:modelValue', false)
  } catch (error) {
    console.error('保存失败', error)
  } finally {
    saveLoading.value = false
  }
}
</script>

<style scoped lang="scss">
// 遵循 Art 框架风格，无需自定义样式
</style>
