// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xygo/internal/model/input/adminin"
)

type (
	IUranContact interface {
		// List 悠然联系人列表
		List(ctx context.Context, in *adminin.UranContactListInp) (*adminin.UranContactListModel, error)
		// View 悠然联系人详情
		View(ctx context.Context, id uint64) (*adminin.UranContactViewModel, error)
		// Edit 保存悠然联系人
		Edit(ctx context.Context, in *adminin.UranContactEditInp) error
		// Delete 删除悠然联系人
		Delete(ctx context.Context, id uint64) error
	}
)

var (
	localUranContact IUranContact
)

func UranContact() IUranContact {
	if localUranContact == nil {
		panic("implement not found for interface IUranContact, forgot register?")
	}
	return localUranContact
}

func RegisterUranContact(i IUranContact) {
	localUranContact = i
}
