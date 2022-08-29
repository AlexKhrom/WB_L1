package main

import (
	"fmt"
	"sort"
)

func newSlices(sl []float32) [][]float32 {
	res := [][]float32{}
	var ten int

	sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
	fmt.Println(sl)

	for i := 0; i < len(sl); i++ {
		newSl := []float32{}
		ten = int(sl[i] / 10)
		for i < len(sl) {
			if int(sl[i]/10) == ten {
				newSl = append(newSl, sl[i])
				i++
			} else {
				break
			}
		}
		res = append(res, newSl)
		if i <= len(sl)-1 {
			i--
		}
	}

	return res
}

func main() {
	sl := []float32{-25000000000.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 320000000000.5}
	sls := newSlices(sl)

	fmt.Println("===============")
	for _, it := range sls {
		fmt.Println(it)
	}

}
