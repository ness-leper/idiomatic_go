package main

import (
	"fmt"
	"math/rand"
)

func main() {
  ex4_5()
}

func ex4_1(){
  x := 10
  if x > 5 {
    fmt.Println(x)
    x := 5
    fmt.Println(x)
  }
  fmt.Println(x)
}

func ex4_5(){
  n := rand.Intn(10)
  if n == 0 {
    fmt.Println("Too Low") 
  } else if n > 5 {
    fmt.Println("That's too big") 
  } else {
    fmt.Println("That's perfect")
  }
}
