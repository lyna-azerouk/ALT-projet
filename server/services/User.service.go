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
