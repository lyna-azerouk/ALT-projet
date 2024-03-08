package handlers

import (
	"log"
	"net/http"
	"serveur/server/services"

	"github.com/gin-gonic/gin"
)

func UserDetailsHandler(c *gin.Context) {
	userId := c.Param("clientId")
	log.Println("user details by id: " + userId)
	user, err := services.GetUserDetails(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": 0, "message": "Failed to get user details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"body": user})
}
