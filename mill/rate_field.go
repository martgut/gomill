package main

// Evaluate the value of a move
type RateField interface {
	rate(stones Fields) int
}

type HighestFieldsRater struct {
}

func (HighestFieldsRater) rate(stones Fields) int {

	// Just sum up the field values of each stone
	score := 0
	for _, stone := range stones {
		score += stone
	}
	return score
}

type CountStonesRater struct {
}

func (CountStonesRater) rate(stones Fields) int {
	// Count the number of stones on the field
	return len(stones) * 10
}
