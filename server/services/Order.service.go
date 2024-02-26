package services

import (
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
	// 1st step: insert order in the database
	query := requests.InsertNewOrderRequestTemplate
	var orderId string

	err = db.QueryRow(
		query,
		orderRequest.ClientId,
		orderRequest.RestaurantId, 0, "PENDING", time.Now()).Scan(&orderId)

	if err != nil {
		return err
	}

	fmt.Println("id: " + (orderId))

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
