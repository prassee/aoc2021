package day6

import (
	"fmt"
	"strconv"
	"strings"
)

var days = 80
var input = []int{}
var popDist = make(map[int]int)

/*part 1 & 2 are same except for change in `days`*/
func Part1() {
	// strings.Split("2,3,1,3,4,4,1,5,2,3,1,1,4,5,5,3,5,5,4,1,2,1,1,1,1,1,1,4,1,1,1,4,1,3,1,4,1,1,4,1,3,4,5,1,1,5,3,4,3,4,1,5,1,3,1,1,1,3,5,3,2,3,1,5,2,2,1,1,4,1,1,2,2,2,2,3,2,1,2,5,4,1,1,1,5,5,3,1,3,2,2,2,5,1,5,2,4,1,1,3,3,5,2,3,1,2,1,5,1,4,3,5,2,1,5,3,4,4,5,3,1,2,4,3,4,1,3,1,1,2,5,4,3,5,3,2,1,4,1,4,4,2,3,1,1,2,1,1,3,3,3,1,1,2,2,1,1,1,5,1,5,1,4,5,1,5,2,4,3,1,1,3,2,2,1,4,3,1,1,1,3,3,3,4,5,2,3,3,1,3,1,4,1,1,1,2,5,1,4,1,2,4,5,4,1,5,1,5,5,1,5,5,2,5,5,1,4,5,1,1,3,2,5,5,5,4,3,2,5,4,1,1,2,4,4,1,1,1,3,2,1,1,2,1,2,2,3,4,5,4,1,4,5,1,1,5,5,1,4,1,4,4,1,5,3,1,4,3,5,3,1,3,1,4,2,4,5,1,4,1,2,4,1,2,5,1,1,5,1,1,3,1,1,2,3,4,2,4,3,1", ",")
	txtInput := strings.Split("3,4,3,1,2", ",")
	/*for i := 0; i <= 8; i++ {
		popDist[i] = 0
	}*/
	for _, v := range txtInput {
		x := toInt(v)
		_, hasKey := popDist[x]
		if hasKey {
			popDist[x] = popDist[x] + 1
		} else {
			popDist[x] = 1
		}
	}
	fmt.Printf("on day 0 given map is %v \n", popDist)
	// prevLen := len(input)
	// for i := 1; i <= days; i++ {
	// 	for i := 0; i < prevLen; i++ {
	// 		newV := input[i] - 1
	// 		if newV < 0 {
	// 			newV = 6
	// 			input = append(input, 8)
	// 		}
	// 		input[i] = newV
	// 	}
	// 	prevLen = len(input)
	// 	fmt.Printf("on day %v popln %v \n", i, prevLen)
	// }
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}
