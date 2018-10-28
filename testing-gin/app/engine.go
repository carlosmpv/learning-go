package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	e := gin.Default()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	e.Use(cors.New(config))

	return e
}
