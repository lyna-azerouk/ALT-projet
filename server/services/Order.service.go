package services

import (
	"database/sql"
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
func CreateNewOrder(orderRequest models.OrderDetailsRequest) (models.OrderDetailsRequest, error) {
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
		return models.OrderDetailsRequest{}, err
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
			return models.OrderDetailsRequest{}, err
		}
	}

	db.Close()
	order := GetOrderDetails(orderId)

	return order, nil
}

/*
*Get all the details of an order
 */
func GetOrderDetails(id_order string) models.OrderDetailsRequest {
	db, _ := database.ConnectDB()
	var order models.OrderDetailsRequest
	var id string

	query_order := requests.GetOrderRequestTemplate
	db.QueryRow(query_order, id_order).Scan(&id, &order.RestaurantId, &order.ClientId, &order.Price, &order.Date, &order.Status)

	query_items := requests.GetOrderItemsRequestTemplate
	rows_order_items, _ := db.Query(query_items, id_order)

	for rows_order_items.Next() {
		var order_item models.OrderItem

		rows_order_items.Scan(&order_item.MenuId, order_item.Count)

		order.OrderItems = append(order.OrderItems, order_item)
	}
	return order
}

/*
*
Update the status of the order from Pending to In-Progress
*/
func UpdateStatusOrder(id_order string, status string) models.OrderDetailsRequest {
	db, _ := database.ConnectDB()

	var order models.OrderDetailsRequest

	query := requests.UpdateStatusOrderRequestTemplate
	_, err := db.Exec(query, status, id_order)

	if err != nil {
		fmt.Println(err)
	}

	if status == "DECLINED" {
		delete_query1 := requests.DeleteOrderItemsRequestTemplate
		delete_query2 := requests.DeleteOrderDetailsRequestTemplate
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
func VerfyOrderCode(id_order string, code string) (models.OrderDetailsRequest, error) {

	db, _ := database.ConnectDB()

	query := requests.GetOrderCodeTemplate
	num_order, _ := strconv.Atoi(id_order)
	num_code, _ := strconv.Atoi(code)
	row, err := db.Exec(query, num_order, num_code)
	fmt.Println(row)

	if err != nil {
		fmt.Println(err)
		log.Fatal("error with the database")
		return models.OrderDetailsRequest{}, err
	}

	order := UpdateStatusOrder(id_order, "COMPLETED")

	return order, nil
}

/*
Function that return all the orders of a user
*/

func GetUserOrdersDetails(userId string) ([]models.OrderDetailsRequest, error) {
	var userIdNumber int
	userIdNumber, _ = strconv.Atoi(userId)

	db, _ := database.ConnectDB()
	query := requests.GetUserOrdersTemplate
	rows, err := db.Query(query, userIdNumber)

	if err != nil {
		return []models.OrderDetailsRequest{}, err
	}

	var orders []models.OrderDetailsRequest

	for rows.Next() {
		var orderid string
		var order models.OrderDetailsRequest
		err := rows.Scan(&orderid)

		if err != nil {
			return []models.OrderDetailsRequest{}, err
		}
		order = GetOrderDetails(orderid)
		orders = append(orders, order)
	}

	return orders, nil
}

func GetRestaurantOrdersDetails(restaurantId string) ([]models.OrderDetailsRequest, error) {
	var restaurantIdNumber int
	restaurantIdNumber, _ = strconv.Atoi(restaurantId)

	db, _ := database.ConnectDB()
	query := requests.GetRestaurantOrdersTemplate
	rows, err := db.Query(query, restaurantIdNumber)

	if err != nil {
		return []models.OrderDetailsRequest{}, err
	}

	var orders []models.OrderDetailsRequest

	for rows.Next() {
		var orderid string
		var order models.OrderDetailsRequest
		err := rows.Scan(&orderid)

		if err != nil {
			return []models.OrderDetailsRequest{}, err
		}
		order = GetOrderDetails(orderid)
		orders = append(orders, order)
	}

	return orders, nil
}