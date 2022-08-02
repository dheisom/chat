package main

import (
	"log"
	"os"
	"server/api"
	"server/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
)

func main() {
	godotenv.Load(".env")
	database.Start(sqlite.Open("database.db"))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := gin.Default()
	api.RegistryAll(server.Group("/api"))
	log.Printf("Starting server at port %s...\n", port)
	server.Run(":" + port)
}
