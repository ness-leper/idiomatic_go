package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ex4_21()
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
  for _,word := range words {
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
