package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maru1005/lulu-de-go/internal/repository"
	"github.com/maru1005/lulu-de-go/pkg/config"
)

func main() {
	cfg := config.Load()
	db := repository.NewDB(cfg.DatabaseURL)
	defer db.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	log.Println("server starting on :8080")
	r.Run(":8080")
}
