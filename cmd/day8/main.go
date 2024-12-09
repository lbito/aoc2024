package main

import (
	utils "aoc-24-lbit/internal"

	"github.com/kr/pretty"
)

type City struct {
	width   int
	height  int
	signals map[rune][]Cell
}

type Cell struct {
	x, y   int
	signal rune
}

func main() {
	lines, _ := utils.LoadData("8.txt")
	signals := make(map[rune][]Cell)
	for i, line := range lines {
		for j, cell := range line {
			if cell != '.' {
				signals[cell] = append(signals[cell], Cell{i, j, cell})
			}
		}
	}

	city := City{len(lines), len(lines[0]), signals}
	partOneSol := city.partOne()
	pretty.Println(partOneSol)

	partTwoSol := city.partTwo()
	pretty.Println(partTwoSol)

}

func (city City) partOne() int {
	signals := city.signals
	antinodes := make(map[Cell]bool)
	for _, cells := range signals {
		pairs := getSignalPairs(cells)
		for _, pair := range pairs {
			currentAntiNodes := city.getAntiNodes(pair)
			for _, antipode := range currentAntiNodes {
				antinodes[antipode] = true
			}
		}
	}
	return len(antinodes)
}

func (city City) partTwo() int {
	antinodes := make(map[Cell]bool)
	for _, cells := range city.signals {
		if len(cells) < 2 {
			continue
		}
		pairs := getSignalPairs(cells)
		for _, pair := range pairs {
			currentAntiNodes := city.resonantAntiNodes(pair)
			for _, antinode := range currentAntiNodes {
				antinodes[antinode] = true
			}
		}
	}
	return len(antinodes)
}

func (city City) IsInBounds(x, y int) bool {
	return x >= 0 && x < city.width && y >= 0 && y < city.height
}

// NcR: combinations of len(Cells) choose 2
func getSignalPairs(cells []Cell) [][2]Cell {
	var pairs [][2]Cell
	for i := 0; i < len(cells); i++ {
		for j := i + 1; j < len(cells); j++ {
			pairs = append(pairs, [2]Cell{cells[i], cells[j]})
		}
	}
	return pairs
}

func (city City) getAntiNodes(cells [2]Cell) []Cell {

	deltaX := cells[1].x - cells[0].x
	deltaY := cells[1].y - cells[0].y

	potentialAntiNodes := []Cell{
		{cells[0].x - deltaX, cells[0].y - deltaY, '#'},
		{cells[1].x + deltaX, cells[1].y + deltaY, '#'},
	}

	var antiNodes []Cell
	for _, antiNode := range potentialAntiNodes {
		if city.IsInBounds(antiNode.x, antiNode.y) {
			antiNodes = append(antiNodes, antiNode)
		}
	}
	return antiNodes
}

func (city City) resonantAntiNodes(cells [2]Cell) []Cell {
	antiNodes := []Cell{}

	deltaX := utils.AbsVal(cells[1].x - cells[0].x)
	deltaY := utils.AbsVal(cells[1].y - cells[0].y)

	directions := [][2]int{
		{-deltaX, -deltaY},
		{deltaX, deltaY},
	}

	for _, dir := range directions {
		x, y := cells[0].x, cells[0].y
		for city.IsInBounds(x, y) {
			antiNodes = append(antiNodes, Cell{x, y, '#'})
			x += dir[0]
			y += dir[1]
		}
	}

	return antiNodes
}
