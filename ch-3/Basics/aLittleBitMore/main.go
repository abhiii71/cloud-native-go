package main

import "fmt"

func main() {
	//
	x := 5
	zeroByValue(x)
	fmt.Println(x)

	zeroByReference(&x)
	fmt.Println(x)
	//

	//
	m := map[string]int{"a": 0, "b": 1}
	fmt.Println(m)
	update(m)
	fmt.Println(m)

	// VARIADIC FUNCTION
	const name, age = "abhi", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	fmt.Println(product(2, 2, 2))

	//
	variadic := []int{3, 3, 3}
	fmt.Println(product(variadic...))

	//

	// ANNONYMOUS FUNCTION
	//
	var f func(int, int) int // func var have  types

	f = sum
	fmt.Println(f(69, 2)) // 71

	f = productt
	fmt.Println(f(3, 2))
	//
}

func zeroByValue(x int) {
	x = 0
}

func zeroByReference(x *int) {
	*x = 0
}

func update(m map[string]int) {
	m["c"] = 2
}

// VARIADIC FUNC
func product(factors ...int) int {
	p := 1

	for _, n := range factors {
		p *= n
	}
	return p
}

// ANNONYMOUS FUNCTION
func sum(x, y int) int      { return x + y }
func productt(x, y int) int { return x * y }
