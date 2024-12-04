package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	rawString, _ := LoadRaw("3.txt")
	opRegEx := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	argRegEx := regexp.MustCompile(`\((\d+),(\d+)\)`)
	matches := opRegEx.FindAllString(rawString, -1)
	ops := make([]Operation, len(matches))

	for i, match := range matches {
		argMatches := argRegEx.FindStringSubmatch(match)
		a, _ := strconv.Atoi(argMatches[1])
		b, _ := strconv.Atoi(argMatches[2])
		ops[i] = Operation{op: "mul", a: a, b: b}
	}
	// fmt.Print(ops)
	partOneSolution := partOne(ops)
	fmt.Println("Part One Solution:", partOneSolution)
}

func LoadRaw(fname string) (string, error) {
	fPath := filepath.Join("data", fname)
	content, err := os.ReadFile(fPath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(content), nil
}

func partOne(ops []Operation) int {
	sum := 0
	for _, op := range ops {
		sum += (op.a * op.b)
	}
	return sum
}
