package controller

import (
	"net/http"
	"strconv"

	"github.com/Mekawy5/chatapp/domain/message"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	Srv     *gin.Engine
	Service *message.MessageService
}

func NewMessageController(srv *gin.Engine, s *message.MessageService) *MessageController {
	mc := &MessageController{
		Srv:     srv,
		Service: s,
	}

	mc.Srv.GET("/app/:app_token/chat/:chat_num/messages", mc.All)

	return mc
}

func (mc *MessageController) All(c *gin.Context) {
	chatNum, _ := strconv.Atoi(c.Param("chat_num"))
	apps := mc.Service.GetAll(c.Param("app_token"), uint(chatNum))

	c.JSON(http.StatusOK, gin.H{"apps": message.GetMessages(apps)})
}
