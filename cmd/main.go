package main

import (
	"github.com/Mekawy5/chatapp/conf"
	"github.com/Mekawy5/chatapp/registry"
	"github.com/Mekawy5/chatapp/workers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := conf.InitDB()
	defer db.Close()

	go workers.Consume(db)

	r := gin.Default()
	registry.Init(db, r)

	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}
