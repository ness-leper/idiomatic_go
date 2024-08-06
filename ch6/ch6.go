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
  ex2()
  ex3()
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

func ex2(){
  src := []string{"Hello","World"}

  fmt.Println(src)
  UpdateSlice(src, "SallY!")
  fmt.Println(src)

  GrowSlice(src, "Johnny")
  fmt.Println(src)
}

func UpdateSlice(src []string, input string) {
  src[len(src)-1] = input
  fmt.Println(src)
}

func GrowSlice(src []string, input string) {
  src = append(src, input)
  fmt.Println(src)
}

func ex3(){
  var p []Person

  pRecord := Person {
    FirstName: "John",
    LastName: "Smith",
    Age: 10,
  }
  for x:=0;x<10000000;x++ {
    p = append(p, pRecord)
  }

  fmt.Println(len(p))
  fmt.Println(p[356543])
}
