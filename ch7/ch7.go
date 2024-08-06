package main

import "fmt"

func main() {
	fmt.Println("Hello Chapter 7!")
	p := Person{
		FirstName: "John",
		LastName:  "Smith",
		Age:       25,
	}

  fmt.Println(p.String())
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
