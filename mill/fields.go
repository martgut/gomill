package main

import "fmt"

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
