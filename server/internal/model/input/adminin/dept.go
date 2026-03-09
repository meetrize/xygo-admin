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

// DeptListInp 部门列表查询入参
type DeptListInp struct {
	Name   string `p:"name"   json:"name"   dc:"按部门名称模糊搜索"`
	Status int    `p:"status" d:"-1" json:"status" dc:"状态过滤:1启用,0禁用,-1全部"`
}

// DeptListItem 部门列表项（树形结构）
type DeptListItem struct {
	Id         uint64          `json:"id"         dc:"部门ID"`
	ParentId   uint64          `json:"parentId"   dc:"上级部门ID"`
	Name       string          `json:"name"       dc:"部门名称"`
	Sort       int             `json:"sort"       dc:"排序"`
	Status     int             `json:"status"     dc:"状态:0禁用,1启用"`
	Remark     string          `json:"remark"     dc:"备注"`
	CreateTime int             `json:"create_time" dc:"创建时间"`
	UpdateTime int             `json:"update_time" dc:"更新时间"`
	Children   []*DeptListItem `json:"children,omitempty" dc:"子部门"`
}

// DeptListModel 部门列表响应模型
type DeptListModel struct {
	List []*DeptListItem `json:"list" dc:"数据列表"`
}

// DeptDetailInp 部门详情入参
type DeptDetailInp struct {
	Id uint `p:"id" v:"required#部门ID不能为空" json:"id" dc:"部门ID"`
}

// DeptSaveInp 部门新增/编辑入参
type DeptSaveInp struct {
	Id       uint64 `p:"id"       json:"id"       dc:"部门ID（为空表示新增）"`
	ParentId uint64 `p:"parentId" d:"0"           json:"parentId" dc:"上级部门ID，0表示根部门"`
	Name     string `p:"name"     v:"required#部门名称不能为空" json:"name"     dc:"部门名称"`
	Sort     int    `p:"sort"     d:"0"           json:"sort"     dc:"排序"`
	Status   int    `p:"status"   d:"1"           json:"status"   dc:"状态:0禁用,1启用"`
	Remark   string `p:"remark"                   json:"remark"   dc:"备注"`
}

// DeptDeleteInp 部门删除入参
type DeptDeleteInp struct {
	Id uint64 `p:"id" v:"required#部门ID不能为空" json:"id" dc:"部门ID"`
}
