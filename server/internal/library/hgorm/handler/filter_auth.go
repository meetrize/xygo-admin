// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package handler

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/library/contexts"
	"xygo/internal/model/entity"
)

// FilterAuth 过滤数据权限
// 通过上下文中的用户角色权限和表中是否含有需要过滤的字段附加查询条件
func FilterAuth(m *gdb.Model) *gdb.Model {
	var (
		needAuth    bool
		filterField string
		fields      = getModelFields(m)
	)

	// 优先级：created_by > member_id
	if gstr.InArray(fields, "created_by") {
		needAuth = true
		filterField = "created_by"
	}

	if !needAuth && gstr.InArray(fields, "member_id") {
		needAuth = true
		filterField = "member_id"
	}

	if !needAuth {
		return m
	}
	return m.Handler(FilterAuthWithField(filterField))
}

// FilterAuthWithField 过滤数据权限，设置指定字段
func FilterAuthWithField(filterField string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		var (
			role *entity.AdminRole
			ctx  = m.GetCtx()
			user = contexts.GetUser(ctx) // ✨ 直接从上下文获取完整用户信息
		)

		if user == nil {
			return m
		}

		// ✅ 超级管理员豁免（只判断角色标识，对齐 HotGo）
		if consts.IsSuperRole(user.RoleKey) {
			return m
		}

		// 查询角色详情（获取数据范围配置）
		err := dao.AdminRole.Ctx(ctx).Where(dao.AdminRole.Columns().Id, user.RoleId).Scan(&role)
		if err != nil {
			g.Log().Warningf(ctx, "FilterAuth: failed to get role information, err:%+v", err)
			return m
		}

		if role == nil {
			g.Log().Warning(ctx, "FilterAuth: role not found")
			return m
		}

		// 辅助函数：根据部门ID获取用户ID列表
		getDeptUserIds := func(deptIds interface{}) []gdb.Value {
			ds, err := dao.AdminUser.Ctx(ctx).
				Fields(dao.AdminUser.Columns().Id).
				Where(dao.AdminUser.Columns().DeptId, deptIds).
				Array()
			if err != nil {
				g.Log().Warningf(ctx, "FilterAuth: failed to get dept users, err:%+v", err)
				return []gdb.Value{}
			}
			return ds
		}

		// 根据数据范围过滤
		switch role.DataScope {
		case consts.RoleDataAll: // 全部权限
			// 不做任何过滤
		case consts.RoleDataNowDept: // 当前部门
			userIds := getDeptUserIds(user.DeptId)
			if len(userIds) > 0 {
				m = m.WhereIn(filterField, userIds)
			} else {
				// 如果部门没有用户，返回空结果
				m = m.Where("1 = 0")
			}
		case consts.RoleDataDeptAndSub: // 当前部门及以下部门
			deptIds := GetDeptAndSub(ctx, user.DeptId)
			userIds := getDeptUserIds(deptIds)
			if len(userIds) > 0 {
				m = m.WhereIn(filterField, userIds)
			} else {
				m = m.Where("1 = 0")
			}
		case consts.RoleDataDeptCustom: // 自定义部门
			var customDepts []int64
			if role.CustomDepts != "" {
				if err := gjson.DecodeTo(role.CustomDepts, &customDepts); err != nil {
					g.Log().Warningf(ctx, "FilterAuth: failed to parse custom_depts, err:%+v", err)
				}
			}
			if len(customDepts) > 0 {
				userIds := getDeptUserIds(customDepts)
				if len(userIds) > 0 {
					m = m.WhereIn(filterField, userIds)
				} else {
					m = m.Where("1 = 0")
				}
			} else {
				m = m.Where("1 = 0")
			}
		case consts.RoleDataSelf: // 仅自己
			m = m.Where(filterField, user.Id)
		case consts.RoleDataSelfAndSub: // 自己和直属下级
			userIds := GetSelfAndSub(ctx, user.Id)
			if len(userIds) > 0 {
				m = m.WhereIn(filterField, userIds)
			} else {
				m = m.Where(filterField, user.Id)
			}
		case consts.RoleDataSelfAndAllSub: // 自己和全部下级
			userIds := GetSelfAndAllSub(ctx, user.Id)
			if len(userIds) > 0 {
				m = m.WhereIn(filterField, userIds)
			} else {
				m = m.Where(filterField, user.Id)
			}
		default:
			g.Log().Warningf(ctx, "FilterAuth: unknown dataScope: %d", role.DataScope)
		}

		return m
	}
}

