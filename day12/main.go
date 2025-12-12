package main

import (
	"AoC2025"
	"fmt"
	"strconv"
	"strings"
)

type Shape struct {
	area int
}

type Area struct {
	area   int
	shapes []int
}

func main() {
	inputAsSlice := AoC2025.ReadFile("./day12/input.txt")
	shapesAsSlice := AoC2025.ReadFile("./day12/shape.txt")
	var shapes []Shape
	var areas []Area

	area := 0
	for _, element := range shapesAsSlice {
		if element == "" {
			shapes = append(shapes, Shape{area: area})
			area = 0
		}
		row := strings.Split(element, "")
		count := 0
		for _, rowElement := range row {
			if rowElement == "#" {
				count++
			}
		}
		area += count
	}
	fmt.Println(shapes)

	for _, element := range inputAsSlice {
		firstSplit := strings.Split(element, ":")
		are := strings.Split(firstSplit[0], "x")
		x, _ := strconv.Atoi(are[0])
		y, _ := strconv.Atoi(are[1])
		areaObj := Area{area: x * y}

		secondSplit := strings.Split(firstSplit[1], " ")
		var arrInt []int
		for _, ele := range secondSplit {
			val, _ := strconv.Atoi(ele)
			arrInt = append(arrInt, val)
		}
		areaObj.shapes = arrInt
		areas = append(areas, areaObj)
	}
	fmt.Println(areas)

	goodAreas := 0
	for _, element := range areas {
		tmp := 0
		for _, shape := range element.shapes {
			tmp += shape * 7
		}
		if tmp <= element.area {
			goodAreas++
		}
	}

	fmt.Println("Good Areas:", goodAreas)
}
