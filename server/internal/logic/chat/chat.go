// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

package chat

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"xygo/internal/dao"
	"xygo/internal/model/entity"
	"xygo/internal/model/input/adminin"
	"xygo/internal/service"
	"xygo/internal/websocket"
)

type sChat struct{}

func init() {
	service.RegisterChat(&sChat{})
}

// Sessions 获取当前用户的会话列表
func (s *sChat) Sessions(ctx context.Context, userId uint64) ([]adminin.ChatSessionItem, error) {
	// 1. 查询用户参与的、未删除的会话成员记录
	var members []entity.AdminChatSessionMember
	err := dao.AdminChatSessionMember.Ctx(ctx).
		Where("user_id", userId).
		Where("is_deleted", 0).
		Scan(&members)
	if err != nil {
		return nil, err
	}
	if len(members) == 0 {
		return []adminin.ChatSessionItem{}, nil
	}

	// 收集会话ID和未读数映射
	sessionIds := make([]uint64, 0, len(members))
	memberMap := make(map[uint64]*entity.AdminChatSessionMember)
	for i := range members {
		sessionIds = append(sessionIds, members[i].SessionId)
		memberMap[members[i].SessionId] = &members[i]
	}

	// 2. 查询会话信息
	var sessions []entity.AdminChatSession
	err = dao.AdminChatSession.Ctx(ctx).
		WhereIn("id", sessionIds).
		OrderDesc("last_message_time").
		Scan(&sessions)
	if err != nil {
		return nil, err
	}

	// 3. 单聊时需要查询对方的信息
	items := make([]adminin.ChatSessionItem, 0, len(sessions))
	for _, s := range sessions {
		mem := memberMap[s.Id]
		item := adminin.ChatSessionItem{
			Id:              s.Id,
			Type:            s.Type,
			Name:            s.Name,
			Avatar:          s.Avatar,
			LastMessage:     s.LastMessage,
			LastMessageTime: s.LastMessageTime,
			UnreadCount:     mem.UnreadCount,
			MemberCount:     s.MemberCount,
			IsMuted:         mem.IsMuted,
		}

		// 单聊时用对方的信息
		if s.Type == 1 {
			otherUser, err := getOtherUser(ctx, s.Id, userId)
			if err == nil && otherUser != nil {
				item.Name = otherUser.Username
				if item.Name == "" {
					item.Name = otherUser.RealName
				}
				item.Avatar = otherUser.Avatar
			}
		}

		items = append(items, item)
	}

	return items, nil
}

// CreateSession 创建会话
func (s *sChat) CreateSession(ctx context.Context, in *adminin.ChatSessionCreateInp, creatorId uint64) (*adminin.ChatSessionCreateModel, error) {
	now := uint64(time.Now().Unix())

	// 单聊：检查是否已存在
	if in.Type == 1 {
		if len(in.UserIds) != 1 {
			return nil, fmt.Errorf("单聊只能选择一个联系人")
		}
		targetId := in.UserIds[0]
		if targetId == creatorId {
			return nil, fmt.Errorf("不能和自己聊天")
		}

		existId, err := findExistingPrivateSession(ctx, creatorId, targetId)
		if err != nil {
			return nil, err
		}
		if existId > 0 {
			// 恢复已删除的会话
			_, _ = dao.AdminChatSessionMember.Ctx(ctx).
				Where("session_id", existId).
				Where("user_id", creatorId).
				Data(g.Map{"is_deleted": 0}).
				Update()
			return &adminin.ChatSessionCreateModel{SessionId: existId}, nil
		}
	}

	// 群聊验证
	if in.Type == 2 && in.Name == "" {
		in.Name = "群聊"
	}

	// 创建会话
	allUsers := append(in.UserIds, creatorId)
	allUsers = uniqueUint64(allUsers)

	result, err := dao.AdminChatSession.Ctx(ctx).Data(g.Map{
		"type":              in.Type,
		"name":              in.Name,
		"creator_id":        creatorId,
		"member_count":      len(allUsers),
		"last_message":      "",
		"last_message_time": now,
		"created_at":        now,
		"updated_at":        now,
	}).Insert()
	if err != nil {
		return nil, err
	}
	sessionId, _ := result.LastInsertId()

	// 添加成员
	for _, uid := range allUsers {
		role := 1
		if in.Type == 2 && uid == creatorId {
			role = 3 // 群主
		}
		_, _ = dao.AdminChatSessionMember.Ctx(ctx).Data(g.Map{
			"session_id": sessionId,
			"user_id":    uid,
			"role":       role,
			"joined_at":  now,
		}).Insert()
	}

	// 群聊时发送系统消息
	if in.Type == 2 {
		var creator entity.AdminUser
		_ = dao.AdminUser.Ctx(ctx).Where("id", creatorId).Fields("real_name,username").Scan(&creator)
		name := creator.RealName
		if name == "" {
			name = creator.Username
		}
		sysContent := fmt.Sprintf("%s 创建了群聊", name)
		insertSystemMessage(ctx, uint64(sessionId), sysContent, now)
	}

	return &adminin.ChatSessionCreateModel{SessionId: uint64(sessionId)}, nil
}

