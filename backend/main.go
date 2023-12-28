// main.go
package main

import (
	"backend/routes" // Update with your actual package structure
	"github.com/gin-contrib/cors"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// Use the routes defined in the routes package
	routes.SetupRoutes(router)

	// Start the server
	err := router.Run(":3089")
	if err != nil {
		log.Fatal(err)
	}
}
