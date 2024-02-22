package handlers

import (
	"net/http"
	"serveur/server/models"

	"github.com/gin-gonic/gin"
)

var restaurants = []models.Restaurant{
	{Id: 1, Name: "Restaurant A", Address: "paris"},
	{Id: 2, Name: "Restaurant B", Address: "Adresse B"},
}

// mette en plca une structure address
func Restaurants(c *gin.Context) {
	var localisation = c.Param("localisation")
	var restaurantsInLocation []models.Restaurant

	for _, restaurant := range restaurants {
		if restaurant.Address == localisation {
			restaurantsInLocation = append(restaurantsInLocation, restaurant)
		}
	}
	c.JSON(http.StatusOK, restaurantsInLocation)
}
