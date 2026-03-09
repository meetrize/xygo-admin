// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package memberloginlog

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"xygo/internal/dao"
	"xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

type sMemberLoginLog struct{}

func init() {
	service.RegisterMemberLoginLog(New())
}

func New() *sMemberLoginLog {
	return &sMemberLoginLog{}
}

// List 登录日志列表
func (s *sMemberLoginLog) List(ctx context.Context, in *adminin.MemberLoginLogListInp) (*adminin.MemberLoginLogListModel, error) {
	model := dao.MemberLoginLog.Ctx(ctx).As("t")
	// 关联表 LeftJoin
	model = model.LeftJoin("xy_member member", "member.id = t.member_id")
	if in.Status != nil {
		model = model.Where("t.status", *in.Status)
	}
	// 关联表搜索条件
	if in.MemberUsername != "" {
		model = model.WhereLike("member.username", "%"+in.MemberUsername+"%")
	}
	// 先计数（不带 Fields，避免 COUNT + 字段别名冲突）
	count, err := model.Clone().Count()
	if err != nil {
		return nil, err
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 20
	}
	// 计数后添加 Fields
	model = model.Fields("t.*")
	model = model.Fields("member.username as member_username")
	var list []adminin.MemberLoginLogListItem
	err = model.Page(in.Page, in.PageSize).OrderDesc("t.id").Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.MemberLoginLogListItem{}
	}

	return &adminin.MemberLoginLogListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    count,
		},
	}, nil
}

// View 登录日志详情
func (s *sMemberLoginLog) View(ctx context.Context, id uint64) (*adminin.MemberLoginLogViewModel, error) {
	var item adminin.MemberLoginLogViewModel
	err := dao.MemberLoginLog.Ctx(ctx).Where("id", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, gerror.New("记录不存在")
	}
	return &item, nil
}

// Delete 删除登录日志
func (s *sMemberLoginLog) Delete(ctx context.Context, id uint64) error {
	_, err := dao.MemberLoginLog.Ctx(ctx).Where("id", id).Delete()
	return err
}
