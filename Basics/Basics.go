package main

import (
	"fmt"
)

/*
func main()  {
	fmt.Println("Hello world!")
}
*/

func add(x, y float32) float32 {
	return x + y
}

func printRange(x, count int) {
	for index := x; index < count; index++ {
		fmt.Println("-->", index)
	}
}

func main() {
	/*	num1, num2 := 4.1, 5.9
		//fmt.Println("Random number from 1 to 100 is : ", rand.Intn(100))
		fmt.Println("Addition of two Float numbers are ", add(num1, num2)) */

	/* Day 3
	var a int = 10
	var b float64 = float64(a)

	//x := a // x should be type Int

	fmt.Println("value of a is ", b) */
	/*
		/* Day 4 Pointers */
	/*a := 10
	b := &a
	fmt.Println(b)  // & memory address
	fmt.Println(*b) // value at that memory address

	*b = 5
	fmt.Println(a)  */
	a := 1
	b := 1001
	printRange(a, b)
}
