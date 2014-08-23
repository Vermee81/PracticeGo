package main

import (
	"fmt"
	"math"
)

//Sqrt returns a square root of a parameter.
// ex. Sqrt(8) will be 2
func Sqrt(x float64) float64 {
	z := float64(1)
	zBefore := float64(1.1)
	diff := float64(100)
	for diff > 0.000001 {
		zBefore = z
		z = z - (z*z-x)/(2*z)
		diff = math.Abs(z - zBefore)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
