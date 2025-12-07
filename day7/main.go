package main

import (
	"AoC2025"
	"fmt"
	"strings"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day7/input.txt")

	var grid [][]string
	for _, line := range inputAsSlice {
		lineArr := strings.Split(line, "")
		grid = append(grid, lineArr)
	}

	splits := 0
	for row, _ := range grid {
		for col, colElement := range grid[row] {
			// skip first row
			if row == 0 {
				continue
			}
			// draw beam
			if colElement == "." && (grid[row-1][col] == "S" || grid[row-1][col] == "|") {
				grid[row][col] = "|"
			}
			if colElement == "^" && grid[row-1][col] == "|" {
				splits += 1
				if col-1 >= 0 {
					grid[row][col-1] = "|"
				}
				if col+1 < len(grid[row]) {
					grid[row][col+1] = "|"
				}
			}
		}
	}

	fmt.Println("Total splits:", splits)
}
