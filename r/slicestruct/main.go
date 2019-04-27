package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
	vals []int
}

func (v Vertex) Abs() float64 {
	v.vals = append(v.vals, 1)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{X: 3, Y: 4, vals: []int{}}
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(v)
}
