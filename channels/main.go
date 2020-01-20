package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Works just like any other variable type
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	// this is a blocking call
	// 	fmt.Println(<-c)
	// }

	// This is an infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative syntax for an infinite for loop
	// Equivalent to the above for loop
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// This checks one URL at a time
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
