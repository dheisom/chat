package main

import (
	"log"
	"os"
	"server/api"
	"server/database"

	"github.com/gin-contrib/static"
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
	web := os.Getenv("WEB_PATH")
	if web == "" {
		web = "../web/build"
	}
	server := gin.Default()
	server.Use(static.Serve("/", static.LocalFile(web, false)))
	api.RegistryAll(server.Group("/api"))
	log.Printf("Starting server at port %s...\n", port)
	server.Run(":" + port)
}
