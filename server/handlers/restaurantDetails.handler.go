package handlers

import (
	"fmt"
	"net/http"
	"serveur/server/models"

	"github.com/gin-gonic/gin"
)

var menus = []models.Restaurant_Menu{
	{Id: 1, Name: "Menu A", Price: 15},
	{Id: 2, Name: "Menu B", Price: 20},
	{Id: 2, Name: "Menu B", Price: 20},
}

func Restaurant_details(c *gin.Context) {
	var id_restaurant = c.Param("id")
	fmt.Print(id_restaurant)
	//Get the menus of the restaurant from the databsae
	// si le id n'existe pas en bdd renvoie uen erreur.

	c.JSON(http.StatusOK, menus)
}
