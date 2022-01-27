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

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

// PrintInfo es la interface
type PrintInfo interface {
	getMessage() string
}

// getMessage implementa la interface con el getMessage del tipo correspondiente
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

// getMessage para el FullTimeEmployee
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

// getMessage para el TemporaryEmployee
func (temporaryEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

// interfaces se utilizan en casi todos los patrones de diseño y también se usan para hacer  en Go/Golang
func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Francisco"
	ftEmployee.age = 39
	ftEmployee.id = 8

	temporaryEmployee := TemporaryEmployee{}

	// GetMessage
	getMessage(ftEmployee)
	getMessage(temporaryEmployee)
}
