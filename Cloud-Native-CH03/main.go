package main

import "fmt"

func main() {

	// fmt.Println("containers: ")
	// containers()
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

}
