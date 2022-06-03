package main

import "fmt"

func main() {
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Println("--------------------------------")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println("--------------------------------")

	number := 0
	for {
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&number)
		if number == 100 {
			break

		}
	}

	var j = 0

	for j < 10 {
		fmt.Printf("\n Valor de j: %d", j)
		if j == 5 {
			fmt.Printf(" multiplicamos por 2\n")
			j = j * 2
			continue
		}
		fmt.Printf("	Paso por aquí\n")
		j++
	}

	var k int = 0

ROUTINE:
	for k < 10 {
		if k == 4 {
			k = k + 2
			fmt.Println("voy a RUTINA sumando 2 a k")
			goto ROUTINE
		}
		fmt.Printf("Valor de i: %d\n", k)
		k++
	}
}
