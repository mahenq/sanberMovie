package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sanberMovie/config"
	"sanberMovie/models"
)

func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	config.DB.Create(&bioskop)
	c.JSON(http.StatusOK, gin.H{"data": bioskop})
}
