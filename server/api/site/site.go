// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package site

import "github.com/gogf/gf/v2/frame/g"

// IndexReq 站点基础信息（公共接口）
type IndexReq struct {
	g.Meta `path:"/site/index" method:"get" tags:"Site" summary:"站点基础信息"`
}

// IndexRes 站点基础信息返回
type IndexRes struct {
	Group            string            `json:"group"`            // 配置分组
	Items            map[string]string `json:"items"`            // 原始键值对
	SiteName         string            `json:"siteName"`         // 站点名称
	SiteSubtitle     string            `json:"siteSubtitle"`     // 副标题
	Icp              string            `json:"icp"`              // ICP 备案
	Timezone         string            `json:"timezone"`         // 时区
	Description      string            `json:"description"`      // 描述
	ThemeColor       string            `json:"themeColor"`       // 主题色
	Logo             string            `json:"logo"`             // Logo 地址
	Closed           string            `json:"closed"`           // 关闭状态
	OpenMemberCenter bool              `json:"openMemberCenter"` // 是否开启会员中心
}
