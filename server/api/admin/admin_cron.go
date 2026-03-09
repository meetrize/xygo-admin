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
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/model/input/adminin"
)

// ==================== 定时任务 ====================

type CronListReq struct {
	g.Meta `path:"/admin/cron/list" method:"get" tags:"Cron" summary:"定时任务列表"`
	adminin.CronListInp
}
type CronListRes struct {
	*adminin.CronListModel
}

type CronSaveReq struct {
	g.Meta `path:"/admin/cron/save" method:"post" tags:"Cron" summary:"保存定时任务"`
	adminin.CronSaveInp
}
type CronSaveRes struct {
	Id uint64 `json:"id"`
}

type CronDeleteReq struct {
	g.Meta `path:"/admin/cron/delete" method:"post" tags:"Cron" summary:"删除定时任务"`
	adminin.CronDeleteInp
}
type CronDeleteRes struct{}

type CronStatusReq struct {
	g.Meta `path:"/admin/cron/status" method:"post" tags:"Cron" summary:"修改任务状态"`
	adminin.CronStatusInp
}
type CronStatusRes struct{}

type CronOnlineExecReq struct {
	g.Meta `path:"/admin/cron/onlineExec" method:"post" tags:"Cron" summary:"在线执行一次"`
	adminin.CronOnlineExecInp
}
type CronOnlineExecRes struct {
	Output string `json:"output"`
}

type CronRegisteredTasksReq struct {
	g.Meta `path:"/admin/cron/registeredTasks" method:"get" tags:"Cron" summary:"获取已注册任务列表"`
}
type CronRegisteredTasksRes struct {
	List []string `json:"list"`
}

// ==================== 定时任务分组 ====================

type CronGroupListReq struct {
	g.Meta `path:"/admin/cronGroup/list" method:"get" tags:"CronGroup" summary:"分组列表"`
	adminin.CronGroupListInp
}
type CronGroupListRes struct {
	*adminin.CronGroupListModel
}

type CronGroupSaveReq struct {
	g.Meta `path:"/admin/cronGroup/save" method:"post" tags:"CronGroup" summary:"保存分组"`
	adminin.CronGroupSaveInp
}
type CronGroupSaveRes struct {
	Id uint64 `json:"id"`
}

type CronGroupDeleteReq struct {
	g.Meta `path:"/admin/cronGroup/delete" method:"post" tags:"CronGroup" summary:"删除分组"`
	adminin.CronGroupDeleteInp
}
type CronGroupDeleteRes struct{}

type CronGroupSelectReq struct {
	g.Meta `path:"/admin/cronGroup/select" method:"get" tags:"CronGroup" summary:"分组下拉选项"`
}
type CronGroupSelectRes struct {
	List []adminin.CronGroupSelectItem `json:"list"`
}

// ==================== 执行日志 ====================

type CronLogListReq struct {
	g.Meta `path:"/admin/cronLog/list" method:"get" tags:"CronLog" summary:"执行日志列表"`
	adminin.CronLogListInp
}
type CronLogListRes struct {
	*adminin.CronLogListModel
}

type CronLogClearReq struct {
	g.Meta `path:"/admin/cronLog/clear" method:"post" tags:"CronLog" summary:"清空执行日志"`
	adminin.CronLogClearInp
}
type CronLogClearRes struct{}
