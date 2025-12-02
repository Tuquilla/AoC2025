package main

import (
	"AoC2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day1/input.txt")
	position := 50
	zeroCounter := 0
	zeroCounterTotal := 0
	zeroMultiplier := 0

	for _, input := range inputAsSlice {
		lineAsSlice := strings.SplitN(input, "", 2)
		zeroMultiplier = 0
		direction := lineAsSlice[0]
		clicks, _ := strconv.Atoi(lineAsSlice[1])
		if clicks > 100 {
			zeroMultiplier = clicks / 100
			clicks = clicks % 100
		}

		if direction == "L" {
			tmp := position - clicks
			if tmp < 0 {
				if position != 0 {
					zeroMultiplier++
				}
				position = 100 + tmp

			} else {
				position = tmp
			}
		} else {
			tmp := position + clicks
			if tmp > 100 {
				position = 0 + (tmp - 100)
				zeroMultiplier++
			} else {
				position = tmp
				if position == 100 {
					position = 0
				}
			}
		}
		if position == 0 {
			zeroCounter++
			zeroMultiplier++
		}
		zeroCounterTotal += zeroMultiplier
		fmt.Println(zeroCounterTotal)
	}
	fmt.Println("Dial points at 0:", zeroCounter)
	fmt.Println("Dial points at 0 at rotation and end", zeroCounterTotal)
}
