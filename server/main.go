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

	setUpUserRoutes(router)

	setUpAffluenceRoutes(router)

	return router
}

func setUpAffluenceRoutes(router *gin.Engine) {
	// getting affluences info
	router.GET("affluence/:restaurantId", middlewares.AuthMiddleware, handlers.GetAffluenceHandler)

	// updating affluences info
	router.PATCH("client/affluence/:restaurantId/:level", middlewares.AuthMiddleware, handlers.UpdateAffluenceWithClientVoteHandler)
	router.PATCH("restaurant/affluence/:restaurantId/:level", middlewares.AuthRestaurant, handlers.UpdateAffluenceWithRestaurantVoteHandler)
}

func setUpUserRoutes(router *gin.Engine) {
	router.GET("/client/:clientId", middlewares.AuthMiddleware, handlers.UserDetailsHandler)
}

func setUpRestaurantRoutes(router *gin.Engine) {
	router.GET("/restaurants/:long/:lal/:radius", middlewares.AuthMiddleware, handlers.RestaurantsHandler)
	router.GET("/bouffluence/restaurants/", middlewares.AuthMiddleware, handlers.AllBouffluenceRestaurantsHandler)
	router.GET("/restaurant/:restaurantId", middlewares.AuthMiddleware, handlers.RestaurantDetailsHandler)

	router.POST("/restaurant/menu/:restaurantId/", middlewares.AuthRestaurant, handlers.AddMenuItemHandler)
}

func setUpAuthRoutes(router *gin.Engine) {
	router.POST("/signup/client", handlers.ClientRegistrationHandler)
	router.POST("/signup/restaurant", handlers.RestaurantRegistrationHandler)
	router.POST("/auth/client", handlers.ClientLoginHandler)
	router.POST("/auth/restaurant", handlers.RestaurantLoginHandler)
}

func setUpOrderRoutes(router *gin.Engine) {
	// cree une commande
	router.POST("/order", middlewares.AuthMiddleware, handlers.InitOrderHandler)
	// details commandes
	router.GET("/order/:orderId", middlewares.OrderAuth, handlers.GetOrderHandler)
	// accepter commande (restaurant)
	router.PATCH("/order/accept/:orderId", middlewares.VerifyOrderMiddleware, handlers.UpdatePendingOrderHandler)
	// commande prête (restaurant)
	router.PATCH("/order/ready/:orderId", middlewares.VerifyOrderMiddleware, handlers.UpdateInProgressOrderHandler)
	// commande terminée (restaurant)
	router.PATCH("/order/complete/:orderId", middlewares.VerifyOrderMiddleware, handlers.UpdatCompletedOrderHandler)
	// supprimer commande (????)
	router.PATCH("/order/delete/:orderId", middlewares.OrderAuth, handlers.UpdatDeleteOrderHandler)
	// récupérer les commandes d'un utilisateur (client)
	router.GET("/order/pick/:orderId", middlewares.OrderClientAuth, handlers.PickOrder)
	// vérifier le code de confirmation (restaurant)
	router.POST("/order/pick/:orderId/:code", middlewares.VerifyOrderMiddleware, handlers.VerfyOrderCode)
	router.GET("order/user/:userId", middlewares.AuthMiddleware, handlers.GetOrdersHandler)
	router.GET("order/restaurant/:restaurantId", middlewares.AuthMiddleware, handlers.GetRestaurantOrdersHandler)
}

func main() {
	router := SetUpRouter()
	err := router.Run(":8080")
	if err == nil {
		log.Fatal("Error while starting the server: " + err.Error())
		return
	}
}
