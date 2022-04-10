package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

// en Go/Golang no existe la herencia como tal, sino el principio que debe existir composición sobre la herencia
// utilizando un campo anónimo en un struct puede "heredar" todas las propiedades y métodos
func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Francisco"
	ftEmployee.age = 39
	ftEmployee.id = 1
	fmt.Printf("%v", ftEmployee)

	fmt.Println("")
	GetMessage(ftEmployee.Person)
}
