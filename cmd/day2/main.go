package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Day 2")
	reportData, _ := utils.LoadDataAsInts("2.txt")
	reports := make([]Report, len(reportData))
	for i, levels := range reportData {
		reports[i] = Report{levels: levels, valid: false, row: i}
	}

	partOneSolution := partOne(reports)
	fmt.Println("Part One Solution:", partOneSolution)
	partTwoSolution := partTwo(reports)
	fmt.Println("Part Two Solution:", partTwoSolution)
}

const (
	Same = iota
	Increasing
	Decreasing
)

type Report struct {
	levels []int
	valid  bool
	row    int
}

func checkReport(report *Report, wg *sync.WaitGroup) {
	defer wg.Done()
	prevLevel := (*report).levels[0]
	prevDelta := deltaLevels(prevLevel, (*report).levels[1])

	for i := 1; i < len((*report).levels); i++ {
		level := (*report).levels[i]
		delta := deltaLevels(prevLevel, level)
		rule1 := ((prevDelta == delta) && (delta != Same))
		rule2 := (abs(prevLevel-level) >= 1) && (abs(prevLevel-level) <= 3)
		valid := rule1 && rule2

		if valid {
			prevLevel = level
			prevDelta = delta
			(*report).valid = true
		} else {
			(*report).valid = false
			break
		}
	}
}

func partOne(reports []Report) int {
	var wg sync.WaitGroup

	for i := 0; i < len(reports); i++ {
		report := &reports[i]
		wg.Add(1)
		go checkReport(report, &wg)
	}

	wg.Wait()
	countValidLevels := 0
	for _, report := range reports {
		if report.valid {
			countValidLevels++
		}
	}
	return countValidLevels
}

func partTwo(reports []Report) int {
	var wg sync.WaitGroup
	//generate damped reports
	dampedReports := make([]Report, 0)
	for _, report := range reports {
		dampedReports = append(dampedReports, report)
		dampened := dampenedReports(report.levels)
		for _, levels := range dampened {
			dampedReports = append(dampedReports, Report{levels: levels, valid: false, row: report.row})
		}
	}
	reports = dampedReports
	for _, report := range reports {
		if report.row == 0 {
			fmt.Println(report)
		}
	}

	for i := 0; i < len(reports); i++ {
		report := &reports[i]
		wg.Add(1)
		go checkReport(report, &wg)
	}
	wg.Wait()
	validRows := make(map[int]bool)
	for _, report := range reports {
		if report.valid {
			validRows[report.row] = true
		}
	}
	return len(validRows)

}

func dampenedReports(report []int) [][]int {
	var reports [][]int
	for i := 0; i < len(report); i++ {
		// Create a new slice excluding the element at index i
		var combination []int
		for j := 0; j < len(report); j++ {
			if j != i {
				combination = append(combination, report[j])
			}
		}
		reports = append(reports, combination)
	}
	return reports
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
