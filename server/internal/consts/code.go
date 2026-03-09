// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package consts

import "github.com/gogf/gf/v2/errors/gcode"

// Code 约定：
// - 0            : 成功（gcode.CodeOK）
// - <0           : GoFrame 内置错误码（gcode 内部使用）
// - 50xx/60xx    : 常见通用错误（未登录、无权限、未找到等，复用 gcode 内置）
// - >= 10000     : 我们自定义的业务错误码
//
// 说明：
// - 优先使用这些别名，而不要在业务里直接写数字常量，便于统一维护和注释。

var (
	// CodeOK 成功
	CodeOK = gcode.CodeOK // 0

	// CodeFailed 通用失败（默认错误）
	CodeFailed = gcode.CodeNil // -1

	// CodeNotAuthorized 未登录或登录失效
	CodeNotAuthorized = gcode.CodeNotAuthorized

	// CodeNoPermission 无访问权限（登录了但没有该接口/资源权限）
	CodeNoPermission = gcode.CodeSecurityReason

	// CodeNotFound 路由/资源不存在
	CodeNotFound = gcode.CodeNotFound
)

// 自定义业务错误码示例（>=10000）
// 根据项目需要逐步补充，使用时通过 gerror.NewCode(consts.CodeXXX, "错误提示") 返回给前端。
var (
	// CodeInvalidParam 参数错误
	CodeInvalidParam = gcode.New(10001, "参数错误", nil)

	// CodeDuplicateData 数据重复（唯一键冲突等）
	CodeDuplicateData = gcode.New(10002, "数据已存在", nil)

	// CodeDataNotFound 业务数据不存在
	CodeDataNotFound = gcode.New(10003, "数据不存在", nil)

	// CodeBusinessError 业务逻辑错误（如密码错误、用户已存在等）
	CodeBusinessError = gcode.New(10004, "业务错误", nil)

	// CodeKickedOut 被踢下线（SSO单点登录：其他设备登录 / 管理员强制下线）
	CodeKickedOut = gcode.New(10010, "账号已在其他设备登录", nil)

	// CodeServerError 服务器内部错误
	CodeServerError = gcode.New(50000, "服务器错误", nil)
)
