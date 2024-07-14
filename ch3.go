package main

import "fmt"

func main() {
	nilSlice()
  ex3_1()
  ex3_4()
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
