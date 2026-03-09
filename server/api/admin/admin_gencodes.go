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
	"github.com/gogf/gf/v2/frame/g"
	"xygo/internal/model/input/adminin"
)

// ==================== 选项 ====================

type GenCodesSelectsReq struct {
	g.Meta `path:"/admin/genCodes/selects" method:"get" tags:"GenCodes" summary:"生成器选项"`
	adminin.GenCodesSelectsInp
}
type GenCodesSelectsRes struct {
	*adminin.GenCodesSelectsModel
}

// ==================== 数据库表选项 ====================

type GenCodesTableSelectReq struct {
	g.Meta `path:"/admin/genCodes/tableSelect" method:"get" tags:"GenCodes" summary:"数据库表列表"`
	adminin.GenCodesTableSelectInp
}
type GenCodesTableSelectRes struct {
	*adminin.GenCodesTableSelectModel
}

// ==================== 表字段列表 ====================

type GenCodesColumnListReq struct {
	g.Meta `path:"/admin/genCodes/columnList" method:"get" tags:"GenCodes" summary:"表字段列表"`
	adminin.GenCodesColumnListInp
}
type GenCodesColumnListRes struct {
	*adminin.GenCodesColumnListModel
}

// ==================== 列表 ====================

type GenCodesListReq struct {
	g.Meta `path:"/admin/genCodes/list" method:"get" tags:"GenCodes" summary:"生成记录列表"`
	adminin.GenCodesListInp
}
type GenCodesListRes struct {
	*adminin.GenCodesListModel
}

// ==================== 详情 ====================

type GenCodesViewReq struct {
	g.Meta `path:"/admin/genCodes/view" method:"get" tags:"GenCodes" summary:"生成配置详情"`
	adminin.GenCodesViewInp
}
type GenCodesViewRes struct {
	*adminin.GenCodesViewModel
}

// ==================== 保存 ====================

type GenCodesEditReq struct {
	g.Meta `path:"/admin/genCodes/edit" method:"post" tags:"GenCodes" summary:"保存生成配置"`
	adminin.GenCodesEditInp
}
type GenCodesEditRes struct {
	*adminin.GenCodesEditModel
}

// ==================== 删除 ====================

type GenCodesDeleteReq struct {
	g.Meta `path:"/admin/genCodes/delete" method:"post" tags:"GenCodes" summary:"删除生成配置"`
	adminin.GenCodesDeleteInp
}
type GenCodesDeleteRes struct{}

// ==================== 预览 ====================

type GenCodesPreviewReq struct {
	g.Meta `path:"/admin/genCodes/preview" method:"post" tags:"GenCodes" summary:"预览生成代码"`
	adminin.GenCodesPreviewInp
}
type GenCodesPreviewRes struct {
	*adminin.GenCodesPreviewModel
}

// ==================== 生成 ====================

type GenCodesBuildReq struct {
	g.Meta `path:"/admin/genCodes/build" method:"post" tags:"GenCodes" summary:"执行生成"`
	adminin.GenCodesBuildInp
}
type GenCodesBuildRes struct{}

// ==================== 发布前端文件 ====================

type GenCodesPublishFrontendReq struct {
	g.Meta `path:"/admin/genCodes/publishFrontend" method:"post" tags:"GenCodes" summary:"发布前端文件（从临时目录移到正式目录）"`
}
type GenCodesPublishFrontendRes struct{}

// ==================== 同步字段到数据库 ====================

type GenCodesSyncFieldsReq struct {
	g.Meta `path:"/admin/genCodes/syncFields" method:"post" tags:"GenCodes" summary:"预览字段变更"`
	adminin.GenCodesSyncFieldsInp
}
type GenCodesSyncFieldsRes struct {
	*adminin.GenCodesSyncFieldsModel
}

type GenCodesExecuteDDLReq struct {
	g.Meta `path:"/admin/genCodes/executeDDL" method:"post" tags:"GenCodes" summary:"执行字段同步"`
	adminin.GenCodesExecuteDDLInp
}
type GenCodesExecuteDDLRes struct{}

// ==================== 创建数据表 ====================

type GenCodesCreateTableReq struct {
	g.Meta `path:"/admin/genCodes/createTable" method:"post" tags:"GenCodes" summary:"创建数据表"`
	adminin.GenCodesCreateTableInp
}
type GenCodesCreateTableRes struct {
	*adminin.GenCodesCreateTableModel
}
