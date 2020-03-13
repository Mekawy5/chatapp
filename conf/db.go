package conf

import (
	"os"

	"github.com/Mekawy5/chatapp/domain/application"
	"github.com/Mekawy5/chatapp/domain/chat"
	"github.com/Mekawy5/chatapp/domain/message"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	var dbUrl string
	if url := os.Getenv("DB_URL"); url == "" {
		dbUrl = "root:123@tcp(db:3306)/chat?charset=utf8&parseTime=True&loc=Local"
	} else {
		dbUrl = url
	}

	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&application.ApplicationModel{})
	db.AutoMigrate(&chat.ChatModel{}).AddForeignKey("application_id", "applications(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&message.MessageModel{}).AddForeignKey("chat_id", "chats(id)", "CASCADE", "CASCADE")

	return db
}
