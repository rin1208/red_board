package handler

import (
	"red_board/pkg/repositry"

	"github.com/gin-gonic/gin"
)

func Request(c *gin.Context) {
	rep, _ := repositry.NewDb()
	defer rep.Close()

}
