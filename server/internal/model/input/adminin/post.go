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

// PostListInp 岗位列表查询入参
type PostListInp struct {
	form.PageReq
	Name   string `p:"name"   json:"name"   dc:"按岗位名称模糊搜索"`
	Code   string `p:"code"   json:"code"   dc:"按岗位编码模糊搜索"`
	Status int    `p:"status" d:"-1" json:"status" dc:"状态过滤:1启用,0禁用,-1全部"`
}

// PostListItem 岗位列表项
type PostListItem struct {
	Id         uint64 `json:"id"         dc:"岗位ID"`
	Code       string `json:"code"       dc:"岗位编码"`
	Name       string `json:"name"       dc:"岗位名称"`
	Sort       int    `json:"sort"       dc:"排序"`
	Status     int    `json:"status"     dc:"状态:0禁用,1启用"`
	Remark     string `json:"remark"     dc:"备注"`
	CreateTime int    `json:"create_time" dc:"创建时间"`
	UpdateTime int    `json:"update_time" dc:"更新时间"`
}

// PostListModel 岗位列表响应模型
type PostListModel struct {
	List []PostListItem `json:"list" dc:"数据列表"`
	form.PageRes
}

// PostDetailInp 岗位详情入参
type PostDetailInp struct {
	Id uint `p:"id" v:"required#岗位ID不能为空" json:"id" dc:"岗位ID"`
}

// PostSaveInp 岗位新增/编辑入参
type PostSaveInp struct {
	Id     uint64 `p:"id"     json:"id"     dc:"岗位ID（为空表示新增）"`
	Code   string `p:"code"   v:"required#岗位编码不能为空" json:"code"   dc:"岗位编码"`
	Name   string `p:"name"   v:"required#岗位名称不能为空" json:"name"   dc:"岗位名称"`
	Sort   int    `p:"sort"   d:"0"         json:"sort"   dc:"排序"`
	Status int    `p:"status" d:"1"         json:"status" dc:"状态:0禁用,1启用"`
	Remark string `p:"remark"               json:"remark" dc:"备注"`
}

// PostDeleteInp 岗位删除入参
type PostDeleteInp struct {
	Id uint64 `p:"id" v:"required#岗位ID不能为空" json:"id" dc:"岗位ID"`
}
