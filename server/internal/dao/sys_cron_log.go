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

var SysCronLog = &sysCronLogDao{
	table: "xy_sys_cron_log",
}

type sysCronLogDao struct {
	table string
}

func (d *sysCronLogDao) Table() string { return d.table }

func (d *sysCronLogDao) Ctx(ctx context.Context) *gdb.Model {
	return g.DB().Model(d.table).Safe().Ctx(ctx)
}

func (d *sysCronLogDao) Columns() *sysCronLogColumns { return &sysCronLogCols }

type sysCronLogColumns struct {
	Id, CronId, Name, Title, Params string
	Status, Output, ErrMsg, TakeMs  string
	CreatedAt                       string
}

var sysCronLogCols = sysCronLogColumns{
	Id: "id", CronId: "cron_id", Name: "name", Title: "title",
	Params: "params", Status: "status", Output: "output",
	ErrMsg: "err_msg", TakeMs: "take_ms", CreatedAt: "created_at",
}
