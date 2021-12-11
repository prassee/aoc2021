/*
https://adventofcode.com/2021/day/3
*/
package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("../day3-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	grid := &Grid{}
	for reader.Scan() {
		row := []int{}
		for _, v := range reader.Bytes() {
			row = append(row, int(v-48))
		}
		grid.addRow(row)
	}
	grid.calcValues()
	log.Printf("power consumption is %v \n", grid.gamma*grid.epsilon)
}

func Part2() {
	file, err := os.Open("../day3-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	grid := &Grid{}
	for reader.Scan() {
		row := []int{}
		for _, v := range reader.Bytes() {
			row = append(row, int(v-48))
		}
		grid.addRow(row)
	}
	var o2Rating, co2Rating int64
	if len(grid.onesig) > len(grid.zerosig) {
		o2Rating = arrToBnry(reduceByBit(grid.onesig, 1, func(x, y int) bool { return x >= y }, func(x, y [][]int) [][]int { return y }), "")
		co2Rating = arrToBnry(reduceByBit(grid.zerosig, 1, func(x, y int) bool { return x <= y }, func(x, y [][]int) [][]int { return x }), "")
	} else {
		o2Rating = arrToBnry(reduceByBit(grid.zerosig, 1, func(x, y int) bool { return x >= y }, func(x, y [][]int) [][]int { return x }), "")
		co2Rating = arrToBnry(reduceByBit(grid.onesig, 1, func(x, y int) bool { return x <= y }, func(x, y [][]int) [][]int { return y }), "")
	}
	log.Printf("%v %v life support %v \n", o2Rating, co2Rating, o2Rating*co2Rating)
}

type Grid struct {
	matrix          [][]int
	cols, rows      int
	gamma, epsilon  int
	onesig, zerosig [][]int
}

func (g *Grid) addRow(row []int) {
	g.matrix = append(g.matrix, row)
	if g.cols == 0 {
		g.cols = int(len(g.matrix[0]))
	}
	if row[0] == 0 {
		g.zerosig = append(g.zerosig, row)
	} else {
		g.onesig = append(g.onesig, row)
	}
	g.rows += 1
}

func (g *Grid) calcValues() {
	msigRow := ""
	lsigRow := ""
	for j := 0; j < int(g.cols); j++ {
		var intersum = 0
		for i := range g.matrix {
			intersum = intersum + g.matrix[i][j]
		}
		if intersum >= g.rows-intersum {
			msigRow = msigRow + "1"
			lsigRow = lsigRow + "0"
		} else {
			msigRow = msigRow + "0"
			lsigRow = lsigRow + "1"
		}
	}
	gamma, _ := strconv.ParseInt(msigRow, 2, 64)
	epsilon, _ := strconv.ParseInt(lsigRow, 2, 64)
	g.gamma = int(gamma)
	g.epsilon = int(epsilon)
}

func reduceByBit(report [][]int, col int, sigFn func(x, y int) bool, tieBreak func(x, y [][]int) [][]int) []int {
	counts1 := [][]int{}
	counts0 := [][]int{}
	if len(report) == 1 {
		log.Printf("%v \n", report[0])
		return report[0]
	} else {
		for row := range report {
			if report[row][col] == 1 {
				counts1 = append(counts1, report[row])
			} else {
				counts0 = append(counts0, report[row])
			}
		}
		if len(counts0) == len(counts1) {
			return reduceByBit(tieBreak(counts0, counts1), col+1, sigFn, tieBreak)
		} else {
			if sigFn(len(counts1), len(counts0)) {
				return reduceByBit(counts1, col+1, sigFn, tieBreak)
			} else {
				return reduceByBit(counts0, col+1, sigFn, tieBreak)
			}
		}
	}
}

func arrToBnry(a []int, delim string) int64 {
	str := strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	asBin, _ := strconv.ParseInt(str, 2, 64)
	return asBin
}
