package main

import (
	"fmt"
)

// Attributes when a stone is moved
type Move struct {
	// Mode
	mode generatorMode

	// Stone which is moved, or removed
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

// ScoreNotSet contains special value to indicate it was not calculated before
const ScoreNotSet = 0xFFFF

func (move *Move) reset(stones stoneT) {
	move.stoneIndex = 0
	move.fromField = 0
	move.toField = 0
	move.valid = false
	move.score = ScoreNotSet
}

func (move Move) String() string {

	var result string

	switch move.mode {
	case placeStone:
		result = fmt.Sprintf("place:      +  %2d ", move.toField)
	case removeStone:
		result = fmt.Sprintf("remove:     -  %2d ", move.toField)
	case moveStone:
		result = fmt.Sprintf("move:    %2d -> %2d ", move.fromField, move.toField)
	}
	return result
}
