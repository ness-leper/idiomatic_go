package main

import "fmt"

func main() {
	exercise1()
}

func printSpaces() {
	fmt.Println("\n")
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

func ex3_4() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
	fmt.Println("d", d)
	fmt.Println("e", e)
}

func ex3_5() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func ex3_6() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, "z")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func copySlice() {
	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x)
	fmt.Println(d)
}

func ex3_10() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittins"])

	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])

	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	fmt.Println(m, len(m))
	clear(m)
	fmt.Println(m, len(m))
}

func structExperiment() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	fmt.Println(person)
}

func exercise1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Приві"}
	fmt.Println(greetings)

	a := greetings[:2]
	b := greetings[1:4]
	c := greetings[3:]

	fmt.Println(a, b, c)
}
