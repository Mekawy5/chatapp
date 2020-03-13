package application

import (
	"github.com/Mekawy5/chatapp/domain/chat"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ApplicationModel struct {
	gorm.Model
	Name       string
	ChatsCount uint
	Token      string           `gorm:"unique_index;not null"`
	Chats      []chat.ChatModel `gorm:"foreignkey:ApplicationID"`
}

func (ApplicationModel) TableName() string {
	return "applications"
}
