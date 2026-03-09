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

import "github.com/gogf/gf/v2/text/gstr"

// concealErrorSlice 需要对外隐藏真实错误、统一成友好提示的错误关键字
var concealErrorSlice = []string{ErrorORM}

// ErrorMessage 用于统一对外错误描述（非 debug 环境可使用）
// - 如果是我们标记的内部错误类型，则返回统一的“操作失败，请稍后重试！”
// - 否则直接返回 err.Error()
func ErrorMessage(err error) (message string) {
	if err == nil {
		return "操作失败！"
	}
	message = err.Error()
	for _, e := range concealErrorSlice {
		if gstr.Contains(message, e) {
			return "操作失败，请稍后重试！"
		}
	}
	return
}
