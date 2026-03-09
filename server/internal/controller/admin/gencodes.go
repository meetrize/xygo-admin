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
	"xygo/internal/service"
)

// GenCodesSelects 获取选项
func (c *ControllerV1) GenCodesSelects(ctx context.Context, req *api.GenCodesSelectsReq) (res *api.GenCodesSelectsRes, err error) {
	result, err := service.GenCodes().Selects(ctx)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesSelectsRes{GenCodesSelectsModel: result}, nil
}

// GenCodesTableSelect 获取数据库表列表
func (c *ControllerV1) GenCodesTableSelect(ctx context.Context, req *api.GenCodesTableSelectReq) (res *api.GenCodesTableSelectRes, err error) {
	result, err := service.GenCodes().TableSelect(ctx)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesTableSelectRes{GenCodesTableSelectModel: result}, nil
}

// GenCodesColumnList 获取表字段列表
func (c *ControllerV1) GenCodesColumnList(ctx context.Context, req *api.GenCodesColumnListReq) (res *api.GenCodesColumnListRes, err error) {
	result, err := service.GenCodes().ColumnList(ctx, &req.GenCodesColumnListInp)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesColumnListRes{GenCodesColumnListModel: result}, nil
}

// GenCodesList 生成记录列表
func (c *ControllerV1) GenCodesList(ctx context.Context, req *api.GenCodesListReq) (res *api.GenCodesListRes, err error) {
	result, err := service.GenCodes().List(ctx, &req.GenCodesListInp)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesListRes{GenCodesListModel: result}, nil
}

// GenCodesView 查看详情
func (c *ControllerV1) GenCodesView(ctx context.Context, req *api.GenCodesViewReq) (res *api.GenCodesViewRes, err error) {
	result, err := service.GenCodes().View(ctx, &req.GenCodesViewInp)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesViewRes{GenCodesViewModel: result}, nil
}

// GenCodesEdit 保存配置
func (c *ControllerV1) GenCodesEdit(ctx context.Context, req *api.GenCodesEditReq) (res *api.GenCodesEditRes, err error) {
	result, err := service.GenCodes().Edit(ctx, &req.GenCodesEditInp)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesEditRes{GenCodesEditModel: result}, nil
}

// GenCodesDelete 删除配置
func (c *ControllerV1) GenCodesDelete(ctx context.Context, req *api.GenCodesDeleteReq) (res *api.GenCodesDeleteRes, err error) {
	err = service.GenCodes().Delete(ctx, &req.GenCodesDeleteInp)
	return &api.GenCodesDeleteRes{}, err
}

// GenCodesPreview 预览代码
func (c *ControllerV1) GenCodesPreview(ctx context.Context, req *api.GenCodesPreviewReq) (res *api.GenCodesPreviewRes, err error) {
	result, err := service.GenCodes().Preview(ctx, &req.GenCodesPreviewInp)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesPreviewRes{GenCodesPreviewModel: result}, nil
}

// GenCodesBuild 执行生成
func (c *ControllerV1) GenCodesBuild(ctx context.Context, req *api.GenCodesBuildReq) (res *api.GenCodesBuildRes, err error) {
	err = service.GenCodes().Build(ctx, &req.GenCodesBuildInp)
	return &api.GenCodesBuildRes{}, err
}

// GenCodesPublishFrontend 发布前端文件
func (c *ControllerV1) GenCodesPublishFrontend(ctx context.Context, req *api.GenCodesPublishFrontendReq) (res *api.GenCodesPublishFrontendRes, err error) {
	err = service.GenCodes().PublishFrontend(ctx)
	return &api.GenCodesPublishFrontendRes{}, err
}

// GenCodesSyncFields 预览字段变更（对比设计器与数据库）
func (c *ControllerV1) GenCodesSyncFields(ctx context.Context, req *api.GenCodesSyncFieldsReq) (res *api.GenCodesSyncFieldsRes, err error) {
	result, err := service.GenCodes().SyncFields(ctx, &req.GenCodesSyncFieldsInp)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesSyncFieldsRes{GenCodesSyncFieldsModel: result}, nil
}

// GenCodesExecuteDDL 执行字段同步DDL
func (c *ControllerV1) GenCodesExecuteDDL(ctx context.Context, req *api.GenCodesExecuteDDLReq) (res *api.GenCodesExecuteDDLRes, err error) {
	err = service.GenCodes().ExecuteDDL(ctx, &req.GenCodesExecuteDDLInp)
	return &api.GenCodesExecuteDDLRes{}, err
}

// GenCodesCreateTable 创建数据表
func (c *ControllerV1) GenCodesCreateTable(ctx context.Context, req *api.GenCodesCreateTableReq) (res *api.GenCodesCreateTableRes, err error) {
	result, err := service.GenCodes().CreateTable(ctx, &req.GenCodesCreateTableInp)
	if err != nil {
		return nil, err
	}
	return &api.GenCodesCreateTableRes{GenCodesCreateTableModel: result}, nil
}
