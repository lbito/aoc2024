package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"time"

	"github.com/kr/pretty"
)

const (
	Up = iota
	Right
	Down
	Left
)

const (
	Nothing = iota
	Obstacle
	Guard
)

type MazeResult struct {
	steps    int
	cyclical bool
	visited  map[Position]bool
}

type Position struct {
	x, y      int
	direction int
}

func main() {
	fmt.Println("Day 6")
	// Read the input file
	lines, _ := utils.LoadData("6.txt")
	// Create a maze
	maze := newMaze(lines)
	startTime := time.Now()
	partOneSolution := partOne(maze)
	pretty.Println("Part 1: ", partOneSolution, "Time: ", time.Since(startTime).Seconds(), "s")
	startTime = time.Now()
	partTwoSolution := partTwo(maze)
	pretty.Println("Part 2: ", partTwoSolution, "Time: ", time.Since(startTime).Seconds(), "s")

}

func partOne(maze [][]int) int {
	result := walkMaze(maze)
	return result.steps
}

func walkMaze(maze [][]int) MazeResult {
	x, y := getGuardPosition(maze)
	direction := Up
	visitedSquares := make(map[Position]bool) //include direction as a visited square, with same direction = cyclical
	for imminentCollision(maze, x, y, direction) {
		direction = (direction + 1) % 4
	}
	visitedSquares[Position{x, y, direction}] = true // add the starting position

	for !escaped(maze, x, y) {
		for imminentCollision(maze, x, y, direction) {
			direction = (direction + 1) % 4
		}
		x, y = stepMaze(x, y, direction)
		if _, exists := visitedSquares[Position{x, y, direction}]; exists {
			//cyclical
			return MazeResult{
				steps:    len(visitedSquares),
				cyclical: true,
				visited:  visitedSquares,
			}
		}
		visitedSquares[Position{x, y, direction}] = true
	}

	uniqueSquares := make(map[[2]int]bool)
	for square := range visitedSquares {
		uniqueSquares[[2]int{square.x, square.y}] = true
	}

	return MazeResult{
		steps:    len(uniqueSquares),
		cyclical: false,
		visited:  visitedSquares,
	}

}

func partTwo(maze [][]int) int {
	//find how many possible positions a single obstacle can be placed in the maze such that the guard never escapes

	//walk maze first to find all possible obstacle placements (filter unreachable squares)
	solvedDefault := walkMaze(maze)
	potentialObstaclePlacements := make(map[[2]int]bool)
	for k, _ := range solvedDefault.visited {
		if maze[k.x][k.y] != Nothing {
			continue
		}
		potentialObstaclePlacements[[2]int{k.x, k.y}] = true
	}

	cyclesFound := 0

	for placement := range potentialObstaclePlacements {
		maze[placement[0]][placement[1]] = Obstacle
		result := walkMaze(maze)
		maze[placement[0]][placement[1]] = Nothing
		if result.cyclical {
			cyclesFound++
		}
	}
	return cyclesFound
}

func newMaze(lines []string) [][]int {
	rows := len(lines)
	cols := len(lines[0])

	maze := make([][]int, rows)
	for i := 0; i < rows; i++ {
		maze[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			switch lines[i][j] {
			case '#':
				maze[i][j] = Obstacle
			case '^':
				maze[i][j] = Guard
			default:
				maze[i][j] = Nothing
			}
		}
	}
	return maze
}

func imminentCollision(maze [][]int, x, y, direction int) bool {
	up := maze[x-1][y]
	right := maze[x][y+1]
	down := maze[x+1][y]
	left := maze[x][y-1]

	dirSquares := []int{up, right, down, left}
	return dirSquares[direction] == Obstacle
}

func getGuardPosition(Maze [][]int) (int, int) {
	for i, row := range Maze {
		for j, cell := range row {
			if cell == Guard {
				return i, j
			}
		}
	}
	panic("Guard not found")
}

// walk a step in a single direction
func stepMaze(x, y, direction int) (int, int) {
	switch direction {
	case Up:
		return x - 1, y
	case Right:
		return x, y + 1
	case Down:
		return x + 1, y
	case Left:
		return x, y - 1
	}
	return x, y
}

// True if Guard is at the edge of the maze
func escaped(Maze [][]int, x, y int) bool {
	edge := x == 0 || y == 0 || x == len(Maze)-1 || y == len(Maze[0])-1
	emptySquare := Maze[x][y] == Nothing
	return edge && emptySquare
}
