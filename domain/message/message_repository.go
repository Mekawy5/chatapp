package message

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		DB: db,
	}
}

func (m *MessageRepository) GetAll(appToken string, chatNum uint) []MessageModel {
	var msgs []MessageModel
	chatId := m.GetChatId(appToken, chatNum)

	if chatId == 0 {
		return []MessageModel{}
	}

	m.DB.Where("chat_id = ?", chatId).Find(&msgs)
	return msgs
}

func (a *MessageRepository) Save(msg MessageModel) MessageModel {
	a.DB.Save(&msg)
	return msg
}

func (a *MessageRepository) GetChatId(appToken string, chatNum uint) uint {
	var chat struct{ ID uint }
	a.DB.Table("chats").Select("chats.id").Joins("left join applications on chats.application_id = applications.id").Where("applications.token = ?", appToken).Where("chats.number = ?", chatNum).Scan(&chat)
	return chat.ID
}
