package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Day 2")
	reports, _ := utils.LoadDataAsInts("2.txt")
	partOneSolution := partOne(reports)
	fmt.Println(partOneSolution)
	// partTwoSolution := partTwo(lines)

}

const (
	Same = iota
	Increasing
	Decreasing
)

var validLevels []bool

func partOne(reports [][]int) int {
	// use goroutines to check each row independently
	// rule 1: the levels are all increasing or all descreasing
	// rule 2: any two adjacent levels (x1,x2) must satisfy  1  <= |x1 - x2| <= 3

	validLevels = make([]bool, len(reports))
	var wg sync.WaitGroup

	checkReport := func(i int) {
		defer wg.Done()
		prevLevel := reports[i][0]
		prevDelta := deltaLevels(prevLevel, reports[i][1])

		for j := 1; j < len(reports[i]); j++ {
			delta := deltaLevels(prevLevel, reports[i][j])
			level := reports[i][j]
			valid := ((prevDelta == delta) && (delta != Same)) && // rule 1
				(abs(prevLevel-level) >= 1) && (abs(prevLevel-level) <= 3) //rule 2
			if valid {
				validLevels[i] = true
				prevLevel = level
				prevDelta = delta
			} else {
				validLevels[i] = false
				break
			}
		}
	}

	for i := 0; i < len(reports); i++ {
		wg.Add(1)
		go checkReport(i)

	}

	wg.Wait()
	countValidLevels := 0
	for _, v := range validLevels {
		if v {
			countValidLevels++
		}
	}
	return countValidLevels
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func deltaLevels(a int, b int) int {
	if a == b {
		return Same
	}
	if a < b {
		return Increasing
	} else {
		return Decreasing
	}
}

// func partTwo() {
// 	fmt.Println("Part 2")
// 	return
// }
