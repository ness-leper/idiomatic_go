package main

import (
	"fmt"
)

func main(){
  fmt.Println("Hello World")
  x := 10
  pointerToX := &x
  fmt.Println(pointerToX)
  fmt.Println(*pointerToX)
  z := 5 + *pointerToX
  fmt.Println(z)

  ex1()
}

type Person struct {
  FirstName string
  LastName string
  Age int
}

func ex1(){
  p := MakePerson("Fred", "Williamson", 25)
  fmt.Println(p)
  p2 := MakePersonPointer("Wilma", "Fredson", 32)
  fmt.Println(p2)
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
  return &Person {
    FirstName: firstName,
    LastName: lastName,
    Age: age,
  }
}

func MakePerson(firstName string, lastName string, age int) Person {
  return Person {
    FirstName: firstName,
    LastName: lastName,
    Age: age,
  }
}
