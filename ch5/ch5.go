package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

type person struct {
  age int
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
}

func addToBase(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
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
