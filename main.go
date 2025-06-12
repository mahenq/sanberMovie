package main

import (
	"github.com/gin-gonic/gin"
	"sanberMovie/config"
	"sanberMovie/models"
	"sanberMovie/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Bioskop{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
