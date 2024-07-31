package main

import (
	"fmt"
	"math/rand"
)

func main() {
	exercise3()
}

func exercise1() {
	stored := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		stored = append(stored, rand.Intn(100))
	}
	fmt.Println(stored)
}

func exercise2() {
	stored := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		stored = append(stored, rand.Intn(100))
	}

	for _, v := range stored {
		switch {
		case v%6 == 0:
			fmt.Println("six")
		case v%3 == 0:
			fmt.Println("three")
		case v%2 == 0:
			fmt.Println("two")
		default:
			fmt.Println("never mind")
		}
	}
}

func exercise3(){
  var total int

  // Question: What is the likely bug?
  // The bug is that we shadowed total - it's being set to the zero value over an dover so it is only showing the value of i
  for i := 0; i <=10; i++ {
    total := total + i
    fmt.Println(total)
  }
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
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Too Low")
	} else if n > 5 {
		fmt.Println("That's too big")
	} else {
		fmt.Println("That's perfect")
	}
}

func ex4_8() {
	for i := 0; i < 10; i++ {
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
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
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

func ex4_13() {
	evenVals := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func ex4_15() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func ex4_19() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "Is the exact right length", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
}

func ex4_21() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word")
		case wordLen > 10:
			fmt.Println(word, "is a long word")
		default:
			fmt.Println(word, "Is the perfect length!")
		}
	}
}

func ex4_22() {
	for i := 1; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "Fizz")
		case i%5 == 0:
			fmt.Println(i, "Buzz")
		default:
			fmt.Println(i)
		}
	}
}
