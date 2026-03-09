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

// ===================== 资源列表 =====================

// ResourceListModel 资源列表响应模型
type ResourceListModel struct {
	List []ResourceItem `json:"list" dc:"资源列表"`
}

// ResourceItem 资源项
type ResourceItem struct {
	Code  string `json:"code" dc:"资源编码（表名）"`
	Label string `json:"label" dc:"资源中文名"`
}
