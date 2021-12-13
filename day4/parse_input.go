/*
https://adventofcode.com/2021/day/4
*/
package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseBingoFile(inputFile string, gridSize int) *BingoBoard {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	lines := []string{}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		x := reader.Text()
		if len(x) > 0 {
			lines = append(lines, x)
		}
	}
	return NewBingoBoard(lines[0], lines[1:], gridSize)
}

func NewBingoBoard(draws string, grids []string, gridSize int) *BingoBoard {
	inputDraws := strings.Split(draws, ",")
	intDraws := []int{}
	for _, v := range inputDraws {
		x, _ := strconv.Atoi(v)
		intDraws = append(intDraws, x)
	}
	ordNum := 1
	gridNum := BingoGrid{}
	bgrids := []*BingoGrid{}
	for _, v := range grids {
		gvalues := strings.Split(strings.ReplaceAll(v, " ", ","), ",")
		gridNums := []int{}
		for i := range gvalues {
			if gvalues[i] != "" {
				x, _ := strconv.Atoi(gvalues[i])
				gridNums = append(gridNums, x)
			}
		}
		gridNum.grid = append(gridNum.grid, gridNums)
		if ordNum%gridSize == 0 {
			x := &BingoGrid{
				grid:     gridNum.grid,
				marked:   initMarked(gridSize),
				gridSize: gridSize,
			}
			bgrids = append(bgrids, x)
			gridNum = BingoGrid{}
		}
		ordNum += 1
	}
	return &BingoBoard{
		draws: intDraws, grids: bgrids,
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

type BingoGrid struct {
	grid, marked [][]int
	gridSize     int
}

func (grid *BingoGrid) markOnGrid(x int) {
	gSize := grid.gridSize
	for i := 0; i < gSize; i++ {
		for j := 0; j < gSize; j++ {
			if grid.grid[i][j] == x {
				if x == 0 {
					x = -1
				}
				grid.marked[i][j] = x
			}
		}
	}
}

func (grid *BingoGrid) isBingo() bool {
	gSize := grid.gridSize
	for i := 0; i < gSize; i++ {
		rowsFilled := 0
		for j := 0; j < gSize; j++ {
			if grid.marked[j][i] != 0 {
				rowsFilled += 1
			}
		}
		if rowsFilled == gSize {
			return true
		}
	}
	for i := 0; i < gSize; i++ {
		colsFilled := 0
		for j := 0; j < gSize; j++ {
			if grid.marked[i][j] != 0 {
				colsFilled += 1
			}
		}
		if colsFilled == gSize {
			return true
		}
	}
	return false
}

func (grid *BingoGrid) boardSum() (sum int) {
	sum = 0
	for i := 0; i < grid.gridSize; i++ {
		for j := 0; j < grid.gridSize; j++ {
			if grid.grid[i][j] != grid.marked[i][j] {
				sum += grid.grid[i][j]
			}
		}
	}
	return sum
}

type BingoBoard struct {
	draws, bingoed []int
	grids          []*BingoGrid
}

func (b *BingoBoard) markDrawOnBoards(x int) (isfilled bool, sum []int) {
	sum = []int{}
	isfilled = false
	for i, v := range b.grids {
		if !b.isBingoed(i) {
			v.markOnGrid(x)
			if v.isBingo() {
				log.Printf("for draw %v board  index %v value %v", x, i, v.marked)
				b.bingoed = append(b.bingoed, i)
				sum = append(sum, v.boardSum())
				isfilled = isfilled || true
				// return true, v.boardSum()
			}
		}
	}
	return isfilled, sum
}

func (b BingoBoard) isBingoed(x int) bool {
	for i := range b.bingoed {
		if b.bingoed[i] == x {
			return true
		}
	}
	return false
}
