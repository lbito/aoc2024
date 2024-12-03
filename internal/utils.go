package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadRaw(fname string) (string, error) {
	fPath := filepath.Join("data", fname)
	fmt.Println("fPath", fPath)
	content, err := os.ReadFile(fPath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(content), nil
}

// LoadLines loads a file and returns a slice of strings, one for each line in the file.
func LoadLines(rawString string) []string {
	lines := strings.Split(rawString, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func LoadData(fname string) ([]string, error) {
	rawString, _ := LoadRaw(fname)
	return LoadLines(rawString), nil
}
