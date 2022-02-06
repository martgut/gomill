package main

import (
	"fmt"
	"unicode"
)

// List of Fields
type Fields []int

func (fields *Fields) contains(ff int) bool {
	for _, value := range *fields {
		if value == ff {
			return true
		}
	}
	return false
}

// Check whether both fields are the same
func (fields *Fields) same(fieldsOther *Fields) bool {

	if len(*fields) != len(*fieldsOther) {
		return false
	}
	for i := 0; i < len(*fields); i++ {
		if (*fields)[i] != (*fieldsOther)[i] {
			return false
		}
	}
	return true
}

func (fields *Fields) asChar(f int) rune {
	if fields.contains(f) {
		return 'o'
	} else {
		return '.'
	}
}

func (fields *Fields) printPlayingField() {
	fmt.Printf("%c     %c     %c\n", fields.asChar(0), fields.asChar(1), fields.asChar(2))
	fmt.Printf("  %c   %c   %c  \n", fields.asChar(3), fields.asChar(4), fields.asChar(5))
	fmt.Printf("    %c %c %c    \n", fields.asChar(6), fields.asChar(7), fields.asChar(8))
	fmt.Printf("%c %c %c   %c %c %c\n", fields.asChar(9), fields.asChar(10), fields.asChar(11), fields.asChar(12), fields.asChar(13), fields.asChar(14))
	fmt.Printf("    %c %c %c    \n", fields.asChar(15), fields.asChar(16), fields.asChar(17))
	fmt.Printf("  %c   %c   %c  \n", fields.asChar(18), fields.asChar(19), fields.asChar(20))
	fmt.Printf("%c     %c     %c\n", fields.asChar(21), fields.asChar(22), fields.asChar(23))
}

func (fields Fields) cp() Fields {
	dst := make(Fields, len(fields))
	copy(dst, fields)
	return dst
}

// Apply move on stones and return new one
func (fields Fields) applyMove(move Move) Fields {

	var dstStones Fields

	switch move.mode {
	case placeStone:
		// Note: A new stone is added with value: "toField" for the new position
		dstStones = make(Fields, len(fields), len(fields)+1)
		copy(dstStones, fields)
		dstStones = append(dstStones, move.toField)

	case moveStone:
		// Note: The stone with index: "stoneIndex" will move to: "toField"
		dstStones = make(Fields, len(fields))
		copy(dstStones, fields)
		dstStones[move.stoneIndex] = move.toField

	case removeStone:
		// Note: The stone with index: "stoneIndex" will be deleted

		// We copy all from source, but miss the last one
		dstStones = make(Fields, len(fields)-1)
		copy(dstStones, fields)

		// Recover the last one to the stone which should be deleted
		if move.stoneIndex < len(dstStones) {
			dstStones[move.stoneIndex] = fields[len(fields)-1]
		}
		return dstStones
	}
	return dstStones
}

type playFieldT struct {
	stonesA    Fields
	stonesB    Fields
	index      int
	printLarge bool
}

func (pf *playFieldT) asChar() rune {
	c := '+'
	if pf.stonesA.contains(pf.index) {
		c = 'x'
	} else if pf.stonesB.contains(pf.index) {
		c = 'o'
	}
	pf.index += 1
	if pf.printLarge {
		return unicode.ToUpper(c)
	}
	return c
}

func (pf *playFieldT) printField() {
	if pf.printLarge {
		pf.printFieldLarge()
	} else {
		pf.printFieldSmall()
	}
}

func (pf *playFieldT) printFieldSmall() {
	fmt.Printf("%c     %c     %c\n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("  %c   %c   %c  \n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("    %c %c %c    \n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("%c %c %c   %c %c %c\n", pf.asChar(), pf.asChar(), pf.asChar(), pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("    %c %c %c    \n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("  %c   %c   %c  \n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("%c     %c     %c\n", pf.asChar(), pf.asChar(), pf.asChar())
}

func (pf *playFieldT) printFieldLarge() {

	fmt.Printf("%c--------%c--------%c\n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("|        |        |\n")
	fmt.Printf("|  %c-----%c-----%c  |\n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("|  |     |     |  |\n")
	fmt.Printf("|  |  %c--%c--%c  |  |\n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("|  |  |     |  |  |\n")
	fmt.Printf("%c--%c--%c     %c--%c--%c\n", pf.asChar(), pf.asChar(), pf.asChar(), pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("|  |  |     |  |  |\n")
	fmt.Printf("|  |  %c--%c--%c  |  |\n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("|  |     |     |  |\n")
	fmt.Printf("|  %c-----%c-----%c  |\n", pf.asChar(), pf.asChar(), pf.asChar())
	fmt.Printf("|        |        |\n")
	fmt.Printf("%c--------%c--------%c\n", pf.asChar(), pf.asChar(), pf.asChar())

}
