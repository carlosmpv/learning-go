package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/learning-go/testing-gin/database"
	"github.com/learning-go/testing-gin/models"
)

func Sell(c *gin.Context) {
	price, _ := strconv.ParseFloat(c.PostForm("price"), 10)

	product := &models.Product{
		Name:  c.PostForm("name"),
		Price: price,
	}

	database.DB.Create(product)
	c.Redirect(302, "/list")
}
