package main

import (
	"fmt"
	"strconv"
)

var number int
var text string
var status bool = true

func main() {
	number2, number3, number4, number5, status := 2, 3, 5, "Hola Mundo", false

	var currency float64 = 3.14

	number2 = int(currency)
	text = strconv.Itoa(int(currency))

	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)
	fmt.Println(number5)
	fmt.Println(status)

	ShowStatus()

}

func ShowStatus() {
	fmt.Println(status)
}
