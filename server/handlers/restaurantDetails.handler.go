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


func AffluenceHandler( c *gin.Context) {
	var restaurantId = c.Param("restaurantId")
    restaurantID, err := strconv.ParseUint(restaurantId, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid restaurant id"})
		return
	}
	affluence := services.GetAffluence(restaurantID)

	c.JSON(200, gin.H{"Affluence": affluence})
}
