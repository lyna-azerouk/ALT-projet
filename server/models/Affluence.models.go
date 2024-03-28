package models

type AffluenceVote struct {
	RestaurantId   int    `json:"restaurantId"`
	AffluenceLevel string `json:"affluenceLevel"`
	Vote           int    `json:"vote"`
}
