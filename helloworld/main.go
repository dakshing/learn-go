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

	printAnything(42)
	printAnything("Hello")
	printAnything(true)
}

func printAnything(a any) {
	switch v := a.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}
