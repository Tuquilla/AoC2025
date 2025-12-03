package main

import (
	"AoC2025"
	"fmt"
	"strconv"
)

type battery struct {
	first  int
	second int
}

func main() {
	inputAsSlice := AoC2025.ReadFile("./day3/input.txt")
	sum := 0

	for _, number := range inputAsSlice {
		bat := battery{first: 0, second: 0}
		for index, character := range number {
			digit, _ := strconv.Atoi(string(character))
			if digit > bat.first && index != len(number)-1 {
				bat.first = digit
				bat.second = 0
			} else if digit > bat.second {
				bat.second = digit
			}
		}
		sum += bat.first*10 + bat.second
	}

	fmt.Println("Total output joltage:", sum)
}
