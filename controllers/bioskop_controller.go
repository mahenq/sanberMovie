package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sanberMovie/config"
	"sanberMovie/models"
)

// POST /bioskop (Sudah ada)

// GET /bioskop
func GetAllBioskop(c *gin.Context) {
	var bioskops []models.Bioskop
	config.DB.Find(&bioskops)
	c.JSON(http.StatusOK, gin.H{"data": bioskops})
}

// GET /bioskop/:id
func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bioskop})
}

// PUT /bioskop/:id
func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	bioskop.Nama = input.Nama
	bioskop.Lokasi = input.Lokasi
	bioskop.Rating = input.Rating

	config.DB.Save(&bioskop)
	c.JSON(http.StatusOK, gin.H{"data": bioskop})
}

// DELETE /bioskop/:id
func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	config.DB.Delete(&bioskop)
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
