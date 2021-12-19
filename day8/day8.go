package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("../day8-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	sum := 0
	for reader.Scan() {
		sampleInput := reader.Text()
		segmentSignals := strings.Split(strings.Split(sampleInput, " | ")[1], " ")
		fmt.Printf("segmentSignal %v \n", segmentSignals)
		for _, v := range segmentSignals {
			sum += testFor1478(v)
		}
	}
	println(sum)
}

func Part2() {
	file, err := os.Open("../day8-input.txt")
	if err != nil {
		log.Fatal("cannot find input file ", err.Error())
	}
	reader := bufio.NewScanner(file)
	sum := 0
	for reader.Scan() {
		sampleLine := reader.Text()
		// "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"
		wiring := strings.Split(sampleLine, " | ")
		segMap := createWiringMap(wiring[0])
		digits := ""
		for _, v := range strings.Split(wiring[1], " ") {
			for k, seg := range segMap {
				if len(v) == len(k) && strSim(v, k) == v {
					digits += fmt.Sprint(seg)
					break
				}
			}
		}
		s, _ := strconv.Atoi(digits)
		sum += s
	}
	println(sum)
}

func createWiringMap(wiring string) (mapping map[string]int) {
	// part1Input := "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb"
	seg235 := []string{}
	seg069 := []string{}
	mapping = make(map[string]int)
	seg4 := ""
	seg9 := ""
	for _, v := range strings.Split(wiring, " ") {
		x := len(v)
		switch x {
		case 2:
			mapping[v] = 1
		case 4:
			mapping[v] = 4
			seg4 = v
		case 3:
			mapping[v] = 7
		case 7:
			mapping[v] = 8
		case 6:
			seg069 = append(seg069, v)
		case 5:
			seg235 = append(seg235, v)
		}
	}
	// from 069 arr any one should contain 4
	for i, v := range seg069 {
		// fmt.Println(seg4, " ", v, " ", strSim(seg4, v))
		if strSim(seg4, v) == seg4 {
			mapping[v] = 9
			seg9 = v
			// seg069[i] = ""
			seg069 = remove(seg069, i)
			break
		}
	}
	// eliminate a letter from 069 and that should match only one from 235
	for i, s5 := range seg235 {
		for j, s6 := range seg069 {
			if strSim(s5, s6) == s5 {
				mapping[s5] = 5
				mapping[s6] = 6
				// seg069[j] = ""
				seg069 = remove(seg069, j)
				mapping[seg069[0]] = 0
				seg069 = remove(seg069, 0)
				seg235[i] = ""
				seg235 = remove(seg235, i)
				break
			}
		}

	}

	// to find 3 & 2 -
	for i, v := range seg235 {
		if strSim(v, seg9) == v {
			mapping[v] = 3
			seg235 = remove(seg235, i)
			break
		}
	}
	mapping[seg235[0]] = 2
	return mapping
}

func testFor1478(s string) int {
	x := len(s)
	switch x {
	case 7:
		return 1 // 8
	case 4:
		return 1 // 4
	case 3:
		return 1 // 7
	case 2:
		return 1 // 1
	default:
		return 0
	}
}

// func bothDiff(a, b string) (d string) {
// 	return fmt.Sprintf("%v%v", strDiff(a, b), strDiff(b, a))
// }

// func strDiff(a, b string) (diff string) {
// 	for _, r1 := range a {
// 		if !strings.Contains(b, string(r1)) {
// 			diff = string(r1)
// 			break
// 		}
// 	}
// 	return diff
// }

func strSim(a, b string) (sim string) {
	for _, r1 := range a {
		if strings.Contains(b, string(r1)) {
			sim = sim + string(r1)
		}
	}
	return sim
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
