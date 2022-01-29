package main

import "fmt"

type playerT int

const (
	playerA playerT = 1
	playerB playerT = -1
)

func (pl playerT) String() string {
	if pl == playerA {
		return "A"
	}
	return "B"
}

// Which stones are used for playing
type stoneT int

const (
	stoneA stoneT = 1
	stoneB stoneT = -1
)

func (pl stoneT) String() string {
	if pl == stoneA {
		return "A"
	} else {
		return "B"
	}
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

	// Which stones are used for playing (A/B)?
	stones stoneT

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

// Return Score factor according to involved players
func (mg *MoveGenerator) scoreFactor(mg2 *MoveGenerator) int {
	if mg.player == mg2.player {
		// Keep the score
		return 1
	} else {
		// Invert the score
		return -1
	}
}

func (mg *MoveGenerator) reset(mode generatorMode, stonesA Fields, stonesB Fields) {
	mg.player = playerA
	mg.stones = stoneA
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
	mg.player = playerA
	mg.stones = stoneA
	mg.stonesA = stonesA
	mg.stonesB = stonesB
	mg.moveIndex = -1
	mg.stoneIndex = 0
	mg.freeStones = freeStones
}

// Compares two scores and returns whether second is better
func (mg *MoveGenerator) evalScore(newScore int, currentScore int) bool {
	return newScore > currentScore
}

func (mg *MoveGenerator) setup(lastMove Move, stonesA Fields, mgPrev *MoveGenerator) {

	mg.moveIndex = -1
	mg.stoneIndex = 0

	if lastMove.isMill {
		// Keep current player
		mg.player = mgPrev.player
		mg.mode = removeStone

		// Still have to switch stones
		mg.stones = mgPrev.stones * -1
		mg.stonesA = mgPrev.stonesB
		mg.stonesB = stonesA
	} else {
		// Switch player
		mg.stones = mgPrev.stones * -1
		mg.player = mgPrev.player * -1
		mg.stonesA = mgPrev.stonesB
		mg.stonesB = stonesA

		if mgPrev.freeStones > 0 {
			mg.mode = placeStone
			mg.freeStones = mgPrev.freeStones - 1
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
		// Next move option
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

	for {
		// Focus on stone, but advance for the next
		stone := mg.stoneIndex
		mg.stoneIndex += 1

		// Check whether we reached last stone
		if stone >= len(mg.stonesA) {
			return Move{valid: false}
		}
		// Check whether stone is part of a mill
		if true { // TODO implement mill check
			return Move{stoneIndex: stone, toField: mg.stonesA[stone], valid: true, isMill: false, mode: removeStone}
		}
	}
}

// Calculate best next move and apply it
func (mg *MoveGenerator) nextApplyMove(srcStones Fields) (Move, Fields) {

	// Find and apply best move
	move := mg.nextMove()
	dstStones := mg.applyMove(srcStones, move)

	// For debugging purpose
	if move.valid {
		pp := move.String()
		fmt.Printf("%s%s[%d] %s: src: %2v dst: %2v", mg.player, mg.stones, mg.level, pp, srcStones, dstStones)
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
		// Note: A new stone is added with value: "toField" for the new position
		dstStones = make(Fields, len(srcStones), len(srcStones)+1)
		copy(dstStones, srcStones)
		dstStones = append(dstStones, move.toField)

	case moveStone:
		// Note: The stone with index: "stoneIndex" will move to: "toField"
		dstStones = make(Fields, len(srcStones))
		copy(dstStones, srcStones)
		dstStones[move.stoneIndex] = move.toField

	case removeStone:
		// Note: The stone with index: "stoneIndex" will be deleted

		// We copy all from source, but miss the last one
		dstStones = make(Fields, len(srcStones)-1)
		copy(dstStones, srcStones)

		// Recover the last one to the stone which should be deleted
		if move.stoneIndex < len(dstStones) {
			dstStones[move.stoneIndex] = srcStones[len(srcStones)-1]
		}
		return dstStones
	}
	return dstStones
}
