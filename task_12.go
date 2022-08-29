package main

import (
	"fmt"
)

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, it := range sl {
		set[it] = struct{}{}
	}

	fmt.Println(set)

}
