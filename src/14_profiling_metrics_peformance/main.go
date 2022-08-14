package main

import "fmt"

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	x := Sum(5, 5)
	fmt.Println(x)

	y := GetMax(3, 11)
	fmt.Println(y)

	f := Fibonacci(50)
	fmt.Println(f)
}
