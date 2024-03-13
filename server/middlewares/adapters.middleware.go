package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"serveur/server/adapters"
	"serveur/server/models"

	"github.com/gin-gonic/gin"
)

func OrderDetailsAdapterMiddleware(c *gin.Context) {
	var orderRequest models.OrderDetailsRequest
	var adaptedOrderRequest models.OrderDetails
	log.Println("Adapting order details request")

	if err := c.BindJSON(&orderRequest); err != nil {
		println(c.Request.Body)
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid order request"})
		return
	}

	adaptedOrderRequest = adapters.OrderRequestToOrderDetailsMapper(orderRequest)

	// Stocke le nouveau corps de la requête adaptée dans le contexte Gin
	c.Set("adaptedRequestBody", adaptedOrderRequest)

	// Réinitialise le corps de la requête avec le corps adapté
	adaptedRequestBodyBytes, err := json.Marshal(adaptedOrderRequest)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"success": 0, "message": "Error marshalling adapted request body"},
		)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(adaptedRequestBodyBytes))
	c.Next()
}
