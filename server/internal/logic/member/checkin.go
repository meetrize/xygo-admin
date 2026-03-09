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
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"xygo/internal/dao"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/memberin"
)

type sMemberCheckin struct{}

func NewMemberCheckin() *sMemberCheckin {
	return &sMemberCheckin{}
}

// GetCheckinInfo 获取签到信息（7天日历+连续天数+今日状态）
func (s *sMemberCheckin) GetCheckinInfo(ctx context.Context, memberId uint64) (out *memberin.CheckinInfoOutput, err error) {
	out = &memberin.CheckinInfoOutput{}

	// 查最近7天的签到记录（checkin_date 为 Unix 时间戳）
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekAgoStart := todayStart.AddDate(0, 0, -6) // 含今天共7天

	var records []entity.MemberCheckin
	err = dao.MemberCheckin.Ctx(ctx).
		Where("member_id", memberId).
		WhereGTE("checkin_date", weekAgoStart.Unix()).
		WhereLTE("checkin_date", todayStart.AddDate(0, 0, 1).Unix()-1).
		OrderAsc("checkin_date").
		Scan(&records)
	if err != nil {
		return nil, err
	}

	// 构建已签到日期 map（按日期字符串索引）
	checkedMap := make(map[string]*entity.MemberCheckin)
	for i := range records {
		t := time.Unix(int64(records[i].CheckinDate), 0)
		dateStr := t.Format("2006-01-02")
		checkedMap[dateStr] = &records[i]
	}

	// 构建7天日历
	out.WeekDays = make([]memberin.CheckinDayItem, 7)
	for i := 0; i < 7; i++ {
		day := todayStart.AddDate(0, 0, i-6)
		dateStr := day.Format("2006-01-02")
		item := memberin.CheckinDayItem{
			Date: dateStr,
		}
		if rec, ok := checkedMap[dateStr]; ok {
			item.Checked = true
			item.Score = rec.Score
		}
		out.WeekDays[i] = item
	}

	// 今日是否已签到
	todayStr := todayStart.Format("2006-01-02")
	if rec, ok := checkedMap[todayStr]; ok {
		out.TodayChecked = true
		out.TodayScore = rec.Score
	}

	// 查连续签到天数（从最近一条向前连续计算）
	out.ContinuousDays = s.calcContinuousDays(ctx, memberId)

	return out, nil
}

// DoCheckin 执行签到
func (s *sMemberCheckin) DoCheckin(ctx context.Context, memberId uint64) (out *memberin.DoCheckinOutput, err error) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrowStart := todayStart.AddDate(0, 0, 1)

	// 检查今日是否已签到（时间戳在今天范围内）
	count, err := dao.MemberCheckin.Ctx(ctx).
		Where("member_id", memberId).
		WhereGTE("checkin_date", todayStart.Unix()).
		WhereLT("checkin_date", tomorrowStart.Unix()).
		Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("今日已签到，请明天再来")
	}

	// 计算连续天数
	continuousDays := s.calcContinuousDays(ctx, memberId)

	// 检查昨天是否签到了（决定是否连续）
	yesterdayStart := todayStart.AddDate(0, 0, -1)
	yesterdayCount, _ := dao.MemberCheckin.Ctx(ctx).
		Where("member_id", memberId).
		WhereGTE("checkin_date", yesterdayStart.Unix()).
		WhereLT("checkin_date", todayStart.Unix()).
		Count()
	if yesterdayCount > 0 {
		continuousDays++
	} else {
		continuousDays = 1
	}

	// 随机积分 1-5，连续签到加成
	score := rand.Intn(5) + 1
	if continuousDays >= 7 {
		score += 3 // 连续7天额外+3
	} else if continuousDays >= 3 {
		score += 1 // 连续3天额外+1
	}

	// 插入签到记录
	_, err = dao.MemberCheckin.Ctx(ctx).Data(g.Map{
		"member_id":       memberId,
		"checkin_date":    gtime.Now().Unix(),
		"score":           score,
		"continuous_days": continuousDays,
	}).Insert()
	if err != nil {
		return nil, err
	}

	// 更新会员积分 + 写积分流水
	err = s.addScore(ctx, memberId, score, "每日签到奖励")
	if err != nil {
		g.Log().Warningf(ctx, "签到积分写入失败: %v", err)
	}

	return &memberin.DoCheckinOutput{
		Score:          score,
		ContinuousDays: continuousDays,
	}, nil
}

// calcContinuousDays 计算连续签到天数（从昨天/今天开始往前数）
func (s *sMemberCheckin) calcContinuousDays(ctx context.Context, memberId uint64) int {
	var records []entity.MemberCheckin
	_ = dao.MemberCheckin.Ctx(ctx).
		Where("member_id", memberId).
		OrderDesc("checkin_date").
		Limit(30). // 最多查30天
		Scan(&records)

	if len(records) == 0 {
		return 0
	}

	days := 0
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	expect := today

	for _, rec := range records {
		recTime := time.Unix(int64(rec.CheckinDate), 0)
		recDate := time.Date(recTime.Year(), recTime.Month(), recTime.Day(), 0, 0, 0, 0, recTime.Location())
		if recDate.Equal(expect) {
			days++
			expect = expect.AddDate(0, 0, -1)
		} else if recDate.Equal(expect.AddDate(0, 0, -1)) && days == 0 {
			// 今天还没签但昨天签了
			expect = recDate
			days++
			expect = expect.AddDate(0, 0, -1)
		} else {
			break
		}
	}
	return days
}

// addScore 增加积分 + 写积分流水
func (s *sMemberCheckin) addScore(ctx context.Context, memberId uint64, score int, memo string) error {
	// 查当前积分
	var member entity.Member
	err := dao.Member.Ctx(ctx).Where("id", memberId).Scan(&member)
	if err != nil {
		return err
	}

	before := member.Score
	after := before + score

	// 更新积分
	_, err = dao.Member.Ctx(ctx).Where("id", memberId).Data(g.Map{
		"score": after,
	}).Update()
	if err != nil {
		return err
	}

	// 写积分流水
	_, err = dao.MemberScoreLog.Ctx(ctx).Data(g.Map{
		"member_id":  memberId,
		"score":      score,
		"before":     before,
		"after":      after,
		"memo": memo,
	}).Insert()

	return err
}
