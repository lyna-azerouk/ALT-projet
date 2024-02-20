package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"serveur/server/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello word")
	})
	router.GET("/restaurants/:localisation", handlers.Restaurants)
	// configure routes
	//router.POST("/login", handlers.LoginHandler)
	//router.POST("/login_validation/:code_validation", handlers.Login_validation)
	router.Run(":8080")
}
