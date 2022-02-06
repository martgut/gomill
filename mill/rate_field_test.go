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

	rater := CountStonesRater{10}
	status := rater.rate(Fields{}) == 0
	status = status && rater.rate(Fields{1}) == 10
	status = status && rater.rate(Fields{1, 23}) == 20

	if !status {
		t.Errorf("Error rating field!")
	}
}

func TestEvenRater(t *testing.T) {
	rater := EvenFieldRater{}
	status := rater.rate(Fields{}) == 0
	status = status && rater.rate(Fields{0}) == 1
	status = status && rater.rate(Fields{0, 1, 2}) == 2

	if !status {
		t.Errorf("Wrong rating!")
	}
}

func TestMultiplexRater(t *testing.T) {
	rater := MultiplexRater{raters: []RateField{EvenFieldRater{}, EvenFieldRater{}}}
	status := rater.rate(Fields{}) == 0
	status = status && rater.rate(Fields{0}) == 2
	status = status && rater.rate(Fields{0, 1, 2}) == 4

	if !status {
		t.Errorf("Wrong rating!")
	}
}

func TestBestCrossingRater(t *testing.T) {
	rater := BestCrossingRaster{}
	status := rater.rate(Fields{}) == 0
	status = status && rater.rate(Fields{0}) == 0
	status = status && rater.rate(Fields{1}) == 1
	status = status && rater.rate(Fields{4}) == 2

	if !status {
		t.Errorf("Wrong rating!")
	}
}