// DeleteSession 删除会话（软删除，仅对当前用户生效）
func (s *sChat) DeleteSession(ctx context.Context, in *adminin.ChatSessionDeleteInp, userId uint64) error {
	_, err := dao.AdminChatSessionMember.Ctx(ctx).
		Where("session_id", in.SessionId).
		Where("user_id", userId).
		Data(g.Map{"is_deleted": 1, "unread_count": 0}).
		Update()
	return err
}

// Messages 获取消息列表
func (s *sChat) Messages(ctx context.Context, in *adminin.ChatMessageListInp, userId uint64) (*adminin.ChatMessageListModel, error) {
	// 验证用户是否在会话中
	count, err := dao.AdminChatSessionMember.Ctx(ctx).
		Where("session_id", in.SessionId).
		Where("user_id", userId).
		Where("is_deleted", 0).
		Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("无权访问该会话")
	}

	pageSize := in.PageSize
	if pageSize <= 0 {
		pageSize = 30
	}

	m := dao.AdminChatMessage.Ctx(ctx).Where("session_id", in.SessionId)
	if in.LastMsgId > 0 {
		m = m.Where("id <", in.LastMsgId)
	}

	var msgs []entity.AdminChatMessage
	err = m.OrderDesc("id").Limit(pageSize + 1).Scan(&msgs)
	if err != nil {
		return nil, err
	}

	hasMore := len(msgs) > pageSize
	if hasMore {
		msgs = msgs[:pageSize]
	}

	// 查询发送者信息
	senderMap := make(map[uint64]*entity.AdminUser)
	senderIds := make([]uint64, 0)
	for _, msg := range msgs {
		if msg.SenderId > 0 {
			senderIds = append(senderIds, msg.SenderId)
		}
	}
	senderIds = uniqueUint64(senderIds)
	if len(senderIds) > 0 {
		var users []entity.AdminUser
		_ = dao.AdminUser.Ctx(ctx).WhereIn("id", senderIds).Fields("id,username,real_name,avatar").Scan(&users)
		for i := range users {
			senderMap[users[i].Id] = &users[i]
		}
	}

	items := make([]adminin.ChatMessageItem, 0, len(msgs))
	for i := len(msgs) - 1; i >= 0; i-- { // 倒序输出（时间正序）
		msg := msgs[i]
		item := adminin.ChatMessageItem{
			Id:        msg.Id,
			SessionId: msg.SessionId,
			SenderId:  msg.SenderId,
			Type:      msg.Type,
			Content:   msg.Content,
			CreatedAt: msg.CreatedAt,
		}
		if u, ok := senderMap[msg.SenderId]; ok {
			item.SenderName = u.Username
			if item.SenderName == "" {
				item.SenderName = u.RealName
			}
			item.SenderAvatar = u.Avatar
		}
		items = append(items, item)
	}

	return &adminin.ChatMessageListModel{
		List:    items,
		HasMore: hasMore,
	}, nil
}

