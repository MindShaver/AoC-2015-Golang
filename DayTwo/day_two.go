package main

import (
	"fmt"
	"../Utilities/LineReader"
	"../Utilities/MathHelpers"
	"strings"
	"strconv"
)

func main() {
	input := LineReader.LineByLine("DayTwo/input")
	fmt.Println("The solution to part one is ", solvePartOne(input))
	fmt.Println("The solution to part two is ", solvePartTwo(input))
}

func solvePartOne(input []string) int {
	runningTotal := 0
	for _, line := range input {
		runningTotal += findWrappingPaper(line)
	}

	return runningTotal
}

func solvePartTwo(input []string) int {
	runningTotal := 0
	for _, line := range input {
		runningTotal += findRibbonLength(line)
	}

	return runningTotal
}

func findWrappingPaper(input string) int {
	parsedLine := ParseInput(input)
	length, width, height := parsedLine[0], parsedLine[1], parsedLine[2]

	sides := getSurfaceAreas(length, width, height)
	smallestSide := MathHelpers.Min(getAreas(length, width, height))
	total := append(sides, smallestSide)

	return  MathHelpers.Sum(total)
}

func findRibbonLength(input string) int {
	parsedInput := ParseInput(input)
	length, width, height := parsedInput[0], parsedInput[1], parsedInput[2]

	smallestPerimeter := MathHelpers.Min(getPerimeters(length, width, height))
	volume := MathHelpers.Volume(length, width, height)

	return smallestPerimeter + volume
}

func ParseInput(input string) []int {
	parsedLine := strings.Split(input, "x")
	length, _ := strconv.Atoi(parsedLine[0])
	width, _ := strconv.Atoi(parsedLine[1])
	height, _ := strconv.Atoi(parsedLine[2])

	return []int {length, width, height}
}

func getAreas(length, width, height int) []int {
	areaOne := MathHelpers.Area(length, width)
	areaTwo := MathHelpers.Area(width, height)
	areaThree := MathHelpers.Area(length, height)

	return []int {areaOne, areaTwo, areaThree}
}

func getSurfaceAreas(length, width, height int) []int {
	sideOne := MathHelpers.SurfaceArea(length, width)
	sideTwo := MathHelpers.SurfaceArea(width, height)
	sideThree := MathHelpers.SurfaceArea(length, height)

	return []int {sideOne, sideTwo, sideThree}
}

func getPerimeters(length, width, height int) []int {
	perimOne := MathHelpers.Perimeter(length, width)
	perimTwo := MathHelpers.Perimeter(width, height)
	perimThree := MathHelpers.Perimeter(length, height)

	return []int {perimOne, perimTwo, perimThree}
}