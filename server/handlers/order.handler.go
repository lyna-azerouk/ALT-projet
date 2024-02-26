package handlers

import (
	"net/http"
	"serveur/server/models"
	"serveur/server/services"

	"github.com/gin-gonic/gin"
)

/*
Create a new order
*/
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

/*
Update status Order
*/

func UpdatOrderHandler(c *gin.Context) {

	var id_order = c.Param("orderId")

	order := services.UpdateStatusOrder(id_order)

	if order {
		// get the information of the order
		// make a get request ot get all the details of a restaurent from its id
		c.JSON(http.StatusOK, gin.H{"message": "Order in progress"})
	} else {
		c.JSON(400, gin.H{"message": "Order not in progress"})
	}
}
