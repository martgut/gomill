package main

import (
	"fmt"
	"math"
)

type MoveOptimizer struct {

	// Rating instance
	rater Rater // TODO remove

	// Instance which rates the field and provides scores
	rateField RateField

	// Move generator for each level
	moveGenerator []MoveGenerator

	// Best move for each level when traversing the move tree
	// Note: when going down the tree, bestMove is always reset
	bestMove []Move

	// Planned best move: contains all moves which lead
	// to the calculated score.
	perfectMove [][]Move

	// Number of calculated moves
	moveCounter int
}

// Return perfect move at given level
func (mo *MoveOptimizer) pMove(level int) *Move {
	return &mo.perfectMove[0][level]
}

func (mo *MoveOptimizer) initBestMove(levelMax int) {
	// Maintain a MoveGenerator for each level
	mo.moveGenerator = make([]MoveGenerator, levelMax)
	for idx := 0; idx < levelMax; idx++ {
		mo.moveGenerator[idx].level = idx
	}
	mo.bestMove = make([]Move, levelMax)
	mo.bestMove[0].score = math.MinInt32
	mo.moveCounter = 0

	// Each level has to store the list of perfect moves
	mo.perfectMove = make([][]Move, levelMax)
	for idx := 0; idx < levelMax; idx++ {
		mo.perfectMove[idx] = make([]Move, 0)
	}
}

// Calculate the best move on the current level - Multi player mode
func (mo *MoveOptimizer) calcBestMoveDouble(stonesA Fields, stonesB Fields, freeStones int, levelMax int) Move {

	fmt.Printf("\ncalc best move for stones player  A: %v stones player B: %v freeStones: %d level_max: %d\n", stonesA, stonesB, freeStones, levelMax)
	level := 0
	mo.initBestMove(levelMax)
	mo.moveGenerator[0].init(stonesA, stonesB, freeStones)

	// Iterate over all possible moves
	for {
		// In this level generate a new move
		mg := &mo.moveGenerator[level]
		move, dstStoneA := mg.nextApplyMove(mg.stonesA)
		if move.valid {
			mo.moveCounter += 1

			if level+1 == levelMax {
				// On last level evaulate the result
				// Player A always coounts positive; B negative
				score := mo.rateField.rate(dstStoneA) - mo.rateField.rate(mg.stonesB)
				if score > mo.bestMove[level].score {
					// Found a better move -> save it
					move.score = score
					mo.bestMove[level] = move
					mo.perfectMove[level] = mo.perfectMove[level][:0]
					mo.perfectMove[level] = append(mo.perfectMove[level], move)
				}
				fmt.Printf("   score: %2d\n", score)
			} else {
				// DOWN one level: Prepare the move generator for the new level
				// Note: switch to other player - it's his turn
				level += 1
				mo.moveGenerator[level].setup(move, dstStoneA, mg.stonesB, mg.freeStones)

				// No best move yet, therefore reset the new level
				mo.bestMove[level].reset()
				mo.perfectMove[level] = mo.perfectMove[level][:0]
				fmt.Printf("\n")
			}
		} else {
			// No more move possible on this level
			if level == 0 {
				// On level zero we are done
				break
			}

			// UP one level: Use the score from child branch and go one level up
			// The best for the child is the worst for the parent -> Inversion
			score := -mo.bestMove[level].score
			level -= 1
			fmt.Printf(" [%d] up:     score: %2d > %2d\n", level, score, mo.bestMove[level].score)
			if score > mo.bestMove[level].score {
				// Use the best from the worst -> Save move in this level with score from below
				currentMove := mo.moveGenerator[level].current()
				currentMove.score = score
				mo.bestMove[level] = currentMove

				// Store perfect moves for this node
				mo.perfectMove[level] = append([]Move{currentMove}, mo.perfectMove[level+1]...)
			}
		}
	}
	fmt.Printf("perfect move (total=%d):\n", mo.moveCounter)
	for idx, move := range mo.perfectMove[0] {
		fmt.Printf("[%d] %v\n", idx, move)
	}
	return mo.perfectMove[0][0]
}
