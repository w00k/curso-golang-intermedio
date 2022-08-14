package main

import "testing"

func TestGetFullEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id:       1,
						position: "CEO",
					}, nil
				}

				GetPersonByDni = func(id string) (Person, error) {
					return Person{
						name: "Francisco",
						age:  39,
						dni:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{},
		},
	}

	for _, item := range table {
		employeeById := item
		if employeeById.expectedEmployee.name == "Francisco" && employeeById.expectedEmployee.position == "CEO" {
			t.Errorf("GetMax was incorrect, get %s and expected %s", employeeById.expectedEmployee.name, "Francisco")
			t.Errorf("GetMax was incorrect, get %s and expected %s", employeeById.expectedEmployee.position, "CEO")
		}
	}
}
