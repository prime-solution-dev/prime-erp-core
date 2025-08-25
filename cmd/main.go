package main

import (
	"log"

	"github.com/prime-solution-dev/prime-erp-core/internal/middleware"
	"github.com/prime-solution-dev/prime-erp-core/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}

	ginEngine := gin.Default()

	middleware.RegisterMiddlewares(ginEngine)

	routes.RegisterRoutes(ginEngine)

	port := "9199"
	log.Printf("Starting server on port %s\n", port)
	if err := ginEngine.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
