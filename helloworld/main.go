package main

import "fmt"

func main() {
	slice := []string{"one", "two", "three"}
	slice = append(slice, "four", "five")

	for i, element := range slice {
		fmt.Println(i, element)
	}
	fmt.Println("Hello, World!")
}
