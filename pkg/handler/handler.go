package handler

import (
	"net/url"
	"red_board/pkg/api"
	"red_board/pkg/repositry"
	"red_board/pkg/domain"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"

)

func Request(c *gin.Context){
	var shell []domain.Shelter
	rep, _ := repositry.NewDb()
	defer rep.Close()


}