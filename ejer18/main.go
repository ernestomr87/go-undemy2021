package main

import "fmt"

var result int

func main() {
	fmt.Println("Starting...")

	result = operationsMiddleware(sum)(2, 3)
	fmt.Println(result)

	result = operationsMiddleware(rest)(10, 6)
	fmt.Println(result)

	result = operationsMiddleware(mult)(2, 4)
	fmt.Println(result)
}

func sum(a, b int) int {
	return a + b
}

func rest(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func operationsMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Starting Operation...")
		return f(a, b)
	}
}
