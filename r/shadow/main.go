package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	zz := func() int {
		z := 0
		return z + 1
	}()
	fmt.Printf("%v %T\n", zz, zz)
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == 0 {
	} else {
		x = 123 + x
		fmt.Println(x)
		fmt.Printf("%g >= %g\n", v, lim)
	}
	fmt.Printf("x (f) = %v\n", x)
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
