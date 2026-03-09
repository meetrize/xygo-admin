// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package system

import "github.com/gogf/gf/v2/frame/g"

// HealthReq 健康检查请求
type HealthReq struct {
	g.Meta `path:"/health" method:"get" tags:"System" summary:"Health check"`
}

// HealthRes 健康检查响应
type HealthRes struct {
	Status string `json:"status" example:"ok"`
}
