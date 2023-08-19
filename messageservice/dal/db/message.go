package db

import "context"

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"to_user_id"`
	ReceiverID int    `json:"from_user_id"`
	Content    string `json:"content"`
	CreatedAt  int    `json:"create_time"`
}

func (message *Message) TableName() string {
	return "message"
}

func CreateMessage(ctx context.Context, message *Message) error {
	return DB.WithContext(ctx).Create(message).Error
}

func GetChatRecords(ctx context.Context, to_user_id, from_user_id int) ([]*Message, error) {
	messages := make([]*Message, 0)
	result := DB.WithContext(ctx).Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", to_user_id, from_user_id, from_user_id, to_user_id).Order("create_at").Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}
