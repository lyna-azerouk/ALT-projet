package models

type OrderDetails struct {
	Id           int         `json:"id"`
	ClientId     int         `json:"clientId"`
	ClientEmail  string      `json:"clientEmail"`
	RestaurantId int         `json:"restaurantId"`
	OrderItems   []OrderItem `json:"items"`
	Status       string      `json:"status"`
	Price        float64     `json:"price"`
	Date         string      `json:"date"`
}


type OrderItem struct {
	MenuId uint64 `json:"menuId"`
	Count  int    `json:"count"`
}


// order response: en fonction des besoins du client ????
