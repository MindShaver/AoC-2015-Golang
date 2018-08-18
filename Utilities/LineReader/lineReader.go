package LineReader

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

func ReadLine(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}

func LineByLine(fileName string) []string {
	file, _ := os.Open(fileName)
	fscanner := bufio.NewScanner(file)
	var lines []string
	for fscanner.Scan() {
		lines = append(lines, fscanner.Text())
	}

	return lines
}