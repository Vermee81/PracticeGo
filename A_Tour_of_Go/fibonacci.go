package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibo1 := 0
	fibo2 := 0
	fibo3 := 0
	return func() int{
		switch fibo1 {
		case 0:
			fibo1 = 1
			fibo3 = fibo1
		default:
			fibo1 = fibo1 + fibo2
			fibo2 = fibo3
			fibo3 = fibo1
		}
		return fibo1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}