package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("vim-go")

	// variables
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	// manejo de errores
	myVar, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Println("error al transformar el valor 7 a int, ", err)
	}
	fmt.Println(myVar)

	// map
	m := make(map[string]int)
	m["key"] = 6

	fmt.Println(m["key"])

	// slice y como recorrerlos
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println("indice:", index, ", valor:", value)
	}
	fmt.Println(" ")
	// como agregar un nuevo elemento en el slice
	s = append(s, 8)
	for index, value := range s {
		fmt.Println("indice:", index, ", valor:", value)
	}

	// channel
	c := make(chan int)
	// goroutine
	go doSomething(c)
	<-c

	// referencias y apuntadores
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

// doSomething función de ejemplo, que recibe un channel
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