// GetDeptAndSub 获取指定部门的所有下级部门ID，含本部门
func GetDeptAndSub(ctx context.Context, deptId uint64) []int64 {
	if deptId == 0 {
		return []int64{}
	}

	// 查询当前部门信息
	var dept *entity.AdminDept
	err := dao.AdminDept.Ctx(ctx).Where(dao.AdminDept.Columns().Id, deptId).Scan(&dept)
	if err != nil || dept == nil {
		g.Log().Warningf(ctx, "GetDeptAndSub: failed to get dept, deptId=%d, err:%+v", deptId, err)
		return []int64{int64(deptId)}
	}

	// 方式1：如果有 tree 字段，使用 LIKE 查询
	// 假设 tree 字段格式为 "0,1,3"（逗号分隔的父级ID链）
	// 查询所有 tree 包含当前部门ID的部门

	// 方式2：使用闭包表（如果有 admin_dept_closure 表）
	// 查询 admin_dept_closure 表获取所有子部门
	closureArray, err := dao.AdminDeptClosure.Ctx(ctx).
		Where(dao.AdminDeptClosure.Columns().Ancestor, deptId).
		Fields(dao.AdminDeptClosure.Columns().Descendant).
		Array()
	if err == nil && len(closureArray) > 0 {
		// 转换为 int64 切片
		var ids []int64
		for _, v := range closureArray {
			ids = append(ids, v.Int64())
		}
		return ids
	}

	// 方式3：递归查询（性能较低，但兜底）
	ids := recursiveGetSubDepts(ctx, deptId)
	ids = append(ids, int64(deptId)) // 包含自己
	return ids
}

// recursiveGetSubDepts 递归查询子部门（兜底方案）
func recursiveGetSubDepts(ctx context.Context, parentId uint64) []int64 {
	var ids []int64

	array, err := dao.AdminDept.Ctx(ctx).
		Where(dao.AdminDept.Columns().ParentId, parentId).
		Fields(dao.AdminDept.Columns().Id).
		Array()
	if err != nil {
		return ids
	}

	for _, v := range array {
		childId := v.Uint64()
		ids = append(ids, int64(childId))
		// 递归查询子部门的子部门
		subIds := recursiveGetSubDepts(ctx, childId)
		ids = append(ids, subIds...)
	}

	return ids
}

// GetSelfAndSub 获取直属下级用户ID，包含自己
func GetSelfAndSub(ctx context.Context, userId uint64) []int64 {
	if userId == 0 {
		return []int64{}
	}

	array, err := dao.AdminUser.Ctx(ctx).
		Where(dao.AdminUser.Columns().Pid, userId).
		Fields(dao.AdminUser.Columns().Id).
		Array()
	if err != nil {
		g.Log().Warningf(ctx, "GetSelfAndSub: failed to get sub users, err:%+v", err)
		return []int64{int64(userId)}
	}

	var ids []int64
	for _, v := range array {
		ids = append(ids, v.Int64())
	}
	ids = append(ids, int64(userId)) // 包含自己
	return ids
}

// GetSelfAndAllSub 获取全部下级用户ID，包含自己
func GetSelfAndAllSub(ctx context.Context, userId uint64) []int64 {
	if userId == 0 {
		return []int64{}
	}

	var getAllSub func(pid uint64) []int64
	getAllSub = func(pid uint64) []int64 {
		var ids []int64
		array, err := dao.AdminUser.Ctx(ctx).
			Where(dao.AdminUser.Columns().Pid, pid).
			Fields(dao.AdminUser.Columns().Id).
			Array()
		if err != nil {
			return ids
		}

		for _, v := range array {
			childId := v.Uint64()
			ids = append(ids, int64(childId))
			// 递归查询子用户的子用户
			subIds := getAllSub(childId)
			ids = append(ids, subIds...)
		}
		return ids
	}

	ids := getAllSub(userId)
	ids = append(ids, int64(userId)) // 包含自己
	return ids
}

// GetUserFromContext 从上下文获取当前用户信息
// 注意：此函数已不再使用，保留是为了兼容
// 推荐直接使用 contexts.GetUser(ctx) 获取完整的 AuthUser 信息
func GetUserFromContext(ctx context.Context) *entity.AdminUser {
	authUser := contexts.GetUser(ctx)
	if authUser == nil {
		return nil
	}

	// 返回简化的用户信息（只包含 entity.AdminUser 有的字段）
	return &entity.AdminUser{
		Id:     authUser.Id,
		DeptId: authUser.DeptId,
		// 注意：已不使用 is_super 字段，超管判断基于 RoleKey
	}
}

// getModelFields 获取模型的字段列表
func getModelFields(m *gdb.Model) []string {
	fieldsStr := m.GetFieldsStr()
	if fieldsStr == "" || fieldsStr == "*" {
		// 如果是 SELECT *，返回空（表示需要检查所有字段）
		// 实际使用时会自动检测表是否有 created_by 等字段
		return []string{}
	}

	// 解析字段字符串（逗号分隔，去除空格和引号）
	fields := gstr.SplitAndTrim(fieldsStr, ",")
	var cleanFields []string
	for _, f := range fields {
		// 去除可能的表名前缀和反引号
		f = gstr.Trim(f, "`")
		if gstr.Contains(f, ".") {
			parts := gstr.Split(f, ".")
			f = parts[len(parts)-1]
		}
		f = gstr.Trim(f, "`")
		cleanFields = append(cleanFields, f)
	}
	return cleanFields
}
