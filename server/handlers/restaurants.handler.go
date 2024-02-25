package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Restaurants(c *gin.Context) {
	latStr := c.Param("lal")
	longStr := c.Param("long")

	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid altitude"})
		return
	}

	longitude, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid longitude"})
		return
	}

	query := fmt.Sprintf(`[out:json];
        node["amenity"="restaurant"](bbox:%f,%f,%f,%f);
        out;`, latitude-0.1, longitude-0.1, latitude+0.1, longitude+0.1)

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
