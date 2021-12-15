package day6

import (
	"fmt"
	"strconv"
	"strings"
)

var days = 256
var popDist = make(map[int]int)

/*part 1 & 2 are same except for change in `days`*/
func Part1() {
	txtInput :=
		//strings.Split("3,4,3,1,2", ",")
		strings.Split("2,3,1,3,4,4,1,5,2,3,1,1,4,5,5,3,5,5,4,1,2,1,1,1,1,1,1,4,1,1,1,4,1,3,1,4,1,1,4,1,3,4,5,1,1,5,3,4,3,4,1,5,1,3,1,1,1,3,5,3,2,3,1,5,2,2,1,1,4,1,1,2,2,2,2,3,2,1,2,5,4,1,1,1,5,5,3,1,3,2,2,2,5,1,5,2,4,1,1,3,3,5,2,3,1,2,1,5,1,4,3,5,2,1,5,3,4,4,5,3,1,2,4,3,4,1,3,1,1,2,5,4,3,5,3,2,1,4,1,4,4,2,3,1,1,2,1,1,3,3,3,1,1,2,2,1,1,1,5,1,5,1,4,5,1,5,2,4,3,1,1,3,2,2,1,4,3,1,1,1,3,3,3,4,5,2,3,3,1,3,1,4,1,1,1,2,5,1,4,1,2,4,5,4,1,5,1,5,5,1,5,5,2,5,5,1,4,5,1,1,3,2,5,5,5,4,3,2,5,4,1,1,2,4,4,1,1,1,3,2,1,1,2,1,2,2,3,4,5,4,1,4,5,1,1,5,5,1,4,1,4,4,1,5,3,1,4,3,5,3,1,3,1,4,2,4,5,1,4,1,2,4,1,2,5,1,1,5,1,1,3,1,1,2,3,4,2,4,3,1", ",")
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
	// on 2 days of swapping
	for i := 1; i <= days; i++ {
		dayDist := make(map[int]int)
		for k, v := range popDist {
			dayDist[k-1] = v
		}
		_, hask := dayDist[-1]
		if hask {
			_, has6 := dayDist[6]
			if has6 {
				dayDist[6] = dayDist[6] + dayDist[-1]
				dayDist[8] = dayDist[8] + dayDist[-1]
			} else {
				dayDist[6] = dayDist[-1]
				dayDist[8] = dayDist[8] + dayDist[-1]
			}
			delete(dayDist, -1)
		}
		popDist = dayDist
	}
	sum := 0
	for _, v := range popDist {
		sum += v
	}
	fmt.Printf("after %v days of %v sum %v \n", days, popDist, sum)
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return i
}
