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

type EvenFieldRater struct {
	value int
}

func (efr EvenFieldRater) rate(stones Fields) int {

	// Just sum up the field values of each stone
	score := 0
	for _, stone := range stones {
		if stone%2 == efr.value {
			score += 1
		}
	}
	return score
}

type CountStonesRater struct {
	// Value of one stone
	valueStone int
}

func (csr CountStonesRater) rate(stones Fields) int {
	// Count the number of stones on the field
	return len(stones) * csr.valueStone
}

func NewCountStonesRaster() *CountStonesRater {
	return &CountStonesRater{valueStone: 10}
}

type MultiplexRater struct {
	raters []RateField
}

func (mr MultiplexRater) rate(stones Fields) int {
	score := 0
	for _, rater := range mr.raters {
		score += rater.rate(stones)
	}
	return score
}

// Put stones on crossing with more move options
type BestCrossingRaster struct {
}

func (BestCrossingRaster) rate(stones Fields) int {

	// Just sum up the field values of each stone
	score := 0
	for _, stone := range stones {

		// Each position has at minimum 2 move options
		// -> account for crossing with 3, or 4 options
		score += len(allMoves[stone]) - 2

	}
	return score
}
