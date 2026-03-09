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

// ===================== 会员分组列表 =====================

// MemberGroupListInp 会员分组列表查询入参
type MemberGroupListInp struct {
	form.PageReq
	Name   string `p:"name" json:"name" dc:"按分组名称模糊搜索"`
	Status int    `p:"status" d:"-1" json:"status" dc:"状态过滤:1启用,0禁用,-1全部"`
}

// MemberGroupListItem 会员分组列表项
type MemberGroupListItem struct {
	Id        uint64 `json:"id" dc:"分组ID"`
	Name      string `json:"name" dc:"分组名称"`
	Rules     string `json:"rules" dc:"权限规则"`
	Status    int    `json:"status" dc:"状态:0禁用,1启用"`
	Sort      int    `json:"sort" dc:"排序"`
	Remark    string `json:"remark" dc:"备注"`
	CreatedAt string `json:"createdAt" dc:"创建时间"`
	UpdatedAt string `json:"updatedAt" dc:"更新时间"`
}

// MemberGroupListModel 会员分组列表响应模型
type MemberGroupListModel struct {
	List []MemberGroupListItem `json:"list" dc:"数据列表"`
	form.PageRes
}

// ===================== 会员分组保存 =====================

// MemberGroupSaveInp 会员分组新增/编辑入参
type MemberGroupSaveInp struct {
	Id     uint64 `p:"id" json:"id" dc:"分组ID（为空表示新增）"`
	Name   string `p:"name" v:"required#分组名称不能为空" json:"name" dc:"分组名称"`
	Rules  string `p:"rules" json:"rules" dc:"权限规则（菜单ID列表，逗号分隔）"`
	Sort   int    `p:"sort" d:"0" json:"sort" dc:"排序"`
	Status int    `p:"status" d:"1" json:"status" dc:"状态:0禁用,1启用"`
	Remark string `p:"remark" json:"remark" dc:"备注"`
}

// ===================== 会员分组删除 =====================

// MemberGroupDeleteInp 会员分组删除入参
type MemberGroupDeleteInp struct {
	Id uint64 `p:"id" v:"required#分组ID不能为空" json:"id" dc:"分组ID"`
}
