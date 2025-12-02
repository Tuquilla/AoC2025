package main

import (
	"AoC2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day2/input.txt")
	sumOfInvalidIDs := 0

	for _, input := range inputAsSlice {
		idRanges := strings.Split(input, ",")
		for _, idRange := range idRanges {
			start, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
			end, _ := strconv.Atoi(strings.Split(idRange, "-")[1])
			for i := start; i <= end; i++ {
				variable := strconv.Itoa(i)
				varLengthHalf := len(variable) / 2
				if variable[0:varLengthHalf] == variable[varLengthHalf:] {
					sumOfInvalidIDs += i
				}
			}
		}
	}
	fmt.Println("Sum of all invalid IDs:", sumOfInvalidIDs)
}
