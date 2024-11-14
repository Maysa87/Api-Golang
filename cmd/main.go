package main

import (
	_ "database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	dsn := "postgresql://MaysaLeocadio:0807@localhost:5432/postgres?sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.AutoMigrate(&album{})
	if err != nil {
		return
	}
}

type album struct {
	gorm.Model
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
	var albums []album
	db.Find(&albums)
	c.JSON(200, albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	db.Create(&newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var IdAlbum album
	if result := db.First(&IdAlbum, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}
	c.JSON(http.StatusOK, IdAlbum)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var deleteAlbum album
	db.Delete(&deleteAlbum, id)
	c.Status(204)
}
