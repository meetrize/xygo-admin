package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"xygo/internal/library/contexts"
)

// TenantResolve 租户识别中间件
//
// 从请求中解析当前租户ID并注入 context，优先级：
//  1. Header: X-Tenant-Id（前端主动传递）
//  2. 已登录用户的 AuthUser.TenantId（从 Token 中解析）
//
// 注意：域名匹配功能需安装多租户扩展后由扩展覆盖此文件实现
func TenantResolve(r *ghttp.Request) {
	var tenantId uint64
	ctx := r.GetCtx()

	headerTenantId := r.GetHeader("X-Tenant-Id")
	if headerTenantId != "" {
		tenantId = gconv.Uint64(headerTenantId)
	}

	if tenantId == 0 {
		user := contexts.GetUser(ctx)
		if user != nil && user.TenantId > 0 {
			tenantId = user.TenantId
		}
	}

	if tenantId > 0 {
		contexts.SetTenantId(ctx, tenantId)
		g.Log().Debugf(ctx, "TenantResolve: tenantId=%d", tenantId)
	}

	r.Middleware.Next()
}
