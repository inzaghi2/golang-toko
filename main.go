package main

import (
	"golang-toko/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Rute CRUD Produk
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	router.Run(":8080")
}
