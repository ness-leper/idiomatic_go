package main

import (
	"fmt"
)

func main(){
  ex1()
  ex2()
  ex3()
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

func ex3(){
  var b byte = 255
  var smallI int32 = 2147483647
  var bigI uint64 = 18446744073709551615

  b = b + 1
  smallI = smallI + 1
  bigI = bigI + 1
  fmt.Println(b, smallI, bigI)
}
