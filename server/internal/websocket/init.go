// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package websocket

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"

	"xygo/internal/library/contexts"
	"xygo/internal/library/monitor"
	"xygo/internal/library/token"
)

// Start 启动 WebSocket 服务
func Start() {
	registerBuiltinEvents()
	Manager.Start()
	go startMonitorPush()

	// 注册踢人 WS 通知回调（避免 token 包循环引用 websocket 包）
	token.WsKickNotifier = func(userType string, userId uint64) {
		SendToUser(userType, userId, NewResponse("kicked", g.Map{
			"message": "您的账号已在其他设备登录",
		}))
	}
}

// startMonitorPush 定时推送服务器指标给订阅了 monitor_server 标签的客户端
func startMonitorPush() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 只有有人订阅时才采集
			if !hasSubscribers("monitor_server") {
				continue
			}
			info, err := monitor.GetServerInfo()
			if err != nil {
				continue
			}
			SendToTag("monitor_server", NewResponse("monitor/server", info))
		case <-Manager.stopCh:
			return
		}
	}
}

// hasSubscribers 检查某个标签是否有订阅者
func hasSubscribers(tag string) bool {
	Manager.clientsLock.RLock()
	defer Manager.clientsLock.RUnlock()
	for client := range Manager.clients {
		if client.HasTag(tag) {
			return true
		}
	}
	return false
}

// Stop 停止 WebSocket 服务
func Stop() {
	Manager.Stop()
}

// WsHandler WebSocket 连接处理（GoFrame 路由绑定用）
func WsHandler(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(r.GetCtx(), "WebSocket upgrade failed:", err)
		r.Exit()
		return
	}

	// 从 context 获取鉴权后的用户信息
	var (
		userID   uint64
		username string
		userType string
	)

	user := contexts.GetUser(r.GetCtx())
	if user != nil {
		userID = user.Id
		username = user.Username
		userType = "admin"
	}

	ip := r.GetClientIp()
	ua := r.Header.Get("User-Agent")

	client := NewClient(guid.S(), ws.Conn, userID, username, userType, ip, ua)

	Manager.Register <- client

	// 发送连接成功消息
	client.SendMsg(NewResponse("connected", g.Map{
		"clientId": client.ID,
		"userId":   client.UserID,
		"username": client.Username,
	}))

	// 启动读写循环
	go client.WritePump()
	client.ReadPump() // 阻塞直到断开
}
