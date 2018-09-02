package main

import (
	"fmt"
)

func main() {

	// for loop
	index := 0
	count := 100
	for index <= count {
		fmt.Println("Value of Index --> ", index)
		index += 10
	}

	// do while
	index = 0

	for {
		fmt.Println("Value of Index --> ", index)
		index += 10

		if index > 100 {
			break
		}
	}
}
