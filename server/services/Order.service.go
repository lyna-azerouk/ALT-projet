package services

import (
	"database/sql"
	"fmt"
	"log"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
	"sync"
	"time"
)

func CreateNewOrder(orderRequest models.OrderDetailsRequest) error {
	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	/**
	Get the price of the order
	*/
	price := GetPrice(db, orderRequest.OrderItems)

	// 1st step: insert order in the database
	query := requests.InsertNewOrderRequestTemplate
	var orderId string

	err = db.QueryRow(
		query,
		orderRequest.ClientId,
		orderRequest.RestaurantId, price, "PENDING", time.Now()).Scan(&orderId)

	if err != nil {
		return err
	}

	// Utilisation d'un WaitGroup pour attendre la fin de toutes les goroutines
	var wg sync.WaitGroup
	errChan := make(chan error, len(orderRequest.OrderItems))

	// Boucle pour chaque item dans orderRequest.OrderItems
	for _, item := range orderRequest.OrderItems {
		wg.Add(1)
		go func(item models.OrderItem) {
			defer wg.Done()

			query := requests.InsertNewOrderItemRequestTemplate
			_, err := db.Exec(query, orderId, item.MenuId, item.Count)
			if err != nil {
				errChan <- err
				return
			}
		}(item)
	}

	// Attendez que toutes les goroutines se terminent
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Récupérez les erreurs des goroutines
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	db.Close()
	return nil
}

func GetPrice(db *sql.DB, liste []models.OrderItem) float64 {
	var order_price float64 = 0
	var price float64

	for _, item := range liste {
		query := requests.SelectMenuByIdTemplate

		err := db.QueryRow(query, item.MenuId).Scan(&price)

		if err != nil {
			fmt.Print(err)
		}
		order_price = price + order_price
	}
	return order_price
}

func UpdateStatusOrder(id_order string) bool {
	db, err := database.ConnectDB()
	if err != nil {
		return false
	}

	var order models.OrderDetailsRequest

	query := requests.UpdateStatusOrderRequestTemplate
	row, _ := db.Exec(query, id_order)
	fmt.Print(row)

	err = db.QueryRow("SELECT client_id, restaurant_id FROM order_details WHERE id = $1", id_order).Scan(&order.ClientId, &order.RestaurantId)

	if err != nil {
		return false
	}

	return true
}
