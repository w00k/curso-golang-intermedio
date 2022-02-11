package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // instanciamos waitgroup, que funciona como un contador

	for i := 0; i < 10; i++ {
		wg.Add(1) // sumamos 1
		go doSomething(i, &wg)
	}
	wg.Wait() // espera hasta que el waitgroup sea 0, no es necesario bloquear un canal
}

//doSomething simula un proceso que se demora arto tiempoS
func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() // restamos 1
	fmt.Printf("Started, index: %d \n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("End")
}
