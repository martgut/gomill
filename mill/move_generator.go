package main

import "fmt"

type playerT int64

const (
	playerA playerT = iota
	playerB
)

func (pl playerT) String() string {
	return []string{"A", "B"}[pl]
}

type generatorMode int64

const (
	// Start & end game jumping
	placeStone generatorMode = iota

	// Middle game, moving stones
	moveStone

	// Remove stone after close mill
	removeStone
)

func (gt generatorMode) String() string {
	return []string{"place", "move", "remove"}[gt]
}

type MoveGenerator struct {

	// Level
	level int

	// Mode of generation
	mode generatorMode

	// Current player
	player playerT

	// List of stones to generate moves for player A
	stonesA Fields

	// List of stones of other player B
	stonesB Fields

	// For current move which stone
	stoneIndex int

	// For current move which index of move option
	moveIndex int

	// Number available stones to place
	freeStones int

	// Current evaluated move
	currentMove Move
}

func (mg *MoveGenerator) reset(mode generatorMode, stonesA Fields, stonesB Fields) {
	mg.mode = mode
	mg.stonesA = stonesA
	mg.stonesB = stonesB
	mg.moveIndex = -1
	mg.stoneIndex = 0
}

func (mg *MoveGenerator) init(stonesA Fields, stonesB Fields, freeStones int) {

	if freeStones > 0 {
		mg.mode = placeStone
	} else {
		mg.mode = moveStone
	}
	mg.stonesA = stonesA
	mg.stonesB = stonesB
	mg.moveIndex = -1
	mg.stoneIndex = 0
	mg.freeStones = freeStones
}

func (mg *MoveGenerator) setup(lastMove Move, stonesA Fields, stonesB Fields, freeStones int) {

	mg.moveIndex = -1
	mg.stoneIndex = 0

	if lastMove.isMill {
		// Keep current player
		fmt.Println("WASMILL")
		mg.stonesA = stonesA
		mg.stonesB = stonesB
		mg.mode = removeStone
	} else {
		// Switch player
		mg.player = (mg.player + 1) % 2
		mg.stonesA = stonesB
		mg.stonesB = stonesA

		if freeStones > 0 {
			mg.mode = placeStone
			mg.freeStones = freeStones - 1
		} else {
			mg.mode = moveStone
		}
	}
}

// Return the current move
func (mg *MoveGenerator) current() Move {
	return mg.currentMove
}

// Place a stone on the field
func (mg *MoveGenerator) placeStone() Move {

	for {
		// Same stone but next move option
		mg.moveIndex += 1

		// Check whether we reached last field
		if mg.moveIndex == 24 {
			return Move{valid: false}
		}
		// Check for free field
		to := mg.moveIndex
		if !mg.stonesA.contains(to) && !mg.stonesB.contains(to) {
			// Check for closing mill
			mill := allMills.isMill(to, mg.stonesA) > 0
			if mill {
				fmt.Println("MILL")
			}
			return Move{toField: to, valid: true, mode: placeStone, isMill: mill}
		}
	}
}

// Move a stone on the field
func (mg *MoveGenerator) moveStone() Move {

	for {
		// Same stone but next move option
		mg.moveIndex += 1

		// Check whether we reached last stone
		if mg.stoneIndex >= len(mg.stonesA) {
			return Move{fromField: 0, toField: 0, valid: false}
		}

		// Determine next possible move
		from := mg.stonesA[mg.stoneIndex]
		to, found := allMoves.getMove(from, mg.moveIndex)

		if found {
			// Check whether a stone is already placed on that field
			if !mg.stonesA.contains(to) && !mg.stonesB.contains(to) {
				DebugLogger.Printf("Found: from: %v to: %v", from, to)

				// Was a mill closed with this move?
				isMill := allMills.isMill(to, mg.stonesA) > 0
				return Move{stoneIndex: mg.stoneIndex, fromField: from, toField: to, valid: true, isMill: isMill, mode: moveStone}
			}
		} else {
			// Try with next stone
			mg.stoneIndex += 1
			mg.moveIndex = -1
		}
	}
}

// Remove stone from the field
func (mg *MoveGenerator) removeStone() Move {
	panic("Not implemented!!!")
}

// Calculate best next move and apply it
func (mg *MoveGenerator) nextApplyMove(srcStones Fields) (Move, Fields) {

	// Find and apply best move
	move := mg.nextMove()
	dstStones := mg.applyMove(srcStones, move)

	// For debugging purpose
	if move.valid {
		pp := move.String()
		fmt.Printf("%s[%d] %s: src: %2v dst: %2v", mg.player, mg.level, pp, srcStones, dstStones)
	}
	return move, dstStones
}

// Return the next move if found
func (mg *MoveGenerator) nextMove() Move {

	switch mg.mode {
	case placeStone:
		mg.currentMove = mg.placeStone()
	case moveStone:
		mg.currentMove = mg.moveStone()
	case removeStone:
		mg.currentMove = mg.removeStone()
	}
	return mg.currentMove
}

// Apply move on stones and return new one
func (mg *MoveGenerator) applyMove(srcStones Fields, move Move) Fields {

	var dstStones Fields

	switch move.mode {
	case placeStone:
		dstStones = make(Fields, len(srcStones), len(srcStones)+1)
		copy(dstStones, srcStones)
		dstStones = append(dstStones, move.toField)

	case moveStone:
		dstStones = make(Fields, len(srcStones))
		copy(dstStones, srcStones)
		dstStones[move.stoneIndex] = move.toField

	case removeStone:
		dstStones = make(Fields, len(srcStones)-1)
		copy(dstStones, srcStones)
		dstStones[move.toField] = srcStones[len(srcStones)-1]
		return dstStones
	}
	return dstStones
}
