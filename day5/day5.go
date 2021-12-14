/*
https://adventofcode.com/2021/day/5
*/
package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("../day5-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	plane := NewPlane(1000)
	for reader.Scan() {
		text := reader.Text()
		pointPair := strings.Split(text, " -> ")
		plane.markOnPlane(strToPoint(pointPair[0]), strToPoint(pointPair[1]))
	}
	fmt.Printf("count %v ", plane.countIntersections())
}

type Plane struct {
	grid [][]int
}

type Point struct {
	c, r int
}

func NewPoint(c, r int) Point {
	return Point{c: c, r: r}
}

func NewPlane(gridSize int) *Plane {
	return &Plane{
		grid: initMarked(gridSize),
	}
}

func initMarked(gridSize int) [][]int {
	grid := make([][]int, gridSize)
	// Initialize the gSize empty slices
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}
	return grid
}

func (plane *Plane) markOnPlane(p1, p2 Point) {
	// horizontal and vertical lines - todo instead of multiple for loops swap out the points
	if p1.c == p2.c || p1.r == p2.r {
		if p1.c == p2.c {
			if p1.r < p2.r {
				for i := p1.r; i <= p2.r; i++ {
					plane.grid[p1.c][i] += 1
				}
			} else if p1.r > p2.r {
				for i := p1.r; i >= p2.r; i-- {
					plane.grid[p1.c][i] += 1
				}
			}
		} else if p1.r == p2.r {
			if p1.c < p2.c {
				for i := p1.c; i <= p2.c; i++ {
					plane.grid[i][p1.r] += 1
				}
			} else if p1.c > p2.c {
				for i := p1.c; i >= p2.c; i-- {
					plane.grid[i][p1.r] += 1
				}
			}
		}
	} else {
		// diagonal lines
		// first swap p1 & p2 based on asc order
		fmt.Printf("pointing %v %v \n", p1, p2)
		if p1.c > p2.c {
			p3 := p1
			p1 = p2
			p2 = p3
		}
		fmt.Printf("after swap %v %v \n", p1, p2)
		if p1.c == p1.r && p2.c == p2.r {
			// logic to mark the points
			for i := p1.c; i <= p2.c; i++ {
				plane.grid[i][i] += 1
			}
		} else if math.Abs(float64(p1.c)-float64(p2.c)) == math.Abs(float64(p1.r)-float64(p2.r)) {
			if p1.c < p2.c && p1.r < p2.r {
				c, r := p1.c, p1.r
				plane.grid[c][r] += 1
				for {
					c += 1
					r += 1
					plane.grid[c][r] += 1
					if c == p2.c && r == p2.r {
						break
					}
				}
			} else {
				c, r := p1.c, p1.r
				plane.grid[c][r] += 1
				for {
					c += 1
					r -= 1
					plane.grid[c][r] += 1
					if c == p2.c && r == p2.r {
						break
					}
				}
			}
		}
	}

}

func (plane *Plane) countIntersections() (count int) {
	for i := 0; i < len(plane.grid); i++ {
		for j := 0; j < len(plane.grid); j++ {
			if plane.grid[i][j] >= 2 {
				count += 1
			}
		}
	}
	return count
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}

func strToPoint(s string) (p Point) {
	x := strings.Split(strings.TrimSpace(s), ",")
	p.c = toInt(x[0])
	p.r = toInt(x[1])
	return p
}
