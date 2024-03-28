// services/affluence

package services

import (
	"fmt"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
	"serveur/utiles"
)

/*
fUNCTION THAT gets the affluence of the restaurant (low, hiengh, medium)
*/
func GetAffluence(restaurantId uint64) (string, error) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	query1 := requests.SelectAffluenceSubmitedByRestaurantRequestTemplate
	query2 := requests.SelectAffluenceSubmitedByClientsRequestTemplate
	row1 := db.QueryRow(query1, restaurantId)
	voteRows, err := db.Query(query2, restaurantId)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	var restaurantVote string
	var votes []models.AffluenceVote

	err = row1.Scan(&restaurantVote)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	for voteRows.Next() {
		var vote models.AffluenceVote
		err := voteRows.Scan(&vote.AffluenceLevel, &vote.Vote)
		if err != nil {
			fmt.Println("Error:", err)
			return "", err
		}
		votes = append(votes, vote)
	}

	return utiles.VoteAggregation(restaurantVote, votes), nil
}

func UpdateAffluenceForRestaurantVote(restaurantId int, vote string) (string, error) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	query := requests.UpdateAffluenceForRestaurantVoteRequestTemplate
	row := db.QueryRow(query, vote, restaurantId)
	var affluence string

	err = row.Scan(&affluence)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return affluence, nil

}

func SubmitClientVoteForAffluence(restaurantId uint64, vote string) (int, error) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	query := requests.SubmitVoteForRestaurantAffluenceRequestTemplate
	row := db.QueryRow(query, restaurantId, vote)
	var affluence int

	err = row.Scan(&affluence)
	if err != nil {
		fmt.Println("Error:", err)
		return (0), err
	}

	return affluence, nil
}
