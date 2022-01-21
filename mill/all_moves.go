package main

import "fmt"

// List of AllMoves for each field
type AllMoves [24]Fields

// TODO check moves
var allMoves = AllMoves{
	{1, 9},           //  0
	{0, 2, 4},        //  1
	{1, 14},          //  2
	{4, 10},          //  3
	{1, 3, 5, 7},     //  4
	{4, 13},          //  5
	{7, 11},          //  6
	{4, 6, 8},        //  7
	{7, 12},          //  8
	{0, 10, 21},      //  9
	{3, 9, 11, 18},   // 10
	{6, 10, 15},      // 11
	{8, 13, 17},      // 12
	{5, 12, 14, 20},  // 13
	{2, 13, 23},      // 14
	{11, 16},         // 15
	{15, 17, 19},     // 16
	{12, 16},         // 17
	{10, 19},         // 18
	{16, 18, 20, 22}, // 19
	{13, 19},         // 20
	{9, 22},          // 21
	{19, 21, 23},     // 22
	{14, 22},         // 23
}

// For a specific field, get move by index
func (m *AllMoves) getMove(from int, toIndx int) (to int, found bool) {

	if toIndx >= len(m[from]) {
		return 0, false
	}
	return m[from][toIndx], true
}

// Pretty print all moves
func (m *AllMoves) print() {
	fmt.Println("Print all possible moves:")
	for from, target := range m {
		fmt.Printf("%2d -> ", from)
		for _, v := range target {
			fmt.Printf("%3d ", v)
		}
		fmt.Println()
	}
}
