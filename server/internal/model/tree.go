// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package model

// 通用树形工具，参考 BuildAdmin/HotGo 的树封装，方便在角色/菜单/部门等场景复用。
//
// 设计思路：
// - 不强依赖具体结构体类型，通过泛型 + 回调函数拿到 id/pid，并设置 children。
// - 这样任何包含 id/pid/children 字段的结构体都可以直接复用本工具。

// BuildTree 将一维切片按 id/pid 关系组装成树形结构，返回所有根节点。
//
// 参数说明：
// - list        ：原始节点列表（通常为 []*T，T 为业务自定义结构体）；
// - getId       ：获取节点 ID 的函数；
// - getPid      ：获取父节点 ID 的函数；
// - setChildren ：为节点设置 children 切片的函数。
//
// 使用示例（以 RoleTree 为例）：
//   roots := model.BuildTree(
//       roles,
//       func(n *RoleTree) uint { return n.Id },
//       func(n *RoleTree) uint { return n.Pid },
//       func(n *RoleTree, children []*RoleTree) { n.Children = children },
//   )
func BuildTree[T any](
	list []*T,
	getId func(*T) uint,
	getPid func(*T) uint,
	setChildren func(*T, []*T),
) []*T {
	if len(list) == 0 {
		return nil
	}

	// 记录所有节点，方便判断某个 pid 是否有对应父节点
	nodeById := make(map[uint]*T, len(list))
	// pid -> children 列表
	childrenMap := make(map[uint][]*T)

	for _, item := range list {
		if item == nil {
			continue
		}
		id := getId(item)
		pid := getPid(item)
		nodeById[id] = item
		childrenMap[pid] = append(childrenMap[pid], item)
	}

	// 找出所有根节点：pid 为 0 或者在 nodeById 中不存在（视为顶层）
	var roots []*T
	for _, item := range list {
		if item == nil {
			continue
		}
		pid := getPid(item)
		if pid == 0 {
			roots = append(roots, item)
			continue
		}
		if _, ok := nodeById[pid]; !ok {
			roots = append(roots, item)
		}
	}

	// 递归挂载 children
	var attachChildren func(nodes []*T)
	attachChildren = func(nodes []*T) {
		for _, n := range nodes {
			if n == nil {
				continue
			}
			id := getId(n)
			cs := childrenMap[id]
			if len(cs) > 0 {
				setChildren(n, cs)
				attachChildren(cs)
			}
		}
	}

	attachChildren(roots)
	return roots
}

// FlattenTree 将树形节点拍平成一维切片（先序遍历），便于树表展示或导出。
//
// 参数说明：
// - roots       ：根节点列表；
// - getChildren ：获取节点 children 的函数。
func FlattenTree[T any](roots []*T, getChildren func(*T) []*T) []*T {
	if len(roots) == 0 {
		return nil
	}

	var result []*T
	var walk func(nodes []*T)

	walk = func(nodes []*T) {
		for _, n := range nodes {
			if n == nil {
				continue
			}
			result = append(result, n)
			if cs := getChildren(n); len(cs) > 0 {
				walk(cs)
			}
		}
	}

	walk(roots)
	return result
}
