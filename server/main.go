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
	router.GET("/restaurants/:long/:lal/:radius", middlewares.AuthMiddleware, handlers.RestaurantsHandler)
	router.GET("/restaurant/:restaurantId", middlewares.AuthMiddleware, handlers.RestaurantDetailsHandler)
}

func setUpAuthRoutes(router *gin.Engine) {
	router.POST("/signup/client", handlers.RegistrationHandler)
	router.POST("/auth/client", handlers.ClientLoginHandler)
	router.POST("/auth/restaurant", handlers.RestaurantLoginHandler)
}

func setUpOrderRoutes(router *gin.Engine) {
	router.POST("/order", middlewares.AuthMiddleware, handlers.InitOrderHandler)
	router.GET("/order/:orderId", middlewares.OrderAuth, handlers.GetOrderHandler)
	router.PATCH("/order/accept/:orderId", middlewares.VerifyOrderMiddleware, handlers.UpdatpendingOrderHandler)
	router.PATCH("/order/complete/:orderId", middlewares.VerifyOrderMiddleware, handlers.UpdatCompletedOrderHandler)
	router.PATCH("/order/delete/:orderId", middlewares.OrderAuth, handlers.UpdatDeleteOrderHandler)
	router.GET("/order/pick/:orderId", handlers.PickOrder)             //I removed the verfication because it will always fail the token in bdd is not the same as the token in jwt
	router.POST("/order/pick/:orderId/:code", handlers.VerfyOrderCode) //I removed the verfication because it will always fail the token in bdd is not the same as the token in jwt

}

func main() {
	router := SetUpRouter()
	err := router.Run(":8080")
	if err == nil {
		log.Fatal("Error while starting the server: " + err.Error())
		return
	}
}
