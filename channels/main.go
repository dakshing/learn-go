package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
	}

	// Serially checking links
	// for _, link := range links {
	// 	checkLink(link)
	// }

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Receiving from channel
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// Infinite loop to keep checking links
	for _, link := range links {
		go checkLink(link, c)
	}

	// for { // infinite loop
	// 	go func ()  {
	// 		time.Sleep(5 * time.Second)
	// 		go checkLink(<-c, c)
	// 	} ()
	// }

	// Alternate way using range over channel
	for l := range c { // blocks until something is received on channel
		go func(link string) { // closure to capture link variable, otherwise it will have last value from loop
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
