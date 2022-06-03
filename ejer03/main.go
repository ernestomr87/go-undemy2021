package main

import "fmt"

var status bool

func main() {

	if status = false; status == true {
		fmt.Println("the status is true")
	} else {
		fmt.Println("the status is false")
	}

	switch number := 7; number {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Menor que 1 o mayor 5")

	}
}
