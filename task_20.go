package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() { // first second third fourth, snow dog sun
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	str := myscanner.Text()
	sl := strings.Split(str, " ")

	var reverSL []string
	for i := len(sl) - 1; i >= 0; i-- {
		reverSL = append(reverSL, sl[i])
	}

	fmt.Println(strings.Join(reverSL, " "))
}
