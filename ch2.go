package main

import (
	"fmt"
)

func main(){
  ex1()
  ex2()
}

func ex1(){
  i := 20
  var f float64
  f = float64(i)
  fmt.Println(i, f)
}

func ex2(){
  const value = 420
  var i int = value
  var f float64 = value
  fmt.Println(i, f)
}
