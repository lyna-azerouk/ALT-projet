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
