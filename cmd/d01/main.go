package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")
	lines, _ := utils.LoadData("1.txt")
	rows := len(lines)

	groupOne := make([]int, rows)
	groupTwo := make([]int, rows)

	for row := 0; row < rows; row++ {
		cols := strings.Fields(lines[row])
		groupOne[row], _ = strconv.Atoi(cols[0])
		groupTwo[row], _ = strconv.Atoi(cols[1])
	}
	fmt.Printf("Part 1: %d\n", partOne(groupOne, groupTwo))
	fmt.Printf("Part 2: %d\n", partTwo(groupOne, groupTwo))
}

func partOne(groupOne []int, groupTwo []int) int {
	//sort groupOne and groupTwo
	sort.Ints(groupOne)
	sort.Ints(groupTwo)
	sum := 0
	for i := 0; i < len(groupOne); i++ {
		sum += int(math.Abs(float64(groupOne[i] - groupTwo[i])))
	}
	return sum
}

func partTwo(groupOne []int, groupTwo []int) int {
	//sort groupOne and groupTwo
	sort.Ints(groupOne)
	sort.Ints(groupTwo)

	similarityScore := make(map[int]int)
	sum := 0
	for i := 0; i < len(groupOne); i++ {
		//if similarityScore is already in the map, continue
		if similarityScore[groupOne[i]] != 0 {
			sum += similarityScore[groupOne[i]]
			continue
		}
		count := 0
		for j := 0; j < len(groupTwo); j++ {
			if groupTwo[j] == groupOne[i] {
				count++
			}
		}
		similarityScore[groupOne[i]] = count * groupOne[i]
		sum += similarityScore[groupOne[i]]
	}
	return sum
}
