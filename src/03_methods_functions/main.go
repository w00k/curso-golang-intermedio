package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)

	e.id = 1
	e.name = "Francisco"
	fmt.Printf("%v\n", e)

	e.SetId(8)
	e.SetName("w00k")
	fmt.Printf("%v\n", e)

	fmt.Println("\nvalores desde métodos get")
	fmt.Println("id:", e.GetId())
	fmt.Println("name:", e.GetName())
}
