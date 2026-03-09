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

// ===================== 通知管理（管理端） =====================

type NoticeListReq struct {
	g.Meta `path:"/admin/notice/list" method:"post" tags:"AdminNotice" summary:"通知列表"`
	adminin.NoticeListInp
}
type NoticeListRes struct {
	*adminin.NoticeListModel
}

type NoticeEditReq struct {
	g.Meta `path:"/admin/notice/edit" method:"post" tags:"AdminNotice" summary:"编辑通知"`
	adminin.NoticeEditInp
}
type NoticeEditRes struct {
	Id uint64 `json:"id"`
}

type NoticeDeleteReq struct {
	g.Meta `path:"/admin/notice/delete" method:"post" tags:"AdminNotice" summary:"删除通知"`
	adminin.NoticeDeleteInp
}
type NoticeDeleteRes struct{}

// ===================== 用户端消息 =====================

type NoticePullReq struct {
	g.Meta `path:"/admin/notice/pull" method:"get" tags:"AdminNotice" summary:"拉取消息"`
}
type NoticePullRes struct {
	*adminin.PullMessagesModel
}

type NoticeUnreadCountReq struct {
	g.Meta `path:"/admin/notice/unreadCount" method:"get" tags:"AdminNotice" summary:"未读数"`
}
type NoticeUnreadCountRes struct {
	List  []adminin.UnreadCountItem `json:"list"`
	Total int                       `json:"total"`
}

type NoticeReadReq struct {
	g.Meta `path:"/admin/notice/read" method:"post" tags:"AdminNotice" summary:"标记已读"`
	adminin.ReadNoticeInp
}
type NoticeReadRes struct{}

type NoticeReadAllReq struct {
	g.Meta `path:"/admin/notice/readAll" method:"post" tags:"AdminNotice" summary:"全部已读"`
	adminin.ReadAllNoticeInp
}
type NoticeReadAllRes struct{}
