package main

import (
	"github.com/FeelDat/go-testtask/internals/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	log.Println("Starting the app")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Run(":" + port)
}
