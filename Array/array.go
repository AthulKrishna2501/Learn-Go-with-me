package main

import "fmt"

func main() {
	var a [5]int   // An array with capacity 5
	fmt.Println(a) // By deafault an array is 0 valued

	a[4] = 130 // Set value to array a of index 4

	l := len(a) // Build in method to find the length of array
	c := cap(a) // Build in method to find the Capacity of array
	fmt.Println(l, c)
	// OR
	fmt.Println(len(a), cap(a))

	b := [5]int{1, 2, 3, 4, 5} // Declare and Initialize an array in one line
	fmt.Println(b)

	b = [...]int{1, 2, 3, 4, 5} // Spread operation. Compiler count the number of elements
	fmt.Println(b)

}
