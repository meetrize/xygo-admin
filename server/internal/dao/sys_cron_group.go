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

var SysCronGroup = &sysCronGroupDao{
	table: "xy_sys_cron_group",
}

type sysCronGroupDao struct {
	table string
}

func (d *sysCronGroupDao) Table() string { return d.table }

func (d *sysCronGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB().Model(d.table).Safe().Ctx(ctx)
}

func (d *sysCronGroupDao) Columns() *sysCronGroupColumns { return &sysCronGroupCols }

type sysCronGroupColumns struct {
	Id, Name, Sort, Remark, Status string
	CreatedAt, UpdatedAt           string
}

var sysCronGroupCols = sysCronGroupColumns{
	Id: "id", Name: "name", Sort: "sort", Remark: "remark", Status: "status",
	CreatedAt: "created_at", UpdatedAt: "updated_at",
}
