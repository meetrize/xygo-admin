// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package admin

import (
	"context"

	api "xygo/api/admin"
	"xygo/internal/field"
	"xygo/internal/model/input/adminin"
)

// ResourceList 获取所有已注册的资源列表
func (c *ControllerV1) ResourceList(ctx context.Context, req *api.ResourceListReq) (res *api.ResourceListRes, err error) {
	resources := field.GetAll()

	items := make([]adminin.ResourceItem, 0, len(resources))
	for _, r := range resources {
		items = append(items, adminin.ResourceItem{
			Code:  r.Code,
			Label: r.Label,
		})
	}

	res = new(api.ResourceListRes)
	res.ResourceListModel = &adminin.ResourceListModel{
		List: items,
	}
	return res, nil
}
