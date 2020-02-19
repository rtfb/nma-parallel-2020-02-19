package main

import (
	"fmt"
	"net/http"
	"sync"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://news.ycombinator.com/", // HL
		"https://google.com",            // HL
		"https://github.com",            // HL
	}
	wg.Add(len(urls))
	for _, u := range urls {
		go func(url string) { // HL
			defer wg.Done()
			resp, err := http.Get(url) // HL
			if err != nil {
				fmt.Printf("ERR - %s\n", url)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("%s - %s\n", resp.Status, url) // HL
		}(u)
	}
	wg.Wait()
}

// END OMIT
