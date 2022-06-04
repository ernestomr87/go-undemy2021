package main

import (
	"fmt"
	"log"
	"os"
)

func examplePanic() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Fatalf("generated a panic: %v\n", rec)
		}
	}()
	a := 1
	if a == 1 {
		panic("Find the value 1")
	}
}

func main() {
	file := "text.txt"
	f, err := os.Open(file)

	examplePanic()

	defer f.Close()

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

}
