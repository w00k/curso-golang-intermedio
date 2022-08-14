package main

import (
	"fmt"
	"sync"
	"time"
)

// c := [][]
// c := [goRoutine1][]
// c := [goRoutine1][goRoutine2]
// c := [][goRoutine2]
// c := [goRoutine3][goRoutine2]
// c := [goRoutine3][]
func main() {
	c := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait() // bloqueo del programa
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)
	<-c
}
