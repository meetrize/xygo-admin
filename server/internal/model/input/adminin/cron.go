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

import "xygo/internal/model/input/form"

// ==================== 定时任务 ====================

// CronListInp 任务列表入参
type CronListInp struct {
	form.PageReq
	GroupId uint64 `p:"groupId" json:"groupId" dc:"分组ID"`
	Status  int    `p:"status"  d:"-1" json:"status" dc:"状态:-1全部,0禁用,1启用"`
	Name    string `p:"name"    json:"name"    dc:"任务标识模糊搜索"`
}

// CronListItem 任务列表项
type CronListItem struct {
	Id        uint64 `json:"id"`
	GroupId   uint64 `json:"groupId"`
	GroupName string `json:"groupName"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Params    string `json:"params"`
	Pattern   string `json:"pattern"`
	Policy    int    `json:"policy"`
	Count     int    `json:"count"`
	Sort      int    `json:"sort"`
	Remark    string `json:"remark"`
	Status    int    `json:"status"`
	CreatedAt uint   `json:"createdAt"`
	UpdatedAt uint   `json:"updatedAt"`
}

// CronListModel 任务列表出参
type CronListModel struct {
	List []CronListItem `json:"list"`
	form.PageRes
}

// CronSaveInp 任务新增/编辑入参
type CronSaveInp struct {
	Id      uint64 `p:"id"      json:"id"      dc:"任务ID"`
	GroupId uint64 `p:"groupId" json:"groupId"  dc:"分组ID"`
	Title   string `p:"title"   json:"title"    v:"required#任务标题不能为空" dc:"任务标题"`
	Name    string `p:"name"    json:"name"     v:"required#任务标识不能为空" dc:"任务标识"`
	Params  string `p:"params"  json:"params"   dc:"任务参数"`
	Pattern string `p:"pattern" json:"pattern"  v:"required#Cron表达式不能为空" dc:"Cron表达式"`
	Policy  int    `p:"policy"  json:"policy"   d:"1" dc:"策略"`
	Count   int    `p:"count"   json:"count"    d:"0" dc:"固定次数"`
	Sort    int    `p:"sort"    json:"sort"     d:"0" dc:"排序"`
	Remark  string `p:"remark"  json:"remark"   dc:"备注"`
	Status  int    `p:"status"  json:"status"   d:"1" dc:"状态"`
}

// CronStatusInp 修改任务状态入参
type CronStatusInp struct {
	Id     uint64 `p:"id"     json:"id"     v:"required#任务ID不能为空" dc:"任务ID"`
	Status int    `p:"status" json:"status"  v:"required|in:0,1#状态不能为空|状态值无效" dc:"状态"`
}

// CronDeleteInp 删除任务入参
type CronDeleteInp struct {
	Id uint64 `p:"id" json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

// CronOnlineExecInp 在线执行入参
type CronOnlineExecInp struct {
	Id uint64 `p:"id" json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

// ==================== 定时任务分组 ====================

// CronGroupListInp 分组列表入参
type CronGroupListInp struct {
	form.PageReq
	Status int `p:"status" d:"-1" json:"status" dc:"状态"`
}

// CronGroupListItem 分组列表项
type CronGroupListItem struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	Remark    string `json:"remark"`
	Status    int    `json:"status"`
	CreatedAt uint   `json:"createdAt"`
	UpdatedAt uint   `json:"updatedAt"`
}

// CronGroupListModel 分组列表出参
type CronGroupListModel struct {
	List []CronGroupListItem `json:"list"`
	form.PageRes
}

// CronGroupSaveInp 分组保存入参
type CronGroupSaveInp struct {
	Id     uint64 `p:"id"     json:"id"     dc:"分组ID"`
	Name   string `p:"name"   json:"name"   v:"required#分组名称不能为空" dc:"分组名称"`
	Sort   int    `p:"sort"   json:"sort"   d:"0" dc:"排序"`
	Remark string `p:"remark" json:"remark" dc:"备注"`
	Status int    `p:"status" json:"status" d:"1" dc:"状态"`
}

// CronGroupDeleteInp 分组删除入参
type CronGroupDeleteInp struct {
	Id uint64 `p:"id" json:"id" v:"required#分组ID不能为空" dc:"分组ID"`
}

// CronGroupSelectItem 分组下拉项
type CronGroupSelectItem struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// ==================== 执行日志 ====================

// CronLogListInp 执行日志入参
type CronLogListInp struct {
	form.PageReq
	CronId uint64 `p:"cronId" json:"cronId" dc:"任务ID"`
	Status int    `p:"status" d:"-1" json:"status" dc:"状态"`
}

// CronLogListItem 执行日志项
type CronLogListItem struct {
	Id        uint64 `json:"id"`
	CronId    uint64 `json:"cronId"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Params    string `json:"params"`
	Status    int    `json:"status"`
	Output    string `json:"output"`
	ErrMsg    string `json:"errMsg"`
	TakeMs    int    `json:"takeMs"`
	CreatedAt uint   `json:"createdAt"`
}

// CronLogListModel 执行日志出参
type CronLogListModel struct {
	List []CronLogListItem `json:"list"`
	form.PageRes
}

// CronLogClearInp 清空日志入参
type CronLogClearInp struct {
	CronId uint64 `p:"cronId" json:"cronId" dc:"任务ID（为空清全部）"`
}
