package main

import "fmt"

func calculate(a int, b int) []float64 {
	var sum float64 = (float64) (a + b)
	var difference float64 = (float64) (a - b)
	var product float64 = (float64) (a * b)
	var quotient float64 = (float64) (a / b)
	results := []float64{sum, difference, product, quotient}
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}