package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string // Uninitialized slice with legth 0
	fmt.Println(s)

	s = make([]string, 3) // Another way of declaring a slice.
	// 3 is the size

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])

	s = append(s, "d") // Adding new values to slice s
	s = append(s, "e", "f")
	fmt.Println("New values", s)

	c := make([]string, len(s)) // Creating another slice with length of slice s.
	copy(c, s)                  // Copying values in slice s to slice c
	fmt.Println(c)

	l := s[2:5] // Get the slice of elements s[2],s[3],[4] (excluding s[5]).
	fmt.Println(l)

	l = s[:5] // Get the slice of elements upto 5 (excluding 5).
	l = s[2:] // Get the slice of elements from s[2] (includin s[2]).

	t := []string{"a", "b", "c", "d"}
	fmt.Println(t)

	t2 := []string{"a", "b", "c", "d"}
	if slices.Equal(t, t2) { // Method to check the two slices are equal or not.
		fmt.Println("Slice t is same as Slice t2")
	}

}
