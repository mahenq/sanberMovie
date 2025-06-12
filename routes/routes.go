package routes

import (
	"github.com/gin-gonic/gin"
	"sanberMovie/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/bioskop", controllers.CreateBioskop)
}
