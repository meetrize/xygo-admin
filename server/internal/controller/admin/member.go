// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package admin

import (
	"context"

	api "xygo/api/admin"
	"xygo/internal/model/input/adminin"
	"xygo/internal/model/input/form"
	"xygo/internal/service"
)

// MemberList 会员列表
func (c *ControllerV1) MemberList(ctx context.Context, req *api.MemberListReq) (res *api.MemberListRes, err error) {
	list, total, err := service.AdminMember().List(ctx, &req.MemberListInp)
	if err != nil {
		return nil, err
	}

	res = new(api.MemberListRes)
	res.MemberListModel = &adminin.MemberListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return
}

// MemberDetail 会员详情
func (c *ControllerV1) MemberDetail(ctx context.Context, req *api.MemberDetailReq) (res *api.MemberDetailRes, err error) {
	out, err := service.AdminMember().Detail(ctx, &req.MemberDetailInp)
	if err != nil {
		return nil, err
	}

	res = new(api.MemberDetailRes)
	res.MemberDetailModel = out
	return
}

// MemberAdd 添加会员
func (c *ControllerV1) MemberAdd(ctx context.Context, req *api.MemberAddReq) (res *api.MemberAddRes, err error) {
	out, err := service.AdminMember().Add(ctx, &req.MemberAddInp)
	if err != nil {
		return nil, err
	}

	res = new(api.MemberAddRes)
	res.MemberAddModel = out
	return
}

// MemberEdit 编辑会员
func (c *ControllerV1) MemberEdit(ctx context.Context, req *api.MemberEditReq) (res *api.MemberEditRes, err error) {
	err = service.AdminMember().Edit(ctx, &req.MemberEditInp)
	if err != nil {
		return nil, err
	}
	res = &api.MemberEditRes{}
	return
}

// MemberDelete 删除会员
func (c *ControllerV1) MemberDelete(ctx context.Context, req *api.MemberDeleteReq) (res *api.MemberDeleteRes, err error) {
	err = service.AdminMember().Delete(ctx, &req.MemberDeleteInp)
	if err != nil {
		return nil, err
	}
	res = &api.MemberDeleteRes{}
	return
}

// MemberStatus 修改会员状态
func (c *ControllerV1) MemberStatus(ctx context.Context, req *api.MemberStatusReq) (res *api.MemberStatusRes, err error) {
	err = service.AdminMember().Status(ctx, &req.MemberStatusInp)
	if err != nil {
		return nil, err
	}
	res = &api.MemberStatusRes{}
	return
}

// MemberResetPassword 重置会员密码
func (c *ControllerV1) MemberResetPassword(ctx context.Context, req *api.MemberResetPasswordReq) (res *api.MemberResetPasswordRes, err error) {
	err = service.AdminMember().ResetPassword(ctx, &req.MemberResetPasswordInp)
	if err != nil {
		return nil, err
	}
	res = &api.MemberResetPasswordRes{}
	return
}

// MemberGroupOptions 会员分组选项
func (c *ControllerV1) MemberGroupOptions(ctx context.Context, req *api.MemberGroupOptionsReq) (res *api.MemberGroupOptionsRes, err error) {
	out, err := service.AdminMember().GroupOptions(ctx)
	if err != nil {
		return nil, err
	}

	res = new(api.MemberGroupOptionsRes)
	res.MemberGroupOptionsModel = out
	return
}

// ===================== 会员分组管理 =====================

// MemberGroupList 会员分组列表
func (c *ControllerV1) MemberGroupList(ctx context.Context, req *api.MemberGroupListReq) (res *api.MemberGroupListRes, err error) {
	list, total, err := service.AdminMemberGroup().List(ctx, &req.MemberGroupListInp)
	if err != nil {
		return nil, err
	}

	res = new(api.MemberGroupListRes)
	res.MemberGroupListModel = &adminin.MemberGroupListModel{
		List: list,
		PageRes: form.PageRes{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return
}

// MemberGroupSave 保存会员分组
func (c *ControllerV1) MemberGroupSave(ctx context.Context, req *api.MemberGroupSaveReq) (res *api.MemberGroupSaveRes, err error) {
	id, err := service.AdminMemberGroup().Save(ctx, &req.MemberGroupSaveInp)
	if err != nil {
		return nil, err
	}

	res = &api.MemberGroupSaveRes{Id: id}
	return
}

// MemberGroupDelete 删除会员分组
func (c *ControllerV1) MemberGroupDelete(ctx context.Context, req *api.MemberGroupDeleteReq) (res *api.MemberGroupDeleteRes, err error) {
	err = service.AdminMemberGroup().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.MemberGroupDeleteRes{}
	return
}

// ===================== 会员菜单管理 =====================

// MemberMenuTree 会员菜单树
func (c *ControllerV1) MemberMenuTree(ctx context.Context, req *api.MemberMenuTreeReq) (res *api.MemberMenuTreeRes, err error) {
	tree, err := service.AdminMemberMenu().Tree(ctx, &req.MemberMenuTreeInp)
	if err != nil {
		return nil, err
	}

	res = new(api.MemberMenuTreeRes)
	res.MemberMenuTreeModel = &adminin.MemberMenuTreeModel{
		List: tree,
	}
	return
}

// MemberMenuSave 保存会员菜单
func (c *ControllerV1) MemberMenuSave(ctx context.Context, req *api.MemberMenuSaveReq) (res *api.MemberMenuSaveRes, err error) {
	id, err := service.AdminMemberMenu().Save(ctx, &req.MemberMenuSaveInp)
	if err != nil {
		return nil, err
	}

	res = &api.MemberMenuSaveRes{Id: id}
	return
}

// MemberMenuDelete 删除会员菜单
func (c *ControllerV1) MemberMenuDelete(ctx context.Context, req *api.MemberMenuDeleteReq) (res *api.MemberMenuDeleteRes, err error) {
	err = service.AdminMemberMenu().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.MemberMenuDeleteRes{}
	return
}
