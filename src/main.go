package main

import (
	"fmt"
	"time"
)

func DoSomeThing(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomeThing(d1, c1, 1)
	go DoSomeThing(d2, c2, 2)

	// fmt.Println(<-c1) // queda bloqueado, independientemente de quien llegue primero
	// fmt.Println(<-c2) // termina primero y queda a la espera que el primer print termina primero

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}
