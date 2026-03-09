// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var AdminNoticeRead = &adminNoticeReadDao{
	table: "xy_admin_notice_read",
}

type adminNoticeReadDao struct {
	table string
}

func (d *adminNoticeReadDao) Table() string { return d.table }

func (d *adminNoticeReadDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB().Model(d.table).Safe().Ctx(ctx)
}
