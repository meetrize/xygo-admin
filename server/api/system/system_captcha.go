// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package system

import "github.com/gogf/gf/v2/frame/g"

// ==================== 点选验证码（公共接口，前后台通用） ====================

// ClickCaptchaReq 获取点选验证码
type ClickCaptchaReq struct {
	g.Meta `path:"/captcha/click" method:"get" tags:"公共" summary:"获取点选验证码"`
}

// ClickCaptchaRes 获取点选验证码响应
type ClickCaptchaRes struct {
	Id     string   `json:"id"`
	Text   []string `json:"text"`
	Base64 string   `json:"base64"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

// CheckClickCaptchaReq 校验点选验证码
type CheckClickCaptchaReq struct {
	g.Meta `path:"/captcha/checkClick" method:"post" tags:"公共" summary:"校验点选验证码"`
	Id     string `json:"id" v:"required#请提供验证码ID"`
	Info   string `json:"info" v:"required#请提供点击坐标"`
}

// CheckClickCaptchaRes 校验响应
type CheckClickCaptchaRes struct{}
