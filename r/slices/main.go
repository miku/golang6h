package main

import "fmt"

func main() {
	fmt.Println("vim-go")
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0] // intuition: cap is also zero
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	s = s[:0] // panic: runtime error: slice bounds out of range
	printSlice(s)

	// Drop its first two values.
	s = s[2:] // destructive
	printSlice(s)

	s = s[:0]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
}
