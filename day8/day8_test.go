package day8

import (
	"fmt"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	Part1()
}

func TestPart2(t *testing.T) {
	Part2()
}

func plotOnSegment(s string, wire map[rune]int) {
	missing := make([]int, 7)
	l := len(s)
	digit := -1
	switch l {
	case 7:
		digit = 8
	case 4:
		digit = 4
	case 3:
		digit = 7
	case 2:
		digit = 1
	case 5:
		if !strings.ContainsRune(s, 'a') && !strings.ContainsRune(s, 'g') {
			digit = 5
		}
		if !strings.ContainsRune(s, 'e') && !strings.ContainsRune(s, 'g') {
			digit = 3
		}
		if !strings.ContainsRune(s, 'e') && !strings.ContainsRune(s, 'b') {
			digit = 2
		}
	case 6:
		if missing[3] == 0 {
			digit = 0
		}
		if missing[2] == 0 {
			digit = 6
		}
		if missing[4] == 0 {
			digit = 9
		}
	}

	fmt.Printf("digit %v \n", digit)
}
