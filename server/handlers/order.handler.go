package handlers

import (
	"net/http"
	"strconv"
	"serveur/server/models"
	"serveur/server/services"
	"fmt"
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

	code, err := services.GenerateCode(id_order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": 0, "message": "Failed to generate a code"})
	}
	c.JSON(http.StatusOK, gin.H{"Code": code})

}

/*
	Verficiation du code
*/

func VerfyOrderCode(c *gin.Context) {
	var id_order = c.Param("orderId")
	var code = c.Param("code")
	order, err := services.VerfyOrderCode(id_order, code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": 0, "message": "Incorrect code"})
	}

	c.JSON(http.StatusOK, gin.H{"Code": order})
}

/*
Get menu details by iD
*/

func GetMenuDetailsHandler(c *gin.Context) {
	menu_id_str := c.Param("menu_id")
	fmt.Println(menu_id_str)
	menu_id, err := strconv.ParseUint(menu_id_str, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid menu ID"})
		return
	}
	menu := services.GetMenuDetails(menu_id)

	c.JSON(http.StatusOK, gin.H{"Menu": menu})
}