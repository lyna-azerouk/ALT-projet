package models

type OrderDetails struct {
	Id           uint64      `json:"id"`
	ClientId     uint64      `json:"clientId"`
	ClientEmail  string      `json:"clientEmail"`
	RestaurantId uint64      `json:"restaurantId"`
	OrderItems   []OrderItem `json:"items"`
	Status       string      `json:"status"`
	Price        float64     `json:"price"`
	Date         string      `json:"date"`
}

type OrderDetailsRequest struct {
	Id           string             `json:"id"`
	ClientId     string             `json:"clientId"`
	RestaurantId string             `json:"restaurantId"`
	OrderItems   []OrderItemRequest `json:"items"`
	Status       string             `json:"status"`
	Price        string             `json:"price"`
	Date         string             `json:"date"`
}

type OrderItem struct {
	MenuId uint64 `json:"menuId"`
	Count  int    `json:"count"`
}

type OrderItemRequest struct {
	MenuId string `json:"menuId"`
	Count  string `json:"count"`
}

// order response: en fonction des besoins du client ????
