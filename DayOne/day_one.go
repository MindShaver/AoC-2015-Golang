package main

import (
	"fmt"
	"../Utilities/LineReader"
	)

func main() {
	input := LineReader.ReadLine("DayOne/input")
	fmt.Println("The solution to part one is ", solvePartOne(input))
	fmt.Println("The solution to part two is ", solvePartTwo(input))
}

func solvePartOne(input string) int {
	floor := 0
	for _, r := range input {
		direction := string(r)
		if direction == "(" {
			floor++
		} else {
			floor--
		}
	}

	return floor
}

func solvePartTwo(input string) int {
	floor, position := 0, 1
	for _, r := range input {
		direction := string(r)
		if direction == "(" {
			floor++
			position++
		} else {
			floor--
			if floor == -1 {
				return position
			}
			position++;
		}
	}

	return position
}
