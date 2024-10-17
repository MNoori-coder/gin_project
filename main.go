package main

import (
	"fmt"
	"log"
	"os"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	server.Run(addr)
}
