package LineReader

import (
	"io/ioutil"
	"fmt"
)

func ReadLine(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}
