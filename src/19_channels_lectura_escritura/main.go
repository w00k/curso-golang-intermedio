package main

import (
	"fmt"
)

// Generator genera números de 0 a 10, con el '<-' a la derecha de 'chan' el channel solo para escribir
func Generator(c chan<- int) {
	for i := 0; i <= 10; i++ {
		c <- i
	}
	close(c) // bloquea el canal, deshabilitar el canal
}

// Double lee el channel y retorna en otro channel 'out' el resultado,
// el '<-' a la izquiera de 'chan' lo deja solo de lectura, con esto se mitiga que los channel sean usados para lectura y escritura
func Double(in <-chan int, out chan int) {
	for value := range in {
		out <- 2 * value
		//in <- 1 deadlock, porque 'in' solo es de lectura en esta clase
	}
	close(out)
}

func Print(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	// pipelines: funciones que llaman a más funciones en orden, existe dependencia de dichas funciones
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
