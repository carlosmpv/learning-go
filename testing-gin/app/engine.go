package app

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	e := gin.Default()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(static.Serve("/", static.LocalFile("testing-gin/client", true)))

	// config := cors.DefaultConfig()
	// config.AllowMethods = []string{"GET, POST"}
	// config.AllowAllOrigins = true

	// e.Use(cors.New(config))

	return e
}
