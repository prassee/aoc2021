package day6

import (
	"fmt"
	"strconv"
	"strings"
)

var days = 256

func Part1() {
	txtInput := strings.Split("2,3,1,3,4,4,1,5,2,3,1,1,4,5,5,3,5,5,4,1,2,1,1,1,1,1,1,4,1,1,1,4,1,3,1,4,1,1,4,1,3,4,5,1,1,5,3,4,3,4,1,5,1,3,1,1,1,3,5,3,2,3,1,5,2,2,1,1,4,1,1,2,2,2,2,3,2,1,2,5,4,1,1,1,5,5,3,1,3,2,2,2,5,1,5,2,4,1,1,3,3,5,2,3,1,2,1,5,1,4,3,5,2,1,5,3,4,4,5,3,1,2,4,3,4,1,3,1,1,2,5,4,3,5,3,2,1,4,1,4,4,2,3,1,1,2,1,1,3,3,3,1,1,2,2,1,1,1,5,1,5,1,4,5,1,5,2,4,3,1,1,3,2,2,1,4,3,1,1,1,3,3,3,4,5,2,3,3,1,3,1,4,1,1,1,2,5,1,4,1,2,4,5,4,1,5,1,5,5,1,5,5,2,5,5,1,4,5,1,1,3,2,5,5,5,4,3,2,5,4,1,1,2,4,4,1,1,1,3,2,1,1,2,1,2,2,3,4,5,4,1,4,5,1,1,5,5,1,4,1,4,4,1,5,3,1,4,3,5,3,1,3,1,4,2,4,5,1,4,1,2,4,1,2,5,1,1,5,1,1,3,1,1,2,3,4,2,4,3,1", ",")
	input := []int{}
	for _, v := range txtInput {
		input = append(input, toInt(v))
	}
	recurse(1, input)
	// recurse(1, []int{3, 4, 3, 1, 2})
}

func recurse(dayCount int, popln []int) {
	newPop := []int{}
	for i, v := range popln {
		newV := v - 1
		if newV < 0 {
			newV = 6
			newPop = append(newPop, 8)
		}
		popln[i] = newV
	}
	popln = append(popln, newPop...)
	if dayCount == days {
		fmt.Printf("day count %v total popln %v sum population %v \n", dayCount, len(popln), func(x []int) (sum int) {
			for v := range x {
				sum += v
			}
			return sum
		}(popln))
		return
	} else {
		recurse(dayCount+1, popln)
	}
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}
