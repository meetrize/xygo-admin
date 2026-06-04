package urantestable

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	adminin "xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

type sUranTestable struct{}

func init() {
	service.RegisterUranTestable(New())
}

func New() *sUranTestable {
	return &sUranTestable{}
}

// List 测试记帐列表
func (s *sUranTestable) List(ctx context.Context, in *adminin.UranTestableListInp) (*adminin.UranTestableListModel, error) {
	model := dao.UranTestable.Ctx(ctx)
	if in.Status != nil {
		model = model.Where("status", *in.Status)
	}
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
	var list []adminin.UranTestableListItem
	err = model.Page(in.Page, in.PageSize).OrderDesc("id").Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.UranTestableListItem{}
	}

	return &adminin.UranTestableListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    count,
		},
	}, nil
}

// View 测试记帐详情
func (s *sUranTestable) View(ctx context.Context, id uint64) (*adminin.UranTestableViewModel, error) {
	var item adminin.UranTestableViewModel
	err := dao.UranTestable.Ctx(ctx).Where("id", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, gerror.New("记录不存在")
	}
	return &item, nil
}

// Edit 保存测试记帐
func (s *sUranTestable) Edit(ctx context.Context, in *adminin.UranTestableEditInp) error {
	data := g.Map{
		"status": in.Status,
		"sort": in.Sort,
		"stuff": in.Stuff,
		"price": in.Price,
	}

	if in.Id == 0 {
		// 新增（created_at/updated_at 由 GoFrame 自动维护）
		_, err := dao.UranTestable.Ctx(ctx).Data(data).Insert()
		return err
	}

	// 更新（updated_at 由 GoFrame 自动维护）
	_, err := dao.UranTestable.Ctx(ctx).Where("id", in.Id).Data(data).Update()
	return err
}

// Delete 删除测试记帐
func (s *sUranTestable) Delete(ctx context.Context, id uint64) error {
	_, err := dao.UranTestable.Ctx(ctx).Where("id", id).Delete()
	return err
}
