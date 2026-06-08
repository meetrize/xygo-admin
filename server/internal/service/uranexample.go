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
	IUranExample interface {
		// List 悠然示例列表
		List(ctx context.Context, in *adminin.UranExampleListInp) (*adminin.UranExampleListModel, error)
		// View 悠然示例详情
		View(ctx context.Context, id uint64) (*adminin.UranExampleViewModel, error)
		// Edit 保存悠然示例
		Edit(ctx context.Context, in *adminin.UranExampleEditInp) error
		// Delete 删除悠然示例
		Delete(ctx context.Context, id uint64) error
	}
)

var (
	localUranExample IUranExample
)

func UranExample() IUranExample {
	if localUranExample == nil {
		panic("implement not found for interface IUranExample, forgot register?")
	}
	return localUranExample
}

func RegisterUranExample(i IUranExample) {
	localUranExample = i
}
