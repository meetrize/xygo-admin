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
    title="菜单权限"
    width="520px"
    align-center
    class="el-dialog-border"
    @close="handleClose"
  >
    <ElScrollbar height="70vh">
      <ElTree
        ref="treeRef"
        :data="processedMenuList"
        show-checkbox
        node-key="id"
        :default-expand-all="isExpandAll"
        :props="defaultProps"
        @check="handleTreeCheck"
      >
        <template #default="{ data }">
          <div style="display: flex; align-items: center">
            <span>{{ data.label }}</span>
          </div>
        </template>
      </ElTree>
    </ElScrollbar>
    <template #footer>
      <ElButton @click="toggleExpandAll">{{ isExpandAll ? '全部收起' : '全部展开' }}</ElButton>
      <ElButton @click="toggleSelectAll">{{
        isSelectAll ? '取消全选' : '全部选择'
      }}</ElButton>
      <ElButton type="primary" @click="savePermission">保存</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import { getMemberMenuTree } from '@/api/backend/member/menu'
  import type { MemberGroupItem } from '@/api/backend/member/group'

  interface Props {
    modelValue: boolean
    groupData?: MemberGroupItem
  }

  interface Emits {
    (e: 'update:modelValue', value: boolean): void
    (e: 'success'): void
  }

  const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
    groupData: undefined
  })

  const emit = defineEmits<Emits>()

  const menuList = ref<any[]>([])
  const treeRef = ref()
  const isExpandAll = ref(true)
  const isSelectAll = ref(false)
  const loading = ref(false)

  const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
  })

  /**
   * 处理菜单树数据
   */
  const processedMenuList = computed(() => {
    const processNode = (node: any): any => {
      return {
        id: node.id,
        title: node.title,
        name: node.name,
        label: node.title || node.name || '',
        type: node.type,
        children: node.children ? node.children.map(processNode) : undefined
      }
    }

    return (menuList.value || []).map(processNode)
  })

  const defaultProps = {
    children: 'children',
    label: (data: any) => data.label || data.title || data.name || ''
  }

  /**
   * 监听弹窗打开
   */
  watch(
    () => props.modelValue,
    async (newVal) => {
      if (newVal && props.groupData) {
        loading.value = true
        try {
          // 1. 加载会员菜单树
          const res = await getMemberMenuTree()
          menuList.value = res.list || []
          console.log('[会员分组菜单权限] 加载菜单树:', menuList.value)
          
          // 2. 解析已绑定的菜单ID (从rules字段)
          const menuIds = parseMenuIds(props.groupData.rules)
          console.log('[会员分组菜单权限] 已绑定的菜单ID:', menuIds)
          
          // 3. 设置选中的菜单
          nextTick(() => {
            treeRef.value?.setCheckedKeys(menuIds)
          })
        } catch (error) {
          console.error('加载菜单权限失败:', error)
          ElMessage.error('加载菜单权限失败')
        } finally {
          loading.value = false
        }
      }
    }
  )

  /**
   * 解析rules字段中的菜单ID（逗号分隔）
   */
  const parseMenuIds = (rules: string): number[] => {
    if (!rules || !rules.trim()) return []
    return rules.split(',').map(id => parseInt(id.trim())).filter(id => !isNaN(id))
  }

  /**
   * 树节点选中变化
   */
  const handleTreeCheck = () => {
    // 可选：添加逻辑
  }

  /**
   * 切换全部展开/收起
   */
  const toggleExpandAll = () => {
    isExpandAll.value = !isExpandAll.value
  }

  /**
   * 切换全选/取消全选
   */
  const toggleSelectAll = () => {
    if (isSelectAll.value) {
      treeRef.value?.setCheckedKeys([])
    } else {
      const allKeys = getAllNodeKeys(processedMenuList.value)
      treeRef.value?.setCheckedKeys(allKeys)
    }
    isSelectAll.value = !isSelectAll.value
  }

  /**
   * 获取所有节点ID
   */
  const getAllNodeKeys = (nodes: any[]): number[] => {
    const keys: number[] = []
    const traverse = (node: any) => {
      keys.push(node.id)
      if (node.children?.length) {
        node.children.forEach(traverse)
      }
    }
    nodes.forEach(traverse)
    return keys
  }

  /**
   * 保存菜单权限
   */
  const savePermission = async () => {
    if (!props.groupData) return

    try {
      // 获取选中的菜单ID（包含半选状态）
      const checkedKeys = treeRef.value?.getCheckedKeys() || []
      const halfCheckedKeys = treeRef.value?.getHalfCheckedKeys() || []
      const allKeys = [...checkedKeys, ...halfCheckedKeys]
      
      // 转换为逗号分隔的字符串
      const rulesStr = allKeys.join(',')
      
      // TODO: 调用保存接口
      console.log('保存会员分组菜单权限:', {
        groupId: props.groupData.id,
        rules: rulesStr,
        menuIds: allKeys
      })

      // 临时使用saveMemberGroup更新rules字段
      const { saveMemberGroup } = await import('@/api/backend/member/group')
      await saveMemberGroup({
        id: props.groupData.id,
        name: props.groupData.name,
        rules: rulesStr,
        sort: props.groupData.sort,
        status: props.groupData.status,
        remark: props.groupData.remark
      })
      
      ElMessage.success('保存成功')
      emit('success')
      visible.value = false
    } catch (error) {
      console.error('保存菜单权限失败:', error)
      ElMessage.error('保存失败')
    }
  }

  /**
   * 关闭弹窗
   */
  const handleClose = () => {
    emit('update:modelValue', false)
  }
</script>
