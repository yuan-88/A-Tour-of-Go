package main

import (
	"fmt"
	"math"
)

// Return the result if Calculate 10 times or the result is almost no change
func Sqrt(x float64) float64 {
	var z float64 = 1.0
	var d float64

	for n := 0; n < 10; n++ {
		d = (z*z - x) / (2*z)
		z -= d
		if (d*d < 1e-20) {
			break
		}
	}
	return z
}

func main() {
	fmt.Printf("My answer: %v\n",Sqrt(2))
	fmt.Printf("math.Sqrt: %v", math.Sqrt(2))
}