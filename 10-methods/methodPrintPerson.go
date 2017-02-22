package main

import "fmt"

// BEGIN OMIT

type Person struct {
	Age  int
	Name string
}

func (p Person) print() {
	fmt.Printf("%s is %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{Age: 15, Name: "Strans"}
	person.print()
}

// END OMIT
