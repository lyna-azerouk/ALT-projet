package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	//Insert into bdd?
	c.Data(http.StatusOK, "application/json", body)
}
