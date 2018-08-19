package main

import (
	"../Utilities/LineReader"
	"fmt"
)

func main() {
	input := LineReader.LineByLine("DayFive/input")
	fmt.Println("The solution to part one is ", solvePartOne(input))
}

func solvePartOne(input []string) int {
	niceCounter := 0
	for _, line := range input {
		if containsDuplicate(line) && containsThreeVowels(line) && hasNoForbidden(line) {
			niceCounter++
		}
	}

	return niceCounter
}

func containsThreeVowels(input string) bool {
	vowels := map[rune] bool {'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	counter := 0

	for _, c := range input {
		if vowels[c] {
			counter++
		}
	}

	return counter >= 3
}

func hasNoForbidden(input string) bool {
	forbidden := map[string] bool {"ab": true, "cd": true, "pq": true, "xy": true}

	for i := 0; i < len(input) - 2; i++ {
		subStr := input[i: i+2]
		if forbidden[subStr]{
			return false
		}
	}

	return true
}

func containsDuplicate(input string) bool {
	for i := 0; i < len(input) - 1; i++ {
		if input[i] == input[i + 1] {
			return true;
		}
	}

	return false
}