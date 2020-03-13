package controller

import (
	"net/http"
	"strconv"

	"github.com/Mekawy5/chatapp/domain/chat"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	Srv     *gin.Engine
	Service *chat.ChatService
}

func NewChatController(srv *gin.Engine, s *chat.ChatService) *ChatController {
	cc := &ChatController{
		Srv:     srv,
		Service: s,
	}

	cc.Srv.GET("/chats", cc.All)
	cc.Srv.GET("/chats/:number", cc.Find)
	cc.Srv.DELETE("/chats/:number", cc.Delete)

	return cc
}

func (cc *ChatController) All(c *gin.Context) {
	chats := cc.Service.GetAll()
	c.JSON(http.StatusOK, gin.H{"apps": chat.GetChats(chats)})
}

func (cc *ChatController) Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ch := cc.Service.Get(uint(id))
	c.JSON(http.StatusOK, gin.H{"app": chat.GetChat(ch)})
}

func (cc *ChatController) Delete(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("number"))
	del := cc.Service.Delete(uint(num))
	if del {
		c.JSON(http.StatusOK, gin.H{"message": "Chat Deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Error deleting chat"})
	}
}
