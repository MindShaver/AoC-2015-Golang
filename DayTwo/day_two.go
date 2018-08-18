package main

import (
	"fmt"
	"../Utilities/LineReader"
	"../Utilities/MathHelper"
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
	parsedLine := parseInput(input)
	length, width, height := parsedLine[0], parsedLine[1], parsedLine[2]

	sides := getSurfaceAreas(length, width, height)
	smallestSide := MathHelper.Min(getAreas(length, width, height))
	total := append(sides, smallestSide)

	return  MathHelper.Sum(total)
}

func findRibbonLength(input string) int {
	parsedInput := parseInput(input)
	length, width, height := parsedInput[0], parsedInput[1], parsedInput[2]

	smallestPerimeter := MathHelper.Min(getPerimeters(length, width, height))
	volume := MathHelper.Volume(length, width, height)

	return smallestPerimeter + volume
}

func parseInput(input string) []int {
	parsedLine := strings.Split(input, "x")
	length, _ := strconv.Atoi(parsedLine[0])
	width, _ := strconv.Atoi(parsedLine[1])
	height, _ := strconv.Atoi(parsedLine[2])

	return []int {length, width, height}
}

func getAreas(length, width, height int) []int {
	areaOne := MathHelper.Area(length, width)
	areaTwo := MathHelper.Area(width, height)
	areaThree := MathHelper.Area(length, height)

	return []int {areaOne, areaTwo, areaThree}
}

func getSurfaceAreas(length, width, height int) []int {
	sideOne := MathHelper.SurfaceArea(length, width)
	sideTwo := MathHelper.SurfaceArea(width, height)
	sideThree := MathHelper.SurfaceArea(length, height)

	return []int {sideOne, sideTwo, sideThree}
}

func getPerimeters(length, width, height int) []int {
	perimOne := MathHelper.Perimeter(length, width)
	perimTwo := MathHelper.Perimeter(width, height)
	perimThree := MathHelper.Perimeter(length, height)

	return []int {perimOne, perimTwo, perimThree}
}