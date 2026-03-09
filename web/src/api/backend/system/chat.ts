// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

/**
 * 聊天 API
 */
import { adminRequest } from '@/utils/http'

// ==================== 类型定义 ====================

export interface ChatSessionItem {
  id: number
  type: number // 1=单聊, 2=群聊
  name: string
  avatar: string
  lastMessage: string
  lastMessageTime: number
  unreadCount: number
  memberCount: number
  isMuted: number
}

export interface ChatMessageItem {
  id: number
  sessionId: number
  senderId: number
  senderName: string
  senderAvatar: string
  type: number // 1=文字, 2=图片, 3=系统消息
  content: string
  createdAt: number
}

export interface ChatContactItem {
  id: number
  username: string
  realName: string
  deptId: number
  deptName: string
  postName: string
  avatar: string
  isOnline: boolean
}

// ==================== 会话接口 ====================

/** 获取会话列表 */
export function fetchChatSessions() {
  return adminRequest.get<{ list: ChatSessionItem[] }>({ url: '/chat/sessions' })
}

/** 创建会话 */
export function fetchChatSessionCreate(params: { type: number; name?: string; userIds: number[] }) {
  return adminRequest.post<{ sessionId: number }>({ url: '/chat/session/create', params })
}

/** 删除会话 */
export function fetchChatSessionDelete(sessionId: number) {
  return adminRequest.post<any>({ url: '/chat/session/delete', params: { sessionId } })
}

// ==================== 消息接口 ====================

/** 获取消息列表 */
export function fetchChatMessages(params: { sessionId: number; lastMsgId?: number; page?: number; pageSize?: number }) {
  return adminRequest.post<{ list: ChatMessageItem[]; hasMore: boolean }>({ url: '/chat/messages', params })
}

/** 发送消息 */
export function fetchChatSend(params: { sessionId: number; type: number; content: string }) {
  return adminRequest.post<{ id: number; createdAt: number }>({ url: '/chat/send', params })
}

/** 标记已读 */
export function fetchChatRead(sessionId: number) {
  return adminRequest.post<any>({ url: '/chat/read', params: { sessionId } })
}

// ==================== 联系人接口 ====================

/** 获取联系人列表 */
export function fetchChatContacts() {
  return adminRequest.get<{ list: ChatContactItem[] }>({ url: '/chat/contacts' })
}

// ==================== 群聊管理 ====================

/** 编辑群聊 */
export function fetchChatGroupUpdate(params: { sessionId: number; name?: string; addUsers?: number[]; delUsers?: number[] }) {
  return adminRequest.post<any>({ url: '/chat/group/update', params })
}

// ==================== 未读总数 ====================

/** 获取聊天未读总数 */
export function fetchChatUnreadTotal() {
  return adminRequest.get<{ total: number }>({ url: '/chat/unread-total' })
}
