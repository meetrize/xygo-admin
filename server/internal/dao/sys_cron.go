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

var SysCron = &sysCronDao{
	table: "xy_sys_cron",
}

type sysCronDao struct {
	table string
}

func (d *sysCronDao) Table() string { return d.table }

func (d *sysCronDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB().Model(d.table).Safe().Ctx(ctx)
}

func (d *sysCronDao) Columns() *sysCronColumns { return &sysCronCols }

type sysCronColumns struct {
	Id, GroupId, Title, Name, Params, Pattern string
	Policy, Count, Sort, Remark, Status       string
	CreatedAt, UpdatedAt                      string
}

var sysCronCols = sysCronColumns{
	Id: "id", GroupId: "group_id", Title: "title", Name: "name",
	Params: "params", Pattern: "pattern", Policy: "policy", Count: "count",
	Sort: "sort", Remark: "remark", Status: "status",
	CreatedAt: "created_at", UpdatedAt: "updated_at",
}
