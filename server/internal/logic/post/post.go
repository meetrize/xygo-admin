// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package post

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/model/do"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/adminin"
	"xygo/internal/service"
)

type sAdminPost struct{}

func init() {
	service.RegisterAdminPost(New())
}

// New 构造岗位服务
func New() *sAdminPost {
	return &sAdminPost{}
}

// List 获取岗位列表
func (s *sAdminPost) List(ctx context.Context, in *adminin.PostListInp) (list []adminin.PostListItem, total int, err error) {
	model := dao.AdminPost.Ctx(ctx)

	// 按名称模糊搜索
	if in.Name != "" {
		model = model.WhereLike("name", "%"+in.Name+"%")
	}

	// 按编码模糊搜索
	if in.Code != "" {
		model = model.WhereLike("code", "%"+in.Code+"%")
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
		Fields("id, code, name, sort, status, remark, create_time, update_time").
		OrderAsc("sort, id").
		Page(in.Page, in.PageSize).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}

// Detail 获取岗位详情
func (s *sAdminPost) Detail(ctx context.Context, id uint64) (*adminin.PostListItem, error) {
	var post *entity.AdminPost

	if err := dao.AdminPost.Ctx(ctx).
		Where("id", id).
		Scan(&post); err != nil {
		return nil, err
	}

	if post == nil {
		return nil, gerror.NewCode(consts.CodeDataNotFound, "岗位不存在")
	}

	return &adminin.PostListItem{
		Id:         post.Id,
		Code:       post.Code,
		Name:       post.Name,
		Sort:       post.Sort,
		Status:     post.Status,
		Remark:     post.Remark,
		CreateTime: int(post.CreateTime),
		UpdateTime: int(post.UpdateTime),
	}, nil
}

// Save 保存岗位（新增/编辑）
func (s *sAdminPost) Save(ctx context.Context, in *adminin.PostSaveInp) (uint, error) {
	// 校验编码唯一性
	count, err := dao.AdminPost.Ctx(ctx).
		Where("code", in.Code).
		WhereNot("id", in.Id).
		Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, gerror.NewCode(consts.CodeInvalidParam, "岗位编码已存在")
	}

	data := do.AdminPost{
		Code:   in.Code,
		Name:   in.Name,
		Sort:   in.Sort,
		Status: in.Status,
		Remark: in.Remark,
	}

	if in.Id == 0 {
		// 新增
		r, err := dao.AdminPost.Ctx(ctx).Data(data).OmitNil().Insert()
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
	_, err = dao.AdminPost.Ctx(ctx).
		Data(data).
		OmitNil().
		Where("id", in.Id).
		Update()
	if err != nil {
		return 0, err
	}
	return uint(in.Id), nil
}

// Delete 删除岗位
func (s *sAdminPost) Delete(ctx context.Context, id uint64) error {
	// 检查是否有用户绑定该岗位
	userCount, err := dao.AdminUserPost.Ctx(ctx).
		Where("post_id", id).
		Count()
	if err != nil {
		return err
	}
	if userCount > 0 {
		return gerror.NewCode(consts.CodeInvalidParam, "该岗位下还有用户，无法删除")
	}

	_, err = dao.AdminPost.Ctx(ctx).
		Where("id", id).
		Delete()
	return err
}
