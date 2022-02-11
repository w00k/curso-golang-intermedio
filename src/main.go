package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
	dni  string
}

type Employee struct {
	id       int
	position string
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

// para testcase
// GetPersonByDni simula la búsqueda en bdd de una persona por el dni
var GetPersonByDni = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM ...
	return Person{}, nil
}

// para testcase
// GetEmployeeById simula la búsqueda en bdd de un empleado por el id
var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM ...
	return Employee{}, nil
}

// para el testcase
// GetFullTimeEmployeeById obtiene los datos de la persona y del empleado
func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e

	p, err := GetPersonByDni(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p

	return ftEmployee, nil
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
