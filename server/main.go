package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"serveur/server/handlers"
	"serveur/server/middlewares"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", middlewares.AuthMiddleware, func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello word")
	})

	// configure routes
	router.GET("/restaurants/:localisation", handlers.Restaurants)
	router.GET("/restaurant/:id", handlers.Restaurant_details)

	router.POST("/signup", handlers.RegistrationHandler)
	router.POST("/auth/client", handlers.ClientLoginHandler)
	router.POST("/auth/restaurant", handlers.RestaurantLoginHandler)
	router.Run(":8080")

	secret := os.Getenv("TOKEN_SECRET")
	fmt.Print("secret: " + secret)
}
