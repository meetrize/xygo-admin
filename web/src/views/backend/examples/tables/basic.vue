<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 基础表格 -->
<template>
  <div class="user-page art-full-height">
    <ElCard class="art-table-card" shadow="never" style="margin-top: 0">
      <!-- 表格 -->
      <ArtTable
        rowKey="id"
        :show-table-header="false"
        :loading="loading"
        :data="data"
        :columns="columns"
        :pagination="pagination"
        @pagination:size-change="handleSizeChange"
        @pagination:current-change="handleCurrentChange"
      >
      </ArtTable>
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import { useTable } from '@/hooks/core/useTable'
  import { fetchGetUserList } from '@/api/backend/system'

  defineOptions({ name: 'UserMixedUsageExample' })

  const { data, columns, loading, pagination, handleSizeChange, handleCurrentChange } = useTable({
    core: {
      apiFn: fetchGetUserList,
      apiParams: {
        current: 1,
        size: 20,
        userName: '',
        userPhone: '',
        userEmail: ''
      },
      columnsFactory: () => [
        {
          prop: 'id',
          label: 'ID'
        },
        {
          prop: 'nickName',
          label: '昵称'
        },
        {
          prop: 'userGender',
          label: '性别',
          sortable: true,
          formatter: (row: any) => row.userGender || '未知'
        },
        {
          prop: 'userPhone',
          label: '手机号'
        },
        {
          prop: 'userEmail',
          label: '邮箱'
        }
      ]
    }
  })
</script>
