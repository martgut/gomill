package main

import (
	"fmt"
	"testing"
)

func TestMoveOptimizerSingle(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rater: EvalHighestField{}}

	// Evaluate with one stone; one level
	move := mo.calcBestMoveSingle(Fields{0}, Fields{}, 1)
	success := move.toField == 9
	move = mo.calcBestMoveSingle(Fields{4}, Fields{}, 1)
	success = success && move.toField == 7
	move = mo.calcBestMoveSingle(Fields{22}, Fields{}, 1)
	success = success && move.toField == 23
	move = mo.calcBestMoveSingle(Fields{23}, Fields{}, 1)
	success = success && move.toField == 22
	move = mo.calcBestMoveSingle(Fields{17}, Fields{}, 1)
	success = success && move.toField == 16
	if !success {
		t.Errorf("Wrong move (evauluate one stone to highest field)!")
	}

	// Evaluate with two stones; one level
	move = mo.calcBestMoveSingle(Fields{0, 2}, Fields{}, 1)
	success = move.toField == 14 && mo.moveCounter == 4
	move = mo.calcBestMoveSingle(Fields{20, 9}, Fields{}, 1)
	success = success && move.toField == 21
	success = success && mo.moveCounter == 5
	if !success {
		t.Errorf("Wrong move (evaluate two stones to highest field")
	}

	// Evaluate with one stone; more levels
	move = mo.calcBestMoveSingle(Fields{0}, Fields{}, 2)
	success = move.score == 21 && mo.moveCounter == 8
	move = mo.calcBestMoveSingle(Fields{0}, Fields{}, 3)
	success = success && move.score == 22
	move = mo.calcBestMoveSingle(Fields{0}, Fields{}, 4)
	success = success && move.score == 23
	if !success {
		t.Errorf("Wrong move (evaluate one stones multiple levels")
	}

	// Evaluate with one stone; from every field; until stone is on 23
	for i := 0; i < 24; i++ {
		for j := 1; j < 20; j++ {
			move = mo.calcBestMoveSingle(Fields{i}, Fields{}, j)
			if move.score == 23 {
				// fmt.Printf("found: from: %v level: %v\n", i, j)
				break
			}
		}
	}
}

func TestMoveOptimizerB(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rater: EvalHighestField{}}

	// Evaluate with one stone; with player B; one level
	move := mo.calcBestMoveSingle(Fields{0}, Fields{1}, 1)
	success := move.toField == 9
	move = mo.calcBestMoveSingle(Fields{0}, Fields{9}, 1)
	success = success && move.toField == 1
	if !success {
		t.Errorf("Wrong move (evauluate one stone to highest field)!")
	}

	// Evaluate with one stone; with player B; more levels
	move = mo.calcBestMoveSingle(Fields{0}, Fields{9}, 2)
	success = move.toField == 1 && move.score == 4
	move = mo.calcBestMoveSingle(Fields{17}, Fields{22}, 3)
	success = success && move.toField == 12 && move.score == 20
	move = mo.calcBestMoveSingle(Fields{0}, Fields{22}, 4)
	success = success && move.toField == 1 && move.score == 23
	if !success {
		t.Errorf("Wrong move (evauluate one stone to highest field)! %v", move)
	}
}

func TestMoveOptimizerMulti(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rateField: HighestFieldsRater{}}

	// Evaluate with one stone; one level
	move := mo.calcBestMoveDouble(Fields{0}, Fields{1}, 0, 1)
	success := move.score == 8 && move.toField == 9

	// Evaluate with one stone; one level
	move = mo.calcBestMoveDouble(Fields{0}, Fields{1}, 0, 2)
	success = success && move.toField == 9 && move.score == 5
	move = mo.calcBestMoveDouble(Fields{0}, Fields{2}, 0, 2)
	success = success && move.toField == 9 && move.score == -5

	if !success {
		t.Errorf("Wrong move (evauluate one stone to highest field)!")
	}

	// TODO
	move = mo.calcBestMoveDouble(Fields{0}, Fields{1}, 0, 3)
}

func TestPlaceStones(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rateField: HighestFieldsRater{}}

	move := mo.calcBestMoveDouble(Fields{}, Fields{}, 1, 1)
	fmt.Printf("move: %v\n", move)
	success := move.score == 23 && move.toField == 23

	move = mo.calcBestMoveDouble(Fields{}, Fields{}, 2, 2)
	fmt.Printf("move: %v\n", move)
	success = success && move.score == 1 && move.toField == 23

	move = mo.calcBestMoveDouble(Fields{}, Fields{}, 3, 3)
	fmt.Printf("move: %v\n", move)
	success = success && move.score == 22 && move.toField == 23

	if !success {
		t.Errorf("Wrong move (evauluate one stone to highest field)!")
	}
}
