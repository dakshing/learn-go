package main

import "fmt"

func main() {
	// colors := make(map[int]string)

	m := map[int]string{
		1: "abc",
		2: "xyz",
	}

	fmt.Println(m)

	m[3] = "def"
	fmt.Println(m)

	delete(m, 2)
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}
}
