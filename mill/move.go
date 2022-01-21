package main

import (
	"fmt"
	"math"
)

// Attributes when a stone is moved
type Move struct {
	// Mode
	mode generatorMode

	// Stone which is moved
	stoneIndex int

	// Stone is moved from field
	fromField int

	// Stone is move to field
	toField int

	// Is this a valid move?
	valid bool

	// Value of this move
	score int

	// Was a mill closed with this move?
	isMill bool
}

func (move *Move) reset() {
	move.stoneIndex = 0
	move.fromField = 0
	move.toField = 0
	move.valid = false
	move.score = math.MinInt32
}

func (move Move) String() string {
	return fmt.Sprintf("%2d -> %2d score: %2d", move.fromField, move.toField, move.score)
}
