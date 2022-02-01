package main

import (
	"fmt"
	"testing"
)

func TestMoveGeneratorStoneA(t *testing.T) {

	mg := new(MoveGenerator)
	mg.reset(moveStone, Fields{}, Fields{})
	move := mg.nextMove()
	if move.valid {
		t.Errorf("Wrong move! %v %v", move.fromField, move.toField)
	}

	// Place a stone on first field
	mg.reset(moveStone, Fields{0}, Fields{})
	move = mg.nextMove()
	success := move.fromField == 0 && move.toField == 1 && move.valid

	move = mg.nextMove()
	success = success && move.fromField == 0 && move.toField == 9 && move.valid

	move = mg.nextMove()
	success = success && move.fromField == 0 && move.toField == 0 && !move.valid
	if !success {
		t.Errorf("Wrong move!")
	}

	// List of possible moves for field 7 and 19
	results := [][2]int{
		{7, 4},
		{7, 6},
		{7, 8},
		{19, 16},
		{19, 18},
		{19, 20},
		{19, 22},
	}

	// Place two stones
	mg.reset(moveStone, Fields{7, 19}, Fields{})
	for _, v := range results {

		move = mg.nextMove()
		if move.fromField != v[0] || move.toField != v[1] {
			t.Errorf("from: %v to: %v found: %v", move.fromField, move.toField, move.valid)
		}
		if !move.valid {
			break
		}
	}

}

func TestMoveGeneratorStoneB(t *testing.T) {
	mg := new(MoveGenerator)

	// With one stone
	mg.reset(moveStone, Fields{0}, Fields{1})
	move := mg.nextMove()
	success := move.fromField == 0 && move.toField == 9 && move.valid
	move = mg.nextMove()
	success = success && !move.valid

	mg.reset(moveStone, Fields{0}, Fields{1, 9})
	move = mg.nextMove()
	success = success && !move.valid

	if !success {
		t.Errorf("Wrong move! %v", move)
	}
}

func TestPlaceStone(t *testing.T) {

	mg := new(MoveGenerator)

	// No stone placed
	success := true
	mg.reset(placeStone, Fields{}, Fields{})
	for i := 0; i < 9; i++ {
		move := mg.nextMove()
		success = success && move.toField == i && move.valid
		// fmt.Printf("place: %v\n", move)
	}

	if !success {
		t.Errorf("Wrong move!")
	}
}
func TestRemoveStone(t *testing.T) {

	mg := new(MoveGenerator)

	// Test with one stone on the field
	mg.reset(removeStone, Fields{0}, Fields{})
	move := mg.nextMove()
	success := move.stoneIndex == 0 && move.valid
	move = mg.nextMove()
	success = success && !move.valid

	// Test with two stones on the field
	mg.reset(removeStone, Fields{3, 5}, Fields{})
	move = mg.nextMove()
	success = success && move.stoneIndex == 0 && move.valid
	move = mg.nextMove()
	success = success && move.stoneIndex == 1 && move.valid
	move = mg.nextMove()
	success = success && !move.valid

	if !success {
		t.Errorf("Wrong stone removal!")
	}
}

func TestApplyMove(t *testing.T) {

	src := Fields{1, 4, 7}
	move := Move{mode: placeStone, toField: 10}
	dst := src.applyMove(move)
	success := dst.same(&Fields{1, 4, 7, 10})

	move = Move{mode: moveStone, stoneIndex: 0, toField: 10}
	dst = src.applyMove(move)
	success = success && dst.same(&Fields{10, 4, 7})

	move = Move{mode: removeStone, stoneIndex: 0}
	dst = src.applyMove(move)
	fmt.Println(dst)
	success = success && dst.same(&Fields{7, 4})

	move = Move{mode: removeStone, stoneIndex: 1}
	dst = src.applyMove(move)
	success = success && dst.same(&Fields{1, 7})

	move = Move{mode: removeStone, stoneIndex: 2}
	dst = src.applyMove(move)
	success = success && dst.same(&Fields{1, 4})

	if !success {
		t.Errorf("Error applying change!")
	}
}
