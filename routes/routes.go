package routes

import (
	"github.com/gin-gonic/gin"
	"sanberMovie/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/bioskop", controllers.CreateBioskop)
	r.GET("/bioskop", controllers.GetAllBioskop)
	r.GET("/bioskop/:id", controllers.GetBioskopByID)
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)
}
