package main

import (
	"github.com/gin-gonic/gin"

	"SimpleURLShortener/handlers"
	"SimpleURLShortener/linkStorage"
)

func main() {
	linkStorage.InitPostgres()

	r := gin.Default()

	r.POST("/api/short", handlers.CreateShortLink)
	r.GET("/:code", handlers.ConnectLink)

	r.GET("/api/links", handlers.ListLinks)
	r.GET("/api/stats/:code", handlers.Stats)

	//nay la de tao san link bang trinh duyet
	r.GET("/api/shortByLink", handlers.CreateCodeByLink)

	r.Run(":8080")
}
