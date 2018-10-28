package main

import (
	"github.com/learning-go/testing-gin/app"
	"github.com/learning-go/testing-gin/database"
	"github.com/learning-go/testing-gin/models"
)

func main() {
	engine := app.GetEngine()
	database.InitDB("gorm_user", "123", "golang_store")
	database.DB.AutoMigrate(&models.Product{})

	app.InitializeRoutes(engine)

	engine.Run(":8080")
	defer database.DB.Close()
}
