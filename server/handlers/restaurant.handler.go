package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"serveur/server/models"
)

var restaurants = []models.Restaurant{
	{Id: 1, Name: "Restaurant A", Address: "Adresse A"},
	{Id: 2, Name: "Restaurant B", Address: "Adresse B"},
}

func Restaurants(c *gin.Context) {
	var localisation = c.Param("localisation")
	fmt.Print(localisation)
	var restaurantsInLocation []models.Restaurant

	for _, restaurant := range restaurants {
		if restaurant.Address == localisation {
			restaurantsInLocation = append(restaurantsInLocation, restaurant)
		}
	}
	c.JSON(http.StatusOK, restaurants)
}
