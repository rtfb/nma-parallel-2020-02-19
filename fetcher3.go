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
	results := make(chan string)
	go func() { // HL
		for r := range results { // HL
			fmt.Println(r) // HL
		} // HL
	}() // HL
	for _, u := range urls {
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)

			// END OMIT
			if err != nil {
				results <- fmt.Sprintf("ERR - %s", url)
				return
			}
			defer resp.Body.Close()
			results <- fmt.Sprintf("%s - %s", resp.Status, url)
		}(u)
	}
	wg.Wait()
}
