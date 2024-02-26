package handlers

import (
	"net/http"
	"serveur/server/models"
	"serveur/server/services"

	"github.com/gin-gonic/gin"
)

func InitOrderHandler(c *gin.Context) {
	var orderRequest models.OrderDetailsRequest

	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid order request"})
		return
	}

	err := services.CreateNewOrder(orderRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": 0, "message": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"body": orderRequest})
}
