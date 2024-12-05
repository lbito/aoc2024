package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"slices"
)

type Position struct {
	x int
	y int
}

type Cross struct {
	upDown string
	downUp string
	anchor string
}

func main() {
	fmt.Println("Day 4")
	board, _ := utils.LoadData("4.txt")
	partOneSol := partOne(board)
	fmt.Println("Part One Solution:", partOneSol)
	partTwoSol := partTwo(board)
	fmt.Println("Part Two Solution:", partTwoSol)
}

func partOne(board []string) int {
	words := []string{}
	for i, row := range board {
		for j := range row {
			cellWords := getWords(Position{x: i, y: j}, board)
			words = append(words, cellWords[:]...)
		}
	}
	return countValidWords(words)
}

func partTwo(board []string) int {
	crosses := []Cross{}
	for i, row := range board {
		for j := range row {
			cross := genCross(i, j, board)
			crosses = append(crosses, cross)
		}
	}
	return countValidCrosses(crosses)
}

func countValidWords(words []string) int {
	count := 0
	for _, word := range words {
		if word == "XMAS" {
			count++
		}
	}
	return count
}

func countValidCrosses(crosses []Cross) int {
	count := 0
	for _, cross := range crosses {
		validStrings := []string{"MAS", "SAM"}
		if slices.Contains(validStrings, cross.upDown) && slices.Contains(validStrings, cross.downUp) {
			count++
		}
	}
	return count
}

func getWords(pos Position, board []string) [8]string {
	words := [8]string{}
	paths := genPaths()
	for i, path := range paths {
		word := ""
		for _, p := range path {
			x := pos.x + p.x
			y := pos.y + p.y
			if outOfBounds(x, y, board) {
				word += "."
				continue
			}
			word += string(board[x][y])
		}
		words[i] = word
	}
	return words
}

func genCross(x int, y int, board []string) Cross {
	upDown := ""
	downUp := ""
	anchor := string(board[x][y])
	for i := -1; i <= 1; i++ {
		if outOfBounds(x+i, y+i, board) {
			upDown += "."
		} else {
			upDown += string(board[x+i][y+i])
		}
		if outOfBounds(x+i, y-i, board) {
			downUp += "."
		} else {
			downUp += string(board[x+i][y-i])
		}
	}
	return Cross{upDown: upDown, downUp: downUp, anchor: anchor}
}

// generate all possible paths from any given position
func genPaths() [8][4]Position {
	baseDirections := [8]Position{
		{x: 1, y: 0},   // right
		{x: 0, y: 1},   // down
		{x: -1, y: 0},  // left
		{x: 0, y: -1},  // up
		{x: 1, y: 1},   // down-right
		{x: -1, y: 1},  // down-left
		{x: -1, y: -1}, // up-left
		{x: 1, y: -1},  // up-right
	}

	paths := [8][4]Position{}
	for i, dir := range baseDirections {
		paths[i] = genPath(dir)
	}
	return paths
}

// walk four spaces in given direction
func genPath(pos Position) [4]Position {
	path := [4]Position{}
	for i := 0; i <= 3; i++ {
		if i == 0 {
			path[0] = Position{x: 0, y: 0} //starting square
			continue
		}
		path[i] = Position{x: pos.x * i, y: pos.y * i}
	}
	return path
}

func outOfBounds(i int, j int, data []string) bool {
	return i < 0 || j < 0 || i >= len(data) || j >= len(data[i])
}
