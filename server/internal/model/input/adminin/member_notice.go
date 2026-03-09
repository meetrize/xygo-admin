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

import (
	"xygo/internal/model/input/form"
)

// ==================== 会员通知 ====================

// MemberNoticeListInp 会员通知列表入参
type MemberNoticeListInp struct {
	form.PageReq
	Title string `json:"title" dc:"通知标题"`
	Type string `json:"type" dc:"通知类型"`
	Status *int `json:"status" dc:"状态"`
	// 关联表搜索字段
}

// MemberNoticeListItem 会员通知列表项
type MemberNoticeListItem struct {
	Id uint64 `json:"id" dc:""`
	Title string `json:"title" dc:"通知标题"`
	Type string `json:"type" dc:"通知类型"`
	Target string `json:"target" dc:"目标"`
	TargetId uint64 `json:"targetId" dc:"目标分组ID"`
	Sender string `json:"sender" dc:"发送者"`
	Status int `json:"status" dc:"状态"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
	// 关联表字段（来自 LeftJoin）
	TargetName string `json:"target_name" dc:"Targetname"`
}

// MemberNoticeListModel 会员通知列表出参
type MemberNoticeListModel struct {
	List []MemberNoticeListItem `json:"list"`
	form.PageRes
}

// MemberNoticeViewModel 会员通知详情出参
type MemberNoticeViewModel struct {
	Id uint64 `json:"id" dc:""`
	Title string `json:"title" dc:"通知标题"`
	Content string `json:"content" dc:"通知内容"`
	Type string `json:"type" dc:"通知类型"`
	Target string `json:"target" dc:"目标"`
	TargetId uint64 `json:"targetId" dc:"目标分组ID"`
	Sender string `json:"sender" dc:"发送者"`
	Status int `json:"status" dc:"状态"`
	CreatedAt uint64 `json:"createdAt" dc:"创建时间"`
}

// MemberNoticeEditInp 会员通知编辑入参
type MemberNoticeEditInp struct {
	Id uint64 `json:"id" dc:""`
	Title string `json:"title" v:"required#通知标题不能为空" dc:"通知标题"`
	Content string `json:"content" dc:"通知内容"`
	Type string `json:"type" v:"required#通知类型不能为空" dc:"通知类型"`
	Target string `json:"target" v:"required#目标不能为空" dc:"目标"`
	TargetId uint64 `json:"targetId" v:"required#目标分组ID不能为空" dc:"目标分组ID"`
	Sender string `json:"sender" v:"required#发送者不能为空" dc:"发送者"`
	Status int `json:"status" v:"required#状态不能为空" dc:"状态"`
}
