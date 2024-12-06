package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Day 5\n\n")
	//load data
	lines, _ := utils.LoadData("5.txt")
	ruleLines, pageLines := splitLines(lines, "")
	rules := loadRules(ruleLines)
	pages := loadPages(pageLines)

	partOneSolution := partOne(rules, pages)
	fmt.Printf("Part 1: %v\n", partOneSolution)

	partTwoSolution := partTwo(rules, pages)
	fmt.Printf("Part 2: %v\n", partTwoSolution)
}

// split the file into rules and pages
func splitLines(lines []string, split string) ([]string, []string) {
	var splitIndex int
	for i, line := range lines {
		if line == split {
			splitIndex = i
			break
		}
	}
	return lines[:splitIndex], lines[splitIndex+1:]
}

// load the rules about order of pages
func loadRules(lines []string) map[int][]int {
	rules := make(map[int][]int)
	for _, rule := range lines {
		parts := strings.Split(rule, "|")
		k, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		rules[k] = append(rules[k], v)
	}
	return rules
}

// load pages into a 2d array
func loadPages(lines []string) [][]int {
	pages := make([][]int, len(lines))
	for i, line := range lines {
		pageNumbersStr := strings.Split(line, ",")
		pageNumbers := make([]int, len(pageNumbersStr))
		for j, pageNumberStr := range pageNumbersStr {
			pageNumber, _ := strconv.Atoi(pageNumberStr)
			pageNumbers[j] = pageNumber
		}
		pages[i] = pageNumbers
	}
	return pages
}

// Returns true if the single page number violates a rule
func pageNumViolatesRule(rules map[int][]int, pageNumber int, seen map[int]bool) bool {
	pageNumRules := rules[pageNumber]
	if len(pageNumRules) == 0 {
		return false
	}
	for _, rule := range pageNumRules {
		if seen[rule] {
			return true
		}
	}
	return false
}

// Returns true if any of the page numbers violate a rule
func validatePage(rules map[int][]int, page []int) bool {
	seen := make(map[int]bool)
	for _, pageNumber := range page {
		if pageNumViolatesRule(rules, pageNumber, seen) {
			return false
		}
		seen[pageNumber] = true
	}
	return true
}

func partOne(rules map[int][]int, pages [][]int) int {
	//filter out invalid pages
	validPages := make([][]int, len(pages))
	for i, pageLine := range pages {
		if validatePage(rules, pageLine) {
			validPages[i] = pageLine
		}
	}
	//sum middle number of lines consisting of only valid pages
	middleNumberSum := 0
	for _, page := range validPages {
		if page != nil {
			middleNumberSum += page[len(page)/2]
		}
	}

	return middleNumberSum
}

func partTwo(rules map[int][]int, pages [][]int) int {
	//filter out valid pages (only need to reorder invalid pages)
	invalidPages := make([][]int, len(pages))
	for i, pageLine := range pages {
		if !validatePage(rules, pageLine) {
			invalidPages[i] = pageLine
		}
	}
	//reorder invalid pages using topological sort
	reorderedPages := make([][]int, len(invalidPages))
	for i, pageLine := range invalidPages {
		reorderedPages[i] = utils.TopologicalSort(rules, pageLine)
	}

	//sum middle number of lines consisting of only reordered pages
	middleNumberSum := 0
	for _, page := range reorderedPages {
		if len(page) > 0 {
			middleNumberSum += page[len(page)/2]
		}
	}
	return middleNumberSum
}
