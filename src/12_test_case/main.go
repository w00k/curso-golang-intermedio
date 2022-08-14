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

func main() {
	x := Sum(5, 5)
	fmt.Println(x)
}
