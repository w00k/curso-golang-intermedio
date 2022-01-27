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

// NewEmployee constructor
func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func main() {
	// constructor
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

	// constructor
	employee := Employee{
		id:   5,
		name: "Pancho",
	}

	fmt.Printf("\n%v\n", employee)

	// constructor
	newEmployee := new(Employee)
	fmt.Printf("%v\n", *newEmployee)
	newEmployee.id = 10
	newEmployee.name = "Leif"
	fmt.Printf("%v\n", *newEmployee)

	employeeFunction := NewEmployee(11, "Clark")
	fmt.Printf("%v\n", *employeeFunction)
}
