package main

import (
	"fmt"
	"io"
	"net/http"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "¡Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Overloaded function example (not supported in Go)
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// HTTP Get - example
	response, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response status:", response.Status)
	// bs := make([]byte, 99999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, response.Body)

	// Custom Writer example
	lw := logWriter{}
	io.Copy(lw, response.Body)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
