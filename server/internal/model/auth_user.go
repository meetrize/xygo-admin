// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package model

// AuthUser 登录态中携带的用户完整信息（用于权限控制）
type AuthUser struct {
	// 基础信息
	Id       uint64 `json:"id"       description:"用户ID"`
	Username string `json:"username" description:"用户名"`
	Nickname string `json:"nickname" description:"昵称"`
	Avatar   string `json:"avatar"   description:"头像"`
	Email    string `json:"email"    description:"邮箱"`
	Mobile   string `json:"mobile"   description:"手机号"`

	// 权限相关（数据权限、字段权限）
	Pid     uint64 `json:"pid"     description:"上级用户ID（用于数据范围：自己和下级）"`
	DeptId  uint64 `json:"deptId"  description:"部门ID（用于数据范围：部门）"`
	RoleId  uint64 `json:"roleId"  description:"角色ID"`
	RoleKey string `json:"roleKey" description:"角色标识（用于超管判断：super_admin）"`

	// 登录信息
	LoginAt int64 `json:"loginAt" description:"登录时间"`
}
