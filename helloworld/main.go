package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"one", "two", "three"}
	slice = append(slice, "four", "five")

	for i, element := range slice {
		fmt.Println(i, element)
	}
	fmt.Println(strings.Join(slice[:2], ", "))
	fmt.Println("Hello, World!")
}
