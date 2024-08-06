package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Chapter 7!")
	p := Person{
		FirstName: "John",
		LastName:  "Smith",
		Age:       25,
	}

	fmt.Println(p.String())

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

  doUpateWrong(c)
  // doUpdateRight(&c)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
  c.Increment()
  fmt.Println("in doUpdateRight:", c.String())
}
