package services

import (
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
)

func GetUserDetails(userID string) (models.UserDetails, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return models.UserDetails{}, err
	}
	query := requests.GetUserDetailsRequestTemplate
	var user models.UserDetails
	err = db.QueryRow(query, userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return models.UserDetails{}, err
	}
	return user, nil
}


/*
Function that return all the orders of a user
*/

func GetUserOrdersDetails(userId string) ([]models.OrderDetailsRequest, error) {


	db, _ := database.ConnectDB()
	query := requests.GetUserOrdersTemplate
	rows, err := db.Query(query,  userId)

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
