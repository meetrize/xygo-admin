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
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// clientManager Hub 管理器
type clientManager struct {
	clients     map[*Client]bool      // 所有连接
	clientsLock sync.RWMutex          // 连接锁
	users       map[string][]*Client  // 用户 -> 连接列表（多设备）
	usersLock   sync.RWMutex          // 用户锁
	Register    chan *Client           // 注册
	Unregister  chan *Client           // 注销
	stopCh      chan struct{}          // 停止信号
}

// Manager 全局 Hub 实例
var Manager = &clientManager{
	clients:    make(map[*Client]bool),
	users:      make(map[string][]*Client),
	Register:   make(chan *Client, 64),
	Unregister: make(chan *Client, 64),
	stopCh:     make(chan struct{}),
}

// Start 启动 Hub
func (m *clientManager) Start() {
	go m.run()
	go m.heartbeatCleanup()
	g.Log().Info(nil, "WebSocket Hub started")
}

// Stop 停止 Hub
func (m *clientManager) Stop() {
	close(m.stopCh)
	g.Log().Info(nil, "WebSocket Hub stopped")
}

// run Hub 主循环
func (m *clientManager) run() {
	for {
		select {
		case client := <-m.Register:
			m.addClient(client)
		case client := <-m.Unregister:
			m.removeClient(client)
		case <-m.stopCh:
			return
		}
	}
}

// heartbeatCleanup 心跳超时清理
func (m *clientManager) heartbeatCleanup() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.clientsLock.RLock()
			var timeoutClients []*Client
			for client := range m.clients {
				if client.IsHeartbeatTimeout() {
					timeoutClients = append(timeoutClients, client)
				}
			}
			m.clientsLock.RUnlock()

			for _, client := range timeoutClients {
				g.Log().Debugf(nil, "WebSocket client %s heartbeat timeout, closing", client.ID)
				m.Unregister <- client
			}
		case <-m.stopCh:
			return
		}
	}
}

// addClient 注册客户端
func (m *clientManager) addClient(client *Client) {
	m.clientsLock.Lock()
	m.clients[client] = true
	m.clientsLock.Unlock()

	key := client.UserKey()
	m.usersLock.Lock()
	m.users[key] = append(m.users[key], client)
	m.usersLock.Unlock()

	g.Log().Debugf(nil, "WebSocket client registered: id=%s user=%s(%d) total=%d",
		client.ID, client.Username, client.UserID, m.ClientCount())
}

// removeClient 注销客户端
func (m *clientManager) removeClient(client *Client) {
	m.clientsLock.Lock()
	if _, ok := m.clients[client]; !ok {
		m.clientsLock.Unlock()
		return
	}
	delete(m.clients, client)
	m.clientsLock.Unlock()

	close(client.Send)

	key := client.UserKey()
	m.usersLock.Lock()
	if clients, ok := m.users[key]; ok {
		for i, c := range clients {
			if c == client {
				m.users[key] = append(clients[:i], clients[i+1:]...)
				break
			}
		}
		if len(m.users[key]) == 0 {
			delete(m.users, key)
		}
	}
	m.usersLock.Unlock()

	g.Log().Debugf(nil, "WebSocket client unregistered: id=%s user=%s(%d) total=%d",
		client.ID, client.Username, client.UserID, m.ClientCount())
}

// ---- 推送方法 ----

// SendToAll 广播给所有客户端
func SendToAll(resp *WsResponse) {
	data, err := json.Marshal(resp)
	if err != nil {
		return
	}
	Manager.clientsLock.RLock()
	defer Manager.clientsLock.RUnlock()
	for client := range Manager.clients {
		select {
		case client.Send <- data:
		default:
		}
	}
}

// SendToUser 推送给指定用户的所有连接
func SendToUser(userType string, userID uint64, resp *WsResponse) {
	data, err := json.Marshal(resp)
	if err != nil {
		return
	}
	key := userKey(userType, userID)
	Manager.usersLock.RLock()
	defer Manager.usersLock.RUnlock()
	if clients, ok := Manager.users[key]; ok {
		for _, client := range clients {
			select {
			case client.Send <- data:
			default:
			}
		}
	}
}

// SendToTag 推送给指定标签组的客户端
func SendToTag(tag string, resp *WsResponse) {
	data, err := json.Marshal(resp)
	if err != nil {
		return
	}
	Manager.clientsLock.RLock()
	defer Manager.clientsLock.RUnlock()
	for client := range Manager.clients {
		if client.HasTag(tag) {
			select {
			case client.Send <- data:
			default:
			}
		}
	}
}

// ClientCount 当前连接数
func (m *clientManager) ClientCount() int {
	m.clientsLock.RLock()
	defer m.clientsLock.RUnlock()
	return len(m.clients)
}

// IsUserOnline 判断指定用户是否在线
func IsUserOnline(userType string, userID uint64) bool {
	key := userKey(userType, userID)
	Manager.usersLock.RLock()
	defer Manager.usersLock.RUnlock()
	clients, ok := Manager.users[key]
	return ok && len(clients) > 0
}

// GetOnlineUserIDs 获取所有在线用户ID（指定类型）
func GetOnlineUserIDs(userType string) []uint64 {
	prefix := fmt.Sprintf("ws_%s_", userType)
	Manager.usersLock.RLock()
	defer Manager.usersLock.RUnlock()

	ids := make([]uint64, 0)
	seen := make(map[uint64]bool)
	for key := range Manager.users {
		if len(key) > len(prefix) && key[:len(prefix)] == prefix {
			var uid uint64
			if _, err := fmt.Sscanf(key, "ws_"+userType+"_%d", &uid); err == nil && !seen[uid] {
				seen[uid] = true
				ids = append(ids, uid)
			}
		}
	}
	return ids
}

// userKey 生成用户唯一键
func userKey(userType string, userID uint64) string {
	return fmt.Sprintf("ws_%s_%d", userType, userID)
}
