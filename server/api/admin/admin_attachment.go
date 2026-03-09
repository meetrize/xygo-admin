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

// ===================== 附件列表 =====================

type AttachmentListReq struct {
	g.Meta `path:"/admin/attachment/list" method:"get" tags:"AdminAttachment" summary:"附件列表"`
	adminin.AttachmentListInp
}

type AttachmentListRes struct {
	*adminin.AttachmentListModel
}

// ===================== 删除附件 =====================

type AttachmentDeleteReq struct {
	g.Meta `path:"/admin/attachment/delete" method:"post" tags:"AdminAttachment" summary:"删除附件"`
	adminin.AttachmentDeleteInp
}

type AttachmentDeleteRes struct{}
