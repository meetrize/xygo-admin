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

// ===================== 登录日志列表 =====================

// LoginLogListInp 登录日志列表入参
type LoginLogListInp struct {
	form.PageReq
	Username string `json:"username" dc:"登录账号（模糊搜索）"`
	Ip       string `json:"ip" dc:"登录IP"`
	Status   int    `json:"status" d:"-1" dc:"状态：-1=全部 0=失败 1=成功"`
	DateRange []string `json:"dateRange" dc:"时间范围"`
}

// LoginLogListModel 登录日志列表出参
type LoginLogListModel struct {
	List []LoginLogItem `json:"list"`
	form.PageRes
}

// LoginLogItem 登录日志列表项
type LoginLogItem struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"userId"`
	Username  string `json:"username"`
	Ip        string `json:"ip"`
	Location  string `json:"location"`
	Browser   string `json:"browser"`
	Os        string `json:"os"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
}

// LoginLogDeleteInp 删除登录日志入参
type LoginLogDeleteInp struct {
	Ids []uint `json:"ids" v:"required|min-length:1" dc:"日志ID列表"`
}

// LoginLogClearInp 清空登录日志入参
type LoginLogClearInp struct{}

// ===================== 操作日志列表 =====================

// OperationLogListInp 操作日志列表入参
type OperationLogListInp struct {
	form.PageReq
	Username  string   `json:"username" dc:"操作人账号（模糊搜索）"`
	Module    string   `json:"module" dc:"模块名称"`
	Status    int      `json:"status" d:"-1" dc:"状态：-1=全部 0=失败 1=成功"`
	DateRange []string `json:"dateRange" dc:"时间范围"`
}

// OperationLogListModel 操作日志列表出参
type OperationLogListModel struct {
	List []OperationLogItem `json:"list"`
	form.PageRes
}

// OperationLogItem 操作日志列表项
type OperationLogItem struct {
	Id           uint   `json:"id"`
	UserId       uint   `json:"userId"`
	Username     string `json:"username"`
	Module       string `json:"module"`
	Title        string `json:"title"`
	Method       string `json:"method"`
	Url          string `json:"url"`
	Ip           string `json:"ip"`
	Location     string `json:"location"`
	RequestBody  string `json:"requestBody"`
	ResponseBody string `json:"responseBody"`
	ErrorMessage string `json:"errorMessage"`
	Status       int    `json:"status"`
	Elapsed      uint   `json:"elapsed"`
	CreatedAt    string `json:"createdAt"`
}

// OperationLogDetailInp 操作日志详情入参
type OperationLogDetailInp struct {
	Id uint `json:"id" v:"required|min:1" dc:"日志ID"`
}

// OperationLogDeleteInp 删除操作日志入参
type OperationLogDeleteInp struct {
	Ids []uint `json:"ids" v:"required|min-length:1" dc:"日志ID列表"`
}

// OperationLogClearInp 清空操作日志入参
type OperationLogClearInp struct{}
