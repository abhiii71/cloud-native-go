package main

import (
	"fmt"
	"os"
	"time"
)

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
	fmt.Printf("%7T %v\n", b, b)

	// Maps
	fmt.Println("MAPS:")
	mapp := make(map[string]int)
	mapp["abhi"] = 71
	mapp["kd"] = 69
	fmt.Println(mapp)
	delete(mapp, "kd")
	fmt.Println(mapp)

	// initialize and populated as map literals
	mapps := map[string]int{
		"abhi":  71,
		"kd":    22,
		"rajat": 69,
	}
	fmt.Println(mapps)

	///maps membership testing
	foo := mapps["sezal"]
	fmt.Println(foo)
	newton, ok := mapps["sezal"]
	fmt.Println(ok)
	fmt.Println(newton)

	// Pointers
	fmt.Println("POINTERS: ")

	var ar int = 71
	var p *int = &ar
	fmt.Println(p)
	fmt.Println("Value of p: ", *p)
	*p = 69
	fmt.Println("Value of ar: ", ar)

	// One loop only -> for Loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	summ, i := 0, 0
	for i < 10 { // Equivalent to: for; i < 10;
		summ += i
		i++
	}
	fmt.Println(i, summ) // 10 45

	// range over for loop
	abhii := []int{2, 4, 8, 16, 32}
	for val, ind := range abhii {
		fmt.Println(val, "->", ind)
	}

	abhii2 := []int{2, 4, 6, 8, 10}
	for _, vall := range abhii2 { // ommiting the index
		fmt.Println(vall)
	}

	aa := []int{0, 2, 4, 6, 8}
	summm := 0
	for _, v := range aa {
		summm += v
	}
	fmt.Println(summm)

	// loop over maps
	mmap := map[int]string{
		1: "jan",
		2: "feb",
		3: "march",
	}
	for k, v := range mmap {
		fmt.Println(k, "->", v)
	}

	// IF ELSE

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if _, err := os.Open("main.go"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is well")
	}

	_, er := os.Open("main.go")
	if er != nil {
		fmt.Println("err")
	} else {
		fmt.Println("All is well")
	}

	// SWITCH
	varr := 0
	switch varr % 3 {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("huh")
	}

	hour := time.Now().Hour()

	switch {
	case hour >= 5 && hour < 9:
		fmt.Println("I am working")
	case hour >= 9 && hour < 18:
		fmt.Println("I am writing")
	default:
		fmt.Println("I am sleeping")
	}

	// same but different approach
	switch hour := time.Now().Hour(); {
	case hour >= 5 && hour < 9:
		fmt.Println("I am writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I am working")
	default:
		fmt.Println("I am sleeping bruhh ")

	}
}
