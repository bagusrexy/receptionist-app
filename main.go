package main

import (
	"log"

	"github.com/bagusrexy/test-dataon/config"
	"github.com/bagusrexy/test-dataon/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	log.Println("Connecting to PostgreSQL database...")
	db, err := config.CreateConnectionPostgres()
	if err != nil {
		log.Fatalf("Error creating connection to PostgreSQL: %s", err)
	}
	log.Println("Connected to PostgreSQL database")

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.Router(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting DB instance from GORM: %s", err)
	}
	defer postgresDB.Close()
}
