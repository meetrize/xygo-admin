package middleware

import (
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"xygo/internal/consts"
	"xygo/internal/library/contexts"
	"xygo/internal/library/token"
)

// WsTenantAuth 租户 WebSocket 鉴权中间件
func WsTenantAuth(r *ghttp.Request) {
	customCtx := &contexts.Context{
		Module: "tenant",
	}
	contexts.Init(r, customCtx)

	tokenStr := r.Get("token").String()
	if tokenStr == "" {
		authHeader := r.Header.Get("Authorization")
		tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
	}

	if tokenStr == "" {
		r.Response.WriteStatus(401, "Unauthorized")
		r.Exit()
		return
	}

	tenantUser, err := token.ParseTenant(r.Context(), tokenStr)
	if err != nil {
		r.Response.WriteStatus(401, gerror.NewCode(consts.CodeNotAuthorized, "登录已失效").Error())
		r.Exit()
		return
	}

	contexts.SetTenantUser(r.Context(), tenantUser)
	contexts.SetTenantId(r.Context(), tenantUser.TenantId)

	r.Middleware.Next()
}
