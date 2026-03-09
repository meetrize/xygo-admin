// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package membernotice

import (
	"context"
	"time"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	"xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

type sMemberNotice struct{}

func init() {
	service.RegisterMemberNotice(New())
}

func New() *sMemberNotice {
	return &sMemberNotice{}
}

// List 会员通知列表
func (s *sMemberNotice) List(ctx context.Context, in *adminin.MemberNoticeListInp) (*adminin.MemberNoticeListModel, error) {
	model := dao.MemberNotice.Ctx(ctx).As("t")
	// 关联表 LeftJoin
	model = model.LeftJoin("xy_member_group target", "target.id = t.target_id")
	if in.Title != "" {
		model = model.WhereLike("t.title", "%"+in.Title+"%")
	}
	if in.Type != "" {
		model = model.Where("t.type", in.Type)
	}
	if in.Status != nil {
		model = model.Where("t.status", *in.Status)
	}
	// 关联表搜索条件
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
	model = model.Fields("target.name as target_name")
	var list []adminin.MemberNoticeListItem
	err = model.Page(in.Page, in.PageSize).OrderDesc("t.id").Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.MemberNoticeListItem{}
	}

	return &adminin.MemberNoticeListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    count,
		},
	}, nil
}

// View 会员通知详情
func (s *sMemberNotice) View(ctx context.Context, id uint64) (*adminin.MemberNoticeViewModel, error) {
	var item adminin.MemberNoticeViewModel
	err := dao.MemberNotice.Ctx(ctx).Where("id", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, gerror.New("记录不存在")
	}
	return &item, nil
}

// Edit 保存会员通知
func (s *sMemberNotice) Edit(ctx context.Context, in *adminin.MemberNoticeEditInp) error {
	data := g.Map{
		"title": in.Title,
		"content": in.Content,
		"type": in.Type,
		"target": in.Target,
		"target_id": in.TargetId,
		"sender": in.Sender,
		"status": in.Status,
	}

	if in.Id == 0 {
		// 新增
		data["created_at"] = time.Now().Unix()
		_, err := dao.MemberNotice.Ctx(ctx).Data(data).Insert()
		return err
	}

	// 更新
	_, err := dao.MemberNotice.Ctx(ctx).Where("id", in.Id).Data(data).Update()
	return err
}

// Delete 删除会员通知
func (s *sMemberNotice) Delete(ctx context.Context, id uint64) error {
	_, err := dao.MemberNotice.Ctx(ctx).Where("id", id).Delete()
	return err
}
