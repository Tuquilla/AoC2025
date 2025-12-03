package main

import (
	"AoC2025"
	"fmt"
	"math/big"
	"strconv"
)

type battery struct {
	first  int64
	second int64
}

type secondBattery struct {
	first    int64
	second   int64
	third    int64
	fourth   int64
	fifth    int64
	sixth    int64
	seventh  int64
	eighth   int64
	ninth    int64
	tenth    int64
	eleventh int64
	twelve   int64
}

func main() {
	inputAsSlice := AoC2025.ReadFile("./day3/input.txt")
	var sum int64 = 0
	secondSum := big.NewInt(0)

	for _, number := range inputAsSlice {
		bat := battery{first: 0, second: 0}
		secondBat := secondBattery{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for index, character := range number {
			digit, _ := strconv.ParseInt(string(character), 10, 64)

			// part 1
			if digit > bat.first && index != len(number)-1 {
				bat.first = digit
				bat.second = 0
			} else if digit > bat.second {
				bat.second = digit
			}

			// part 2
			if digit > secondBat.first && index <= len(number)-12 {
				secondBat.first = digit
				secondBat.second = 0
				secondBat.third = 0
				secondBat.fourth = 0
				secondBat.fifth = 0
				secondBat.sixth = 0
				secondBat.seventh = 0
				secondBat.eighth = 0
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.second && index <= len(number)-11 {
				secondBat.second = digit
				secondBat.third = 0
				secondBat.fourth = 0
				secondBat.fifth = 0
				secondBat.sixth = 0
				secondBat.seventh = 0
				secondBat.eighth = 0
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.third && index <= len(number)-10 {
				secondBat.third = digit
				secondBat.fourth = 0
				secondBat.fifth = 0
				secondBat.sixth = 0
				secondBat.seventh = 0
				secondBat.eighth = 0
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.fourth && index <= len(number)-9 {
				secondBat.fourth = digit
				secondBat.fifth = 0
				secondBat.sixth = 0
				secondBat.seventh = 0
				secondBat.eighth = 0
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.fifth && index <= len(number)-8 {
				secondBat.fifth = digit
				secondBat.sixth = 0
				secondBat.seventh = 0
				secondBat.eighth = 0
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.sixth && index <= len(number)-7 {
				secondBat.sixth = digit
				secondBat.seventh = 0
				secondBat.eighth = 0
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.seventh && index <= len(number)-6 {
				secondBat.seventh = digit
				secondBat.eighth = 0
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.eighth && index <= len(number)-5 {
				secondBat.eighth = digit
				secondBat.ninth = 0
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.ninth && index <= len(number)-4 {
				secondBat.ninth = digit
				secondBat.tenth = 0
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.tenth && index <= len(number)-3 {
				secondBat.tenth = digit
				secondBat.eleventh = 0
				secondBat.twelve = 0
			} else if digit > secondBat.eleventh && index <= len(number)-2 {
				secondBat.eleventh = digit
				secondBat.twelve = 0
			} else if digit > secondBat.twelve && index <= len(number)-1 {
				secondBat.twelve = digit
			}
		}
		sum += bat.first*10 + bat.second
		twelveSum := big.NewInt(0)
		twelveSum.Add(twelveSum, big.NewInt(secondBat.twelve))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.eleventh).Mul(big.NewInt(secondBat.eleventh), big.NewInt(10)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.tenth).Mul(big.NewInt(secondBat.tenth), big.NewInt(100)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.ninth).Mul(big.NewInt(secondBat.ninth), big.NewInt(1000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.eighth).Mul(big.NewInt(secondBat.eighth), big.NewInt(10000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.seventh).Mul(big.NewInt(secondBat.seventh), big.NewInt(100000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.sixth).Mul(big.NewInt(secondBat.sixth), big.NewInt(1000000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.fifth).Mul(big.NewInt(secondBat.fifth), big.NewInt(10000000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.fourth).Mul(big.NewInt(secondBat.fourth), big.NewInt(100000000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.third).Mul(big.NewInt(secondBat.third), big.NewInt(1000000000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.second).Mul(big.NewInt(secondBat.second), big.NewInt(10000000000)))
		twelveSum.Add(twelveSum, big.NewInt(secondBat.first).Mul(big.NewInt(secondBat.first), big.NewInt(100000000000)))

		secondSum.Add(secondSum, twelveSum)
	}

	fmt.Println("Total output joltage:", sum)
	fmt.Println("Total output joltage twelve digits:", secondSum)
}
