package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5

	// función anónima
	y := func() int {
		return x * 2
	}()

	fmt.Println("y:", y)

	// función concurrente y anónima
	c := make(chan int)
	go func() {
		fmt.Println("Start function")
		time.Sleep(5 * time.Second)
		fmt.Println("End function")
		c <- 1
	}()
	<-c
}
