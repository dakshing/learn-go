package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(add(2.5, 3.5))
}

func add[T int | float64](a, b T) T {
	return a + b
}
