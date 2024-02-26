package models

// order request
type OrderDetailsRequest struct {
	ClientId     int         `json:"clientId"`
	RestaurantId int         `json:"restaurantId"`
	OrderItems   []OrderItem `json:"items"`
}

type OrderItem struct {
	MenuId int `json:"menuId"`
	Count  int `json:"count"`
}

// order response: en fonction des besoins du client ????
