package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-go/testing-gin/database"
	"github.com/learning-go/testing-gin/models"
)

func Sell(c *gin.Context) {

	var product models.Product
	c.Bind(&product)
	// log.Fatal(c.GetRawData())
	database.DB.Create(&product)
	c.JSON(200, product)

	// c.Redirect(http.StatusPermanentRedirect, "/")

}
