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

import (
	"context"

	api "xygo/api/system"
)

// Health 健康检查接口，配合统一响应中间件返回标准结构。
func (c *ControllerV1) Health(ctx context.Context, req *api.HealthReq) (res *api.HealthRes, err error) {
	res = &api.HealthRes{
		Status: "ok",
	}
	return
}
