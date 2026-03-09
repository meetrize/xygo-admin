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
            <span v-if="data.isAuth">
              {{ data.label }}
            </span>
            <span v-else>{{ defaultProps.label(data) }}</span>
          </div>
        </template>
      </ElTree>
    </ElScrollbar>
    <template #footer>
      <ElButton @click="outputSelectedData" style="margin-left: 8px">获取选中数据</ElButton>

      <ElButton @click="toggleExpandAll">{{ isExpandAll ? '全部收起' : '全部展开' }}</ElButton>
      <ElButton @click="toggleSelectAll" style="margin-left: 8px">{{
        isSelectAll ? '取消全选' : '全部选择'
      }}</ElButton>
      <ElButton type="primary" @click="savePermission">保存</ElButton>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import { formatMenuTitle } from '@/utils/router'
  import { fetchRoleMenuIds, fetchRoleBindMenus, fetchGetMenuTree } from '@/api/backend/system'

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

  const menuList = ref<any[]>([])  // ✅ 改用本地state存储菜单树
  const treeRef = ref()
  const isExpandAll = ref(true)
  const isSelectAll = ref(false)
  const loading = ref(false)

  /**
   * 弹窗显示状态双向绑定
   */
  const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
  })

  /**
   * 菜单节点类型
   */
  interface MenuNode {
    id?: string | number
    name?: string
    label?: string
    meta?: {
      title?: string
      authList?: Array<{
        authMark: string
        title: string
        checked?: boolean
      }>
    }
    children?: MenuNode[]
    [key: string]: any
  }

  /**
   * 处理后端菜单树数据
   * 后端字段：id, title, name, type, children...
   * 树组件需要：id, label（显示文本）, children
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

  /**
   * 树形组件配置
   */
  const defaultProps = {
    children: 'children',
    label: (data: any) => data.label || data.title || data.name || ''
  }

  /**
   * 监听弹窗打开，加载菜单树和角色已绑定的菜单权限
   */
  watch(
    () => props.modelValue,
    async (newVal) => {
      if (newVal && props.roleData) {
        loading.value = true
        try {
          // ✅ 1. 加载完整的菜单树（从后端获取）
          const menuTree = await fetchGetMenuTree()
          menuList.value = menuTree
          console.log('[菜单权限] 加载菜单树:', menuTree)
          
          // ✅ 2. 加载角色已绑定的菜单ID
          const res = await fetchRoleMenuIds(props.roleData.id)
          const menuIds = res.menuIds || []
          console.log('[菜单权限] 已绑定的菜单ID:', menuIds)
          
          // ✅ 3. 设置选中的菜单
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
   * 关闭弹窗并清空选中状态
   */
  const handleClose = () => {
    visible.value = false
    treeRef.value?.setCheckedKeys([])
  }

  /**
   * 保存权限配置
   */
  const savePermission = async () => {
    if (!props.roleData) {
      ElMessage.error('角色数据不存在')
      return
    }

    try {
      const tree = treeRef.value
      if (!tree) return

      // 获取选中的节点key（只要完全选中的，不要半选中的）
      const checkedKeys = tree.getCheckedKeys()
      
      console.log('[保存菜单权限] 选中的key:', checkedKeys)
      console.log('[保存菜单权限] 角色ID:', props.roleData.id)
      
      // 只保存完全选中的菜单ID（过滤掉字符串类型的按钮权限ID）
      const menuIds = checkedKeys.filter((key: any) => typeof key === 'number')
      
      console.log('[保存菜单权限] 提交的菜单ID:', menuIds)
      
      await fetchRoleBindMenus({
        roleId: props.roleData.id,
        menuIds: menuIds as number[]
      })
      
      ElMessage.success('菜单权限保存成功')
      emit('success')
      handleClose()
    } catch (error) {
      console.error('保存菜单权限失败:', error)
    }
  }

  /**
   * 切换全部展开/收起状态
   */
  const toggleExpandAll = () => {
    const tree = treeRef.value
    if (!tree) return

    const nodes = tree.store.nodesMap
    // 这里保留 any，因为 Element Plus 的内部节点类型较复杂
    Object.values(nodes).forEach((node: any) => {
      node.expanded = !isExpandAll.value
    })

    isExpandAll.value = !isExpandAll.value
  }

  /**
   * 切换全选/取消全选状态
   */
  const toggleSelectAll = () => {
    const tree = treeRef.value
    if (!tree) return

    if (!isSelectAll.value) {
      const allKeys = getAllNodeKeys(processedMenuList.value)
      tree.setCheckedKeys(allKeys)
    } else {
      tree.setCheckedKeys([])
    }

    isSelectAll.value = !isSelectAll.value
  }

  /**
   * 递归获取所有节点的 key（菜单ID）
   * @param nodes 节点列表
   * @returns 所有节点的 key 数组
   */
  const getAllNodeKeys = (nodes: MenuNode[]): (string | number)[] => {
    const keys: (string | number)[] = []
    const traverse = (nodeList: MenuNode[]): void => {
      nodeList.forEach((node) => {
        if (node.id) keys.push(node.id)
        if (node.children?.length) traverse(node.children)
      })
    }
    traverse(nodes)
    return keys
  }

  /**
   * 处理树节点选中状态变化
   * 同步更新全选按钮状态
   */
  const handleTreeCheck = () => {
    const tree = treeRef.value
    if (!tree) return

    const checkedKeys = tree.getCheckedKeys()
    const allKeys = getAllNodeKeys(processedMenuList.value)

    isSelectAll.value = checkedKeys.length === allKeys.length && allKeys.length > 0
  }

  /**
   * 输出选中的权限数据到控制台
   * 用于调试和查看当前选中的权限配置
   */
  const outputSelectedData = () => {
    const tree = treeRef.value
    if (!tree) return

    const selectedData = {
      checkedKeys: tree.getCheckedKeys(),
      halfCheckedKeys: tree.getHalfCheckedKeys(),
      checkedNodes: tree.getCheckedNodes(),
      halfCheckedNodes: tree.getHalfCheckedNodes(),
      totalChecked: tree.getCheckedKeys().length,
      totalHalfChecked: tree.getHalfCheckedKeys().length
    }

    console.log('=== 选中的权限数据 ===', selectedData)
    ElMessage.success(`已输出选中数据到控制台，共选中 ${selectedData.totalChecked} 个节点`)
  }
</script>
