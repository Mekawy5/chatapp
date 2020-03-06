package conf

import (
	"os"

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

	return db
}
