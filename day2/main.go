package main

import (
	"AoC2025"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputAsSlice := AoC2025.ReadFile("./day2/input.txt")
	sumOfInvalidIDs := 0
	sumOfInvalidIDsSequences := 0

	for _, input := range inputAsSlice {
		idRanges := strings.Split(input, ",")
		for _, idRange := range idRanges {
			start, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
			end, _ := strconv.Atoi(strings.Split(idRange, "-")[1])
			for i := start; i <= end; i++ {
				// part 1
				variable := strconv.Itoa(i)
				varLengthHalf := len(variable) / 2
				if variable[0:varLengthHalf] == variable[varLengthHalf:] {
					sumOfInvalidIDs += i
				}

				// part 2
				maxLength := len(variable)
				// create all possible substrings for a value/string e.g. 1234 --> [1 2 3 4], [1 234], [12 34], [123 4]
				for index := 1; index < maxLength; index++ {
					tmpVariable := variable
					var subStringArray []string
					for len(tmpVariable) > 0 {
						if len(tmpVariable) < index {
							subStringArray = append(subStringArray, tmpVariable)
							break
						}
						subStringArray = append(subStringArray, tmpVariable[:index])
						tmpVariable = tmpVariable[index:]
					}

					// check if substrings of given number ([113 113 113]) are equal with Compact function (just one element should remain in array)
					subStringArray = slices.Compact(subStringArray)
					if len(subStringArray) == 1 {
						sumOfInvalidIDsSequences += i
						break
					}
				}

			}
		}
	}
	fmt.Println("Sum of all invalid IDs:", sumOfInvalidIDs)
	fmt.Println("Sum of all invlaid IDs with sequences:", sumOfInvalidIDsSequences)
}
