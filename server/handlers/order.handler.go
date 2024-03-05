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

	order, err := services.CreateNewOrder(orderRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": 0, "message": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"body": order})
}

/*
Update status Order to IN_PROGRESS
*/
func UpdatpendingOrderHandler(c *gin.Context) {
	var id_order = c.Param("orderId")

	order := services.UpdateStatusOrder(id_order, "IN_PROGRESS")

	c.JSON(http.StatusOK, gin.H{"message": "Order in progress", "order": order})
}

/*
Update status Order to DECLINED and delte the  order from database
*/
func UpdatDeleteOrderHandler(c *gin.Context) {
	var id_order = c.Param("orderId")

	order := services.UpdateStatusOrder(id_order, "DECLINED")

	c.JSON(http.StatusOK, gin.H{"message": "Order Deleted", "order": order})
}

/*
Update status Order to COMPLETED
*/
func UpdatCompletedOrderHandler(c *gin.Context) {
	var id_order = c.Param("orderId")

	order := services.UpdateStatusOrder(id_order, "COMPLETED")

	c.JSON(http.StatusOK, gin.H{"message": "Order  Completed", "order": order})
}

/*
 Générate Code pour le client
*/

func PickOrder(c *gin.Context) {
	var id_order = c.Param("orderId")

	code := services.GenerateCode(id_order)
	c.JSON(http.StatusOK, gin.H{"Code": code})

}

/*
	Verficiation du code
*/

func VerfyOrderCode(c *gin.Context) {
	var id_order = c.Param("orderId")
	var code = 112345678
	order := services.VerfyOrderCode(id_order, code)

	c.JSON(http.StatusOK, gin.H{"Code": order})

}
