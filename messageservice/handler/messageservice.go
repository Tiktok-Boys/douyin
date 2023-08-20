package handler

import (
	"context"
	"strconv"

	"github.com/Tiktok-Boys/douyin/messageservice/dal/db"
	pb "github.com/Tiktok-Boys/douyin/messageservice/proto"
)

type MessageService struct{}

func (s *MessageService) ActMessage(ctx context.Context, in *pb.DouyinMessageActionRequest, out *pb.DouyinMessageActionResponse) error {
	senderId, err := strconv.Atoi(in.UserId)
	if err != nil {
		out.StatusCode = 1
		out.StatusMsg = "用户 ID 不合法"
		return err
	}
	receiverId, err := strconv.Atoi(in.ToUserId)
	if err != nil {
		out.StatusCode = 2
		out.StatusMsg = "对方用户 ID 不合法"
		return err
	}
	err = db.CreateMessage(ctx, &db.Message{
		SenderID:   senderId,
		ReceiverID: receiverId,
		Content:    in.Content,
	})
	if err != nil {
		out.StatusCode = 3
		out.StatusMsg = "创建消息失败：" + err.Error()
		return err
	}
	out.StatusCode = 0
	return nil
}

func (s *MessageService) GetChat(ctx context.Context, in *pb.DouyinMessageChatRequest, out *pb.DouyinMessageChatResponse) error {
	fromUserId, err := strconv.Atoi(in.UserId)
	if err != nil {
		out.StatusCode = 1
		out.StatusMsg = "用户 ID 不合法"
		return err
	}
	toUserId, err := strconv.Atoi(in.ToUserId)
	if err != nil {
		out.StatusCode = 2
		out.StatusMsg = "对方用户 ID 不合法"
		return err
	}
	messages, err := db.GetChatRecords(ctx, toUserId, fromUserId)
	if err != nil {
		out.StatusCode = 3
		out.StatusMsg = "获取聊天记录失败：" + err.Error()
		return err
	}
	out.StatusCode = 0
	out.MessageList = make([]*pb.Message, len(messages))
	for i, m := range messages {
		message := pb.Message{}
		message.Id = int64(m.ID)
		message.ToUserId = int64(m.ReceiverID)
		message.FromUserId = int64(m.SenderID)
		message.Content = m.Content
		message.CreateTime = int64(m.CreatedAt)
		out.MessageList[i] = &message
	}
	return nil
}
