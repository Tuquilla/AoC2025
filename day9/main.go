package main

import (
	"AoC2025"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type Area struct {
	x1   int
	y1   int
	x2   int
	y2   int
	area int
}

func main() {
	inputAsSlice := AoC2025.ReadFile("./day9/input.txt")
	var coordArr []coord

	for _, input := range inputAsSlice {
		coords := strings.Split(input, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		coordArr = append(coordArr, coord{x: x, y: y})
	}

	// part 1
	areas := make([]Area, 0, len(coordArr)*len(coordArr))
	for outerIndex := 0; outerIndex < len(coordArr); outerIndex++ {
		for innerIndex := outerIndex + 1; innerIndex < len(coordArr); innerIndex++ {
			area := calcArea(&coordArr[outerIndex], &coordArr[innerIndex])
			areas = append(areas, Area{x1: coordArr[outerIndex].x, y1: coordArr[outerIndex].y, x2: coordArr[innerIndex].x, y2: coordArr[innerIndex].y, area: area})
		}
	}

	slices.SortFunc(areas, func(a, b Area) int {
		if a.area > b.area {
			return -1
		}
		if a.area < b.area {
			return 1
		} else {
			return 0
		}
	})

	fmt.Println("Max area of rectangle:", areas[0].area)

	// part 2
	fmt.Println(len(areas))

}

func calcArea(coord1, coord2 *coord) int {
	var x1 float64 = float64(coord1.x)
	var y1 float64 = float64(coord1.y)
	var x2 float64 = float64(coord2.x)
	var y2 float64 = float64(coord2.y)
	var areaF float64 = (math.Abs(x1-x2) + 1) * (math.Abs(y1-y2) + 1)
	return int(areaF)
}

//func checkGreenAreas(area1, area2 *Area) int {

//}

func within(outerArea, innerArea *Area) bool {
	xBig := 0
	xSmall := 0
	yBig := 0
	ySmall := 0

	if outerArea.x1 > outerArea.x2 {
		xBig = outerArea.x1
		xSmall = outerArea.x2
	} else {
		xSmall = outerArea.x1
		xBig = outerArea.x2
	}
	if outerArea.y1 > outerArea.y2 {
		yBig = outerArea.y1
		ySmall = outerArea.y2
	} else {
		ySmall = outerArea.y1
		yBig = outerArea.y2
	}

	if innerArea.x1 <= xBig &&
		innerArea.x2 <= xBig &&
		innerArea.x1 >= xSmall &&
		innerArea.x2 >= xSmall &&
		innerArea.y1 <= yBig &&
		innerArea.y2 <= yBig &&
		innerArea.y1 >= ySmall &&
		innerArea.y2 >= ySmall {
		return true
	}
	return false
}
