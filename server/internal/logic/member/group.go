// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// =================================================================================
// 会员分组逻辑层
// =================================================================================

package member

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/model/do"
	"xygo/internal/model/input/adminin"
)

type sAdminMemberGroup struct{}

// NewAdminMemberGroup 构造会员分组服务
func NewAdminMemberGroup() *sAdminMemberGroup {
	return &sAdminMemberGroup{}
}

// List 获取会员分组列表
func (s *sAdminMemberGroup) List(ctx context.Context, in *adminin.MemberGroupListInp) (list []adminin.MemberGroupListItem, total int, err error) {
	model := dao.MemberGroup.Ctx(ctx)

	// 按名称模糊搜索
	if in.Name != "" {
		model = model.WhereLike("name", "%"+in.Name+"%")
	}

	// 状态过滤
	if in.Status == 0 || in.Status == 1 {
		model = model.Where("status", in.Status)
	}

	// 统计总数
	count, err := model.Clone().Count()
	if err != nil {
		return nil, 0, err
	}

	// 分页参数兜底
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 20
	}

	// 查询列表
	err = model.
		Fields("id, name, rules, status, sort, remark, created_at, updated_at").
		OrderAsc("sort, id").
		Page(in.Page, in.PageSize).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

// Save 保存会员分组（新增/编辑）
func (s *sAdminMemberGroup) Save(ctx context.Context, in *adminin.MemberGroupSaveInp) (uint, error) {
	// 校验名称唯一性
	count, err := dao.MemberGroup.Ctx(ctx).
		Where("name", in.Name).
		WhereNot("id", in.Id).
		Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, gerror.NewCode(consts.CodeInvalidParam, "分组名称已存在")
	}

	data := do.MemberGroup{
		Name:   in.Name,
		Rules:  in.Rules,
		Sort:   in.Sort,
		Status: in.Status,
		Remark: in.Remark,
	}

	if in.Id == 0 {
		// 新增
		r, err := dao.MemberGroup.Ctx(ctx).Data(data).OmitNil().Insert()
		if err != nil {
			return 0, err
		}
		lastId, err := r.LastInsertId()
		if err != nil {
			return 0, err
		}
		return uint(lastId), nil
	}

	// 编辑
	_, err = dao.MemberGroup.Ctx(ctx).
		Data(data).
		OmitNil().
		Where("id", in.Id).
		Update()
	if err != nil {
		return 0, err
	}
	return uint(in.Id), nil
}

// Delete 删除会员分组
func (s *sAdminMemberGroup) Delete(ctx context.Context, id uint64) error {
	// 检查是否有会员使用该分组
	memberCount, err := dao.Member.Ctx(ctx).
		Where("group_id", id).
		Count()
	if err != nil {
		return err
	}
	if memberCount > 0 {
		return gerror.NewCode(consts.CodeInvalidParam, "该分组下还有会员，无法删除")
	}

	_, err = dao.MemberGroup.Ctx(ctx).
		Where("id", id).
		Delete()
	return err
}
