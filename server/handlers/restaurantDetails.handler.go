package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"serveur/server/database"
	"serveur/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Restaurant_details(c *gin.Context) {
	var restaurant_id = c.Param("restaurant_id")

	restaurantID, err := strconv.Atoi(restaurant_id)

	query := fmt.Sprintf("[out:json];node(%d);out;", restaurantID)
	apiUrl := "https://overpass-api.de/api/interpreter?data=" + url.QueryEscape(query)

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	/**
	Get all ths menus of the restaurent(restaurent_id) from the table Menus
	**/
	var menus []models.Restaurant_Menu
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM menus WHERE restaurant_id = $1", restaurantID)
	for rows.Next() {
		var menu models.Restaurant_Menu

		err := rows.Scan(&menu.Id, &menu.Name, &menu.RestaurantID, &menu.Price, &menu.Description, &menu.Image)
		if err != nil {
			log.Fatal(err)
		}
		menus = append(menus, menu)
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Invalid credentials"})
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	var bodyMap map[string]interface{}

	_ = json.Unmarshal(body, &bodyMap)

	var firstElement map[string]interface{}

	elements, ok := bodyMap["elements"].([]interface{})

	if ok && len(elements) > 0 {
		firstElement = elements[0].(map[string]interface{})
	}

	if firstElement != nil {
		firstElement["menus"] = menus
	}

	modifiedBody, err := json.Marshal(firstElement)

	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal modified JSON"})
		return
	}

	c.Data(http.StatusOK, "application/json", modifiedBody)
}
