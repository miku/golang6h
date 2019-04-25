package main

import "fmt"

func main() {
	s := "Schöneberg"

	fmt.Printf("%s -- %d\n", s, len(s))
	fmt.Printf("%s -- %d\n", s, len([]rune(s)))
}
