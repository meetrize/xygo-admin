<!-- 悠然示例管理 -->
<template>
  <div class="uran-example-page art-full-height">
    <!-- 搜索栏 -->
    <UranExampleSearch v-model="searchForm" @search="handleSearch" @reset="resetSearchParams" />

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
      <UranExampleDialog
        v-model:visible="dialogVisible"
        :type="dialogType"
        :edit-data="currentRow"
        @submit="handleDialogSubmit"
      />
      <!-- 详情抽屉 -->
      <UranExampleDetailDrawer v-model="detailVisible" :view-id="detailId" />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
  import { useTable } from '@/hooks/core/useTable'
  import { useAuth } from '@/hooks/core/useAuth'
  import { formatTimestamp } from '@/utils/time'
  import {
    fetchUranExampleList,
    fetchUranExampleEdit,
    fetchUranExampleDelete
  } from '@/api/backend/uran-example'
  import UranExampleSearch from './modules/uran-example-search.vue'
  import UranExampleDialog from './modules/uran-example-dialog.vue'
  import UranExampleDetailDrawer from './modules/uran-example-detail-drawer.vue'
  import { ElTag, ElImage, ElMessageBox } from 'element-plus'
  import { DialogType } from '@/types'

  defineOptions({ name: 'UranExample' })
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
      apiFn: fetchUranExampleList,
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
          prop: 'tname',
          label: '标题',
          minWidth: 100,
          formatter: (row: any) => row.tname ?? '-'
        },
        {
          prop: 'timage',
          label: '图片',
          width: 80,
          align: 'center',
          formatter: (row: any) =>
            row.timage
              ? h(ElImage, {
                  src: row.timage,
                  style: 'width:40px;height:40px',
                  fit: 'cover',
                  previewSrcList: [row.timage],
                  previewTeleported: true
                })
              : '-'
        },
        {
          prop: 'ttest',
          label: '测试',
          minWidth: 100,
          formatter: (row: any) => row.ttest ?? '-'
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
      await fetchUranExampleDelete(row.id)
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
        await fetchUranExampleDelete(row.id)
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
      await fetchUranExampleEdit(formData)
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
