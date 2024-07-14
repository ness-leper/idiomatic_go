package main

import (
	"fmt"
)

func main() {
  var x int = 10
  var b byte = 100
  var sum1 int = x + int(b)
  var sum2 byte = byte(x) + b
  fmt.Println(sum1, sum2)
}
