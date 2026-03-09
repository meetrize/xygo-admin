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
// 会员认证逻辑层
// =================================================================================

package member

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/bcrypt"

	"xygo/internal/consts"
	"xygo/internal/dao"
	"xygo/internal/library/captcha"
	"xygo/internal/library/token"
	"xygo/internal/model"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/memberin"
)

type sMemberAuth struct{}

// NewMemberAuth 构造会员认证服务
func NewMemberAuth() *sMemberAuth {
	return &sMemberAuth{}
}

// Login 会员登录
func (s *sMemberAuth) Login(ctx context.Context, in *memberin.LoginInput) (out *memberin.LoginOutput, err error) {
	// 获取客户端信息（用于登录日志）
	ip := getClientIP(ctx)
	ua := ""
	if r := ghttp.RequestFromCtx(ctx); r != nil {
		ua = r.Header.Get("User-Agent")
	}

	// 记录登录日志的辅助函数
	recordLog := func(memberId uint64, username string, status int, message string) {
		go func() {
			_, logErr := dao.MemberLoginLog.Ctx(context.Background()).Data(g.Map{
				"member_id":  memberId,
				"username":   username,
				"ip":         ip,
				"user_agent": ua,
				"status":     status,
				"message":    message,
			}).Insert()
			if logErr != nil {
				g.Log().Errorf(context.Background(), "记录会员登录日志失败: %v", logErr)
			}
		}()
	}

	// 点选验证码校验（验证码与登录强关联，在登录接口内部校验）
	if in.CaptchaId != "" && in.Captcha != "" {
		if !captcha.VerifyClick(ctx, in.CaptchaId, in.Captcha) {
			recordLog(0, in.Username, 0, "验证码错误")
			return nil, gerror.NewCode(consts.CodeBusinessError, "验证码错误或已过期，请重试")
		}
	}

	// 1. 查询会员
	var member *entity.Member
	err = dao.Member.Ctx(ctx).
		Where("username", in.Username).
		Scan(&member)
	if err != nil {
		return nil, err
	}

	if member == nil {
		recordLog(0, in.Username, 0, "用户名或密码错误")
		return nil, gerror.NewCode(consts.CodeDataNotFound, "用户名或密码错误")
	}

	// 2. 检查状态
	if member.Status != 1 {
		recordLog(member.Id, in.Username, 0, "账号已被禁用")
		return nil, gerror.NewCode(consts.CodeBusinessError, "账号已被禁用")
	}

	// 3. 验证密码
	if err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(in.Password)); err != nil {
		recordLog(member.Id, in.Username, 0, "密码错误")
		return nil, gerror.NewCode(consts.CodeBusinessError, "用户名或密码错误")
	}

	// 4. 生成 Token
	memberUser := model.MemberUser{
		Id:       member.Id,
		Username: member.Username,
		Nickname: member.Nickname,
		Avatar:   member.Avatar,
		Email:    member.Email,
		Mobile:   member.Mobile,
		Gender:   member.Gender,
		Level:    uint(member.Level),
		GroupId:  member.GroupId,
		Score:    member.Score,
		Money:    member.Money,
		LoginAt:  gtime.Now().Unix(),
	}

	accessToken, expiresIn, err := token.GenerateMember(ctx, memberUser)
	if err != nil {
		recordLog(member.Id, in.Username, 0, "Token 生成失败")
		return nil, gerror.NewCode(consts.CodeServerError, "登录失败，请稍后重试")
	}

	// 5. 更新登录信息
	_, _ = dao.Member.Ctx(ctx).
		Where("id", member.Id).
		Data(map[string]interface{}{
			"last_login_at": gtime.Now().Unix(),
			"last_login_ip": ip,
			"login_count":   member.LoginCount + 1,
		}).
		Update()

	// 6. 记录登录成功日志
	recordLog(member.Id, member.Username, 1, "登录成功")

	return &memberin.LoginOutput{
		Token:     accessToken,
		ExpiresIn: expiresIn,
	}, nil
}

// Register 会员注册
func (s *sMemberAuth) Register(ctx context.Context, in *memberin.RegisterInput) (out *memberin.RegisterOutput, err error) {
	// 1. 检查用户名是否已存在
	count, err := dao.Member.Ctx(ctx).
		Where("username", in.Username).
		Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.NewCode(consts.CodeBusinessError, "用户名已存在")
	}

	// 2. 检查手机号是否已存在（如果提供了手机号）
	if in.Mobile != "" {
		count, err = dao.Member.Ctx(ctx).
			Where("mobile", in.Mobile).
			Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, gerror.NewCode(consts.CodeBusinessError, "手机号已被注册")
		}
	}

	// 3. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeServerError, "注册失败，请稍后重试")
	}

	// 4. 插入会员记录
	result, err := dao.Member.Ctx(ctx).Data(map[string]interface{}{
		"username":   in.Username,
		"password":   string(hashedPassword),
		"mobile":     in.Mobile,
		"email":      in.Email,
		"nickname":   in.Username, // 默认昵称为用户名
		"status":     1,           // 默认启用
		"group_id":   1,           // 默认分组
		"level":      1,           // 默认等级
		// created_at / updated_at 由 GoFrame 自动维护
	}).Insert()
	if err != nil {
		return nil, gerror.NewCode(consts.CodeServerError, "注册失败，请稍后重试")
	}

	id, _ := result.LastInsertId()

	return &memberin.RegisterOutput{
		Id: uint64(id),
	}, nil
}

// Logout 会员退出登录
func (s *sMemberAuth) Logout(ctx context.Context, tokenStr string) (err error) {
	return token.DeleteByApp(ctx, token.AppMember, tokenStr)
}

// getClientIP 获取客户端 IP（支持代理头）
func getClientIP(ctx context.Context) string {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}

	// 优先从代理头获取
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// 取第一个
		for i, c := range ip {
			if c == ',' {
				ip = ip[:i]
				break
			}
		}
		ip = trimSpace(ip)
	} else {
		ip = r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.GetClientIp()
		}
	}

	if ip == "::1" || ip == "[::1]" {
		ip = "127.0.0.1"
	}

	return ip
}

// trimSpace 简单去首尾空格
func trimSpace(s string) string {
	start, end := 0, len(s)
	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}
