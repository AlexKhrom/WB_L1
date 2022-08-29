package main

import "fmt"

func main() {
	var str = ""
	fmt.Scanf("%s", &str)

	runes := []rune(str)
	var newStr = ""

	for i := range runes {
		newStr = string(runes[i]) + newStr
	}
	
	fmt.Println("res = ", newStr)

}

// ÊÉÆÄ