// SendMessage 发送消息
func (s *sChat) SendMessage(ctx context.Context, in *adminin.ChatSendMessageInp, senderId uint64) (*adminin.ChatSendMessageModel, error) {
	now := int(time.Now().Unix())

	// 验证用户是否在会话中
	count, err := dao.AdminChatSessionMember.Ctx(ctx).
		Where("session_id", in.SessionId).
		Where("user_id", senderId).
		Where("is_deleted", 0).
		Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("无权在该会话中发言")
	}

	// 插入消息
	result, err := dao.AdminChatMessage.Ctx(ctx).Data(g.Map{
		"session_id": in.SessionId,
		"sender_id":  senderId,
		"type":       in.Type,
		"content":    in.Content,
		"created_at": now,
	}).Insert()
	if err != nil {
		return nil, err
	}
	msgId, _ := result.LastInsertId()

	// 更新会话最后消息
	preview := in.Content
	if in.Type == 2 {
		preview = "[图片]"
	}
	if len(preview) > 100 {
		preview = preview[:100]
	}
	_, _ = dao.AdminChatSession.Ctx(ctx).Where("id", in.SessionId).Data(g.Map{
		"last_message":      preview,
		"last_message_time": now,
		"updated_at":        now,
	}).Update()

	// 更新所有其他成员的未读数 +1，并恢复 is_deleted
	// 注意：按 user_id 去重，避免历史脏数据（重复成员行）导致重复累加
	var sessionMembers []entity.AdminChatSessionMember
	_ = dao.AdminChatSessionMember.Ctx(ctx).
		Where("session_id", in.SessionId).
		Scan(&sessionMembers)

	recipientMemberIdMap := make(map[uint64]uint64)
	for _, m := range sessionMembers {
		if m.UserId == senderId {
			continue
		}
		if _, exists := recipientMemberIdMap[m.UserId]; !exists {
			recipientMemberIdMap[m.UserId] = m.Id
		}
	}
	for _, memberId := range recipientMemberIdMap {
		_, _ = dao.AdminChatSessionMember.Ctx(ctx).
			Where("id", memberId).
			Data(g.Map{"is_deleted": 0}).
			Increment("unread_count", 1)
	}

	// 获取发送者信息用于推送
	var sender entity.AdminUser
	_ = dao.AdminUser.Ctx(ctx).Where("id", senderId).Fields("id,username,real_name,avatar").Scan(&sender)
	senderName := sender.Username
	if senderName == "" {
		senderName = sender.RealName
	}

	// WebSocket 推送给会话中所有成员
	go pushChatMessage(sessionMembers, uint64(msgId), in.SessionId, senderId, senderName, sender.Avatar, in.Type, in.Content, uint64(now))

	return &adminin.ChatSendMessageModel{
		Id:        uint64(msgId),
		CreatedAt: now,
	}, nil
}

// MarkRead 标记会话已读
func (s *sChat) MarkRead(ctx context.Context, sessionId, userId uint64) error {
	// 获取该会话最后一条消息ID
	var lastMsg entity.AdminChatMessage
	err := dao.AdminChatMessage.Ctx(ctx).
		Where("session_id", sessionId).
		OrderDesc("id").
		Limit(1).
		Scan(&lastMsg)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	updateData := g.Map{
		"unread_count": 0,
	}
	// 新建会话可能还没有消息，此时不更新 last_read_msg_id，避免 no rows 报错影响前端
	if lastMsg.Id > 0 {
		updateData["last_read_msg_id"] = lastMsg.Id
	}
	_, err = dao.AdminChatSessionMember.Ctx(ctx).
		Where("session_id", sessionId).
		Where("user_id", userId).
		Data(updateData).
		Update()
	if err != nil {
		return err
	}

	// 推送已读回执给会话其他成员
	go pushReadReceipt(sessionId, userId)

	return nil
}

