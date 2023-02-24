package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR:%s -> %s", url, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("ERROR:%s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %v", url, duration)
}

func main() {
	//siteTime("https://google.com")
	var wg = sync.WaitGroup{}

	urls := []string{
		"https://google.com",
		"https://apple.com",
		"https://no-such-site.com",
	}

	wg.Add(len(urls))

	for _, url := range urls {
		url := url

		go func() {
			defer wg.Done()
			siteTime(url)
		}()
	}

	wg.Wait()
}
