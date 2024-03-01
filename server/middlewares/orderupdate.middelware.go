package middlewares

import (
	"net/http"
	"serveur/server/models"
	"serveur/server/services"
	JwtService "serveur/server/services/jwt"

	"github.com/gin-gonic/gin"
)

func VerifyOrderMiddleware(c *gin.Context) {
	token := extractToken(c)

	claims := JwtService.ParseAccessTokenResraurent(token)

	orderID := c.Param("orderId")

	var order models.OrderDetailsRequest

	order = services.GetOrderDetails(orderID)

	if order.RestaurantId != claims.Id {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Unauthorized to update status"})
		c.Abort()
		return
	}
	c.Next()
}
