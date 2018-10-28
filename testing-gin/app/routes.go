package app

import (
	"github.com/gin-gonic/gin"
	"github.com/learning-go/testing-gin/controllers"
)

func InitializeRoutes(engine *gin.Engine) {
	engine.POST("/sell", controllers.Sell)
	engine.GET("/list", controllers.List)
}
