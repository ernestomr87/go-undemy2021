package main

import (
	"bufio"
	"fmt"
	"os"
)

var number1 int
var number2 int
var result int
var legend string

func main() {
	fmt.Println("Ingrese número1 : ")
	fmt.Scanln(&number1)

	fmt.Println("Ingrese número2 : ")
	fmt.Scanln(&number2)

	fmt.Println("Description : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		legend = scanner.Text()

	}

	result = number1 + number2

	fmt.Println(legend, result)

}
