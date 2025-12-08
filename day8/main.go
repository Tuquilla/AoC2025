package main

import (
	"AoC2025"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type jbox struct {
	name string
	x    int
	y    int
	z    int
}

func main() {
	inputAsSlice := AoC2025.ReadFile("./day8/input.txt")
	var jboxes []jbox

	for _, input := range inputAsSlice {
		tmp := strings.Split(input, ",")
		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])
		z, _ := strconv.Atoi(tmp[2])
		jboxes = append(jboxes, jbox{input, x, y, z})
	}

	// part 1
	//maxLoops := 1000
	repeat := true
	loop := 0
	previousDist := 0
	var box1 *jbox
	var box2 *jbox
	var connections []map[string]*jbox
	// part1
	//for loop < maxLoops {
	for repeat {
		minDist := -1
		for outerIndex := 0; outerIndex < len(jboxes); outerIndex++ {
			for innerIndex := outerIndex + 1; innerIndex < len(jboxes); innerIndex++ {
				dist := calcDist(&jboxes[outerIndex], &jboxes[innerIndex])
				if minDist == -1 || (dist < minDist && dist > previousDist) {
					box1 = &jboxes[outerIndex]
					box2 = &jboxes[innerIndex]
					minDist = dist
				}
			}
		}
		previousDist = minDist

		exists1 := false
		exists2 := false
		index1 := -1
		index2 := -1
		for idx, element := range connections {
			if _, ok := element[box1.name]; ok {
				exists1 = true
				index1 = idx
			}

			if _, ok := element[box2.name]; ok {
				exists2 = true
				index2 = idx
			}
		}

		if exists1 == false && exists2 == false {
			connections = append(connections, map[string]*jbox{box1.name: box1, box2.name: box2})
		}
		if exists1 == true && exists2 == false {
			connections[index1][box2.name] = box2
		}
		if exists1 == false && exists2 == true {
			connections[index2][box1.name] = box1
		}
		if exists1 == true && exists2 == true {
			if index1 != index2 {
				for key, value := range connections[index2] {
					connections[index1][key] = value
				}
				connections = append(connections[:index2], connections[index2+1:]...)
			}
		}

		if len(connections) == 1 && len(connections[0]) == len(jboxes) {
			repeat = false
		}
		loop++
	}

	slices.SortFunc(connections, func(a, b map[string]*jbox) int {
		if len(a) > len(b) {
			return -1
		} else if len(a) < len(b) {
			return 1
		} else {
			return 0
		}
	})

	// part 1
	//multOfFirstThree := 1
	//for i := 0; i < 3; i++ {
	//	multOfFirstThree *= len(connections[i])
	//}

	//fmt.Println("Multiply of first three is: ", multOfFirstThree)

	fmt.Println("Last jboxes:")
	fmt.Println(box1.name)
	fmt.Println(box2.name)
	fmt.Println("Multiply x ccordinates: ", box1.x*box2.x)
}

func calcDist(jboxOne *jbox, jboxTwo *jbox) int {
	return ((jboxOne.x - jboxTwo.x) * (jboxOne.x - jboxTwo.x)) + ((jboxOne.y - jboxTwo.y) * (jboxOne.y - jboxTwo.y)) + ((jboxOne.z - jboxTwo.z) * (jboxOne.z - jboxTwo.z))
}
