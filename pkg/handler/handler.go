package handler

import (
	"red_board/pkg/domain"
	"red_board/pkg/repositry"

	"github.com/gin-gonic/gin"
)

func Request(c *gin.Context) {
	var massages []domain.Shelter
	rep, _ := repositry.NewDb()
	defer rep.Close()
	rep.Find(&massages)
	c.JSON(200, &massages)

}
func Massage(c *gin.Context) {
	var massages domain.Shelter
	rep, _ := repositry.NewDb()
	defer rep.Close()
	rep.Create(&massages)

}