package middlewares

import (
	"net/http"
	services "serveur/server/services/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthRestaurant(c *gin.Context) {

	restaurantId, err := strconv.Atoi(c.Param("restaurantId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid restaurant id"})
		c.Abort()
		return
	}

	token := extractToken(c)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "No token provided"})
		c.Abort()
		return
	}

	claims := services.ParseRestaurantAccessToken(token)
	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid token"})
		c.Abort()
		return
	}

	if claims.Id != restaurantId {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
