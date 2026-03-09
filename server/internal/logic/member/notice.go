// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package member

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"xygo/internal/dao"
	"xygo/internal/model/input/memberin"
)

type sFrontendNotice struct{}

func NewFrontendNotice() *sFrontendNotice {
	return &sFrontendNotice{}
}

// List 通知列表（带已读状态，只返回已发布的通知）
func (s *sFrontendNotice) List(ctx context.Context, memberId uint64, in *memberin.NoticeListInput) (out *memberin.NoticeListOutput, err error) {
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.PageSize <= 0 {
		in.PageSize = 20
	}

	// 只查已发布的通知 + 目标为 all 或当前会员所在分组
	model := dao.MemberNotice.Ctx(ctx).As("n").
		Where("n.status", 1).
		Where("n.target = 'all'") // 暂时只处理全员通知

	count, err := model.Clone().Count()
	if err != nil {
		return nil, err
	}

	// 使用子查询带出已读状态
	type noticeRow struct {
		memberin.NoticeItem
	}

	var rawList []noticeRow
	err = model.
		Fields("n.id, n.title, n.content, n.type, n.sender, n.created_at").
		Page(in.Page, in.PageSize).
		OrderDesc("n.id").
		Scan(&rawList)
	if err != nil {
		return nil, err
	}

	// 查询该会员已读的通知ID列表
	readIds := make(map[uint64]bool)
	if len(rawList) > 0 {
		noticeIds := make([]uint64, 0, len(rawList))
		for _, r := range rawList {
			noticeIds = append(noticeIds, r.Id)
		}
		var readRecords []struct {
			NoticeId uint64 `json:"noticeId"`
		}
		_ = dao.MemberNoticeRead.Ctx(ctx).
			Where("member_id", memberId).
			WhereIn("notice_id", noticeIds).
			Fields("notice_id").
			Scan(&readRecords)
		for _, r := range readRecords {
			readIds[r.NoticeId] = true
		}
	}

	// 组装结果
	list := make([]memberin.NoticeItem, 0, len(rawList))
	for _, r := range rawList {
		item := r.NoticeItem
		item.IsRead = readIds[item.Id]
		list = append(list, item)
	}

	// 统计未读数
	unread := 0
	for _, item := range list {
		if !item.IsRead {
			unread++
		}
	}

	return &memberin.NoticeListOutput{
		List:     list,
		Page:     in.Page,
		PageSize: in.PageSize,
		Total:    count,
		Unread:   unread,
	}, nil
}

// MarkRead 标记单条通知已读
func (s *sFrontendNotice) MarkRead(ctx context.Context, memberId uint64, noticeId uint64) error {
	// 检查是否已读
	count, err := dao.MemberNoticeRead.Ctx(ctx).
		Where("member_id", memberId).
		Where("notice_id", noticeId).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return nil // 已读，跳过
	}

	_, err = dao.MemberNoticeRead.Ctx(ctx).Data(g.Map{
		"notice_id": noticeId,
		"member_id": memberId,
		"read_at":   gtime.Now().Unix(),
	}).Insert()

	return err
}

// MarkAllRead 全部通知标记已读
func (s *sFrontendNotice) MarkAllRead(ctx context.Context, memberId uint64) error {
	// 查所有未读通知（已发布 + 未在 read 表中）
	var noticeIds []uint64
	err := dao.MemberNotice.Ctx(ctx).
		Where("status", 1).
		Where("target = 'all'").
		Fields("id").
		Scan(&noticeIds)
	if err != nil {
		return err
	}
	if len(noticeIds) == 0 {
		return nil
	}

	// 查已读
	var readNoticeIds []uint64
	_ = dao.MemberNoticeRead.Ctx(ctx).
		Where("member_id", memberId).
		WhereIn("notice_id", noticeIds).
		Fields("notice_id").
		Scan(&readNoticeIds)

	readSet := make(map[uint64]bool)
	for _, id := range readNoticeIds {
		readSet[id] = true
	}

	// 批量插入未读的
	now := gtime.Now().Unix()
	for _, nid := range noticeIds {
		if readSet[nid] {
			continue
		}
		_, _ = dao.MemberNoticeRead.Ctx(ctx).Data(g.Map{
			"notice_id": nid,
			"member_id": memberId,
			"read_at":   now,
		}).Insert()
	}

	return nil
}
