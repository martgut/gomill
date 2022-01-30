package main

import (
	"fmt"
	"math"
)

type MoveOptimizer struct {

	// Instance which rates the field and provides scores
	rater RateField

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
func (mo *MoveOptimizer) calcBestMove(stonesA Fields, stonesB Fields, freeStones int, levelMax int) Move {

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
				// A stones count positive; B stones count negative
				score := (mo.rater.rate(dstStoneA) - mo.rater.rate(mg.stonesB)) * int(mg.stones)
				if mg.evalScore(score, mo.bestMove[level].score) {
					// Found a better move -> save it
					move.score = score
					mo.bestMove[level] = move
					mo.perfectMove[level] = mo.perfectMove[level][:0]
					mo.perfectMove[level] = append(mo.perfectMove[level], move)
				}
				fmt.Printf("   score: %2d\n", score)
			} else {
				// DOWN one level: Prepare the move generator for the new level
				level += 1
				mgDown := &mo.moveGenerator[level]
				mgDown.setup(move, dstStoneA, mg)

				// No best move yet, therefore reset the new level
				mo.bestMove[level].reset(mgDown.stones)
				mo.perfectMove[level] = mo.perfectMove[level][:0]
				fmt.Printf("\n")
			}
		} else {
			// No more move possible on this level
			if level == 0 {
				// On level zero we are done
				break
			}

			// UP one level: Use the downScore from child branch and go one level up
			downScore := mo.bestMove[level].score
			level -= 1
			upScore := mo.bestMove[level].score
			mgUp := &mo.moveGenerator[level]
			fmt.Printf(" [%d] up:     score: %2d > %2d\n", level, downScore, upScore)
			if mgUp.evalScore(downScore, upScore) {
				// Use the best from the worst -> Save move in this level with score from below
				currentMove := mgUp.current()
				currentMove.score = downScore
				mo.bestMove[level] = currentMove

				// Store perfect moves for this node
				mo.perfectMove[level] = append([]Move{currentMove}, mo.perfectMove[level+1]...)
			}
		}
	}
	fmt.Printf("perfect move: total=%d score=%v:\n", mo.moveCounter, mo.pMove(0).score)
	for idx, move := range mo.perfectMove[0] {
		fmt.Printf("[%d] %v\n", idx, move)
	}
	return mo.perfectMove[0][0]
}
