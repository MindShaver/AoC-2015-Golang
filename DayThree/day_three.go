package main

import (
	"../Utilities/LineReader"
	"fmt"
)

func main() {
	input := LineReader.ReadLine("DayThree/input")
	fmt.Println("The solution to part one is ", solvePartOne(input))
	fmt.Println("The solution to part two is ", solvePartTwo(input))
}

func solvePartOne(input string) int {
	x, y, houseCount := 0, 0, 1
	mapSet := map[Point] bool {Point{x, y}: true}

	for _, r := range input {
		switch r {
			case '^': {
				y++
				currentPoint := makePoint(x, y)
				currentPoint.evaluate(mapSet, &houseCount)
			}
			case '>': {
				x++
				currentPoint := makePoint(x, y)
				currentPoint.evaluate(mapSet, &houseCount)
			}
			case '<': {
				x--
				currentPoint := makePoint(x, y)
				currentPoint.evaluate(mapSet, &houseCount)
			}
			case 'v': {
				y--
				currentPoint := makePoint(x, y)
				currentPoint.evaluate(mapSet, &houseCount)
			}
		}
	}

	return houseCount
}

func solvePartTwo(input string) int {
	houseCount := 1
	santaX, santaY, realSanta := 0, 0, true
	robotX, robotY := 0, 0
	mapSet := map[Point] bool {makePoint(santaX, santaY): true}

	for _, r := range input {
		switch r {
			case '^': {
				if realSanta {
					realSanta = !realSanta
					santaY++
					currentPoint := makePoint(santaX, santaY)
					currentPoint.evaluate(mapSet, &houseCount)
				} else {
					realSanta = !realSanta
					robotY++
					currentPoint := makePoint(robotX, robotY)
					currentPoint.evaluate(mapSet, &houseCount)
				}
			}
			case '>': {
				if realSanta {
					realSanta = !realSanta
					santaX++
					currentPoint := makePoint(santaX, santaY)
					currentPoint.evaluate(mapSet, &houseCount)
				} else {
					realSanta = !realSanta
					robotX++
					currentPoint := makePoint(robotX, robotY)
					currentPoint.evaluate(mapSet, &houseCount)
				}
			}
			case '<': {
				if realSanta {
					realSanta = !realSanta
					santaX--
					currentPoint := makePoint(santaX, santaY)
					currentPoint.evaluate(mapSet, &houseCount)
				} else {
					realSanta = !realSanta
					robotX--
					currentPoint := makePoint(robotX, robotY)
					currentPoint.evaluate(mapSet, &houseCount)
				}
			}
			case 'v': {
				if realSanta {
					realSanta = !realSanta
					santaY--
					currentPoint := makePoint(santaX, santaY)
					currentPoint.evaluate(mapSet, &houseCount)
				} else {
					realSanta = !realSanta
					robotY--
					currentPoint := makePoint(robotX, robotY)
					currentPoint.evaluate(mapSet, &houseCount)
				}
			}
		}
	}

	return houseCount
}

type Point struct {
	x, y int
}

func makePoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) evaluate(mapSet map[Point] bool, houseCount *int) {
	if !mapSet[p] {
		mapSet[p] = true
		*houseCount++
	}
}