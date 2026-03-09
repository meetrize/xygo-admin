<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 从零建表 - 表设计器 -->
<template>
  <div class="create-table-designer">
    <ElForm ref="formRef" :model="form" :rules="rules" label-width="80px">
      <div class="designer-header">
        <ElFormItem label="表名" prop="tableName" class="flex-1">
          <ElInput v-model="form.tableName" placeholder="如 biz_article (自动加 xy_ 前缀)">
            <template #prepend>xy_</template>
          </ElInput>
        </ElFormItem>
        <ElFormItem label="表注释" prop="tableComment" class="flex-1 ml-4">
          <ElInput v-model="form.tableComment" placeholder="如：文章管理" />
        </ElFormItem>
      </div>
    </ElForm>

    <div class="designer-toolbar">
      <h4 class="designer-toolbar__title">
        <ArtSvgIcon icon="ri:list-settings-line" class="text-base mr-1.5" />
        字段列表
      </h4>
      <div class="designer-toolbar__actions">
        <ElButton size="small" @click="addCommonFields">
            <ArtSvgIcon icon="ri:flashlight-line" class="text-sm mr-1" />
          添加公共字段
        </ElButton>
        <ElButton size="small" type="primary" @click="addColumn">
            <ArtSvgIcon icon="ri:add-line" class="text-sm mr-1" />
          添加字段
        </ElButton>
      </div>
    </div>

    <!-- 提示信息 -->
    <div class="designer-tips">
      <span class="designer-tips__item">字段名用 <b>snake_case</b>（如 user_name）</span>
      <span class="designer-tips__sep">|</span>
      <span class="designer-tips__item">注释格式 <b>标题:值=标签,值=标签</b>（如 状态:0=禁用,1=启用）</span>
      <span class="designer-tips__sep">|</span>
      <span class="designer-tips__item">_id 结尾自动识别为关联字段</span>
    </div>

    <ElTable :data="form.columns" border size="small" max-height="400" style="width:100%">
      <ElTableColumn label="字段名" width="140">
        <template #default="{ row }">
          <ElInput v-model="row.name" size="small" placeholder="如 category_id" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="类型" width="200">
        <template #default="{ row }">
          <ElSelect v-model="row.type" size="small" filterable allow-create placeholder="选择或输入类型">
            <ElOptionGroup label="整数">
              <ElOption label="bigint unsigned — 大整数/主键" value="bigint unsigned" />
              <ElOption label="int unsigned — 普通整数" value="int unsigned" />
              <ElOption label="int — 有符号整数" value="int" />
              <ElOption label="tinyint unsigned — 小整数(0-255)" value="tinyint unsigned" />
              <ElOption label="tinyint(1) — 布尔/开关" value="tinyint(1)" />
            </ElOptionGroup>
            <ElOptionGroup label="字符串">
              <ElOption label="varchar(50) — 短文本" value="varchar(50)" />
              <ElOption label="varchar(100) — 常用文本" value="varchar(100)" />
              <ElOption label="varchar(200) — 中等文本" value="varchar(200)" />
              <ElOption label="varchar(500) — 长文本" value="varchar(500)" />
              <ElOption label="char(1) — 单字符/状态码" value="char(1)" />
            </ElOptionGroup>
            <ElOptionGroup label="大文本">
              <ElOption label="text — 文本域" value="text" />
              <ElOption label="longtext — 富文本/编辑器" value="longtext" />
              <ElOption label="json — JSON数据" value="json" />
            </ElOptionGroup>
            <ElOptionGroup label="数值">
              <ElOption label="decimal(10,2) — 金额/浮点" value="decimal(10,2)" />
              <ElOption label="float — 浮点数" value="float" />
            </ElOptionGroup>
            <ElOptionGroup label="日期时间">
              <ElOption label="datetime — 日期时间" value="datetime" />
              <ElOption label="date — 日期" value="date" />
              <ElOption label="time — 时间" value="time" />
              <ElOption label="bigint unsigned — 时间戳(整数)" value="bigint unsigned" />
            </ElOptionGroup>
            <ElOptionGroup label="枚举">
              <ElOption label="enum('a','b') — 枚举(自定义)" value="enum('opt1','opt2')" />
              <ElOption label="set('a','b') — 集合(多选)" value="set('opt1','opt2')" />
            </ElOptionGroup>
          </ElSelect>
        </template>
      </ElTableColumn>
      <ElTableColumn label="注释（CRUD字典）" min-width="180">
        <template #default="{ row }">
          <ElInput v-model="row.comment" size="small" placeholder="如 状态:0=禁用,1=启用" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="主键" width="50" align="center">
        <template #default="{ row }">
          <ElCheckbox v-model="row.isPk" :true-value="1" :false-value="0" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="可空" width="50" align="center">
        <template #default="{ row }">
          <ElCheckbox v-model="row.isNullable" :true-value="1" :false-value="0" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="默认值" width="90">
        <template #default="{ row }">
          <ElInput v-model="row.defaultValue" size="small" placeholder="默认值" />
        </template>
      </ElTableColumn>
      <ElTableColumn label="" width="50" align="center">
        <template #default="{ $index }">
          <ElButton size="small" type="danger" link @click="removeColumn($index)">
            <ArtSvgIcon icon="ri:delete-bin-line" class="text-sm" />
          </ElButton>
        </template>
      </ElTableColumn>
    </ElTable>

    <div class="designer-footer">
      <ElButton type="primary" :loading="loading" @click="handleCreate">
        <ArtSvgIcon icon="ri:database-2-line" class="text-sm mr-1.5" />
        创建数据表
      </ElButton>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { fetchGenCodesCreateTable } from '@/api/backend/develop/genCodes'
  import ArtSvgIcon from '@/components/core/base/art-svg-icon/index.vue'
  import type { FormInstance, FormRules } from 'element-plus'

  const emit = defineEmits<{ (e: 'created', tableName: string): void }>()

  const formRef = ref<FormInstance>()
  const loading = ref(false)

  const form = reactive({
    tableName: '',
    tableComment: '',
    columns: [
      { name: 'id', type: 'bigint unsigned', comment: '主键', isPk: 1, isNullable: 0, defaultValue: '' }
    ] as any[]
  })

  const rules = reactive<FormRules>({
    tableName: [{ required: true, message: '请输入表名', trigger: 'blur' }],
    tableComment: [{ required: true, message: '请输入表注释', trigger: 'blur' }]
  })

  const addColumn = () => {
    form.columns.push({ name: '', type: 'varchar(100)', comment: '', isPk: 0, isNullable: 0, defaultValue: '' })
  }

  const removeColumn = (index: number) => {
    form.columns.splice(index, 1)
  }

  const addCommonFields = () => {
    const commonFields = [
      { name: 'status', type: 'tinyint(1)', comment: '状态:0=禁用,1=启用', isPk: 0, isNullable: 0, defaultValue: '1' },
      { name: 'sort', type: 'int unsigned', comment: '排序', isPk: 0, isNullable: 0, defaultValue: '0' },
      { name: 'created_at', type: 'bigint unsigned', comment: '创建时间', isPk: 0, isNullable: 1, defaultValue: '' },
      { name: 'updated_at', type: 'bigint unsigned', comment: '更新时间', isPk: 0, isNullable: 1, defaultValue: '' }
    ]
    const existNames = new Set(form.columns.map((c: any) => c.name))
    for (const f of commonFields) {
      if (!existNames.has(f.name)) form.columns.push(f)
    }
  }

  const handleCreate = async () => {
    if (!formRef.value) return
    await formRef.value.validate()
    if (form.columns.length === 0) {
      ElMessage.warning('请至少添加一个字段')
      return
    }
    loading.value = true
    try {
      const res = await fetchGenCodesCreateTable({
        tableName: form.tableName,
        tableComment: form.tableComment,
        columns: form.columns
      })
      emit('created', res.tableName || `xy_${form.tableName}`)
    } catch (e) {
      console.error('创建失败:', e)
    } finally {
      loading.value = false
    }
  }
</script>

<style scoped>
  @reference '@styles/core/tailwind.css';

  .create-table-designer {
    max-width: 960px;
    margin: 0 auto;
  }

  .designer-header {
    display: flex;
    gap: 0;
  }

  .designer-tips {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 10px;
    padding: 8px 12px;
    background: var(--el-fill-color-lighter);
    border-radius: 6px;
    font-size: 12px;
    color: var(--el-text-color-secondary);
  }

  .designer-tips__item b {
    color: var(--el-color-primary);
    font-weight: 600;
  }

  .designer-tips__sep {
    color: var(--el-border-color);
  }

  .designer-toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
  }

  .designer-toolbar__title {
    display: flex;
    align-items: center;
    font-size: 14px;
    font-weight: 600;
    color: var(--el-text-color-primary);
  }

  .designer-toolbar__actions {
    display: flex;
    gap: 8px;
  }

  .designer-footer {
    display: flex;
    justify-content: flex-end;
    margin-top: 16px;
  }
</style>
