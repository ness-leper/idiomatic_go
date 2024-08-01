package main

import "fmt"

func main()  {
  fmt.Println(addToBase(3))
  fmt.Println(addToBase(3,2))
}

func addToBase(base int, vals ...int) []int {
  out := make([]int,0,len(vals))
  for _, v := range vals {
    out = append(out, base + v)
  }
  return out
}
