package main

import (
	"fmt"
	"math"
)

func main() {
	// this will cause an error, because pi (lowercase) is not exported
	// fmt.Println(math.pi)

	fmt.Println(math.Pi)
}
