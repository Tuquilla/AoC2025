package main

import (
	"AoC2025"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day6/input.txt")
	var calcArr [][]int
	var calcSign []string

	for _, line := range inputAsSlice {
		r := regexp.MustCompile(` `)
		lineArr := r.Split(line, -1)
		lineArr = slices.DeleteFunc(lineArr, func(n string) bool {
			return n == ""
		})

		var lineNumber []int
		for _, lineVar := range lineArr {

			number, err := strconv.Atoi(lineVar)
			if err != nil {
				calcSign = append(calcSign, lineVar)
			} else {
				lineNumber = append(lineNumber, number)
			}
		}
		if len(lineNumber) > 0 {
			calcArr = append(calcArr, lineNumber)
		}
	}

	res := 0
	for col, _ := range calcArr[0] {
		tmpRes := 0
		if calcSign[col] == "*" {
			tmpRes = 1
		}
		for row, _ := range calcArr {
			if calcSign[col] == "*" {
				tmpRes *= calcArr[row][col]
			} else {
				tmpRes += calcArr[row][col]
			}
		}
		res += tmpRes
	}
	fmt.Println("Sum of all calculations:", res)
}
