package chat

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChatService struct {
	Repository *ChatRepository
}

func NewChatService(r *ChatRepository) *ChatService {
	return &ChatService{
		Repository: r,
	}
}

// func (s *ChatService) Create(chat ChatModel, appToken string) ChatModel {
// 	redisClient := tools.NewRedisClient()
// 	chatNum := redisClient.GetAppChatNumber(appToken)
// 	chat.Number = uint(chatNum + 1)

// 	appId := s.Repository.GetAppId(appToken)
// 	chat.ApplicationID = appId

// 	newChat := s.Repository.Save(chat)

// 	redisClient.SetAppChatNumber(appToken, chat.Number)
// 	return newChat
// }

func (s *ChatService) GetAll() []ChatModel {
	return s.Repository.GetAll()
}

func (s *ChatService) Get(id uint) ChatModel {
	return s.Repository.Get(id)
}

func (s *ChatService) Delete(number uint) bool {
	return s.Repository.Delete(number)
}
