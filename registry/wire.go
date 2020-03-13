//+build wireinject

package registry

import (
	"github.com/Mekawy5/chatapp/controller"
	"github.com/Mekawy5/chatapp/domain/application"
	"github.com/Mekawy5/chatapp/domain/chat"
	"github.com/Mekawy5/chatapp/domain/message"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitApplicationController(db *gorm.DB, srv *gin.Engine) *controller.ApplicationController {
	wire.Build(controller.NewApplicationController, application.NewApplicationService, application.NewApplicationRepository)
	return &controller.ApplicationController{}
}

func InitChatController(db *gorm.DB, srv *gin.Engine) *controller.ChatController {
	wire.Build(controller.NewChatController, chat.NewChatService, chat.NewChatRepository)
	return &controller.ChatController{}
}

func InitMessageController(db *gorm.DB, srv *gin.Engine) *controller.MessageController {
	wire.Build(controller.NewMessageController, message.NewMessageService, message.NewMessageRepository)
	return &controller.MessageController{}
}
