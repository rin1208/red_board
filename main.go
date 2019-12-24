package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))

	r.POST("/", handler.Request_music)
	r.GET("/", handler.Music_List)

	r.Run()
}
