// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package memberscorelog

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

type sMemberScoreLog struct{}

func init() {
	service.RegisterMemberScoreLog(New())
}

func New() *sMemberScoreLog {
	return &sMemberScoreLog{}
}

// List 积分变动日志列表
func (s *sMemberScoreLog) List(ctx context.Context, in *adminin.MemberScoreLogListInp) (*adminin.MemberScoreLogListModel, error) {
	model := dao.MemberScoreLog.Ctx(ctx).As("t")
	// 关联表 LeftJoin
	model = model.LeftJoin("xy_member member", "member.id = t.member_id")
	if in.ScoreStart != "" && in.ScoreEnd != "" {
		model = model.WhereBetween("t.score", in.ScoreStart, in.ScoreEnd)
	}
	// 关联表搜索条件
	if in.MemberUsername != "" {
		model = model.WhereLike("member.username", "%"+in.MemberUsername+"%")
	}
	if in.MemberNickname != "" {
		model = model.WhereLike("member.nickname", "%"+in.MemberNickname+"%")
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
	model = model.Fields("member.nickname as member_nickname")
	model = model.Fields("member.username as member_username")
	model = model.Fields("member.avatar as member_avatar")
	var list []adminin.MemberScoreLogListItem
	err = model.Page(in.Page, in.PageSize).OrderDesc("t.id").Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.MemberScoreLogListItem{}
	}

	return &adminin.MemberScoreLogListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    count,
		},
	}, nil
}

// View 积分变动日志详情
func (s *sMemberScoreLog) View(ctx context.Context, id uint64) (*adminin.MemberScoreLogViewModel, error) {
	var item adminin.MemberScoreLogViewModel
	err := dao.MemberScoreLog.Ctx(ctx).Where("id", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, gerror.New("记录不存在")
	}
	return &item, nil
}

// Edit 保存积分变动日志
func (s *sMemberScoreLog) Edit(ctx context.Context, in *adminin.MemberScoreLogEditInp) error {
	data := g.Map{
		"member_id": in.MemberId,
		"score": in.Score,
		"before": in.Before,
		"after": in.After,
		"memo": in.Memo,
	}

	if in.Id == 0 {
		// 新增
		data["created_at"] = time.Now().Unix()
		_, err := dao.MemberScoreLog.Ctx(ctx).Data(data).Insert()
		return err
	}

	// 更新
	_, err := dao.MemberScoreLog.Ctx(ctx).Where("id", in.Id).Data(data).Update()
	return err
}

// Delete 删除积分变动日志
func (s *sMemberScoreLog) Delete(ctx context.Context, id uint64) error {
	_, err := dao.MemberScoreLog.Ctx(ctx).Where("id", id).Delete()
	return err
}
