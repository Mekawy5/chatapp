package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Init setup app dependencies.
func Init(db *gorm.DB, srv *gin.Engine) {
	InitApplicationController(db, srv)
	InitChatController(db, srv)
	InitMessageController(db, srv)
}
