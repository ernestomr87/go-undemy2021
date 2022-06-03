package main

import (
	"fmt"
)

func main() {
	fmt.Println(one(5))

	number, status := pair(35)
	fmt.Println(number, status)

	fmt.Println(calculate(1, 2, 3, 4, 5, 6))
	fmt.Println(calculate(1, 3, 3))
	fmt.Println(calculate(13, 24, 33, 45, 54, 62))
}

func one(number int) (z int) {
	z = number * 2
	return
}

func pair(number int) (int, bool) {
	if number%2 == 0 {
		return number % 2, true
	} else {
		return number % 2, false
	}
}

func calculate(numbers ...int) int {
	total := 0

	for i, number := range numbers {
		total += number
		fmt.Printf("Item %d, value %d\n", i, number)
	}

	return total
}
