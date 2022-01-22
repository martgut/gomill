package main

import (
	"testing"
)

func TestHighestFieldRaster(t *testing.T) {

	rater := HighestFieldsRater{}
	status := rater.rate(Fields{}) == 0
	status = status && rater.rate(Fields{1}) == 1
	status = status && rater.rate(Fields{1, 23}) == 24

	if !status {
		t.Errorf("Error rating field!")
	}
}

func TestCountStonesRaster(t *testing.T) {

	rater := CountStonesRater{}
	status := rater.rate(Fields{}) == 0
	status = status && rater.rate(Fields{1}) == 1
	status = status && rater.rate(Fields{1, 23}) == 2

	if !status {
		t.Errorf("Error rating field!")
	}
}
