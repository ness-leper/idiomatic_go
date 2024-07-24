package main

import "fmt"

func main() {
	nilSlice()
  printSpaces()
  ex3_1()
  printSpaces()
  ex3_4()
  printSpaces()
  ex3_5()
  printSpaces()
  ex3_6()
  printSpaces()
}

func printSpaces(){
  fmt.Println("\n\n\n")
}

func nilSlice() {
	var test []int

	fmt.Println(test == nil)
}

func ex3_1() {
	var x []int
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

}

func ex3_4(){
  x:=[]string{"a","b","c","d"}
  y:=x[:2]
  z:=x[1:]
  d:=x[1:3]
  e:=x[:]
  fmt.Println("x",x)
  fmt.Println("y",y)
  fmt.Println("z",z)
  fmt.Println("d",d)
  fmt.Println("e",e)
}

func ex3_5(){
  x:=[]string{"a","b","c","d"}
  y:=x[:2]
  z:=x[1:]
  x[1] = "y"
  y[0] = "x"
  z[1] = "z"
  fmt.Println("x:",x)
  fmt.Println("y:",y)
  fmt.Println("z:",z)
}

func ex3_6(){
  x:=[]string{"a","b","c","d"}
  y:=x[:2]
  fmt.Println(cap(x),cap(y))
  y = append(y, "z")
  fmt.Println("x:", x)
  fmt.Println("y:",y)
}
