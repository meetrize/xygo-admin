package urancontact

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	"xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

type sUranContact struct{}

func init() {
	service.RegisterUranContact(New())
}

func New() *sUranContact {
	return &sUranContact{}
}

// List 悠然联系人列表
func (s *sUranContact) List(ctx context.Context, in *adminin.UranContactListInp) (*adminin.UranContactListModel, error) {
	model := dao.UranContact.Ctx(ctx)
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
	var list []adminin.UranContactListItem
	err = model.Page(in.Page, in.PageSize).OrderDesc("id").Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		list = []adminin.UranContactListItem{}
	}

	return &adminin.UranContactListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    count,
		},
	}, nil
}

// View 悠然联系人详情
func (s *sUranContact) View(ctx context.Context, id uint64) (*adminin.UranContactViewModel, error) {
	var item adminin.UranContactViewModel
	err := dao.UranContact.Ctx(ctx).Where("id", id).Scan(&item)
	if err != nil {
		return nil, err
	}
	if item.Id == 0 {
		return nil, gerror.New("记录不存在")
	}
	return &item, nil
}

// Edit 保存悠然联系人
func (s *sUranContact) Edit(ctx context.Context, in *adminin.UranContactEditInp) error {
	data := g.Map{
		"username": in.Username,
		"phone": in.Phone,
		"age": in.Age,
		"avatar": in.Avatar,
		"sort": in.Sort,
		"remark": in.Remark,
		"switch_field": in.SwitchField,
	}

	if in.Id == 0 {
		// 新增（created_at/updated_at 由 GoFrame 自动维护）
		_, err := dao.UranContact.Ctx(ctx).Data(data).Insert()
		return err
	}

	// 更新（updated_at 由 GoFrame 自动维护）
	_, err := dao.UranContact.Ctx(ctx).Where("id", in.Id).Data(data).Update()
	return err
}

// Delete 删除悠然联系人
func (s *sUranContact) Delete(ctx context.Context, id uint64) error {
	_, err := dao.UranContact.Ctx(ctx).Where("id", id).Delete()
	return err
}
