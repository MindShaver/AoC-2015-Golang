package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func main() {
	secretKey := "ckczppom"
	fmt.Println("The solution to part one is ", solvePartOne(secretKey))
	fmt.Println("The solution to part two is ", solvePartTwo(secretKey))
}

func solvePartOne(secretKey string) int {
	currentSequence := 0

	for true {
		secretBytes := []byte(secretKey + strconv.Itoa(currentSequence))
		hash := md5.Sum(secretBytes)
		hashString := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hashString, "00000") {
			return currentSequence
		}

		currentSequence++
	}

	return 0
}

func solvePartTwo(secretKey string) int {
	currentSequence := 0

	for true {
		secretBytes := []byte(secretKey + strconv.Itoa(currentSequence))
		hash := md5.Sum(secretBytes)
		hashString := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hashString, "000000") {
			return currentSequence
		}

		currentSequence++
	}

	return 0
}
