package chat

import (
	"github.com/Mekawy5/chatapp/domain/message"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChatModel struct {
	gorm.Model
	ApplicationID uint
	Number        uint
	MessagesCount uint
	Messages      []message.MessageModel `gorm:"foreignkey:ChatID"`
}

func (ChatModel) TableName() string {
	return "chats"
}
