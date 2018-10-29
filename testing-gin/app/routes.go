package app

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-go/testing-gin/controllers"
)

func InitializeRoutes(r *gin.RouterGroup) {
	r.POST("/sell", controllers.Sell)
	r.GET("/list", controllers.List)
}
