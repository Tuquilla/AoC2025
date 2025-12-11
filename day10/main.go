package main

import (
	"AoC2025"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day10/input.txt")
	var lights [][]string
	var buttons [][][]int

	for _, input := range inputAsSlice {
		var buttonLine [][]int
		variables := strings.Split(input, " ")
		light, _ := strings.CutPrefix(variables[0], "[")
		light, _ = strings.CutSuffix(light, "]")

		var lightArr []string
		for _, element := range strings.Split(light, "") {
			lightArr = append(lightArr, element)
		}
		lights = append(lights, lightArr)

		for i := 1; i < len(variables)-1; i++ {
			button, _ := strings.CutPrefix(variables[i], "(")
			button, _ = strings.CutSuffix(button, ")")
			buttonArr := strings.Split(button, ",")

			buttonArrInt := make([]int, 0)
			for _, buttonElement := range buttonArr {
				intValue, _ := strconv.Atoi(buttonElement)
				buttonArrInt = append(buttonArrInt, intValue)
			}
			buttonLine = append(buttonLine, buttonArrInt)
		}
		buttons = append(buttons, buttonLine)
	}

	resultat := 0
	for index := range buttons {
		fmt.Println("Round", index+1)
		state := make([]string, len(lights[index]))
		for i := range state {
			state[i] = "."
		}
		resultat += recursiveSearch(0, &buttons[index], &state, &lights[index])
	}

	fmt.Println("Sum of min Presses:", resultat)
}

func recursiveSearch(press int, buttons *[][]int, state *[]string, lights *[]string) int {
	if compareLights(state, lights) {
		return press
	}

	if press >= 10 { // Maximalanzahl an Zügen
		return 10 // Signalwert: zu viele Schritte
	}

	minPress := 10
	for i := 0; i < len(*buttons); i++ {
		updateState(&(*buttons)[i], state) // Button drücken
		tmp := recursiveSearch(press+1, buttons, state, lights)
		if tmp < minPress {
			minPress = tmp
		}
		updateState(&(*buttons)[i], state) // Button zurücksetzen
	}

	return minPress
}

func updateState(button *[]int, state *[]string) {
	for _, idx := range *button {
		if (*state)[idx] == "." {
			(*state)[idx] = "#"
		} else {
			(*state)[idx] = "."
		}
	}
}

func compareLights(state, lights *[]string) bool {
	for i, v := range *state {
		if v != (*lights)[i] {
			return false
		}
	}
	return true
}
