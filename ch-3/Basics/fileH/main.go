package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("/tmp/foo.txt") // creating an empty file
	defer closeFile(file)
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(file, "your homework is done")
	if err != nil {
		return
	}
	fmt.Println("File written successfully")
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println("Error closing file: ", err.Error())
	} else {
		fmt.Println("File closed successfully")
	}
}
