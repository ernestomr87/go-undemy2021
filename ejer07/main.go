package main

import "fmt"

var Calculate func(int, int) int = func(i1, i2 int) int {
	return i1 + i2
}

func Operations() {
	result := func() int {
		a := 23
		b := 13
		return a + b
	}
	fmt.Println(result())
}

func Table(value int) func() int {
	number := value
	sequence := 0

	return func() int {
		sequence++
		return number * sequence
	}
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d\n", Calculate(5, 7))

	// Restamos
	Calculate = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Printf("Resta 6 - 4 = %d\n", Calculate(6, 4))

	// Division
	Calculate = func(i1, i2 int) int {
		return i1 / i2
	}
	fmt.Printf("Dividir 6 / 2 = %d\n", Calculate(6, 2))

	Operations()

	//Closures
	tableDel := 2
	MyTable := Table(tableDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MyTable())
	}

}
