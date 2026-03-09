// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

// =================================================================================
// 图形验证码库
// 基于 SVG 的数学验证码，无需外部依赖
// =================================================================================

package captcha

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/util/guid"

	"xygo/internal/library/cache"
)

// 验证码缓存前缀和过期时间
const (
	cachePrefix = "captcha:"
	expireTime  = 5 * time.Minute
)

// Generate 生成数学验证码，返回 (captchaId, base64SvgImage)
func Generate(ctx context.Context) (string, string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 随机两个 1~20 的数做加法
	a := r.Intn(20) + 1
	b := r.Intn(20) + 1
	answer := fmt.Sprintf("%d", a+b)
	expression := fmt.Sprintf("%d + %d = ?", a, b)

	// 生成唯一 ID
	captchaId := guid.S()

	// 存入缓存
	_ = cache.Instance().Set(ctx, cachePrefix+captchaId, answer, expireTime)

	// 生成 SVG
	svg := renderSVG(expression, r)
	imgBase64 := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(svg))

	return captchaId, imgBase64
}

// Verify 校验验证码，校验后立即删除（一次性）
func Verify(ctx context.Context, captchaId, userAnswer string) bool {
	if captchaId == "" || userAnswer == "" {
		return false
	}

	key := cachePrefix + captchaId
	val, err := cache.Instance().Get(ctx, key)
	if err != nil || val.IsEmpty() {
		return false
	}

	// 校验后立即删除，防止重放
	_, _ = cache.Instance().Remove(ctx, key)

	return val.String() == userAnswer
}

// renderSVG 生成带干扰线的 SVG 验证码图片
func renderSVG(text string, r *rand.Rand) string {
	width := 150
	height := 50

	// 生成干扰线
	lines := ""
	for i := 0; i < 4; i++ {
		x1 := r.Intn(width)
		y1 := r.Intn(height)
		x2 := r.Intn(width)
		y2 := r.Intn(height)
		colors := []string{"#d1d9e6", "#b0bec5", "#90a4ae", "#78909c"}
		lines += fmt.Sprintf(
			`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="1" opacity="0.6"/>`,
			x1, y1, x2, y2, colors[i%len(colors)],
		)
	}

	// 生成干扰点
	dots := ""
	for i := 0; i < 20; i++ {
		cx := r.Intn(width)
		cy := r.Intn(height)
		dots += fmt.Sprintf(
			`<circle cx="%d" cy="%d" r="1" fill="#b0bec5" opacity="0.5"/>`,
			cx, cy,
		)
	}

	// 文字微旋转
	rotate := r.Intn(5) - 2 // -2 ~ 2 度

	svg := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
<rect width="%d" height="%d" rx="12" fill="#f0f3f8"/>
%s
%s
<text x="75" y="34" font-size="24" font-weight="bold" text-anchor="middle"
      fill="#32325d" font-family="Arial, sans-serif"
      transform="rotate(%d, 75, 30)">%s</text>
</svg>`, width, height, width, height, lines, dots, rotate, text)

	return svg
}
