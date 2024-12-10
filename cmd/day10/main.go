package main

import (
	utils "aoc-24-lbit/internal"
	"fmt"
	"strconv"
)

type Book struct {
	width  int
	height int
	board  map[int][]Cell
	raw    []string
}

type Cell struct {
	x   int
	y   int
	val int
}

type Trail struct {
	origin  Cell
	current Cell
	path    []Cell
}

func (cell Cell) String() string {
	return fmt.Sprintf("(x:%d, y:%d, val:%d)", cell.x, cell.y, cell.val)
}

func main() {
	boardData, _ := utils.LoadData("10.txt")
	book := Book{
		width:  len(boardData),
		height: len(boardData[0]),
		board:  make(map[int][]Cell),
		raw:    boardData,
	}
	book.loadBoard()
	partOneSolution := book.partOne()
	fmt.Println("Part one solution is:", partOneSolution)

	partTwoSolution := book.partTwo()
	fmt.Println("Part two solution is:", partTwoSolution)
}

func (book *Book) loadBoard() {
	for i := 0; i < book.width; i++ {
		for j := 0; j < book.height; j++ {
			if book.raw[j][i] == '.' {
				continue
			}
			val, _ := strconv.Atoi(string(book.raw[j][i]))
			book.board[val] = append(book.board[val], Cell{x: i, y: j, val: val})
		}
	}
}

func (book *Book) partOne() int {
	trails := make([]Trail, 0)
	for _, cell := range book.board[0] {
		trails = append(trails, Trail{origin: cell, current: cell, path: []Cell{cell}})
	}

	for i := 0; i < 9; i++ {
		trails = book.nextTrail(trails)
	}
	uniquePaths := make(map[string]bool)
	for _, trail := range trails {
		//ignore paths, to only get unique start and ends
		key := fmt.Sprintf("origin:%s,current:%s", trail.origin.String(), trail.current.String())
		uniquePaths[key] = true
	}

	return len(uniquePaths)
}

func (book *Book) partTwo() int {
	trails := make([]Trail, 0)
	for _, cell := range book.board[0] {
		trails = append(trails, Trail{origin: cell, current: cell, path: []Cell{cell}})
	}

	for i := 0; i < 9; i++ {
		trails = book.nextTrail(trails)
	}
	return len(trails)
}

func (book *Book) nextTrail(trails []Trail) []Trail {
	nextTrails := make([]Trail, 0)
	for _, trail := range trails { //for each current trail
		for _, cell := range book.board[trail.current.val+1] { //each potential future trail
			pTrail := Trail{origin: trail.origin, current: cell, path: append(trail.path, cell)}
			if isAdjacent(pTrail.current, trail.current) {
				nextTrails = append(nextTrails, pTrail)
			}
		}
	}
	return nextTrails
}

func isAdjacent(p1 Cell, p2 Cell) bool {
	xDiff := p1.x - p2.x
	yDiff := p1.y - p2.y
	return (xDiff == 0 && (yDiff == 1 || yDiff == -1)) || (yDiff == 0 && (xDiff == 1 || xDiff == -1))
}
