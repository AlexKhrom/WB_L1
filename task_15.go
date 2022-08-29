package main

import "fmt"

var justString string

func createHugeString(strLen int) string {
	var str string
	for i := 0; i < strLen; i++ {
		str += "ÊÉÆÄ"
	}
	return str
}

func cutString(str string, start, end int) string {
	runes := []rune(str)
	var newStr = ""

	for i := range runes {
		if i >= start && i < end {
			newStr = newStr + string(runes[i])

		}
	}
	return newStr
}

func someFunc() {
	fmt.Println(1 << 10)
	v := createHugeString(1 << 10)
	//justString = v[:11]
	justString = cutString(v, 0, 10)
	fmt.Println(justString)
}

func main() {
	someFunc()
}
