package main

import "fmt"

// funciones variádicas, son para poder enviar varios valores
// sin necesidad de crear una función por cantidad de inputs
// es decir sum2, sum3,sum4, etc
// el input es tratado como un slice
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

// función variádica, que imprime los nombres
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// función con retorno con nombre
func getValues(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(sum(2, 3, 4, 5, 6, 7, 8, 9))

	printNames("Francisca", "Simón", "Luca")

	fmt.Println(getValues(2))
}
