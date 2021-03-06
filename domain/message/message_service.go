package message

import (
	"github.com/Mekawy5/chatserv/tools"
	"github.com/Mekawy5/chatserv/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MessageService struct {
	Repository *MessageRepository
}

func NewMessageService(r *MessageRepository) *MessageService {
	return &MessageService{
		Repository: r,
	}
}

func (s *MessageService) Create(msg MessageModel, appToken string, chatNum uint) MessageModel {
	appChatKey := util.GenerateAppChatKey(appToken, chatNum)
	redisClient := tools.NewRedisClient()
	lastMsgNum, chatId := redisClient.GetAppChatInfo(appChatKey)

	if lastMsgNum == 0 {
		msg.Number = 1
	} else {
		msg.Number = lastMsgNum + 1
	}

	if chatId == 0 {
		msg.ChatID = s.Repository.GetChatId(appToken, chatNum)
	} else {
		msg.ChatID = chatId
	}

	redisClient.SetAppChatInfo(appChatKey, msg.Number, msg.ChatID)
	//TODO save to database.
	return msg
}

func (s *MessageService) GetAll(appToken string, chatNum uint) []MessageModel {
	return s.Repository.GetAll(appToken, chatNum)
}
