package handlers

import (
	"fmt"
	"log"
	"net/http"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
	"time"

	"github.com/gin-gonic/gin"
)

func InitOrderHandler(c *gin.Context) {
	var orderRequest models.OrderDetailsRequest

	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": "Invalid order request"})
		return
	}

	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	// 1st step: insert order in the database
	query := requests.InsertNewOrderRequestTemplate
	var orderId string

	err = db.QueryRow(
		query,
		orderRequest.ClientId,
		orderRequest.RestaurantId, 0, "PENDING", time.Now()).Scan(&orderId)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": err.Error()})
		return
	}

	fmt.Println("id: " + (orderId))

	//// 2nd step: insert order details in the database
	query = requests.InsertNewOrderItemRequestTemplate
	for _, item := range orderRequest.OrderItems {
		_, err = db.Exec(query, orderId, item.MenuId, item.Count)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"status": err})
			return
		}
	}
	db.Close()

	c.JSON(http.StatusOK, gin.H{"body": orderRequest})

}
