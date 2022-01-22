package main

import "fmt"

type Mills [24][2][2]int

// List of all variants of mills for each field
var allMills = Mills{
	{ // field 0
		{1, 2},
		{9, 21},
	},
	{ // field 1
		{0, 2},
		{4, 7},
	},
	{ // field 2
		{0, 1},
		{14, 23},
	},
	{ // field 3
		{4, 5},
		{10, 18},
	},
	{ // field 4
		{1, 7},
		{3, 5},
	},
	{ // field 5
		{3, 4},
		{13, 20},
	},
	{ // field 6
		{7, 8},
		{11, 15},
	},
	{ // field 7
		{6, 8},
		{1, 4},
	},
	{ // field 8
		{6, 7},
		{12, 17},
	},
	{ // field 9
		{0, 21},
		{10, 11},
	},
	{ // field 10
		{9, 11},
		{3, 18},
	},
	{ // field 11
		{6, 15},
		{9, 10},
	},
	{ // field 12
		{8, 17},
		{13, 14},
	},
	{ // field 13
		{12, 14},
		{5, 20},
	},
	{ // field 14
		{2, 23},
		{12, 13},
	},
	{ // field 15
		{6, 11},
		{16, 17},
	},
	{ // field 16
		{15, 17},
		{19, 22},
	},
	{ // field 17
		{15, 16},
		{8, 12},
	},
	{ // field 18
		{3, 10},
		{19, 20},
	},
	{ // field 19
		{18, 20},
		{16, 22},
	},
	{ // field 20
		{18, 19},
		{5, 13},
	},
	{ // field 21
		{0, 9},
		{22, 23},
	},
	{ // field 22
		{21, 23},
		{16, 19},
	},
	{ // field 23
		{21, 22},
		{2, 14},
	},
}

func (m *Mills) print() {

	fmt.Println("Print all possible mills:")
	for idx, field := range m {
		fmt.Printf("%2d -> [%2d,%2d] [%2d,%2d]\n", idx, field[0][0], field[0][1], field[1][0], field[1][1])
	}
}

// Check when stone is placed on field whether mills
func (m *Mills) isMill(field int, stones Fields) int {

	// Number of found mills
	numberMills := 0

	// Focus on field for potential mill
	millsForStone := allMills[field]

	// Next iterate over all two variants
	for _, variant := range millsForStone {

		// Are stones at exact that positions?
		if stones.contains(variant[0]) && stones.contains(variant[1]) {
			DebugLogger.Printf("MILL for field: %v with stone: %v\n", field, stones)
			numberMills += 1
		}
	}
	return numberMills
}
