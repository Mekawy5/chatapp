package main

import (
	"github.com/Mekawy5/chatapp/conf"
	"github.com/Mekawy5/chatapp/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	db := conf.InitDB()
	defer db.Close()

	rmqc := tools.NewRabbitClient()
	go rmqc.Setup()

	defer rmqc.Conn.Close()

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Works..",
		})
	})

	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}
