// routes/routes.go
package routes

import (
	"backend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.Use(cors.Default())

	// Define a route to fetch data from the API
	router.GET("/artists", handlers.GetArtistsHandler)

	// Define a route to fetch an artist by ID
	router.GET("/artists/:id", handlers.GetArtistByIDHandler)
}
