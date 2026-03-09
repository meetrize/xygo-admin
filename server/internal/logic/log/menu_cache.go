// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package log

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	"xygo/internal/model/entity"
)

// RouteInfo 路由对应的模块和操作信息
type RouteInfo struct {
	Module string // 所属模块（父级菜单 title）
	Title  string // 操作名称（当前菜单/按钮 title）
}

// menuRouteCache 菜单路由缓存（权限点 -> 路由信息）
var menuRouteCache struct {
	sync.RWMutex
	data    map[string]*RouteInfo // key: "METHOD /path"
	loaded  bool
}

// GetRouteInfoByPerm 根据权限点获取路由信息
// perm 格式: "POST /admin/member/group/list"
func GetRouteInfoByPerm(ctx context.Context, perm string) *RouteInfo {
	menuRouteCache.RLock()
	loaded := menuRouteCache.loaded
	menuRouteCache.RUnlock()

	if !loaded {
		loadMenuRouteCache(ctx)
	}

	menuRouteCache.RLock()
	defer menuRouteCache.RUnlock()

	if info, ok := menuRouteCache.data[perm]; ok {
		return info
	}
	return nil
}

// RefreshMenuRouteCache 刷新菜单路由缓存（菜单变更时调用）
func RefreshMenuRouteCache(ctx context.Context) {
	loadMenuRouteCache(ctx)
}

// loadMenuRouteCache 从数据库加载菜单数据构建缓存
// 使用独立 context 进行数据库查询，避免请求 context 取消导致加载失败
func loadMenuRouteCache(ctx context.Context) {
	menuRouteCache.Lock()
	defer menuRouteCache.Unlock()

	// 使用独立 context，防止请求 context 取消导致缓存加载失败
	bgCtx := context.Background()

	// 查询所有启用的菜单
	var menus []*entity.AdminMenu
	err := dao.AdminMenu.Ctx(bgCtx).
		Where(dao.AdminMenu.Columns().Status, 1).
		OrderAsc(dao.AdminMenu.Columns().Sort).
		Scan(&menus)
	if err != nil {
		g.Log().Errorf(bgCtx, "加载菜单路由缓存失败: %v", err)
		return
	}

	// 构建 id -> menu 的索引
	menuMap := make(map[uint64]*entity.AdminMenu, len(menus))
	for _, m := range menus {
		menuMap[m.Id] = m
	}

	// 构建权限点 -> RouteInfo 映射
	data := make(map[string]*RouteInfo)
	for _, m := range menus {
		if m.Perms == "" {
			continue
		}

		// 解析 perms JSON 数组
		var perms []string
		if err := json.Unmarshal([]byte(m.Perms), &perms); err != nil {
			// 尝试单个字符串
			perms = []string{m.Perms}
		}

		// 查找父级菜单作为模块名
		moduleName := findModuleName(m, menuMap)

		for _, perm := range perms {
			if perm == "" {
				continue
			}
			data[perm] = &RouteInfo{
				Module: moduleName,
				Title:  m.Title,
			}
		}
	}

	menuRouteCache.data = data
	menuRouteCache.loaded = true

	g.Log().Infof(bgCtx, "菜单路由缓存已加载，共 %d 条权限点映射", len(data))
}

// findModuleName 向上查找所属模块名称
// 逻辑：如果当前是按钮(type=3)，取其父菜单(type=2)的 title；
//       如果当前是菜单(type=2)，取其父目录(type=1)的 title；
//       如果找不到父级，用自身 title
func findModuleName(menu *entity.AdminMenu, menuMap map[uint64]*entity.AdminMenu) string {
	current := menu

	// 最多向上查3层，防止死循环
	for i := 0; i < 3; i++ {
		if current.ParentId == 0 {
			return current.Title
		}
		parent, ok := menuMap[current.ParentId]
		if !ok {
			return current.Title
		}
		// 如果父级是目录(type=1)或菜单(type=2)，用父级的 title 作为模块名
		if parent.Type == 1 || parent.Type == 2 {
			return parent.Title
		}
		current = parent
	}

	return menu.Title
}
