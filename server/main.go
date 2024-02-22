package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"serveur/server/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
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

	secret := os.Getenv("TOKEN_SECRET")
	fmt.Print("secret: " + secret)
}
