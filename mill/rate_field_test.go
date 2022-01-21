package main

import (
	"testing"
)

func TestRateField(t *testing.T) {

	rater := HighestFieldsRater{}
	success := rater.rate(Fields{}) == 0
	success = success && rater.rate(Fields{1}) == 1
	success = success && rater.rate(Fields{1, 23}) == 24

	if !success {
		t.Errorf("Error rating field!")
	}
}
