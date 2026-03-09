// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package adminin

// MenuTreeItem 菜单树节点
// 对应表：xy_admin_menu，适配前端菜单/路由配置所需字段。
// 仅用于对外返回的只读结构，新增/编辑的入参在后续的 MenuSaveInp 中单独定义。
type MenuTreeItem struct {
	Id         uint64          `json:"id"         dc:"菜单ID"`
	ParentId   uint64          `json:"parentId"   dc:"上级菜单ID"`
	Type       int             `json:"type"       dc:"类型:1=目录,2=菜单,3=按钮"`
	Title      string          `json:"title"      dc:"菜单名称"`
	Name       string          `json:"name"       dc:"前端路由name"`
	Path       string          `json:"path"       dc:"路由路径"`
	Component  string          `json:"component"  dc:"前端组件路径"`
	Resource   string          `json:"resource"   dc:"关联数据表名"`
	Icon       string          `json:"icon"       dc:"图标"`
	Hidden     int             `json:"hidden"     dc:"是否隐藏:0=否,1=是"`
	KeepAlive  int             `json:"keepAlive"  dc:"是否缓存:0=否,1=是"`
	Redirect   string          `json:"redirect"   dc:"重定向地址"`
	FrameSrc   string          `json:"frameSrc"   dc:"内嵌iframe地址"`
	Perms      string          `json:"perms"      dc:"权限点列表(JSON数组,内容为 METHOD+PATH 字符串)"`
	IsFrame    int             `json:"isFrame"    dc:"是否内嵌:0=否,1=是"`
	Affix      int             `json:"affix"      dc:"是否固定标签:0=否,1=是"`
	ShowBadge  int             `json:"showBadge"  dc:"是否显示徽章:0=否,1=是"`
	BadgeText  string          `json:"badgeText"  dc:"徽章文本"`
	ActivePath string          `json:"activePath" dc:"激活高亮路径"`
	HideTab    int             `json:"hideTab"    dc:"是否隐藏标签:0=否,1=是"`
	IsFullPage int             `json:"isFullPage" dc:"是否全屏页面:0=否,1=是"`
	Sort       int             `json:"sort"       dc:"排序(越大越靠后)"`
	Status     int             `json:"status"     dc:"状态:0=禁用,1=启用"`
	Remark     string          `json:"remark"     dc:"备注"`
	CreateTime int             `json:"create_time"  dc:"创建时间"`
	UpdateTime int             `json:"update_time"  dc:"更新时间"`
	Children   []*MenuTreeItem `json:"children,omitempty" dc:"子菜单"`
}

// MenuTreeModel 菜单树返回结构
type MenuTreeModel struct {
	List []*MenuTreeItem `json:"list" dc:"菜单树列表"`
}

// ===================== 菜单路由 =====================

// MenuRoutesInp 菜单路由入参
type MenuRoutesInp struct {
	RoleId uint `p:"roleId" json:"roleId" dc:"可选：按角色过滤路由；不填返回全量启用目录/菜单"`
}

// ===================== 菜单详情 =====================

// MenuDetailInp 菜单详情入参
type MenuDetailInp struct {
	Id uint64 `p:"id" v:"required#菜单ID不能为空" json:"id" dc:"菜单ID"`
}

// ===================== 菜单保存 =====================

// MenuSaveInp 菜单保存入参（新增/编辑）
type MenuSaveInp struct {
	Id         uint64 `p:"id" json:"id" dc:"菜单ID（0=新增）"`
	ParentId   uint64 `p:"parentId" json:"parentId" dc:"上级菜单ID"`
	Type       int    `p:"type" v:"required|in:1,2,3#类型不能为空|类型值错误" json:"type" dc:"类型:1=目录,2=菜单,3=按钮"`
	Title      string `p:"title" v:"required#菜单名称不能为空" json:"title" dc:"菜单名称"`
	Name       string `p:"name" json:"name" dc:"前端路由name"`
	Path       string `p:"path" json:"path" dc:"路由路径"`
	Component  string `p:"component" json:"component" dc:"前端组件路径"`
	Resource   string `p:"resource" json:"resource" dc:"关联数据表名"`
	Icon       string `p:"icon" json:"icon" dc:"图标"`
	Hidden     int    `p:"hidden" d:"0" json:"hidden" dc:"是否隐藏:0=否,1=是"`
	KeepAlive  int    `p:"keepAlive" d:"0" json:"keepAlive" dc:"是否缓存:0=否,1=是"`
	Redirect   string `p:"redirect" json:"redirect" dc:"重定向地址"`
	FrameSrc   string `p:"frameSrc" json:"frameSrc" dc:"内嵌iframe地址"`
	Perms      string `p:"perms" json:"perms" dc:"权限点列表(JSON)"`
	IsFrame    int    `p:"isFrame" d:"0" json:"isFrame" dc:"是否内嵌:0=否,1=是"`
	Affix      int    `p:"affix" d:"0" json:"affix" dc:"是否固定标签:0=否,1=是"`
	ShowBadge  int    `p:"showBadge" d:"0" json:"showBadge" dc:"是否显示徽章:0=否,1=是"`
	BadgeText  string `p:"badgeText" json:"badgeText" dc:"徽章文本"`
	ActivePath string `p:"activePath" json:"activePath" dc:"激活高亮路径"`
	HideTab    int    `p:"hideTab" d:"0" json:"hideTab" dc:"是否隐藏标签:0=否,1=是"`
	IsFullPage int    `p:"isFullPage" d:"0" json:"isFullPage" dc:"是否全屏页面:0=否,1=是"`
	Sort       int    `p:"sort" d:"0" json:"sort" dc:"排序(越大越靠后)"`
	Status     int    `p:"status" d:"1" json:"status" dc:"状态:0=禁用,1=启用"`
	Remark     string `p:"remark" json:"remark" dc:"备注"`
}

// ===================== 菜单删除 =====================

// MenuDeleteInp 菜单删除入参
type MenuDeleteInp struct {
	Id uint64 `p:"id" v:"required#菜单ID不能为空" json:"id" dc:"菜单ID"`
}
