package middleware

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"xygo/internal/consts"
	"xygo/internal/library/contexts"
	"xygo/internal/library/token"
)

// TenantAdminAuth 租户管理员鉴权中间件
//   - 放行 /tenant/auth/login
//   - 其余 /tenant/** 需携带有效 accessToken
//   - 验证通过后注入 TenantAuthUser + TenantId 到上下文
func TenantAdminAuth(r *ghttp.Request) {
	path := r.URL.Path

	customCtx := &contexts.Context{
		Module: "tenant",
	}
	contexts.Init(r, customCtx)

	if path == "/tenant/auth/login" {
		r.Middleware.Next()
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.SetError(gerror.NewCode(consts.CodeNotAuthorized, "未登录"))
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		r.SetError(gerror.NewCode(consts.CodeNotAuthorized, "未登录"))
		return
	}

	tenantUser, err := token.ParseTenant(r.Context(), tokenStr)
	if err != nil {
		if errors.Is(err, token.ErrTokenKicked) {
			r.SetError(gerror.NewCode(consts.CodeKickedOut, "您的账号已在其他设备登录，请重新登录"))
			return
		}
		r.SetError(gerror.NewCode(consts.CodeNotAuthorized, "登录已失效，请重新登录"))
		return
	}

	contexts.SetTenantUser(r.Context(), tenantUser)
	contexts.SetTenantId(r.Context(), tenantUser.TenantId)

	r.Middleware.Next()
}
