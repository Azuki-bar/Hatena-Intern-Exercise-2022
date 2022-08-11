package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a1, a2 := 0, 1
	count := 0
	return func() int {
		defer func() { count++ }()
		switch count {
		case 0:
			return 0
		case 1:
			return 1
		default:
			fib := a1 + a2
			a1, a2 = a2, fib
			return fib
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
