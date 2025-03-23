package main

import (
	"log"
	"os"
	"strconv"

	"github.com/alexanderlesser/golf-tracker/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portStr := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		log.Fatalf("Invalid DB_PORT value: %s", err)
	}

	mysqlStorage, err := db.DbConnect(db.MysqlConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     uint(port),
		Host:     os.Getenv("DB_HOST"),
	})

	if err != nil {
		log.Fatalf("Failed to create MySQL storage: %s", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "golf-tracker",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Println(mysqlStorage)
	log.Println("Server running on port 8080")
	r.Run("localhost:8080")
}
