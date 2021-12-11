package day2

import (
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	Part1()
}

/*
while writing this test I found that its better to assert with
logic with the problem description.
*/
func TestPart2Eg(t *testing.T) {
	coords := &Coords{}
	coords.move(5, true)
	coords.move(5, false)
	coords.move(8, true)
	coords.move(-3, false)
	coords.move(8, false)
	coords.move(2, true)
	log.Printf("coords details x %v y %v aim %v depth %v  product %v \n", coords.x, coords.y, coords.aim, coords.depth, coords.product())
}

func TestPart2(t *testing.T) {
	Part2()
}
