package main

import (
	"AoC2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day7/input.txt")

	var grid [][]string
	for _, line := range inputAsSlice {
		lineArr := strings.Split(line, "")
		grid = append(grid, lineArr)
	}

	// part 1
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

	// part 2
	inputAsSlice = AoC2025.ReadFile("./day7/input.txt")

	for _, line := range inputAsSlice {
		lineArr := strings.Split(line, "")
		grid = append(grid, lineArr)
	}

	for row, _ := range grid {
		for col, colElement := range grid[row] {
			// skip first row
			if row == 0 {
				continue
			}
			// draw timelinenumber
			if colElement == "." {
				if grid[row-1][col] == "S" {
					grid[row][col] = "1"
				} else if grid[row-1][col] != "^" && grid[row-1][col] != "." {
					grid[row][col] = grid[row-1][col]
				}
			}

		}
		for col2, colElement2 := range grid[row] {
			if colElement2 == "^" && grid[row-1][col2] != "." {
				numAbove, _ := strconv.Atoi(grid[row-1][col2])
				if col2-1 >= 0 {
					numberLeft, _ := strconv.Atoi(grid[row][col2-1])
					grid[row][col2-1] = strconv.Itoa(numberLeft + numAbove)
				}
				if col2+1 < len(grid[row]) {
					number, err := strconv.Atoi(grid[row][col2+1])
					if err != nil {
						grid[row][col2+1] = grid[row-1][col2]
					} else {
						grid[row][col2+1] = strconv.Itoa(number + numAbove)
					}
				}
			}
		}
	}

	timelines := 0
	for _, element := range grid[len(grid)-1] {
		number, _ := strconv.Atoi(element)
		timelines += number
	}

	fmt.Println("Total timelines:", timelines)
}
