package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ex4_12()
}

func ex4_1() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func ex4_5() {
	if n := rand.Intn(10);n == 0 {
		fmt.Println("Too Low")
	} else if n > 5 {
		fmt.Println("That's too big")
	} else {
		fmt.Println("That's perfect")
	}
}

func ex4_8() {
  for i:=0; i<10; i++ {
    fmt.Println(i)
  }
}

func ex4_9() {
  i := 1
  for i < 100 {
    fmt.Println(i)
    i = i * 2
  }
}

func ex4_12() {
  for i:=1;i<100;i++{
    if i%3 == 0 && i%5==0 {
      fmt.Println("FizzBuzz")
      continue
    }
    if i%3 == 0 {
      fmt.Println("Fizz")
      continue
    }
    if i%5 == 0 {
      fmt.Println("Buzz")
      continue
    }
  }
}
