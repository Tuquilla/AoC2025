package main

import (
	"AoC2025"
	"fmt"
	"strings"
)

type coord struct {
	x int
	y int
}

func main() {

	var inputMatrix [][]string
	var validCoords []coord

	inputAsSlice := AoC2025.ReadFile("./day4/input.txt")

	for _, line := range inputAsSlice {
		lineAsArr := strings.Split(line, "")
		inputMatrix = append(inputMatrix, lineAsArr)
	}

	for {
		anythingChanged := false
		for horizontal, _ := range inputMatrix {
			for vertical, _ := range inputMatrix[horizontal] {
				roleCounter := 0

				// skip if position is not a role
				if inputMatrix[horizontal][vertical] != "@" {
					continue
				}

				// check in all eight directions
				// l
				if vertical-1 >= 0 && inputMatrix[horizontal][vertical-1] == "@" {
					roleCounter++
				}
				// ul
				if horizontal-1 >= 0 && vertical-1 >= 0 && inputMatrix[horizontal-1][vertical-1] == "@" {
					roleCounter++
				}
				// u
				if horizontal-1 >= 0 && inputMatrix[horizontal-1][vertical] == "@" {
					roleCounter++
				}
				// ur
				if horizontal-1 >= 0 && vertical+1 < len(inputMatrix[horizontal]) && inputMatrix[horizontal-1][vertical+1] == "@" {
					roleCounter++
				}
				// r
				if vertical+1 < len(inputMatrix[horizontal]) && inputMatrix[horizontal][vertical+1] == "@" {
					roleCounter++
				}
				// lr
				if horizontal+1 < len(inputMatrix) && vertical+1 < len(inputMatrix[horizontal]) && inputMatrix[horizontal+1][vertical+1] == "@" {
					roleCounter++
				}
				// l
				if horizontal+1 < len(inputMatrix) && inputMatrix[horizontal+1][vertical] == "@" {
					roleCounter++
				}
				// ll
				if vertical-1 >= 0 && horizontal+1 < len(inputMatrix) && inputMatrix[horizontal+1][vertical-1] == "@" {
					roleCounter++
				}

				if roleCounter < 4 {
					coordinate := coord{x: horizontal, y: vertical}
					validCoords = append(validCoords, coordinate)
					// remove role in matrix
					inputMatrix[horizontal][vertical] = "."
					anythingChanged = true
				}
			}
		}
		if anythingChanged == false {
			break
		}
	}

	fmt.Println("coords with max of four roles and removal of role: ", len(validCoords))
}
