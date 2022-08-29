package main

import "fmt"

func typeOfVar(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	default:
		fmt.Println("unknown")
	}
	// Вывод: float64: 2.3
}

func main() {
	var boolean bool = true
	var float float64 = 1.41
	var str string = "some str"
	var integer int = 32

	typeOfVar(boolean)
	typeOfVar(float)
	typeOfVar(str)
	typeOfVar(integer)

}
