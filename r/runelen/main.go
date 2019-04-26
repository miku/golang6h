package main

import "fmt"

func main() {
	s := "Schöneberg"

	fmt.Printf("%s -- %d\n", s, len(s))         // 11
	fmt.Printf("%s -- %d\n", s, len([]rune(s))) // 10
}
