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
	"encoding/json"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
)

const (
	// 写入超时
	writeWait = 10 * time.Second
	// 读取 pong 超时
	pongWait = 60 * time.Second
	// ping 间隔（必须小于 pongWait）
	pingPeriod = (pongWait * 9) / 10
	// 最大消息大小
	maxMessageSize = 4096
	// 发送缓冲区大小
	sendBufferSize = 128
	// 心跳超时（秒）
	heartbeatTimeout = 5 * 60
)

// Client 单个 WebSocket 连接
type Client struct {
	ID            string          // 连接唯一 ID
	UserID        uint64          // 用户 ID
	Username      string          // 用户名
	UserType      string          // 用户类型: admin / member
	Conn          *websocket.Conn // 底层连接
	Send          chan []byte     // 发送缓冲区
	Tags          []string        // 已订阅的标签
	tagMu         sync.RWMutex   // 标签锁
	HeartbeatTime int64           // 上次心跳时间戳
	IP            string          // 客户端 IP
	UserAgent     string          // User-Agent
}

// NewClient 创建客户端
func NewClient(id string, conn *websocket.Conn, userID uint64, username, userType, ip, ua string) *Client {
	return &Client{
		ID:            id,
		UserID:        userID,
		Username:      username,
		UserType:      userType,
		Conn:          conn,
		Send:          make(chan []byte, sendBufferSize),
		Tags:          make([]string, 0),
		HeartbeatTime: time.Now().Unix(),
		IP:            ip,
		UserAgent:     ua,
	}
}

// ReadPump 读取消息循环
func (c *Client) ReadPump() {
	defer func() {
		Manager.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		c.refreshHeartbeat()
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				g.Log().Debugf(nil, "WebSocket client %s read error: %v", c.ID, err)
			}
			break
		}

		c.refreshHeartbeat()

		// 解析请求
		var req WsRequest
		if err := json.Unmarshal(message, &req); err != nil {
			c.SendMsg(NewErrorResponse("error", -1, "invalid message format"))
			continue
		}

		// 路由到事件处理器
		handleEvent(c, &req)
	}
}

// WritePump 写入消息循环
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Hub 关闭了 send channel
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// SendMsg 发送响应消息
func (c *Client) SendMsg(resp *WsResponse) {
	data, err := json.Marshal(resp)
	if err != nil {
		return
	}
	select {
	case c.Send <- data:
	default:
		// 缓冲区满，跳过
		g.Log().Debugf(nil, "WebSocket client %s send buffer full, dropping message", c.ID)
	}
}

// JoinTag 加入标签组
func (c *Client) JoinTag(tag string) {
	c.tagMu.Lock()
	defer c.tagMu.Unlock()
	for _, t := range c.Tags {
		if t == tag {
			return
		}
	}
	c.Tags = append(c.Tags, tag)
}

// QuitTag 离开标签组
func (c *Client) QuitTag(tag string) {
	c.tagMu.Lock()
	defer c.tagMu.Unlock()
	for i, t := range c.Tags {
		if t == tag {
			c.Tags = append(c.Tags[:i], c.Tags[i+1:]...)
			return
		}
	}
}

// HasTag 检查是否有某个标签
func (c *Client) HasTag(tag string) bool {
	c.tagMu.RLock()
	defer c.tagMu.RUnlock()
	for _, t := range c.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// refreshHeartbeat 刷新心跳时间
func (c *Client) refreshHeartbeat() {
	c.HeartbeatTime = time.Now().Unix()
}

// IsHeartbeatTimeout 心跳是否超时
func (c *Client) IsHeartbeatTimeout() bool {
	return time.Now().Unix()-c.HeartbeatTime > heartbeatTimeout
}

// UserKey 用户唯一标识
func (c *Client) UserKey() string {
	return userKey(c.UserType, c.UserID)
}
