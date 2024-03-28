package utiles

import (
	"math"
	"serveur/server/const/affluences"
	"serveur/server/models"
)

func VoteAggregation(voteRestaurant string, votes []models.AffluenceVote) string {
	low_vote_client, moderate_vote_restaurant, high_vote_restaurant := getVotes(votes)

	w_c := affluences.PONDERATION_VOTE_CLIENT
	w_r := affluences.PONDERATION_VOTE_RESTAURANT

	low_ponderee := calculateWeightedVote(w_c, w_r, low_vote_client)
	moderate_ponderee := calculateWeightedVote(w_c, w_r, moderate_vote_restaurant)
	high_ponderee := calculateWeightedVote(w_c, w_r, high_vote_restaurant)

	adjustWeightedVote(&low_ponderee, &moderate_ponderee, &high_ponderee, voteRestaurant, w_r)

	moyenne_ponderee := calculateMeanWeightedVote(low_ponderee, moderate_ponderee, high_ponderee)

	return determineCurrentAffluence(moyenne_ponderee, low_ponderee, moderate_ponderee, high_ponderee)
}

func getVotes(votes []models.AffluenceVote) (int, int, int) {
	var low_vote_client, moderate_vote_restaurant, high_vote_restaurant int
	for _, vote := range votes {
		switch vote.AffluenceLevel {
		case affluences.LOW_AFFLUENCE:
			low_vote_client = vote.Vote
		case affluences.MEDIUM_AFFLUENCE:
			moderate_vote_restaurant = vote.Vote
		case affluences.HIGH_AFFLUENCE:
			high_vote_restaurant = vote.Vote
		}
	}
	return low_vote_client, moderate_vote_restaurant, high_vote_restaurant
}

func calculateWeightedVote(w_c, w_r float64, vote int) float64 {
	return w_c * float64(vote)
}

func adjustWeightedVote(low, moderate, high *float64, voteRestaurant string, w_r float64) {
	switch voteRestaurant {
	case affluences.LOW_AFFLUENCE:
		*low += w_r
	case affluences.MEDIUM_AFFLUENCE:
		*moderate += w_r
	case affluences.HIGH_AFFLUENCE:
		*high += w_r
	}
}

func calculateMeanWeightedVote(low, moderate, high float64) float64 {
	return (low + moderate + high) / 3.0
}

func determineCurrentAffluence(mean, low, moderate, high float64) string {
	min_diff := math.Abs(mean - low)
	diffs := []float64{math.Abs(mean - low), math.Abs(mean - moderate), math.Abs(mean - high)}
	for _, diff := range diffs {
		if diff < min_diff {
			min_diff = diff
		}
	}
	switch min_diff {
	case math.Abs(mean - low):
		return affluences.LOW_AFFLUENCE
	case math.Abs(mean - moderate):
		return affluences.MEDIUM_AFFLUENCE
	default:
		return affluences.HIGH_AFFLUENCE
	}
}
