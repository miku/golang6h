package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// row[x] = uint8((x+y)/2) // x*y
			row[x] = uint8(x * y) // x*y
		}
		result[y] = row
	}
	return result
}

func main() {
	pic.Show(Pic)
}
