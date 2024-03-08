package models

// order request
type OrderDetailsRequest struct {
	ClientId     string      `json:"clientId"`
	RestaurantId string      `json:"restaurantId"`
	OrderItems   []OrderItem `json:"items"`
	Status       string      `json:"status"`
	Price        float64     `json:"price"`
	Date         string      `json:"date"`
}

type OrderDetails struct {
	ClientId     string      `json:"clientId"`
	ClientEmail  int         `json:"clientId"`
	RestaurantId int         `json:"restaurantId"`
	OrderItems   []OrderItem `json:"items"`
	Status       string      `json:"status"`
	Price        float64     `json:"price"`
	Date         string      `json:"date"`
}

type OrderItem struct {
	MenuId string `json:"menuId"`
	Count  int    `json:"count"`
}

// order response: en fonction des besoins du client ????
