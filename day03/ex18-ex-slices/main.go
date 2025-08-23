package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// outer slice of length dy
	picture := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		// inner slice of length dx
		row := make([]uint8, dx)

		for x := 0; x < dx; x++ {
			// choose a formula for pixel value
			// options: (x + y) / 2, x * y, x ^ y, etc.
			row[x] = uint8(x ^ y)
		}

		picture[y] = row
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
