package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
	"strconv"
	"sync"
	"time"
)

/*
* Function that creates a new order
 */
func CreateNewOrder(orderRequest models.OrderDetails) (models.OrderDetails, error) {
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
		return models.OrderDetails{}, err
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
			return models.OrderDetails{}, err
		}
	}

	db.Close()
	order := GetOrderDetails(orderId)

	return order, nil
}

/*
*Get all the details of an order
 */
func GetOrderDetails(id_order string) models.OrderDetails {
	db, _ := database.ConnectDB()
	var order models.OrderDetails

	query_order := requests.GetOrderRequestTemplate
	db.QueryRow(query_order, id_order).Scan(&order.Id, &order.RestaurantId, &order.ClientId, &order.Price, &order.Date, &order.Status)

	query_items := requests.GetOrderItemsRequestTemplate
	rows_order_items, _ := db.Query(query_items, id_order)

	for rows_order_items.Next() {
		var order_item models.OrderItem

		rows_order_items.Scan(&order_item.MenuId, &order_item.Count)

		order.OrderItems = append(order.OrderItems, order_item)
	}
	return order
}

/* Update the status of the order from Pending to In-Progress*/
func UpdateStatusOrder(id_order string, status string) models.OrderDetails {
	db, _ := database.ConnectDB()

	var order models.OrderDetails

	query := requests.UpdateStatusOrderRequestTemplate
	_, err := db.Exec(query, status, id_order)

	if err != nil {
		fmt.Println(err)
	}

	if status == "DECLINED" {
		delete_query1 := requests.DeleteOrderItemsRequestTemplate
		delete_query2 := "" // requests.DeleteOrderDetailsTemplate
		_, err = db.Exec(delete_query1, id_order)
		_, err = db.Exec(delete_query2, id_order)
		if err != nil {
			fmt.Println(err)
		}
	}
	order = GetOrderDetails(id_order)

	return order
}

/*
* Function that get the Price of an order:  sum(price of each menu)
 */
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

/*
Function that generate the code
*/

func GenerateCode(id_order string) (int, error) {
	rand.Seed(time.Now().UnixNano())

	code := rand.Intn(9000) + 1000

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query := requests.InsertCodeRequestTemplate
	num, err := strconv.Atoi(id_order)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	_, err = db.Exec(query, num, code)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return code, nil
}

/*
Restaurent: A function that takes a code as a parameter and verifies if the code is valid
*/
func VerfyOrderCode(id_order string, code string) (models.OrderDetails, error) {

	db, _ := database.ConnectDB()

	query := requests.GetOrderCodeTemplate
	num_order, err := strconv.Atoi(id_order)
	if err != nil {
		fmt.Println(err)
		return models.OrderDetails{}, err
	}

	num_code, err := strconv.Atoi(code)

	if err != nil {
		fmt.Println(err)
		return models.OrderDetails{}, err
	}

	row, err := db.Query(query, num_order, num_code)

	if err != nil {
		fmt.Println(err)
		log.Fatal("error with the database")
		return models.OrderDetails{}, err
	}

	if !row.Next() {
		errorMsg := "Invalid code provided or order not found"
		log.Println(errorMsg)
		err = errors.New(errorMsg)
		return models.OrderDetails{}, err
	}

	order := UpdateStatusOrder(id_order, "COMPLETED")

	return order, nil
}

/*
Function that return all the orders of a user
*/

func GetUserOrdersDetails(userId string) ([]models.OrderDetails, error) {
	var userIdNumber int
	userIdNumber, _ = strconv.Atoi(userId)

	db, _ := database.ConnectDB()
	query := requests.GetUserOrdersTemplate
	rows, err := db.Query(query, userIdNumber)

	if err != nil {
		return []models.OrderDetails{}, err
	}

	var orders []models.OrderDetails

	for rows.Next() {
		var orderid string
		var order models.OrderDetails
		err := rows.Scan(&orderid)

		if err != nil {
			return []models.OrderDetails{}, err
		}
		order = GetOrderDetails(orderid)
		orders = append(orders, order)
	}

	return orders, nil
}

func GetRestaurantOrdersDetails(restaurantId string) ([]models.OrderDetails, error) {
	var restaurantIdNumber int
	restaurantIdNumber, _ = strconv.Atoi(restaurantId)

	db, _ := database.ConnectDB()
	query := requests.GetRestaurantOrdersDetailsTemplate
	rows, err := db.Query(query, restaurantIdNumber)
	if err != nil {
		log.Printf("Error while getting restaurant orders: %s", err.Error())
		return []models.OrderDetails{}, err
	}

	var ordersMap = make(map[int]*models.OrderDetails)
	for rows.Next() {
		var orderId, clientId, restaurantId int
		var status, date string
		var price float64
		var menuId int
		var count int

		err := rows.Scan(&orderId, &clientId, &restaurantId, &status, &price, &date, &menuId, &count)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la ligne:", err)
			ordersMap[orderId] = &models.OrderDetails{
				Id:           orderId,
				ClientId:     clientId,
				RestaurantId: restaurantId,
				Status:       status,
				Price:        price,
				Date:         date,
			}
			continue
		}

		// On recupere la commande pour y ajouter l'article
		order, ok := ordersMap[orderId]
		if !ok {
			// Si la commande n'existe pas encore, la créer
			order = &models.OrderDetails{
				Id:           orderId,
				ClientId:     clientId,
				RestaurantId: restaurantId,
				Status:       status,
				Price:        price,
				Date:         date,
			}
			// Ajouter la commande à la carte
			ordersMap[orderId] = order
		}
		// Ajouter l'élément de commande à la commande actuelle
		order.OrderItems = append(order.OrderItems, models.OrderItem{
			MenuId: menuId,
			Count:  count,
		})
	}

	// convertir la carte en tableau
	var orders []models.OrderDetails
	for _, order := range ordersMap {
		orders = append(orders, *order)
	}
	return orders, nil
}
