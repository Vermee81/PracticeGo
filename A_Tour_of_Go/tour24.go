package main

import (
	"fmt"
	"math"
)

var zBefore = float64(1)

func Sqrt(x float64) float64 {
	var z float64
	diff := float64(100)
	for diff > 0.0000001 {
		z = zBefore - (zBefore*zBefore-x)/(2*zBefore)
		diff = math.Abs(z - zBefore)
		zBefore = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
