package models

// order request
type OrderDetailsRequest struct {
	Id           uint64      `json:"id"`
	ClientId     uint64      `json:"clientId"`
	RestaurantId uint64      `json:"restaurantId"`
	OrderItems   []OrderItem `json:"items"`
	Status       string      `json:"status"`
	Price        float64     `json:"price"`
	Date         string      `json:"date"`
}

type OrderDetailsRequestV2 struct {
	Id           string      `json:"id"`
	ClientId     string      `json:"clientId"`
	RestaurantId string      `json:"restaurantId"`
	OrderItems   []OrderItem `json:"items"`
	Status       string      `json:"status"`
	Price        string      `json:"price"`
	Date         string      `json:"date"`
}

type OrderDetails struct {
	ClientId     uint64      `json:"clientId"`
	ClientEmail  int         `json:"clientEmail"`
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
