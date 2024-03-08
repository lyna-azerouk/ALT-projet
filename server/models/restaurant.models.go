package models

import "database/sql"

type Tags struct {
	Amenity string `json:"amenity"`
	Name    string `json:"name"`
	Tourism string `json:"tourism"`
}

type OverPassRestaurant struct {
	Type string  `json:"type"`
	ID   uint64     `json:"id"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Tags Tags    `json:"tags"`
}

type OverPassResponse struct {
	Version   float64              `json:"version"`
	Generator string               `json:"generator"`
	OSM3S     map[string]string    `json:"osm3s"`
	Elements  []OverPassRestaurant `json:"elements"`
}

type BouffluenceRestaurant struct {
	RestaurantDetails    OverPassRestaurant `json:"restaurantDetails"`
	Menu                 []Menu             `json:"menus"`
	OrderAverageDuration int                `json:"order_average_duration"`
}

type Menu struct {
	Id           uint64            `json:"id"`
	Name         string         `json:"name"`
	Price        int            `json:"price"`
	RestaurantID int            `json:"restaurent_id"`
	Description  sql.NullString `json:"description"`
	Image        sql.NullString `json:"url"`
}
