package main

import (
	"AoC2025"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day5/input.txt")
	var ingredientId []int
	var validIdRange [][]int
	emptyLine := false
	sumOfValidIds := 0

	for _, line := range inputAsSlice {
		if line == "" {
			emptyLine = true
			continue
		}

		if emptyLine == false {
			lineArr := strings.Split(line, "-")
			minId, _ := strconv.Atoi(lineArr[0])
			maxId, _ := strconv.Atoi(lineArr[1])
			validIdRange = append(validIdRange, []int{minId, maxId})
		}
		if emptyLine == true {
			lineAsInt, _ := strconv.Atoi(line)
			ingredientId = append(ingredientId, lineAsInt)
		}
	}

	// part 1
	for _, id := range ingredientId {
		for _, idRange := range validIdRange {
			if id >= idRange[0] && id <= idRange[1] {
				sumOfValidIds++
				break
			}
		}
	}

	// part 2
	slices.SortFunc(validIdRange, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})

	tmpValue := 0
	sumOfValidRanges := 0
	for _, idRange := range validIdRange {
		if idRange[0] < tmpValue {
			idRange[0] = tmpValue
		}

		validIds := idRange[1] - idRange[0] + 1
		if validIds < 0 {
			continue
		}
		tmpValue = idRange[1] + 1
		sumOfValidRanges += validIds
	}

	fmt.Println("Total Sum of valid Ids:", sumOfValidIds)
	fmt.Println("Sum of Valid Id ranges:", sumOfValidRanges)
}