// Contacts 获取所有管理员作为联系人
func (s *sChat) Contacts(ctx context.Context, currentUserId uint64) ([]adminin.ChatContactItem, error) {
	var users []entity.AdminUser
	err := dao.AdminUser.Ctx(ctx).
		Where("status", 1).
		Fields("id,username,real_name,dept_id,avatar").
		OrderAsc("id").
		Scan(&users)
	if err != nil {
		return nil, err
	}

	userIds := make([]uint64, 0, len(users))
	for _, u := range users {
		userIds = append(userIds, u.Id)
	}

	// 部门名称映射
	deptIds := make([]uint64, 0)
	deptNameMap := make(map[uint64]string)
	for _, u := range users {
		if u.DeptId > 0 {
			deptIds = append(deptIds, u.DeptId)
		}
	}
	deptIds = uniqueUint64(deptIds)
	if len(deptIds) > 0 {
		var depts []entity.AdminDept
		_ = dao.AdminDept.Ctx(ctx).
			WhereIn("id", deptIds).
			Fields("id,name").
			Scan(&depts)
		for _, d := range depts {
			deptNameMap[d.Id] = d.Name
		}
	}

	// 岗位名称映射：user_id → 岗位名(多个逗号拼接)
	postNameMap := make(map[uint64]string)
	if len(userIds) > 0 {
		var userPosts []entity.AdminUserPost
		_ = dao.AdminUserPost.Ctx(ctx).
			WhereIn("user_id", userIds).
			Scan(&userPosts)
		if len(userPosts) > 0 {
			postIds := make([]uint64, 0)
			for _, up := range userPosts {
				postIds = append(postIds, up.PostId)
			}
			postIds = uniqueUint64(postIds)
			postMap := make(map[uint64]string)
			var posts []entity.AdminPost
			_ = dao.AdminPost.Ctx(ctx).
				WhereIn("id", postIds).
				Fields("id,name").
				Scan(&posts)
			for _, p := range posts {
				postMap[p.Id] = p.Name
			}
			for _, up := range userPosts {
				if name, ok := postMap[up.PostId]; ok {
					if existing := postNameMap[up.UserId]; existing != "" {
						postNameMap[up.UserId] = existing + "," + name
					} else {
						postNameMap[up.UserId] = name
					}
				}
			}
		}
	}

	items := make([]adminin.ChatContactItem, 0, len(users))
	for _, u := range users {
		if u.Id == currentUserId {
			continue // 排除自己
		}
		name := u.RealName
		if name == "" {
			name = u.Username
		}
		items = append(items, adminin.ChatContactItem{
			Id:       u.Id,
			Username: u.Username,
			RealName: name,
			DeptId:   u.DeptId,
			DeptName: deptNameMap[u.DeptId],
			PostName: postNameMap[u.Id],
			Avatar:   u.Avatar,
			IsOnline: websocket.IsUserOnline("admin", u.Id),
		})
	}
	return items, nil
}

// GroupUpdate 编辑群聊
func (s *sChat) GroupUpdate(ctx context.Context, in *adminin.ChatGroupUpdateInp, userId uint64) error {
	// 验证是否为群聊
	var session entity.AdminChatSession
	err := dao.AdminChatSession.Ctx(ctx).Where("id", in.SessionId).Scan(&session)
	if err != nil {
		return err
	}
	if session.Type != 2 {
		return fmt.Errorf("非群聊会话不支持此操作")
	}

	now := uint64(time.Now().Unix())

	// 更新群名
	if in.Name != "" && in.Name != session.Name {
		_, _ = dao.AdminChatSession.Ctx(ctx).Where("id", in.SessionId).Data(g.Map{
			"name":       in.Name,
			"updated_at": now,
		}).Update()
	}

	// 添加成员
	if len(in.AddUsers) > 0 {
		for _, uid := range in.AddUsers {
			count, _ := dao.AdminChatSessionMember.Ctx(ctx).
				Where("session_id", in.SessionId).
				Where("user_id", uid).
				Count()
			if count == 0 {
				_, _ = dao.AdminChatSessionMember.Ctx(ctx).Data(g.Map{
					"session_id": in.SessionId,
					"user_id":    uid,
					"role":       1,
					"joined_at":  now,
				}).Insert()
			} else {
				// 恢复已删除的
				_, _ = dao.AdminChatSessionMember.Ctx(ctx).
					Where("session_id", in.SessionId).
					Where("user_id", uid).
					Data(g.Map{"is_deleted": 0}).
					Update()
			}
		}
	}

	// 移除成员
	if len(in.DelUsers) > 0 {
		for _, uid := range in.DelUsers {
			_, _ = dao.AdminChatSessionMember.Ctx(ctx).
				Where("session_id", in.SessionId).
				Where("user_id", uid).
				Delete()
		}
	}

	// 更新成员数
	memberCount, _ := dao.AdminChatSessionMember.Ctx(ctx).
		Where("session_id", in.SessionId).
		Where("is_deleted", 0).
		Count()
	_, _ = dao.AdminChatSession.Ctx(ctx).Where("id", in.SessionId).Data(g.Map{
		"member_count": memberCount,
		"updated_at":   now,
	}).Update()

	return nil
}

