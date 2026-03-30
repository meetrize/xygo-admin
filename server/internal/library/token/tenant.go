package token

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"

	"xygo/internal/library/cache"
	"xygo/internal/model"
)

// TenantClaims 租户管理员 JWT 载荷
type TenantClaims struct {
	*model.TenantAuthUser
	jwt.RegisteredClaims
}

// GenerateTenant 租户管理员登录生成 token
func GenerateTenant(ctx context.Context, user model.TenantAuthUser) (accessToken string, expiresIn int64, err error) {
	cfg := getConfig(ctx)

	now := time.Now()
	expireAt := now.Add(time.Duration(cfg.Expires) * time.Second)

	claims := &TenantClaims{
		TenantAuthUser: &user,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = tok.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", 0, err
	}

	var (
		authKey  = GetAuthKey(accessToken)
		tokenKey = GetTokenKey(AppTenant, authKey)
		bindKey  = GetBindKey(AppTenant, user.Id)
		duration = time.Second * gconv.Duration(cfg.Expires)
	)

	if !cfg.MultiLogin {
		kickOldSession(ctx, AppTenant, bindKey)
	}

	tokenMeta := &TokenMeta{
		ExpireAt:     expireAt.Unix(),
		RefreshAt:    now.Unix(),
		RefreshCount: 0,
	}

	if err = cache.Instance().Set(ctx, tokenKey, tokenMeta, duration); err != nil {
		return "", 0, err
	}
	if err = cache.Instance().Set(ctx, bindKey, tokenKey, duration); err != nil {
		return "", 0, err
	}

	return accessToken, cfg.Expires, nil
}

// ParseTenant 解析租户管理员 Token
func ParseTenant(ctx context.Context, accessToken string) (*model.TenantAuthUser, error) {
	cfg := getConfig(ctx)

	parsed, err := jwt.ParseWithClaims(accessToken, &TenantClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})
	if err != nil || !parsed.Valid {
		return nil, err
	}

	claims, ok := parsed.Claims.(*TenantClaims)
	if !ok || claims.TenantAuthUser == nil {
		return nil, jwt.ErrTokenMalformed
	}

	var (
		authKey  = GetAuthKey(accessToken)
		tokenKey = GetTokenKey(AppTenant, authKey)
	)

	if err = validateTokenMeta(ctx, tokenKey); err != nil {
		return nil, err
	}

	return claims.TenantAuthUser, nil
}

// DeleteTenant 删除租户管理员 Token（退出登录）
func DeleteTenant(ctx context.Context, accessToken string) error {
	return DeleteByApp(ctx, AppTenant, accessToken)
}
