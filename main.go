package main

import (
	"red_board/pkg/handler"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))

	r.GET("/request", handler.Request)
	r.POST("/massage", handler.Massage)

	r.Run()
}
