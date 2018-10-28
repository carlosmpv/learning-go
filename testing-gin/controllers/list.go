package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-go/testing-gin/database"
	"github.com/learning-go/testing-gin/models"
)

func List(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}
