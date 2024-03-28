// services/affluence

package services

import (
	"fmt"
	"serveur/server/const/affluences"
	"serveur/server/const/requests"
	"serveur/server/database"
	"serveur/server/models"
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

	return aggregation(restaurantVote, votes), nil
}

/*Aggregation des votes des clients et du restaurant
 * Nous avons attribuer des ponderations aux votes des clients et du restaurant
 * 0.7 pour le restaurant et 0.3 pour les clients.
 * Nous prenons le niveau qui a la plus grande valeur pondérée.
 *- @param restaurantVote : vote du restaurant
 *- @param votes : votes des clients
 */
func aggregation(voteRestaurant string, votes []models.AffluenceVote) string {

	var low_vote_client = 0
	var moderate_vote_restaurant = 0
	var high_vote_restaurant = 0
	for _, vote := range votes {
		if vote.AffluenceLevel == affluences.LOW_AFFLUENCE {
			low_vote_client = vote.Vote
		} else if vote.AffluenceLevel == affluences.MEDIUM_AFFLUENCE {
			moderate_vote_restaurant = vote.Vote
		} else if vote.AffluenceLevel == affluences.HIGH_AFFLUENCE {
			high_vote_restaurant = vote.Vote
		}
	}

	var w_c = affluences.PONDERATION_VOTE_CLIENT
	var w_r = affluences.PONDERATION_VOTE_RESTAURANT
	var low_ponderee = w_c * float64(low_vote_client)
	var moderate_ponderee = w_r * float64(moderate_vote_restaurant)
	var high_ponderee = w_r * float64(high_vote_restaurant)
	if voteRestaurant == affluences.LOW_AFFLUENCE {
		low_ponderee += w_r
	} else if voteRestaurant == affluences.MEDIUM_AFFLUENCE {
		moderate_ponderee += w_r
	} else if voteRestaurant == affluences.HIGH_AFFLUENCE {
		high_ponderee += w_r
	}

	if low_ponderee > moderate_ponderee && low_ponderee > high_ponderee {
		return affluences.LOW_AFFLUENCE
	} else if moderate_ponderee > low_ponderee && moderate_ponderee > high_ponderee {
		return affluences.MEDIUM_AFFLUENCE
	} else {
		return affluences.HIGH_AFFLUENCE
	}
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
