package main

import (
	"fmt"
	"testing"
)

// Check that moves array is consistent
func TestMoves(t *testing.T) {

	if len(allMoves) != 24 {
		t.Errorf("Expected 24, but: %v", len(allMoves))
	}

	for from, fields := range allMoves {

		for _, field := range fields {

			if !allMoves[field].contains(from) {
				t.Errorf("Error field is not contained!")
			}
		}

	}

	status := true
	to, found := allMoves.getMove(0, 0)
	status = status && (found && to == 1)

	to, found = allMoves.getMove(0, 1)
	status = status && (found && to == 9)

	to, found = allMoves.getMove(23, 1)
	status = status && (found && to == 22)

	_, found = allMoves.getMove(0, 2)
	status = status && !found

	if !status {
		t.Errorf("Error checking moves!")
	}

	// Statistic: count number of moves
	var count [4]int
	for i := 0; i < 24; i++ {
		n := len(allMoves[i]) - 1
		count[n] += 1
	}
	fmt.Println("move counts: &v", count)
	if count != [4]int{0, 12, 8, 4} {
		t.Errorf("Number of moves does not match!")
	}
}
