package main

import (
	"fmt"
	"strings"
	"time"
)

func mySlowName(name string) {
	words := strings.Split(name, "")
	for _, w := range words {
		time.Sleep(1 * time.Second)
		fmt.Println(w)
	}
}

func main() {
	go mySlowName("Ernesto Rodriguez")
	fmt.Println("I'm here!!!")
	var x string
	fmt.Scanln(&x)
}
