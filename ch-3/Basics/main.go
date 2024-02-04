package main

import "fmt"

// constant
const language string = "Go"

var favorite bool = true

func main() {
	// Boolean
	and := true && false
	fmt.Println(and)
	or := true || false
	fmt.Println(or)
	not := !true
	fmt.Println(not)

	// complex num
	var x complex64 = 3.71i
	fmt.Println(x)

	// variables
	//  var a, b, c = 22, 69, 71
	// fmt.Println(a,b,c)

	// SHORT HAND DECLARATION
	abhi := 71
	fmt.Println("By short hand declaration: ", abhi)

	// zero value
	var a int // just declare it
	fmt.Printf("The zero value of a is: %d \n", a)

	// flags
	// %d -> int, %f -> float , %t -> bool, %q -> string
	//
	// Blank identifer can be written as - _
	str := "word"
	_, err := fmt.Printf("Hello %s\n", str)
	if err != nil {
		panic(err)
	}

	// constant

	const text = "Does %s rule? %t!"
	output := fmt.Sprintf(text, language, favorite)

	fmt.Println(output) // "Does Go rule? true"

	///Container types - Arrays, Slices & Maps
	//
	//Array
	var arr [3]int      // zero value array of type [3] int
	fmt.Println(arr)    // [0,0,0]
	fmt.Println(arr[1]) // 0

	// update the index
	arr[1] = 71
	fmt.Println(arr)

	i := arr[1]
	fmt.Println(i)

	// initialize using array literals
	arr2 := [3]int{71, 69, 22}
	arr3 := [...]int{71, 69, 22} // both are same
	fmt.Println("array2: ", arr2, "& array3: ", arr3)
	fmt.Println("Array length of arr3: ", len(arr3))
	fmt.Println(arr3[len(arr3)-1])

	// Slices

	sllc := make([]int, 3)
	sllc[0] = 71
	fmt.Println("Slices: ", sllc)

	slc := []int{3} // slice literals
	fmt.Println(slc)

	// appending Slices
	slc = append(slc, 71, 69)
	fmt.Println(slc)
	slc = append(slc, slc...) // appending slices to iteslf
	fmt.Println(slc)

	// slice operator
	s0 := []int{1, 2, 3, 4, 5, 6}
	s1 := s0[:3]
	s2 := s0[3:]
	s0[3] = 71
	fmt.Println(s0)
	fmt.Println("S1 from index0-index3", s1)
	fmt.Println("S2 from index3-indexLast", s2)

	// strings as slice
	s := "foö" // unicode: f=0x66 o=0x6F ö=0xC3B6
	r := []rune(s)
	b := []byte(s)
	fmt.Printf("%7T %v\n", s, s) // string foo
	fmt.Printf("%7T %v\n", r, r) // []int32 [102 111 246]
	fmt.Printf("%7T %v\n", b, b) // []uint8 [102 111 195 182]
}
