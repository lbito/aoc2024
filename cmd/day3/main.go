package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"regexp"
	"strconv"
)

type Operation struct {
	op string
	a  int
	b  int
}

func main() {
	fmt.Println("Day 3")
	rawString, _ := utils.LoadRaw("3.txt")
	opRegEx := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	opNameRegEx := regexp.MustCompile(`^([\w']+)\(`)
	mulRegEx := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := opRegEx.FindAllString(rawString, -1)
	ops := make([]Operation, len(matches))

	for i, match := range matches {
		opname := opNameRegEx.FindStringSubmatch(match)
		if opname[1] == "mul" {
			mulArgs := mulRegEx.FindStringSubmatch(match)
			a, _ := strconv.Atoi(mulArgs[1])
			b, _ := strconv.Atoi(mulArgs[2])
			ops[i] = Operation{op: opname[1], a: a, b: b}
		} else {
			ops[i] = Operation{op: opname[1]}
		}
	}
	partOneSolution := partOne(ops)
	fmt.Println("Part One Solution:", partOneSolution)
	partTwoSolution := partTwo(ops)
	fmt.Println("Part Two Solution:", partTwoSolution)
}

func partOne(ops []Operation) int {
	sum := 0
	for _, op := range ops {
		if op.op == "mul" {
			sum += (op.a * op.b)
		}
	}
	return sum
}

func partTwo(ops []Operation) int {
	mul := true
	sum := 0
	for _, op := range ops {
		switch op.op {
		case "do":
			mul = true
		case "don't":
			mul = false
		case "mul":
			if mul {
				sum += (op.a * op.b)
			}
		}
	}
	return sum
}
