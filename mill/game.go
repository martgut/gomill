package main

type Game struct {
	// Stones of player A
	stonesA Fields

	// Stones of player B
	stonesB Fields

	// Think level
	level int

	// Number of stones to place
	freeStones int

	// Optimizer to calculate the best move
	mo MoveOptimizer
}

// Constructor to create new game
func newGame() *Game {
	game := new(Game)
	game.mo.rater = HighestFieldsRater{}
	game.level = 3
	game.freeStones = 18
	return game
}

// Calculate next best move
func (game *Game) calcBestMove() {
	game.mo.calcBestMove(game.stonesA, game.stonesB, game.freeStones, game.level)
	game.mo.printPerfectMove()
}

// Pretty print player field
func (game *Game) print() {
	pf := playFieldT{stonesA: game.stonesA, stonesB: game.stonesB}
	pf.printField()
}
