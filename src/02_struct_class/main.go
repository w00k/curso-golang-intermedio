package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 0
	e.name = "Francisco"
	fmt.Println(e)
}
