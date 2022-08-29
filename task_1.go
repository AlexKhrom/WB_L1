package main

import "fmt"

type Human struct {
	value int
}

func (h *Human) sayHi() {
	fmt.Println("Hi!")
}

func (h *Human) sayAnotherThing() {
	fmt.Println("say another thing ")
}

type Action struct {
	str string
	Human
}

func main() {
	var act Action
	act.sayHi()
	act.sayAnotherThing()
}
