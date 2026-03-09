<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 关联表字段配置表格 -->
<template>
  <div class="rel-field-table">
    <div v-if="loading" class="rel-field-table__loading">加载中...</div>
    <div v-else-if="!columns.length" class="rel-field-table__empty">
      该关联表无字段数据
    </div>
    <ElTable v-else :data="fieldConfigs" size="small" border stripe max-height="480">
      <ElTableColumn prop="field" label="字段名" width="130" />
      <ElTableColumn prop="label" label="描述" width="120">
        <template #default="{ row }">
          <ElInput v-model="row.label" size="small" @change="syncConfig" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="列表" width="60" align="center">
        <template #default="{ row }">
          <ElCheckbox v-model="row.inList" size="small" @change="syncConfig" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="搜索" width="60" align="center">
        <template #default="{ row }">
          <ElCheckbox v-model="row.inSearch" size="small" @change="syncConfig" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="导出" width="60" align="center">
        <template #default="{ row }">
          <ElCheckbox v-model="row.inExport" size="small" @change="syncConfig" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="列渲染" width="110">
        <template #default="{ row }">
          <ElSelect v-if="row.inList" v-model="row.listRender" size="small" @change="syncConfig">
            <ElOption value="text" label="文本" />
            <ElOption value="tag" label="Tag" />
            <ElOption value="tags" label="Tags" />
            <ElOption value="switch" label="开关" />
            <ElOption value="image" label="图片" />
            <ElOption value="images" label="多图" />
            <ElOption value="url" label="URL" />
            <ElOption value="datetime" label="时间日期" />
            <ElOption value="color" label="颜色" />
            <ElOption value="icon" label="图标" />
          </ElSelect>
          <span v-else class="text-gray-400 text-xs">—</span>
        </template>
      </ElTableColumn>
      <ElTableColumn label="搜索方式" width="110">
        <template #default="{ row }">
          <ElSelect v-if="row.inSearch" v-model="row.searchType" size="small" @change="syncConfig">
            <ElOption value="like" label="模糊(LIKE)" />
            <ElOption value="eq" label="精确(=)" />
            <ElOption value="between" label="区间" />
            <ElOption value="in" label="IN" />
          </ElSelect>
          <span v-else class="text-gray-400 text-xs">—</span>
        </template>
      </ElTableColumn>
      <ElTableColumn label="搜索组件" width="130">
        <template #default="{ row }">
          <ElSelect v-if="row.inSearch" v-model="row.searchComponent" size="small" @change="syncConfig">
            <ElOption value="input" label="输入框" />
            <ElOption value="inputTag" label="标签输入" />
            <ElOption value="number" label="数字" />
            <ElOption value="select" label="下拉选择" />
            <ElOption value="switch" label="开关" />
            <ElOption value="date" label="日期" />
            <ElOption value="daterange" label="日期范围" />
            <ElOption value="datetime" label="日期时间" />
            <ElOption value="datetimerange" label="日期时间范围" />
            <ElOption value="timepicker" label="时间" />
            <ElOption value="cascader" label="级联" />
            <ElOption value="treeselect" label="树选择" />
          </ElSelect>
          <span v-else class="text-gray-400 text-xs">—</span>
        </template>
      </ElTableColumn>
    </ElTable>
  </div>
</template>

<script setup lang="ts">
  import { fetchGenCodesColumnList } from '@/api/backend/develop/genCodes'

  interface RelFieldConfig {
    field: string
    label: string
    inList: boolean
    inSearch: boolean
    inExport: boolean
    searchType: string
    searchComponent: string
    listRender: string
  }

  const props = defineProps<{
    field: any        // 主表中的 remoteSelect 字段对象
    remoteTable: string
  }>()

  const loading = ref(false)
  const columns = ref<any[]>([])
  const fieldConfigs = ref<RelFieldConfig[]>([])

  // 从注释提取 label
  const extractLabel = (comment: string, fallback: string): string => {
    if (!comment) return fallback
    const idx = comment.search(/[:：]/)
    return idx > 0 ? comment.substring(0, idx) : comment
  }

  // 将配置同步回字段的 _formProps
  const syncConfig = () => {
    if (!props.field) return
    if (!props.field._formProps) props.field._formProps = {}
    props.field._formProps['relation-fields-config'] = JSON.stringify(fieldConfigs.value)
    // 兼容旧逻辑
    props.field._formProps['relation-fields'] = fieldConfigs.value.filter(f => f.inList).map(f => f.field).join(',')
    props.field._formProps['relation-search-fields'] = fieldConfigs.value.filter(f => f.inSearch).map(f => f.field).join(',')
    props.field._formProps['relation-export-fields'] = fieldConfigs.value.filter(f => f.inExport).map(f => f.field).join(',')
  }

  // 加载关联表字段
  const loadColumns = async () => {
    if (!props.remoteTable) return
    loading.value = true
    try {
      const res = await fetchGenCodesColumnList(props.remoteTable)
      columns.value = (res.list || []).map((col: any) => ({
        columnName: col.name || col.columnName,
        columnComment: col.comment || col.columnComment || '',
        dataType: col.dbType || col.dataType || '',
      }))

      // 恢复已有配置或初始化
      const existingStr = props.field?._formProps?.['relation-fields-config'] || ''
      let existing: RelFieldConfig[] = []
      if (existingStr) {
        try { existing = JSON.parse(existingStr) } catch { /* ignore */ }
      }

      // 以远程表字段为准，合并已有配置
      const existingMap = new Map(existing.map(e => [e.field, e]))
      fieldConfigs.value = columns.value.map((col: any) => {
        const prev = existingMap.get(col.columnName)
        return prev ? { ...prev } : {
          field: col.columnName,
          label: extractLabel(col.columnComment, col.columnName),
          inList: false,
          inSearch: false,
          inExport: false,
          searchType: 'like',
          searchComponent: 'input',
          listRender: 'text',
        }
      })
      // 只有从旧格式迁移（有 relation-fields 但没有 relation-fields-config）时才需要初始同步
      if (!existingStr && fieldConfigs.value.some(f => f.inList || f.inSearch || f.inExport)) {
        syncConfig()
      }
    } catch { /* ignore */ }
    loading.value = false
  }

  // 监听 remoteTable 变化重新加载
  watch(() => props.remoteTable, () => loadColumns(), { immediate: true })
</script>

<style scoped>
  .rel-field-table {
    padding: 8px;
    min-height: 200px;
  }

  .rel-field-table__loading,
  .rel-field-table__empty {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 200px;
    color: var(--el-text-color-placeholder);
    font-size: 14px;
  }
</style>
