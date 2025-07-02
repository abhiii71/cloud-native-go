package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {

	// fmt.Println("containers: ")
	// containers()

	// fmt.Println("pointers: ")
	pointers()

	// closures
	var f func(int, int) int // Function variables have types
	f = sum
	fmt.Println(f(3, 5)) // "8"
	f = product          // Legal: product has same type as sum
	fmt.Println(f(3, 5))
	increment := incrementer()
	fmt.Println(increment()) // "1"
	fmt.Println(increment()) // "2"
	fmt.Println(increment()) // "3"
	newIncrement := incrementer()
	fmt.Println(newIncrement()) // "1"
}

// bool
func boolean() {
	// it is special  1-bit integer type that has two possible values

	and := true && false
	fmt.Println(and) // false

	and = true && true
	fmt.Println(and) // true

	//  Go doesn’t include a logical XOR operator
	// There is a ^ operator, but it’s reserved for bitwise XOR operations

	// two integer types have mnemonic aliases: byte, which is an alias for uint8
	// and rune, which is an alias for uint32.

}

func strings() {

	// Go supports two styles of string literals,

	// The interpreted form
	fmt.Println("Hello\nworld!\n")

	// The raw form
	fmt.Println(`Hello world!`)
}

func constants() {
	// Behind the scenes, a string is actually just a wrapper around a slice of UTF-8 encoded
	// Byte values, so any operation that can be applied to slices and arrays can also be applied to strings.

	// := is a declaration, and = is an assignment

	/*
		Some of the common verb flags used in format strings include:
		%v The value in a default format
		%T A representation of the type of the value
		%% A literal percent sign; consumes no value
		%t Boolean: the word true or false
		%b Integer: base 2
		%d Integer: base 10
		%f Floating point: decimal point but no exponent, e.g. 123.456
		%s String: the uninterpreted bytes of the string or slice
		%q String: a double-quoted string (safely escaped with Go syntax)
	*/

	// CONSTANTS
	/*
	 attempting to modify a constant will generate an error at compile time. Second,
	 constants must be assigned a value at declaration: they have no zero value.
	*/

	const language string = "Go"
	var favorite bool = true
	const text = "Does %s rule? %t!"
	var output = fmt.Sprintf(text, language, favorite)
	fmt.Println(output) // "Does Go rule? true!"
}

func containers() {
	// containers - arrays, slices and maps
	// array - fixed size lenght
	// slice - abstract around an array that can be resized at runtime
	// maps - in key-value pairs type

	// arrays indexed from 0 - n-1
	var a [3]int   // zero value array of type [3]int
	fmt.Println(a) // 3 size array with [0, 0, 0]

	// initialized using array  literals
	bx := [3]int{2, 4, 6}
	fmt.Println(bx)

	// compiler count the array elements here
	c := [...]int{1, 2, 3}
	fmt.Println(c)

	// slices are not more than abstract around traditional arrays
	// it consists 3 components
	// ptr to some element of a backing array that represents first element of slice
	// length, represents no. of elements in slice
	// capacity, represents upper value of the length
	ss := []int{1, 3, 34} // slice with 3 elements
	fmt.Println(ss)
	fmt.Println(len(ss)) // 3; len works for slice and arrays

	// slices extended with append
	ss = append(ss, 71)
	fmt.Println("updated ss : ", ss)
	// append to itself
	ss = append(ss, ss...)
	fmt.Println(ss) // you see updated s with [1,3,34, 71 ....] again same elements

	// string as slices
	// go allow you to cast strings into bytes or rune array
	s := "foö" // Unicode: f=0x66 o=0x6F ö=0xC3B6
	r := []rune(s)
	b := []byte(s)

	fmt.Printf("%7T %v\n", s, s) // string "foo"
	fmt.Printf("%7T %v\n", r, r) // []int32 [102 111 246]
	fmt.Printf("%7T %v\n", b, b) // []uint8 [102 111 195 182] "foo"

	// rune and bytes are type of mnemonic aliases for uint8 and int32

	// MAPS
	// key value format
	// fast lookup O(1) time complexity
	freezing := make(map[string]float32)
	freezing["celcius"] = 0.0
	freezing["fahrenheit"] = 32.0
	freezing["kelvin"] = 273.2

	// deleting
	delete(freezing, "celcius")

	fmt.Println(freezing)
	// if value of key is not present in a map it won't cause an exception to be thrown or return some kind of null value. Rather it returns zero value for the map's value type
	foo := freezing["no-such-key"] // Get non-existent key
	fmt.Println(foo)
}

