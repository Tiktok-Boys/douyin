package handler

import (
	"context"
	"net/http"

	proto "github.com/Tiktok-Boys/douyin/api/proto/message"
	"github.com/gin-gonic/gin"
)

var (
	MessageServiceClient proto.MessageService
)

func ActMessageHandler(c *gin.Context) {
	user_id := c.MustGet("uid").(string)
	to_user_id := c.Query("to_user_id")
	// action_type := c.Query("action_type")
	content := c.Query("content")
	res, err := MessageServiceClient.ActMessage(context.TODO(), &proto.DouyinMessageActionRequest{
		UserId:     user_id,
		ToUserId:   to_user_id,
		ActionType: 1,
		Content:    content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -2,
			"message": "MessageService.ActMessage: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"status_msg":  res.StatusMsg,
	})
}

func GetChatHandler(c *gin.Context) {
	user_id := c.MustGet("uid").(string)
	to_user_id := c.Query("to_user_id")
	res, err := MessageServiceClient.GetChat(context.TODO(), &proto.DouyinMessageChatRequest{
		UserId:   user_id,
		ToUserId: to_user_id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -2,
			"message": "MessageService.GetChat: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":  res.StatusCode,
		"status_msg":   res.StatusMsg,
		"message_list": res.MessageList,
	})
}
