package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	sum := 0
	sum1 := 0
	sum2 := 1
	return func() int {
		sum = sum1 + sum2
		sum1 = sum2
		sum2 = sum
		return sum1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
