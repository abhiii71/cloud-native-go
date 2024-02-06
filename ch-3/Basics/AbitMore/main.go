package main

import (
	"fmt"
)

type error interface {
	Error() string
}

func main() {
	// file, err := os.Open("main.go")
	//	if err != nil {
	//		log.Fatal(err)
	//		return err
	//	}

	//	e1 := errors.New("error 71 ")
	//	e2 := fmt.Errorf("error %d", 71)

	//type NestedError struct {
	//	Message string
	//Err     error
	//}

	//func (e *NestedError) Error() string {
	//return fmt.Sprintf("%s\n contains: %s", e.Message, e.Err.Error())
	//}

	sum := add(10, 61)
	fmt.Println(sum)

	a, b := swap("foo", "bar")
	fmt.Println(a, b)

	defer fmt.Println("Executed after below line")
	fmt.Println("Executed before defer")

	defer fmt.Println("defer 1") // executed 3rd
	defer fmt.Println("defer 2") // executed 2nd
	defer fmt.Println("defer 3") // executed 1st
}

// FUNCTIONS
func add(x, y int) int {
	return x + y
}

// MULTIPLE RETURN
func swap(a, b string) (string, string) {
	return b, a
}
