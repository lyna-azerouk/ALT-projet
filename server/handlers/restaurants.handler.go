package handlers

import (
	"net/http"
	"serveur/server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RestaurantsHandler(c *gin.Context) {
	radius, latitude, longitude, notDone := parseCoordinates(c)
	if notDone {
		return
	}

	restaurants := services.RestaurantsAround(longitude, latitude, radius)

	c.JSON(http.StatusOK, gin.H{"restaurants": restaurants})
}

func parseCoordinates(c *gin.Context) (float64, float64, float64, bool) {
	latStr := c.Param("lal")
	longStr := c.Param("long")
	radiusAsStr := c.Param("radius")
	if radiusAsStr == "" {
		radiusAsStr = "1000"
	}

	radius, err := strconv.ParseFloat(radiusAsStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid radius"})
		return 0, 0, 0, true
	}
	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid altitude"})
		return 0, 0, 0, true
	}

	longitude, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid longitude"})
		return 0, 0, 0, true

	}
	return radius, latitude, longitude, false
}
