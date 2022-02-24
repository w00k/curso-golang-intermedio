package main

import (
	"fmt"
)

// Poder de la concurrencia que posee Go/Golang o también conocido como EORKER POOLS

// Fibonacci genera la secuencia de fibonacci de forma recursiva
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Worker que a la escucha de los channels, calcula la serie y el resultado lo almacena en otro channel
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fic with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func main() {

	// def de tareas
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks)) // crear el channel jobs de tipo int, longitud de canal que es la cantidad de tareas que hay (tipo semáforo)
	results := make(chan int, len(tasks))

	// se crean los workers
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	// lee el valor del slice, lo agrega al channel jobs y se procesa la serie de fibonacci
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs) // cierro el channel

	// imprime los resultados
	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
