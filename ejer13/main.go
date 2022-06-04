package main

import "fmt"

func exponential(num int) {
	if num > 100000000 {
		return
	}
	fmt.Println(num)
	exponential(num * 2)
}

func main() {
	exponential(2)
}
