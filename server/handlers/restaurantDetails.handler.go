package handlers

import (
	"serveur/server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RestaurantDetailsHandler(c *gin.Context) {
	var restaurantIdAsStr = c.Param("restaurantId")

	restaurantID, err := strconv.Atoi(restaurantIdAsStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid restaurant id"})
		return
	}
	restaurantDetails := services.GetRestaurantDetails(restaurantID)

	c.JSON(200, gin.H{"restaurant": restaurantDetails})
}
