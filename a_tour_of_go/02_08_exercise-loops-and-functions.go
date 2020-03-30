package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	loop := 10
	for i := 0; i < loop; i++ {
		z -= ((z*z - x) / (2 * z))
	}
	return z
}

func main() {
	fmt.Println("(1) my Sqrt (2) math.Sqrt")
	for i := 2.0; i < 10; i++ {
		fmt.Println("number:", i)
		fmt.Println("(1)", Sqrt(i))
		fmt.Println("(2)", math.Sqrt(i))
	}
}
