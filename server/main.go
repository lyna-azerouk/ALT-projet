package main

import (
	"log"
	"net/http"
	"serveur/server/handlers"
	"serveur/server/middlewares"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/authtest", middlewares.AuthMiddleware, func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok authentified")
	})

	setUpAuthRoutes(router)

	setUpRestaurantRoutes(router)

	setUpOrderRoutes(router)
	return router
}

func setUpRestaurantRoutes(router *gin.Engine) {
	router.GET("/restaurants/:long/:lal/:radius", handlers.RestaurantsHandler)
	router.GET("/restaurant/:restaurantId", handlers.RestaurantDetailsHandler)
}

func setUpAuthRoutes(router *gin.Engine) {
	router.POST("/signup", handlers.RegistrationHandler)
	router.POST("/auth/client", handlers.ClientLoginHandler)
	router.POST("/auth/restaurant", handlers.RestaurantLoginHandler)
}

func setUpOrderRoutes(router *gin.Engine) {
	router.POST("/order", handlers.InitOrderHandler)
	router.GET("/order/:orderId", handlers.GetOrderHandler)
	router.PATCH("/orderpernding/:orderId", middlewares.VerifyOrderMiddleware, handlers.UpdatpendingOrderHandler)
	router.PATCH("/orderdelete/:orderId", middlewares.VerifyOrderMiddleware, handlers.UpdatDeleteOrderHandler)
	router.PATCH("/ordercompleted/:orderId", middlewares.AuthMiddleware, handlers.UpdatCompletedOrderHandler)
}

func main() {
	router := SetUpRouter()
	err := router.Run(":8080")
	if err == nil {
		log.Fatal("Error while starting the server: " + err.Error())
		return
	}
}
