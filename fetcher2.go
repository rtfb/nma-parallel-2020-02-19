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
		"https://news.ycombinator.com/",
		"https://google.com",
		"https://github.com",
	}
	wg.Add(len(urls))
	results := make(chan string) // HL
	for _, u := range urls {
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				results <- fmt.Sprintf("ERR - %s", url) // HL
				return
			}
			defer resp.Body.Close()
			results <- fmt.Sprintf("%s - %s", resp.Status, url) // HL
		}(u)
	}
	wg.Wait()
}

// END OMIT
