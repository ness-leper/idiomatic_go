package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type person struct {
	age  int
	name string
}

func main() {
	fmt.Println(addToBase(3))
	fmt.Println(addToBase(3, 2))

	result, remainder, err := divAndRemainder(5, 2)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}

	for i := 0; i < 5; i++ {
		f(i)
	}

	fmt.Println(addVals(1, 10, 15, 25))

	SortingSlices()
	funcAsValues()

	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	newS := []int{1, 2, 3}
	modSlice(newS)
	fmt.Println(newS)

	ex1()

  ex2()

  fileLength, err := fileLen("ch5.go")
  if err != nil {
    log.Fatal("File does not exist")
  }

  fmt.Println("The file has % bytes", fileLength)

  helloPrefix := prefixer("Hello")
  fmt.Println(helloPrefix("Nick"))
  fmt.Println(helloPrefix("Jack"))
}

func ex1() {
	var (
		add = func(i, j int) int { return i + j }
		sub = func(i, j int) int { return i - j }
		mul = func(i, j int) int { return i * j }
		div = func(i, j int) (int, error) {
			if j == 0 {
				return 0, errors.New("division by zero")
			}
			return (i / j), nil
		}
	)

	fmt.Println("adding", add(1, 2))
	fmt.Println("subtracting", sub(1, 2))
	fmt.Println("multiplying", mul(1, 2))
	dividing, err := div(1, 0)
	if err != nil {
		fmt.Println("Error: You tried dividing by zero", err)
	} else {
		fmt.Println("dividing", dividing)
	}
}

func addToBase(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func ex2() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func fileLen (i string) (int,error) {
  f, err := os.Open(i)
  if err != nil {
    return 0, err 
  }
  defer f.Close()

  data := make([]byte, 2048)
  total := 0
	for {
		count, err := f.Read(data)
    total += count
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

  return total, nil
}

/*
Add as many integers as you want to be tallied seperated by a comma
*/
func addVals(vals ...int) int {
	var out int
	for _, v := range vals {
		out = out + v
	}
	return out
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

func SortingSlices() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func funcAsValues() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func prefixer(i string) func(string) string {
  return func (v string) string {
    return i + " " + v
  }
}
