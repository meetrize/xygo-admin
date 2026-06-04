// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	adminin "xygo/internal/model/input/adminin"
)

type (
	IUranTestable interface {
		// List 测试记帐列表
		List(ctx context.Context, in *adminin.UranTestableListInp) (*adminin.UranTestableListModel, error)
		// View 测试记帐详情
		View(ctx context.Context, id uint64) (*adminin.UranTestableViewModel, error)
		// Edit 保存测试记帐
		Edit(ctx context.Context, in *adminin.UranTestableEditInp) error
		// Delete 删除测试记帐
		Delete(ctx context.Context, id uint64) error
	}
)

var (
	localUranTestable IUranTestable
)

func UranTestable() IUranTestable {
	if localUranTestable == nil {
		panic("implement not found for interface IUranTestable, forgot register?")
	}
	return localUranTestable
}

func RegisterUranTestable(i IUranTestable) {
	localUranTestable = i
}
