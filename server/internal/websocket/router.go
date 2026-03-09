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
	"github.com/gogf/gf/v2/frame/g"
)

// EventHandler 事件处理函数
type EventHandler func(client *Client, req *WsRequest)

// 事件路由表
var eventHandlers = make(map[string]EventHandler)

// RegisterEvent 注册事件处理器
func RegisterEvent(event string, handler EventHandler) {
	eventHandlers[event] = handler
}

// handleEvent 路由分发
func handleEvent(client *Client, req *WsRequest) {
	handler, ok := eventHandlers[req.Event]
	if !ok {
		client.SendMsg(NewErrorResponse(req.Event, -1, "unknown event: "+req.Event))
		return
	}

	// 安全执行
	func() {
		defer func() {
			if r := recover(); r != nil {
				g.Log().Errorf(nil, "WebSocket event handler panic: event=%s err=%v", req.Event, r)
				client.SendMsg(NewErrorResponse(req.Event, -1, "internal error"))
			}
		}()
		handler(client, req)
	}()
}

// registerBuiltinEvents 注册内置事件
func registerBuiltinEvents() {
	// 心跳
	RegisterEvent("ping", func(client *Client, req *WsRequest) {
		client.SendMsg(NewResponse("pong", nil))
	})

	// 加入标签组
	RegisterEvent("join", func(client *Client, req *WsRequest) {
		tag, _ := req.Data["tag"].(string)
		if tag == "" {
			client.SendMsg(NewErrorResponse("join", -1, "tag is required"))
			return
		}
		client.JoinTag(tag)
		client.SendMsg(NewResponse("join", map[string]string{"tag": tag, "status": "joined"}))
		g.Log().Debugf(nil, "WebSocket client %s joined tag: %s", client.ID, tag)
	})

	// 离开标签组
	RegisterEvent("quit", func(client *Client, req *WsRequest) {
		tag, _ := req.Data["tag"].(string)
		if tag == "" {
			client.SendMsg(NewErrorResponse("quit", -1, "tag is required"))
			return
		}
		client.QuitTag(tag)
		client.SendMsg(NewResponse("quit", map[string]string{"tag": tag, "status": "quited"}))
		g.Log().Debugf(nil, "WebSocket client %s quited tag: %s", client.ID, tag)
	})
}
