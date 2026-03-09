// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package member

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	"xygo/api/member"
	"xygo/internal/model/input/memberin"
	"xygo/internal/service"
)

// Login 会员登录
func (c *ControllerV1) Login(ctx context.Context, req *member.LoginReq) (res *member.LoginRes, err error) {
	input := &memberin.LoginInput{
		Username:  req.Username,
		Password:  req.Password,
		Captcha:   req.Captcha,
		CaptchaId: req.CaptchaId,
	}

	output, err := service.MemberAuth().Login(ctx, input)
	if err != nil {
		return nil, err
	}

	return &member.LoginRes{
		Token:     output.Token,
		ExpiresIn: output.ExpiresIn,
	}, nil
}

// Register 会员注册
func (c *ControllerV1) Register(ctx context.Context, req *member.RegisterReq) (res *member.RegisterRes, err error) {
	input := &memberin.RegisterInput{
		Username: req.Username,
		Password: req.Password,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Code:     req.Code,
	}

	output, err := service.MemberAuth().Register(ctx, input)
	if err != nil {
		return nil, err
	}

	return &member.RegisterRes{
		Id: output.Id,
	}, nil
}

// Logout 会员退出登录
func (c *ControllerV1) Logout(ctx context.Context, req *member.LogoutReq) (res *member.LogoutRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	if r != nil {
		tokenStr := r.Header.Get("Xy-User-Token")
		if tokenStr != "" {
			_ = service.MemberAuth().Logout(ctx, tokenStr)
		}
	}
	return &member.LogoutRes{}, nil
}

// 注：验证码接口已统一到公共 /captcha/click 和 /captcha/checkClick
