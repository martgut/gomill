package main

import (
	"testing"
)

func validateMoveTo(t *testing.T, mo *MoveOptimizer, to []int, escore int) {

	moves := mo.perfectMove[0]

	score := moves[0].score
	if score != escore {
		t.Errorf("Wrong score! score=%v escore=%v  moves=%v emove=%v", score, escore, moves, to)
		return
	}

	if len(moves) != len(to) {
		t.Errorf("Wrong move length! moves=%v emove=%v", moves, to)
		return
	}
	for i := 0; i < len(to); i++ {
		if moves[i].toField != to[i] {
			t.Errorf("Wrong move! moves=%v emove=%v", moves, to)
			return
		}
	}
}

func TestMoveOptimizerSimple(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rater: HighestFieldsRater{}}

	// Evaluate with one stone; one level
	mo.calcBestMove(Fields{0}, Fields{}, 0, 1)
	validateMoveTo(t, &mo, []int{9}, 9)
	mo.calcBestMove(Fields{4}, Fields{}, 0, 1)
	validateMoveTo(t, &mo, []int{7}, 7)
	mo.calcBestMove(Fields{22}, Fields{}, 0, 1)
	validateMoveTo(t, &mo, []int{23}, 23)
	mo.calcBestMove(Fields{23}, Fields{}, 0, 1)
	validateMoveTo(t, &mo, []int{22}, 22)
	mo.calcBestMove(Fields{17}, Fields{}, 0, 1)
	validateMoveTo(t, &mo, []int{16}, 16)

	// Evaluate with two stones; one level
	mo.calcBestMove(Fields{0, 2}, Fields{}, 0, 1)
	validateMoveTo(t, &mo, []int{14}, 14)
	mo.calcBestMove(Fields{20, 9}, Fields{}, 0, 1)
	validateMoveTo(t, &mo, []int{21}, 41)

	// Evaluate with two stones; more levels
	mo.calcBestMove(Fields{0}, Fields{2}, 0, 2)
	validateMoveTo(t, &mo, []int{9, 14}, -5)
	mo.calcBestMove(Fields{0}, Fields{2}, 0, 3)
	validateMoveTo(t, &mo, []int{9, 14, 21}, 7)
	mo.calcBestMove(Fields{0}, Fields{2}, 0, 4)
	validateMoveTo(t, &mo, []int{9, 14, 21, 23}, -2)

	// Evaluate with one stone; from every field; until stone is on 23
	for i := 0; i < 24; i++ {
		for j := 1; j < 20; j++ {
			move := mo.calcBestMove(Fields{i}, Fields{}, 0, j)
			if move.score == 23 {
				// fmt.Printf("found: from: %v level: %v\n", i, j)
				break
			}
		}
	}
}

func TestMoveOptimizerAdvanced(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rater: HighestFieldsRater{}}

	// Evaluate with one stone; with player B; one level
	mo.calcBestMove(Fields{0}, Fields{1}, 0, 1)
	validateMoveTo(t, &mo, []int{9}, 8)
	mo.calcBestMove(Fields{0}, Fields{9}, 0, 1)
	validateMoveTo(t, &mo, []int{1}, -8)

	// Evaluate with one stone; with player B; more levels
	mo.calcBestMove(Fields{0}, Fields{9}, 0, 2)
	validateMoveTo(t, &mo, []int{1, 21}, -20)
	mo.calcBestMove(Fields{17}, Fields{22}, 0, 3)
	validateMoveTo(t, &mo, []int{16, 23, 19}, -4)
	mo.calcBestMove(Fields{0}, Fields{22}, 0, 4)
	validateMoveTo(t, &mo, []int{9, 21, 10, 22}, -12)
}

func TestMoveOptimizerMulti(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rater: HighestFieldsRater{}}

	// Evaluate with one stone; one level
	mo.calcBestMove(Fields{0}, Fields{1}, 0, 1)
	validateMoveTo(t, &mo, []int{9}, 8)

	// Evaluate with one stone; one level
	mo.calcBestMove(Fields{0}, Fields{1}, 0, 2)
	validateMoveTo(t, &mo, []int{9, 4}, 5)
	mo.calcBestMove(Fields{0}, Fields{2}, 0, 2)
	validateMoveTo(t, &mo, []int{9, 14}, -5)

	// Evaluate three levels
	mo.calcBestMove(Fields{0}, Fields{1}, 0, 3)
	validateMoveTo(t, &mo, []int{9, 4, 21}, 17)
}

func TestPlaceStones(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rater: HighestFieldsRater{}}

	mo.calcBestMove(Fields{}, Fields{}, 1, 1)
	validateMoveTo(t, &mo, []int{23}, 23)

	mo.calcBestMove(Fields{}, Fields{}, 2, 2)
	validateMoveTo(t, &mo, []int{23, 22}, 1)

	mo.calcBestMove(Fields{}, Fields{}, 3, 3)
	validateMoveTo(t, &mo, []int{23, 22, 21}, 22)
}

func TestClosingMillSimple(t *testing.T) {

	// Search for highest field move
	mo := MoveOptimizer{rater: CountStonesRater{}}

	// Make a mill and remove stone
	mo.calcBestMove(Fields{0, 1}, Fields{23}, 1, 2)
	validateMoveTo(t, &mo, []int{2, 23}, 30)
}

func TestClosingMillSmart(t *testing.T) {

	// Keep stone on even
	rater := MultiplexRater{raters: []RateField{CountStonesRater{}, EvenFieldRater{value: 0}}}
	mo := MoveOptimizer{rater: &rater}

	// Make a mill and remove stone: The one on the even field must remain!
	// So Player A should remove stone on even place for better score
	mo.calcBestMove(Fields{0, 1}, Fields{22, 23}, 1, 2)
	// TODO validateMoveTo(t, &mo, []int{2, 22}, 99)

	/* 	// Keep stone on odd places
	   	rater = MultiplexRater{raters: []RateField{CountStonesRater{}, EvenFieldRater{value: 1}}}
	   	mo = MoveOptimizer{rater: &rater}

	   	// Make a mill and remove stone: The one on the odd field must remain!
	   	mo.calcBestMove(Fields{0, 1}, Fields{22, 23}, 1, 2)
	   	validateMoveTo(t, &mo, []int{2, 22}, 20)
	*/
}
