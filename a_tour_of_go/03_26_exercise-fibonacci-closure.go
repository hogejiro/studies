package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	ret := 0
	adder := 0
	counter := 0 // use 1st & 2nd iteration only
	return func() int {
		if counter < 2 {
			ret = counter
			counter++
			return ret
		}
		var tmp = ret
		ret = ret + adder // f(n) = f(n-1) + f(n-2)
		adder = tmp
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
