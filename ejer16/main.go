package main

import (
	"fmt"
	"time"
)

func loop(channel1 chan time.Duration) {
	start := time.Now()

	for i := 0; i < 10000000000; i++ {

	}

	end := time.Now().Sub(start)
	channel1 <- end
}

func main() {
	channel1 := make(chan time.Duration)
	go loop(channel1)
	fmt.Println("Stop here...")

	msg := <-channel1

	fmt.Println(msg)

}
