package day4

import "log"

var bingoBoard = parseBingoFile("../day4-input.txt", 5)

func Part1() {
	for _, v := range bingoBoard.draws {
		status, sum := bingoBoard.markDrawOnBoards(v)
		if status {
			log.Printf("sum of board is %v @ draw %v score %v \n", sum, v, sum)
			break
		}
	}
}

func Part2() {
	for _, v := range bingoBoard.draws {
		status, sum := bingoBoard.markDrawOnBoards(v)
		if status && len(bingoBoard.grids) == len(bingoBoard.bingoed) {
			log.Printf("draw %v sum %v score %v", v, sum[0], v*sum[0])
		}
	}
}
