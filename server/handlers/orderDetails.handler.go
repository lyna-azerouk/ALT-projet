package handlers

import (
	"net/http"
	"serveur/server/services"

	"github.com/gin-gonic/gin"
)

func GetOrderHandler(c *gin.Context) {
	var id_order = c.Param("orderId")

	order := services.GetOrderDetails(id_order)

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrdersHandler(c *gin.Context) {
	var uderId = c.Param("userId")

	orders, err := services.GetUserOrdersDetails(uderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": 0, "message": "Failed to get user Orders"})
	}
	c.JSON(http.StatusOK, gin.H{"User orders": orders})
}
