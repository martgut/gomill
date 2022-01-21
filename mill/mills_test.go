package main

import (
	"testing"
)

func TestIsMill(t *testing.T) {

	status := true
	status = status && allMills.isMill(0, Fields{1, 2}) == 1
	status = status && allMills.isMill(0, Fields{2, 1}) == 1
	status = status && allMills.isMill(0, Fields{9, 21}) == 1
	status = status && allMills.isMill(0, Fields{9, 20}) == 0
	if !status {
		t.Error("Error checking mills!")
	}

	// Place no stone of any field
	stones := make([]int, 24)

	// Iterating over all fields must lead to zero mills
	for field := 0; field < 24; field++ {

		result := allMills.isMill(field, stones)
		if result != 0 {
			t.Errorf("Error with mills calculation! field=%v", field)
		}
	}

	// Place a stone on each field
	for i := 0; i < 24; i++ {
		stones[i] = i
	}
	for field := 0; field < 24; field++ {

		result := allMills.isMill(field, stones)
		if result != 2 {
			t.Errorf("Error with mills calculation! field=%v", field)
		}
	}
}
