package main

import (
	"net/http"
	"serveur/server/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello word")
	})

	// configure routes
	router.GET("/restaurants/:localisation", handlers.Restaurants)
	router.GET("/restaurant/:id", handlers.Restaurant_details)

	router.POST("/signup", handlers.RegistrationHandler)
	router.POST("/login", handlers.LoginHandler)
	//router.POST("/login_validation/:code_validation", handlers.Login_validation)
	router.Run(":8080")
}
