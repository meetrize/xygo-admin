package uranexample

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	adminin "xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

type sUranExample struct{}

func init() {
	service.RegisterUranExample(New())
}

func New() *sUranExample {
	return &sUranExample{}
}

// List 悠然示例列表
func (s *sUranExample) List(ctx context.Context, in *adminin.UranExampleListInp) (*adminin.UranExampleListModel, error) {
	model := dao.UranExample.Ctx(ctx)
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
	var list []adminin.UranExampleListItem
	err = model.Page(in.Page, in.PageSize).OrderDesc("id").Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.UranExampleListItem{}
	}

	return &adminin.UranExampleListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    count,
		},
	}, nil
}

// View 悠然示例详情
func (s *sUranExample) View(ctx context.Context, id uint64) (*adminin.UranExampleViewModel, error) {
	var item adminin.UranExampleViewModel
	err := dao.UranExample.Ctx(ctx).Where("id", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, gerror.New("记录不存在")
	}
	return &item, nil
}

// Edit 保存悠然示例
func (s *sUranExample) Edit(ctx context.Context, in *adminin.UranExampleEditInp) error {
	data := g.Map{
		"status": in.Status,
		"sort": in.Sort,
		"tname": in.Tname,
		"timage": in.Timage,
		"ttest": in.Ttest,
	}

	if in.Id == 0 {
		// 新增（created_at/updated_at 由 GoFrame 自动维护）
		_, err := dao.UranExample.Ctx(ctx).Data(data).Insert()
		return err
	}

	// 更新（updated_at 由 GoFrame 自动维护）
	_, err := dao.UranExample.Ctx(ctx).Where("id", in.Id).Data(data).Update()
	return err
}

// Delete 删除悠然示例
func (s *sUranExample) Delete(ctx context.Context, id uint64) error {
	_, err := dao.UranExample.Ctx(ctx).Where("id", id).Delete()
	return err
}
