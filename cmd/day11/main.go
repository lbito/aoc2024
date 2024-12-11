package main

import (
	utils "aoc-24-lbit/internal"
	"math"
	"strconv"
	"strings"
)

type Stones struct {
	memoized map[int]int
	original []int
}

func splitNum(num int) (int, int) {
	digits := int(math.Log10(float64(num))) + 1
	left := num / int(math.Pow(10, float64(digits/2)))
	right := num % int(math.Pow(10, float64(digits/2)))
	return left, right
}

func numberOfDigitsIsEven(num int) bool {
	digits := int(math.Log10(float64(num))) + 1
	return digits%2 == 0
}

func (stones Stones) CountStones() int {
	total := 0
	for _, v := range stones.memoized {
		total += v
	}
	return total
}

func main() {
	lines, _ := utils.LoadData("11.txt")
	data := make([]int, 0)
	inputs := strings.Split(lines[0], " ")
	for _, numStr := range inputs {
		num, _ := strconv.Atoi(numStr)
		data = append(data, num)
	}

	stones := Stones{
		memoized: make(map[int]int),
		original: data,
	}

	for _, num := range data {
		stones.memoized[num]++
	}

	partOneSolution := stones.partOne()
	println("Part one solution is:", partOneSolution)

	stones.reset()

	partTwoSolution := stones.partTwo()
	println("Part two solution is:", partTwoSolution)

}

func (stones *Stones) reset() {
	stones.memoized = make(map[int]int)
	for _, num := range stones.original {
		stones.memoized[num]++
	}
}

func (stones *Stones) partOne() int {
	iters := 25
	for i := 0; i < iters; i++ {
		stones.blinkEach()
	}
	return stones.CountStones()
}

func (stones *Stones) partTwo() int {
	iters := 75
	for i := 0; i < iters; i++ {
		stones.blinkEach()
	}
	return stones.CountStones()
}

func (stones *Stones) blinkEach() {
	newCount := make(map[int]int)
	for num, count := range stones.memoized {
		newStones := stones.blink(num)
		for _, newStone := range newStones {
			newCount[newStone] += count
		}
	}
	stones.memoized = newCount
}

func (stones Stones) blink(num int) []int {
	result := make([]int, 0)
	if num == 0 { //rule 1
		result = append(result, 1)
	} else if numberOfDigitsIsEven(num) { //rule 2
		leftNum, rightNum := splitNum(num)
		result = append(result, leftNum)
		result = append(result, rightNum)
	} else {
		result = append(result, num*2024) //rule 3
	}
	return result
}
