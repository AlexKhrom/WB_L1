package main

import "fmt"

type DuckInterface interface { // этот интерфейс , хотим чтобы dog его реализовывал
	fly()
}

type Duck struct {
}

func (d *Duck) fly() {
	fmt.Println("fly")
}

type Dog struct { // пытаемся привести к интерфейсу Duck с помощью DogAdapter
}

func (d *Dog) makeFly() {
	fmt.Println("buy a plane ticket for a dog")
}

type DogAdapter struct {
	dog *Dog
}

func (adapter DogAdapter) fly() {
	adapter.dog.makeFly()
	fmt.Println("fly")
}

func makeFly(duckInterface DuckInterface) {
	duckInterface.fly()
}

func main() {
	var dog = Dog{}
	var dogAdapter DogAdapter
	dogAdapter.dog = &dog

	makeFly(dogAdapter)
}
