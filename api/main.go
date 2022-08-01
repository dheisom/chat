package main

import (
	"api/database"
	"api/server"
	"log"
	"os"

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
	log.Printf("Starting server at port %s...\n", port)
	server.Start(port)
}
