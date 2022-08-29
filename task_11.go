package main

import "fmt"

func crossSets(m1, m2 map[string]int) map[string]int {
	resMap := make(map[string]int)
	for key := range m1 {
		_, ok := m2[key]
		if ok {
			resMap[key] = 1
		}
	}
	return resMap
}

func main() {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	m1["rabbit"] = 1
	m1["horse"] = 1
	m1["cat"] = 1
	m1["dog"] = 1
	m1["elephant"] = 1

	m2["rabbit"] = 1
	m2["monkey"] = 1
	m2["cat"] = 1
	m2["goat"] = 1
	m2["elephant"] = 1

	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)
	fmt.Println("cross = ", crossSets(m1, m2))

}
