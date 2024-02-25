package models

type Tags struct {
	Amenity string `json:"amenity"`
	Name    string `json:"name"`
	Tourism string `json:"tourism"`
}

type OverPassRestaurant struct {
	Type string  `json:"type"`
	ID   int     `json:"id"`
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
