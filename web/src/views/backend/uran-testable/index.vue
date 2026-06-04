<!-- 测试记帐管理 -->
<template>
  <div class="uran-testable-page art-full-height">
    <!-- 搜索栏 -->
    <UranTestableSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams" />

    <ElCard class="art-table-card" shadow="never">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton v-auth="'add'" @click="showDialog('add')" v-ripple>新增</ElButton>
            <ElButton
              v-auth="'batchDel'"
              type="danger"
              :disabled="selectedRows.length === 0"
              @click="handleBatchDelete"
              v-ripple
              >批量删除</ElButton
            >
            <ElButton v-auth="'export'" @click="handleExport" v-ripple>导出</ElButton>
          </ElSpace>
        </template>
      </ArtTableHeader>

      <!-- 表格 -->
      <ArtTable
        :loading="loading"
        :data="data"
        :columns="columns"
        :pagination="pagination"
        @selection-change="handleSelectionChange"
        @pagination:size-change="handleSizeChange"
        @pagination:current-change="handleCurrentChange"
      />
      <!-- 编辑弹窗 -->
      <UranTestableDialog
        v-model:visible="dialogVisible"
        :type="dialogType"
        :edit-data="currentRow"
        @submit="handleDialogSubmit"
      />
      <!-- 详情抽屉 -->
      <UranTestableDetailDrawer v-model="detailVisible" :view-id="detailId" />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
  import { useTable } from '@/hooks/core/useTable'
  import { useAuth } from '@/hooks/core/useAuth'
  import { formatTimestamp } from '@/utils/time'
  import {
    fetchUranTestableList,
    fetchUranTestableEdit,
    fetchUranTestableDelete
  } from '@/api/backend/uran-testable'
  import UranTestableSearch from './modules/uran-testable-search.vue'
  import UranTestableDialog from './modules/uran-testable-dialog.vue'
  import UranTestableDetailDrawer from './modules/uran-testable-detail-drawer.vue'
  import { ElTag, ElMessageBox } from 'element-plus'
  import { DialogType } from '@/types'

  defineOptions({ name: 'UranTestable' })
  const { hasAuth } = useAuth()
  const dialogType = ref<DialogType>('add')
  const dialogVisible = ref(false)
  const currentRow = ref<any>({})
  const detailVisible = ref(false)
  const detailId = ref<number>()
  const selectedRows = ref<any[]>([])

  const searchForm = ref({
    status: undefined
  })

  const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    getData,
    searchParams,
    resetSearchParams,
    handleSizeChange,
    handleCurrentChange,
    refreshData
  } = useTable({
    core: {
      apiFn: fetchUranTestableList,
      apiParams: {
        page: 1,
        pageSize: 20,
        ...searchForm.value
      },
      paginationKey: { current: 'page', size: 'pageSize' },
      columnsFactory: () => [
        { type: 'selection' },
        {
          prop: 'id',
          label: '主键',
          minWidth: 100,
          formatter: (row: any) => row.id ?? '-'
        },
        {
          prop: 'status',
          label: '状态',
          width: 100,
          align: 'center',
          formatter: (row: any) => {
            const map: Record<string, [string, string]> = {
              '0': ['禁用', 'success'],
              '1': ['启用', 'danger']
            }
            const m = map[String(row.status)]
            return m
              ? h(ElTag, { type: m[1] as any, size: 'small' }, () => m[0])
              : h(ElTag, { size: 'small' }, () => String(row.status ?? '-'))
          }
        },
        {
          prop: 'sort',
          label: '排序',
          minWidth: 100,
          formatter: (row: any) => row.sort ?? '-'
        },
        {
          prop: 'createdAt',
          label: '创建时间',
          width: 180,
          formatter: (row: any) => formatTimestamp(row.createdAt)
        },
        {
          prop: 'updatedAt',
          label: '更新时间',
          width: 180,
          formatter: (row: any) => formatTimestamp(row.updatedAt)
        },
        {
          prop: 'stuff',
          label: '物件名称',
          minWidth: 120,
          formatter: (row: any) => row.stuff ?? '-'
        },
        {
          prop: 'price',
          label: '价格',
          minWidth: 100,
          formatter: (row: any) => row.price ?? '-'
        },
        {
          prop: 'operation',
          label: '操作',
          width: 220,
          fixed: 'right',
          formatter: (row: any) =>
            h(
              'div',
              { class: 'flex items-center gap-1' },
              [
                hasAuth('view')
                  ? h(ArtButtonTable, { type: 'view', onClick: () => handleView(row) })
                  : null,
                hasAuth('edit')
                  ? h(ArtButtonTable, { type: 'edit', onClick: () => showDialog('edit', row) })
                  : null,
                hasAuth('delete')
                  ? h(ArtButtonTable, { type: 'delete', onClick: () => handleDelete(row) })
                  : null
              ].filter(Boolean)
            )
        }
      ]
    }
  })

  const handleSearch = (params: Record<string, any>) => {
    // 先清空旧搜索值（保留分页参数），再写入新值
    const paramsRecord = searchParams as Record<string, unknown>
    Object.keys(paramsRecord).forEach((key) => {
      if (key !== 'page' && key !== 'pageSize') {
        delete paramsRecord[key]
      }
    })
    // 过滤掉空值，避免后端收到空字符串参数
    for (const [k, v] of Object.entries(params)) {
      if (v !== undefined && v !== null && v !== '') {
        paramsRecord[k] = v
      }
    }
    paramsRecord['page'] = 1 // 搜索时回到第一页
    getData()
  }

  const showDialog = (type: DialogType, row?: any) => {
    dialogType.value = type
    currentRow.value = row || {}
    nextTick(() => {
      dialogVisible.value = true
    })
  }

  const handleView = (row: any) => {
    detailId.value = row.id
    detailVisible.value = true
  }

  const handleDelete = async (row: any) => {
    try {
      await ElMessageBox.confirm('确定要删除该记录吗？删除后无法恢复', '删除确认', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      await fetchUranTestableDelete(row.id)
      ElMessage.success('删除成功')
      refreshData()
    } catch (e) {
      if (e !== 'cancel') console.error(e)
    }
  }

  const handleBatchDelete = async () => {
    if (selectedRows.value.length === 0) return
    try {
      await ElMessageBox.confirm(
        `确定要删除选中的 ${selectedRows.value.length} 条记录吗？`,
        '批量删除',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
      for (const row of selectedRows.value) {
        await fetchUranTestableDelete(row.id)
      }
      ElMessage.success('批量删除成功')
      selectedRows.value = []
      refreshData()
    } catch (e) {
      if (e !== 'cancel') console.error(e)
    }
  }

  const handleExport = () => {
    ElMessage.info('导出功能开发中')
  }

  const handleDialogSubmit = async (formData: any) => {
    try {
      await fetchUranTestableEdit(formData)
      ElMessage.success(formData.id ? '编辑成功' : '添加成功')
      dialogVisible.value = false
      refreshData()
    } catch (e) {
      console.error(e)
    }
  }

  const handleSelectionChange = (selection: any[]) => {
    selectedRows.value = selection
  }
</script>
