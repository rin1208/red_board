package handler

import (
	"fmt"
	"red_board/pkg/domain"
	"red_board/pkg/repositry"

	"github.com/gin-gonic/gin"
)

func Request(c *gin.Context) {
	var massages []domain.Data
	rep, _ := repositry.NewDb()
	defer rep.Close()
	rep.Find(&massages)
	c.JSON(200, &massages)

}
func Message(c *gin.Context) {
	var massages domain.Data
	rep, _ := repositry.NewDb()
	defer rep.Close()
	c.BindJSON(&massages)
	fmt.Println(massages)
	rep.Create(&massages)

}
