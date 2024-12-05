package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LoadRaw(fname string) (string, error) {
	fPath := filepath.Join("data", fname)
	content, err := os.ReadFile(fPath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(content), nil
}

// func LoadRunes(rawString string) [][]rune {
// 	//load into 2d array of runes
// 	lines := strings.Split(rawString, "\n")
// 	rows := make([][]rune, len(lines))
// 	for i := range lines {
// 		rows[i] = []rune(lines[i])
// 	}
// 	return rows
// }

// LoadLines loads a file and returns a slice of strings, one for each line in the file.
func LoadLines(rawString string) []string {
	lines := strings.Split(rawString, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func LoadData(fname string) ([]string, error) {
	rawString, err := LoadRaw(fname)
	if err != nil {
		panic("err")
	}
	return LoadLines(rawString), nil
}

func LoadDataAsInts(fname string) ([][]int, error) {
	lines, _ := LoadData(fname)
	data := make([][]int, len(lines))
	for i, line := range lines {
		data[i] = make([]int, len(strings.Fields(line)))
		for j, field := range strings.Fields(line) {
			data[i][j], _ = strconv.Atoi(field)
		}
	}
	return data, nil
}

// assume all rows have the same number of columns
func MatrixLength(data [][]int) int {
	return len(data) * len(data[0])
}

// 1d to 2d co-ord
func CoOrd2D(index int, m int, n int) (int, int) {
	return index / n, index % n
}
