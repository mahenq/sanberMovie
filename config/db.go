package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=sanbermovie_db port=5432 sslmode=disable"
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	fmt.Println("Database berhasil terkoneksi!")
}
