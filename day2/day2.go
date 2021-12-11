/*
https://adventofcode.com/2021/day/2
*/
package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("../day2-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	coords := &Coords{}
	for reader.Scan() {
		inst := strings.Split(reader.Text(), " ")
		axis, step := inst[0], inst[1]
		stepN, _ := strconv.Atoi(step)
		switch axis {
		case "forward":
			coords.move(stepN, true)
		case "up":
			coords.move(-stepN, false)
		case "down":
			coords.move(stepN, false)
		}
	}
	log.Printf("final coord x %v and y %v product %v \n", coords.x, coords.y, coords.product())
}

func Part2() {
	file, err := os.Open("../day2-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	coords := &Coords{}
	for reader.Scan() {
		inst := strings.Split(reader.Text(), " ")
		axis, step := inst[0], inst[1]
		stepN, _ := strconv.Atoi(step)
		switch axis {
		case "forward":
			coords.move(stepN, true)
		case "up":
			coords.move(-stepN, false)
		case "down":
			coords.move(stepN, false)
		}
	}
	log.Printf("final coord x %v and y %v product %v \n", coords.x, coords.y, coords.product())
}

type Coords struct {
	x, y, aim, depth int
}

func (cord *Coords) move(step int, isFwd bool) {
	if isFwd {
		cord.x = cord.x + step
		cord.depth = cord.depth + (cord.aim * step)
	} else {
		cord.y = cord.y + step
		cord.aim = cord.aim + step
	}
}

func (cord *Coords) product() int {
	return cord.x * cord.depth
}