// UnreadTotal 获取聊天未读总数
func (s *sChat) UnreadTotal(ctx context.Context, userId uint64) (int, error) {
	total := 0
	var members []entity.AdminChatSessionMember
	err := dao.AdminChatSessionMember.Ctx(ctx).
		Where("user_id", userId).
		Where("is_deleted", 0).
		Fields("unread_count").
		Scan(&members)
	if err != nil {
		return 0, err
	}
	for _, m := range members {
		total += m.UnreadCount
	}
	return total, nil
}

// ==================== 内部辅助函数 ====================

// getOtherUser 单聊中获取对方用户信息
func getOtherUser(ctx context.Context, sessionId, currentUserId uint64) (*entity.AdminUser, error) {
	var member entity.AdminChatSessionMember
	err := dao.AdminChatSessionMember.Ctx(ctx).
		Where("session_id", sessionId).
		WhereNot("user_id", currentUserId).
		Limit(1).
		Scan(&member)
	if err != nil {
		return nil, err
	}
	if member.UserId == 0 {
		return nil, nil
	}
	var user entity.AdminUser
	err = dao.AdminUser.Ctx(ctx).Where("id", member.UserId).Fields("id,username,real_name,avatar").Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// findExistingPrivateSession 查找两人之间已有的单聊会话
func findExistingPrivateSession(ctx context.Context, userId1, userId2 uint64) (uint64, error) {
	// 查找两人都在的单聊会话
	var sessions []entity.AdminChatSession
	err := dao.AdminChatSession.Ctx(ctx).Where("type", 1).Scan(&sessions)
	if err != nil {
		return 0, err
	}

	for _, s := range sessions {
		var members []entity.AdminChatSessionMember
		_ = dao.AdminChatSessionMember.Ctx(ctx).
			Where("session_id", s.Id).
			Scan(&members)
		if len(members) == 2 {
			hasUser1, hasUser2 := false, false
			for _, m := range members {
				if m.UserId == userId1 {
					hasUser1 = true
				}
				if m.UserId == userId2 {
					hasUser2 = true
				}
			}
			if hasUser1 && hasUser2 {
				return s.Id, nil
			}
		}
	}
	return 0, nil
}

// insertSystemMessage 插入系统消息
func insertSystemMessage(ctx context.Context, sessionId uint64, content string, now uint64) {
	_, _ = dao.AdminChatMessage.Ctx(ctx).Data(g.Map{
		"session_id": sessionId,
		"sender_id":  0,
		"type":       3, // 系统消息
		"content":    content,
		"created_at": now,
	}).Insert()
}

// pushChatMessage WebSocket 推送聊天消息（不推给发送者自己）
func pushChatMessage(members []entity.AdminChatSessionMember, msgId, sessionId, senderId uint64, senderName, senderAvatar string, msgType int, content string, createdAt uint64) {
	resp := &websocket.WsResponse{
		Event: "chat/message",
		Data: g.Map{
			"id":           msgId,
			"sessionId":    sessionId,
			"senderId":     senderId,
			"senderName":   senderName,
			"senderAvatar": senderAvatar,
			"type":         msgType,
			"content":      content,
			"createdAt":    createdAt,
		},
	}

	pushedUsers := make(map[uint64]bool)
	for _, m := range members {
		if m.UserId == senderId {
			continue
		}
		if pushedUsers[m.UserId] {
			continue
		}
		pushedUsers[m.UserId] = true
		websocket.SendToUser("admin", m.UserId, resp)
	}
}

// pushReadReceipt 推送已读回执
func pushReadReceipt(sessionId, readerId uint64) {
	// 获取会话成员
	var members []entity.AdminChatSessionMember
	_ = dao.AdminChatSessionMember.Ctx(context.Background()).
		Where("session_id", sessionId).
		Scan(&members)

	resp := &websocket.WsResponse{
		Event: "chat/read",
		Data: g.Map{
			"sessionId": sessionId,
			"readerId":  readerId,
		},
	}

	pushedUsers := make(map[uint64]bool)
	for _, m := range members {
		if m.UserId == readerId {
			continue
		}
		if pushedUsers[m.UserId] {
			continue
		}
		pushedUsers[m.UserId] = true
		websocket.SendToUser("admin", m.UserId, resp)
	}
}

// uniqueUint64 去重
func uniqueUint64(arr []uint64) []uint64 {
	seen := make(map[uint64]bool)
	result := make([]uint64, 0, len(arr))
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
