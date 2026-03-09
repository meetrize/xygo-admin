// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package form

// PageReq 通用分页请求
type PageReq struct {
	Page     int `p:"page"      d:"1"  v:"min:1#页码至少为1"             json:"page"`
	PageSize int `p:"pageSize"  d:"20" v:"min:1|max:100#每页数量范围为1-100" json:"pageSize"`
}

// PageRes 通用分页响应
type PageRes struct {
	Page     int `json:"page"     dc:"当前页码"`
	PageSize int `json:"pageSize" dc:"每页数量"`
	Total    int `json:"total"    dc:"总记录数"`
}
