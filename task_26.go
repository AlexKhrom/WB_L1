package main

import (
	"fmt"
	"strings"
)

func uniq(str string) bool {
	m := make(map[rune]int)

	runes := []rune(strings.ToLower(str))

	for i := range runes {
		_, ok := m[runes[i]]
		if !ok {
			m[runes[i]] = 1
		} else {
			m[runes[i]] = 2
		}
	}

	for _, val := range m {
		if val == 2 {
			return false
		}
	}

	return true
}

/*abcd — true
abCdefAaf — false
aabcd — false*/

func main() {

	var str = ""
	fmt.Scanf("%s", &str)

	fmt.Println(uniq(str))

}
