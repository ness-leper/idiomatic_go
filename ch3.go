package main

import "fmt"

func main(){
  nilSlice()
}

func nilSlice(){
  var test[] int

  fmt.Println(test == nil)
}
