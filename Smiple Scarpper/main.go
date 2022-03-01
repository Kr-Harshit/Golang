package main

import (
	"fmt"
	"os"
)

func main() {

	foundUrls := make(map[string]bool)
	SeedUrls := os.Args[1:]

	chUrls := make(chan string)
	chFinished := make(chan bool)

	for _, url := range SeedUrls {
		go crawl(url, chUrls, chFinished)
	}

	for c := 0; c <= len(SeedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}

	fmt.Println("\nFound", len(foundUrls), "Unique urls.\n")

	for url, _ := range foundUrls {
		fmt.Println(url)
	}

	close(chUrls)
}

func crawl(url string, chUrls chan string, chFinished chan string) {

}
