<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 会员分组管理页面 -->
<template>
  <div class="art-full-height">
    <GroupSearch
      v-show="showSearchBar"
      v-model="searchForm"
      @search="handleSearch"
      @reset="resetSearchParams"
    ></GroupSearch>

    <ElCard
      class="art-table-card"
      shadow="never"
      :style="{ 'margin-top': showSearchBar ? '12px' : '0' }"
    >
      <ArtTableHeader
        v-model:columns="columnChecks"
        v-model:showSearchBar="showSearchBar"
        :loading="loading"
        @refresh="refreshData"
      >
        <template #left>
          <ElSpace wrap>
            <ElButton @click="showDialog('add')" v-ripple>新增分组</ElButton>
          </ElSpace>
        </template>
      </ArtTableHeader>

      <!-- 表格 -->
      <ArtTable
        ref="tableRef"
        rowKey="id"
        :loading="loading"
        :data="data"
        :columns="columns"
        :pagination="pagination"
        @pagination:size-change="handleSizeChange"
        @pagination:current-change="handleCurrentChange"
      >
      </ArtTable>
    </ElCard>

    <!-- 分组编辑弹窗 -->
    <GroupEditDialog
      v-model="dialogVisible"
      :dialog-type="dialogType"
      :group-data="currentGroupData"
      @success="refreshData"
    />

    <!-- 菜单权限弹窗 -->
    <GroupPermissionDialog
      v-model="permissionDialog"
      :group-data="currentGroupData"
      @success="refreshData"
    />
  </div>
</template>

<script setup lang="ts">
  import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
  import { useTable } from '@/hooks/core/useTable'
  import { formatTimestamp } from '@/utils/time'
  import {
    getMemberGroupList,
    saveMemberGroup,
    deleteMemberGroup,
    type MemberGroupItem
  } from '@/api/backend/member/group'
  import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
  import GroupSearch from './modules/group-search.vue'
  import GroupEditDialog from './modules/group-edit-dialog.vue'
  import GroupPermissionDialog from './modules/group-permission-dialog.vue'
  import { ElTag, ElMessageBox } from 'element-plus'

  defineOptions({ name: 'MemberGroup' })

  // 搜索表单
  const searchForm = ref({
    name: undefined,
    status: undefined
  })

  const showSearchBar = ref(false)
  const tableRef = ref()

  const dialogVisible = ref(false)
  const permissionDialog = ref(false)
  const currentGroupData = ref<MemberGroupItem | undefined>(undefined)

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
    // 核心配置
    core: {
      apiFn: getMemberGroupList,
      apiParams: {
        page: 1,
        pageSize: 20
      },
      paginationKey: {
        current: 'page',
        size: 'pageSize'
      },
      columnsFactory: () => [
        {
          prop: 'id',
          label: '分组ID',
          width: 100
        },
        {
          prop: 'name',
          label: '分组名称',
          minWidth: 150
        },
        {
          prop: 'sort',
          label: '排序',
          width: 80,
          align: 'center'
        },
        {
          prop: 'status',
          label: '状态',
          width: 100,
          align: 'center',
          formatter: (row: MemberGroupItem) => {
            return h(ElTag, {
              type: row.status === 1 ? 'success' : 'danger'
            }, () => row.status === 1 ? '启用' : '禁用')
          }
        },
        {
          prop: 'remark',
          label: '备注',
          minWidth: 150,
          showOverflowTooltip: true,
          formatter: (row: MemberGroupItem) => row.remark || '-'
        },
        {
          prop: 'createdAt',
          label: '创建时间',
          width: 180,
          sortable: true,
          formatter: (row: MemberGroupItem) => formatTimestamp(row.createdAt)
        },
        {
          prop: 'operation',
          label: '操作',
          width: 80,
          fixed: 'right',
          formatter: (row: MemberGroupItem) => {
            const menuList: any[] = []
            
            // 菜单权限按钮
            menuList.push({
              key: 'permission',
              label: '菜单权限',
              icon: 'ri:menu-line'
            })
            
            // 编辑按钮
            menuList.push({
              key: 'edit',
              label: '编辑分组',
              icon: 'ri:edit-2-line'
            })
            
            // 删除按钮
            menuList.push({
              key: 'delete',
              label: '删除分组',
              icon: 'ri:delete-bin-4-line',
              color: '#f56c6c'
            })
            
            return h('div', [
              h(ArtButtonMore, {
                list: menuList,
                onClick: (item: ButtonMoreItem) => buttonMoreClick(item, row)
              })
            ])
          }
        }
      ]
    }
  })

  const dialogType = ref<'add' | 'edit'>('add')

  const showDialog = (type: 'add' | 'edit', row?: MemberGroupItem) => {
    dialogVisible.value = true
    dialogType.value = type
    currentGroupData.value = row
  }

  /**
   * 搜索处理
   * @param params 搜索参数
   */
  const handleSearch = (params: Record<string, any>) => {
    // 搜索参数赋值
    Object.assign(searchParams, params)
    getData()
  }

  const buttonMoreClick = (item: ButtonMoreItem, row: MemberGroupItem) => {
    switch (item.key) {
      case 'permission':
        showPermissionDialog(row)
        break
      case 'edit':
        showDialog('edit', row)
        break
      case 'delete':
        deleteGroup(row)
        break
    }
  }

  const showPermissionDialog = (row?: MemberGroupItem) => {
    permissionDialog.value = true
    currentGroupData.value = row
  }

  const deleteGroup = (row: MemberGroupItem) => {
    ElMessageBox.confirm(`确定删除分组"${row.name}"吗？此操作不可恢复！`, '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async () => {
        await deleteMemberGroup(row.id)
        ElMessage.success('删除成功')
        refreshData()
      })
      .catch(() => {
        ElMessage.info('已取消删除')
      })
  }
</script>
