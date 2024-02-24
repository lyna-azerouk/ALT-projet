package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// mette en plca une structure address
func Restaurants(c *gin.Context) {
	var location = c.Param("localisation")

	query := fmt.Sprintf(`[out:json];
	area["name"="%s"] -> .area;
	node["amenity"="restaurant"](area.area);
	out;`, location)

	apiUrl := fmt.Sprintf("https://overpass-api.de/api/interpreter?data=%s", url.QueryEscape(query))

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

	c.Data(http.StatusOK, "application/json", body)

}
