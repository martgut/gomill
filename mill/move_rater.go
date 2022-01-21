package main

// Evaluate the value of a move
type Rater interface {
	rate(move Move) int
}

type EvalHighestField struct {
}

func (EvalHighestField) rate(move Move) int {
	return move.toField
}

type EvalLowestField struct {
}

func (EvalLowestField) rate(move Move) int {
	return -move.toField
}
