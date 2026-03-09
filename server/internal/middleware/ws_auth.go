// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package middleware

import (
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"xygo/internal/consts"
	"xygo/internal/library/contexts"
	"xygo/internal/library/token"
)

// WsAuth WebSocket 鉴权中间件
// - 支持从 query param (?token=xxx) 或 Authorization header 获取 token
// - 验证通过后将用户信息写入 context，供 WsHandler 使用
func WsAuth(r *ghttp.Request) {
	// 初始化自定义上下文
	customCtx := &contexts.Context{
		Module: "admin",
	}
	contexts.Init(r, customCtx)

	// 优先从 query param 获取 token（WebSocket 握手通常通过 URL 参数传递）
	tokenStr := r.Get("token").String()

	// 回退到 Authorization header
	if tokenStr == "" {
		authHeader := r.Header.Get("Authorization")
		tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
	}

	if tokenStr == "" {
		r.Response.WriteStatus(401, "Unauthorized")
		r.Exit()
		return
	}

	// 解析 Token 并获取完整用户信息
	authUser, err := token.Parse(r.Context(), tokenStr)
	if err != nil {
		r.Response.WriteStatus(401, gerror.NewCode(consts.CodeNotAuthorized, "登录已失效").Error())
		r.Exit()
		return
	}

	// 将用户信息注入到上下文中
	contexts.SetUser(r.Context(), authUser)

	r.Middleware.Next()
}
