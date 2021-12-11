/*
https://adventofcode.com/2021/day/1
*/
package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part1() {
	file, err := os.Open("day1-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	var prevValue = 0
	var increments = 0
	for reader.Scan() {
		currValue, _ := strconv.Atoi(reader.Text())
		log.Printf("prev value %v curr value %v is greater %v \n", prevValue, currValue, currValue > prevValue)
		if currValue > prevValue && prevValue != 0 {
			increments = increments + 1
		}
		prevValue = currValue
	}
	log.Printf("no of increments  %v \n", increments)
}

/*
same as above but introduces Reducer Struct for the sliding window
*/
func Part2() {
	file, err := os.Open("day1-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	reducer := &Reducer{
		size: 3,
	}
	var prevValue = 0
	var increments = 0
	for reader.Scan() {
		obs, _ := strconv.Atoi(reader.Text())
		reducer.add(obs)
		currValue := reducer.reduce()
		log.Printf("prev value %v curr value %v is greater %v window %v \n", prevValue, currValue, currValue > prevValue, len(reducer.window))
		if currValue > prevValue && prevValue != 0 {
			increments = increments + 1
		}
		prevValue = currValue
	}
	log.Printf("no of increments  %v \n", increments)
}

type Reducer struct {
	window []int
	size   int
}

func (wind *Reducer) add(a int) {
	if len(wind.window) == wind.size {
		tail := wind.window[1:]
		wind.window = append(tail, a)
	} else {
		wind.window = append(wind.window, a)
	}
}

func (wind *Reducer) reduce() (sum int) {
	if len(wind.window) == wind.size {
		for _, v := range wind.window {
			sum = sum + v
		}
		return sum
	}
	return sum
}
