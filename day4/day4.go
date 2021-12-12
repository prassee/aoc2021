package day4

import "log"

var bingoBoard = parseBingoFile("../day4-input.txt", 5)

func Part1() {
	for _, v := range bingoBoard.draws {
		status, sum := bingoBoard.markNum(v)
		if status {
			log.Printf("sum of board is %v @ draw %v score %v \n", sum, v, sum*v)
			break
		}
	}
}

func Part2() {

}
