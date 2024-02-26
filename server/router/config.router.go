package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"serveur/server/handlers"
	"serveur/server/middlewares"
)

func setUpRestaurantRoutes(router *gin.Engine) {
	router.GET("/restaurants/:long/:lal/:radius", handlers.RestaurantsHandler)
	router.GET("/restaurant/:restaurantId", handlers.RestaurantDetailsHandler)
}

func setUpAuthRoutes(router *gin.Engine) {
	router.POST("/signup", handlers.RegistrationHandler)
	router.POST("/auth/client", handlers.ClientLoginHandler)
	router.POST("/auth/restaurant", handlers.RestaurantLoginHandler)
}

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/authtest", middlewares.AuthMiddleware, func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok authentified")
	})

	setUpAuthRoutes(router)

	setUpRestaurantRoutes(router)

	return router
}