func pointers() {

	// ptr store the address of the variable
	// ptr allow us to read and update the value of the variable

	// retrieving address of the variable
	aa := 71
	pp := &aa
	// ptr type 'points to' type of *int * indicates that it's ptr type that points to an int.

	// dereferencing a pointer
	// retrieve value of the value a from p, dereference using * before ptr variable name allowing update/read

	var a int = 71
	var p *int = &a
	fmt.Println(pp) // igives address of a

	*p = 69 // updating using ptr
	fmt.Println(a)

}

func controlStruct() {
	// there's only one for loop in the language
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i

	}
	fmt.Println(sum) // 10

	// Go's for statement's init and post statements are entire optional

	sum, i := 0, 0
	for i <= 0 {
		sum += i
		i++
	}
	fmt.Println(i, sum) // 10, 45

	// Looping over arrays and slices
	s := []int{2, 4, 8, 16, 32} // A slice of ints
	for i, v := range s {       // range gets each index/value
		fmt.Println(i, "->", v) // Output index and its value
	}

	// Looping over maps
	m := map[int]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
	}
	for k, v := range m {
		fmt.Println(k, "->", v)
	}

	// if Statement
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// an initialization statement to precede the condition clause in an if statement
	if _, err := os.Open("foo.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All in fine.")
	}
	// same as below
	_, err := os.Open("foo.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is fine.")
	}

	// switch Statement
	// there’s no fallthrough between the cases by default
	j := 0
	switch j % 3 {
	case 0:
		fmt.Println("zero")
		fallthrough
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Heh????")
	}

	// ways you can write a switch statements

	// 1
	hour := time.Now().Hour()
	switch {
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}

	// 2
	switch hour := time.Now().Hour(); { // Empty expression means "true"
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}
}

// Error Handling
func errrors() {
	// file, err := os.Open("somefile.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }

	// The actual implementation of the error type is actually incredibly simple: it’s just a
	// universally visible interface that declares a single method:
	type error interface {
		Error() string
	}

	// Creating an Error

	// there are two ways to create errors
	e1 := errors.New("error 71")
	e2 := fmt.Errorf("error %d", 71)

	fmt.Println("e1: ", e1, "e2: ", e2)
}

func functions() {
	//  Functions: Variadics and Closures

	// functions:
	// can have Multiple return values
	// Defer (keyword used to execute immediately before function return)
	// defer fmt.Println("cruel world")

}

// variadic functions
// It is one that may be called with zero or more trailing arguments.
func Printf(format string, a ...interface{}) (n int, err error) {

	// ... = ellipsis
	// This is the varia‐dic operator, which indicates that the function
	// may be called with any number of arguments of this type.
	const name, age = "abhi", 23
	fmt.Printf("%s is %d years old.\n", name, age)

	return 0, nil
}

func productt(factors ...int) int {
	p := 1
	for _, n := range factors {
		p *= n
	}
	return p
}
func vard() {
	// Passing slices as variadic values
	m := []int{3, 4, 5}
	fmt.Println(productt(m...))
}

// Anonymous Functions and Closures
func sum(x, y int) int     { return x + y }
func product(x, y int) int { return x * y }

// Functions may be created within other functions as anonymous functions
// A closure is a nested function that has access to the variables of its
// parent function, even after the parent has executed.
func incrementer() func() int {
	i := 0
	return func() int { // Return an anonymous function
		i++ // "Closes over" parent function's i
		return i
	}
}

// Structs, Methods, and Interfaces

