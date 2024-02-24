package main

import (
	"fmt"
	"net/http"
	"os"
	"serveur/server/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", handlers.AuthMiddleware, func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello word")
	})

	// configure routes
	router.GET("/restaurants/:localisation", handlers.Restaurants)
	router.GET("/restaurant/:restaurant_id", handlers.Restaurant_details)

	router.POST("/signup", handlers.RegistrationHandler)
	router.POST("/login", handlers.LoginHandler)
	//router.POST("/login_validation/:code_validation", handlers.Login_validation)
	router.Run(":8080")

	secret := os.Getenv("TOKEN_SECRET")
	fmt.Print("secret: " + secret)
}
