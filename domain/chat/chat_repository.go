package chat

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{
		DB: db,
	}
}

func (c *ChatRepository) GetAll() []ChatModel {
	var chats []ChatModel
	c.DB.Preload("Chats").Find(&chats)
	return chats
}

func (c *ChatRepository) Get(id uint) ChatModel {
	var chat ChatModel
	c.DB.Find(&chat, id)
	return chat
}

func (c *ChatRepository) Save(chat ChatModel) ChatModel {
	c.DB.Save(&chat)
	return chat
}

func (c *ChatRepository) Delete(number uint) bool {
	del := c.DB.Where("number = ?", number).Delete(ChatModel{}).RowsAffected
	if del > 0 {
		return true
	}
	return false
}

func (c *ChatRepository) GetAppId(token string) uint {
	var app struct{ ID uint }
	c.DB.Table("applications").Select("id").Where("token = ?", token).Scan(&app)
	return app.ID
}
